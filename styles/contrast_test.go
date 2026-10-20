// Package styles holds flint-ui's Tailwind partial. This test parses its
// token values and checks WCAG AA contrast for every text/background pair
// the components render, in light and dark mode.
package styles

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// rgba is a color in sRGB, components in 0..1.
type rgba struct{ r, g, b, a float64 }

// layer is one paint layer: a token at an opacity modifier (bg-x/15 → 0.15).
type layer struct {
	token string
	alpha float64
}

func solid(token string) layer { return layer{token, 1} }

// pair is text (or a non-text indicator) painted over a stack of layers,
// bottom first. The bottom layer must be opaque.
type pair struct {
	name string
	fg   layer
	bg   []layer
	min  float64
}

const (
	aaText    = 4.5 // WCAG 1.4.3, body text
	aaNonText = 3.0 // WCAG 1.4.11, focus rings and other indicators
)

// pairs lists what the components actually render. When a component adds
// a new text color or background, add the pair here.
func pairs() []pair {
	var ps []pair
	text := func(name string, fg layer, bg ...layer) {
		ps = append(ps, pair{name, fg, bg, aaText})
	}

	// Body copy and secondary text on every neutral background.
	for _, bg := range []string{"background", "surface", "muted", "canvas"} {
		text("foreground on "+bg, solid("foreground"), solid(bg))
		text("muted-foreground on "+bg, solid("muted-foreground"), solid(bg))
	}
	text("surface-foreground on surface", solid("surface-foreground"), solid("surface"))
	// Inputs paint bg-input over the surface they sit on (translucent in dark).
	text("foreground on input", solid("foreground"), solid("surface"), solid("input"))
	text("placeholder on input", solid("muted-foreground"), solid("surface"), solid("input"))
	// Tooltip inverts the neutrals.
	text("background on foreground", solid("background"), solid("foreground"))

	// Solid fills: buttons, checked controls, avatars, menu highlights.
	for _, h := range []string{"primary", "accent", "danger", "success", "warning"} {
		text(h+"-foreground on "+h, solid(h+"-foreground"), solid(h))
	}

	// Hue as text on a plain surface: error messages, "Copied", links.
	for _, h := range []string{"primary", "danger", "success"} {
		for _, bg := range []string{"background", "surface"} {
			text(h+" on "+bg, solid(h), solid(bg))
		}
	}

	// Tinted badges (/15) and alerts (/10, description at 90% opacity),
	// on either page background.
	tinted := map[string]string{
		"primary": "primary",
		"accent":  "accent-text",
		"danger":  "danger-text",
		"success": "success-text",
	}
	for _, base := range []string{"background", "surface"} {
		for h, fg := range tinted {
			for _, a := range []float64{0.10, 0.15} {
				text(fmt.Sprintf("%s on %s/%g over %s", fg, h, a*100, base),
					solid(fg), solid(base), layer{h, a})
			}
			text(fmt.Sprintf("%s/90 on %s/10 over %s", fg, h, base),
				layer{fg, 0.9}, solid(base), layer{h, 0.10})
		}
		text("warning-text on warning/25 over "+base,
			solid("warning-text"), solid(base), layer{"warning", 0.25})
		text("warning-text/90 on warning/25 over "+base,
			layer{"warning-text", 0.9}, solid(base), layer{"warning", 0.25})
	}

	// Focus ring against whatever it sits on.
	for _, bg := range []string{"background", "surface", "canvas"} {
		ps = append(ps, pair{"ring on " + bg, solid("ring"), []layer{solid(bg)}, aaNonText})
	}
	return ps
}

func TestContrast(t *testing.T) {
	src, err := os.ReadFile("flint-ui.css")
	if err != nil {
		t.Fatal(err)
	}
	light, dark, err := parseThemes(string(src))
	if err != nil {
		t.Fatal(err)
	}
	for _, th := range []struct {
		name   string
		tokens map[string]string
	}{{"light", light}, {"dark", dark}} {
		t.Run(th.name, func(t *testing.T) {
			r := resolver{tokens: th.tokens}
			for _, p := range pairs() {
				got, err := r.contrast(p)
				if err != nil {
					t.Errorf("%s: %v", p.name, err)
					continue
				}
				if got < p.min {
					t.Errorf("%s: %.2f:1, want at least %.1f:1", p.name, got, p.min)
				}
			}
		})
	}
}

// TestDarkOverridesEveryColor guards against a new light token with no
// dark value. Derived tokens (the -text mixes) need no dark entry; they
// recompute from the dark values they reference.
func TestDarkOverridesEveryColor(t *testing.T) {
	src, err := os.ReadFile("flint-ui.css")
	if err != nil {
		t.Fatal(err)
	}
	lightOnly, darkOnly, err := themeBlocks(string(src))
	if err != nil {
		t.Fatal(err)
	}
	for name, v := range lightOnly {
		if strings.Contains(v, "var(") {
			continue
		}
		if _, ok := darkOnly[name]; !ok {
			t.Errorf("--color-%s has a light value but no dark value", name)
		}
	}
}

// --- parsing ---

var declRe = regexp.MustCompile(`--color-([a-z-]+):\s*([^;]+);`)

// blockAfter returns the contents of the {...} block that follows marker.
func blockAfter(src, marker string) (string, error) {
	i := strings.Index(src, marker)
	if i < 0 {
		return "", fmt.Errorf("%q not found", marker)
	}
	start := strings.Index(src[i:], "{")
	if start < 0 {
		return "", fmt.Errorf("no block after %q", marker)
	}
	start += i
	depth := 0
	for j := start; j < len(src); j++ {
		switch src[j] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return src[start+1 : j], nil
			}
		}
	}
	return "", fmt.Errorf("unterminated block after %q", marker)
}

func decls(block string) map[string]string {
	m := map[string]string{}
	for _, d := range declRe.FindAllStringSubmatch(block, -1) {
		m[d[1]] = strings.TrimSpace(d[2])
	}
	return m
}

// themeBlocks returns the raw declarations of the @theme block and of the
// dark block.
func themeBlocks(src string) (light, dark map[string]string, err error) {
	// Comments can mention the markers; drop them before searching.
	src = regexp.MustCompile(`(?s)/\*.*?\*/`).ReplaceAllString(src, "")
	lb, err := blockAfter(src, "@theme")
	if err != nil {
		return nil, nil, err
	}
	db, err := blockAfter(src, "@variant dark")
	if err != nil {
		return nil, nil, err
	}
	return decls(lb), decls(db), nil
}

// parseThemes returns the full token set for each theme: dark is the
// light set with the dark block's values laid over it, as the cascade does.
func parseThemes(src string) (light, dark map[string]string, err error) {
	light, darkOnly, err := themeBlocks(src)
	if err != nil {
		return nil, nil, err
	}
	dark = map[string]string{}
	for k, v := range light {
		dark[k] = v
	}
	for k, v := range darkOnly {
		dark[k] = v
	}
	return light, dark, nil
}

// --- color evaluation ---

type resolver struct {
	tokens map[string]string
}

func (r resolver) token(name string) (rgba, error) {
	v, ok := r.tokens[name]
	if !ok {
		return rgba{}, fmt.Errorf("unknown token --color-%s", name)
	}
	c, err := r.eval(v)
	if err != nil {
		return rgba{}, fmt.Errorf("--color-%s: %w", name, err)
	}
	return c, nil
}

var (
	varRe = regexp.MustCompile(`^var\(--color-([a-z-]+)\)$`)
	rgbRe = regexp.MustCompile(`^rgb\(\s*([\d.]+)\s+([\d.]+)\s+([\d.]+)\s*(?:/\s*([\d.]+)(%?))?\s*\)$`)
	mixRe = regexp.MustCompile(`^color-mix\(in oklab,\s*(var\([^)]*\))\s+([\d.]+)%,\s*(var\([^)]*\))\)$`)
)

// eval understands the value forms flint-ui uses: hex, rgb() with optional
// alpha, var(--color-x), and color-mix(in oklab, var() N%, var()).
func (r resolver) eval(v string) (rgba, error) {
	v = strings.TrimSpace(v)
	if strings.HasPrefix(v, "#") {
		return parseHex(v)
	}
	if m := varRe.FindStringSubmatch(v); m != nil {
		return r.token(m[1])
	}
	if m := rgbRe.FindStringSubmatch(v); m != nil {
		c := rgba{a: 1}
		for i, p := range []*float64{&c.r, &c.g, &c.b} {
			f, _ := strconv.ParseFloat(m[i+1], 64)
			*p = f / 255
		}
		if m[4] != "" {
			f, _ := strconv.ParseFloat(m[4], 64)
			if m[5] == "%" {
				f /= 100
			}
			c.a = f
		}
		return c, nil
	}
	if m := mixRe.FindStringSubmatch(v); m != nil {
		a, err := r.eval(m[1])
		if err != nil {
			return rgba{}, err
		}
		b, err := r.eval(m[3])
		if err != nil {
			return rgba{}, err
		}
		p, _ := strconv.ParseFloat(m[2], 64)
		return mixOklab(a, b, p/100), nil
	}
	return rgba{}, fmt.Errorf("can't evaluate %q", v)
}

func parseHex(s string) (rgba, error) {
	h := strings.TrimPrefix(s, "#")
	if len(h) == 3 {
		h = string([]byte{h[0], h[0], h[1], h[1], h[2], h[2]})
	}
	if len(h) != 6 {
		return rgba{}, fmt.Errorf("bad hex color %q", s)
	}
	n, err := strconv.ParseUint(h, 16, 32)
	if err != nil {
		return rgba{}, fmt.Errorf("bad hex color %q", s)
	}
	return rgba{float64(n>>16&0xff) / 255, float64(n>>8&0xff) / 255, float64(n&0xff) / 255, 1}, nil
}

// contrast flattens the background stack, paints the foreground over it,
// and returns the WCAG contrast ratio.
func (r resolver) contrast(p pair) (float64, error) {
	var bg rgba
	for i, l := range p.bg {
		c, err := r.token(l.token)
		if err != nil {
			return 0, err
		}
		c.a *= l.alpha
		if i == 0 {
			if c.a < 1 {
				return 0, fmt.Errorf("bottom layer --color-%s is translucent", l.token)
			}
			bg = c
			continue
		}
		bg = over(c, bg)
	}
	fg, err := r.token(p.fg.token)
	if err != nil {
		return 0, err
	}
	fg.a *= p.fg.alpha
	fg = over(fg, bg)

	l1, l2 := luminance(fg), luminance(bg)
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05), nil
}

// over composites c onto an opaque base.
func over(c, base rgba) rgba {
	return rgba{
		c.r*c.a + base.r*(1-c.a),
		c.g*c.a + base.g*(1-c.a),
		c.b*c.a + base.b*(1-c.a),
		1,
	}
}

func toLinear(c float64) float64 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func fromLinear(c float64) float64 {
	c = math.Max(0, math.Min(1, c))
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return 1.055*math.Pow(c, 1/2.4) - 0.055
}

func luminance(c rgba) float64 {
	return 0.2126*toLinear(c.r) + 0.7152*toLinear(c.g) + 0.0722*toLinear(c.b)
}

// mixOklab is CSS color-mix(in oklab, a p, b) for opaque colors.
func mixOklab(a, b rgba, p float64) rgba {
	la, lb := toOklab(a), toOklab(b)
	var m [3]float64
	for i := range m {
		m[i] = la[i]*p + lb[i]*(1-p)
	}
	c := fromOklab(m)
	c.a = a.a*p + b.a*(1-p)
	return c
}

func toOklab(c rgba) [3]float64 {
	r, g, b := toLinear(c.r), toLinear(c.g), toLinear(c.b)
	l := math.Cbrt(0.4122214708*r + 0.5363325363*g + 0.0514459929*b)
	m := math.Cbrt(0.2119034982*r + 0.6806995451*g + 0.1073969566*b)
	s := math.Cbrt(0.0883024619*r + 0.2817188376*g + 0.6299787005*b)
	return [3]float64{
		0.2104542553*l + 0.7936177850*m - 0.0040720468*s,
		1.9779984951*l - 2.4285922050*m + 0.4505937099*s,
		0.0259040371*l + 0.7827717662*m - 0.8086757660*s,
	}
}

func fromOklab(lab [3]float64) rgba {
	l := lab[0] + 0.3963377774*lab[1] + 0.2158037573*lab[2]
	m := lab[0] - 0.1055613458*lab[1] - 0.0638541728*lab[2]
	s := lab[0] - 0.0894841775*lab[1] - 1.2914855480*lab[2]
	l, m, s = l*l*l, m*m*m, s*s*s
	return rgba{
		fromLinear(4.0767416621*l - 3.3077115913*m + 0.2309699292*s),
		fromLinear(-1.2684380046*l + 2.6097574011*m - 0.3413193965*s),
		fromLinear(-0.0041960863*l - 0.7034186147*m + 1.7076147010*s),
		1,
	}
}
