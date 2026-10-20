package progress

import (
	"math"
	"strconv"

	"github.com/a-h/templ"
)

var variantClasses = map[Variant]string{
	VariantPrimary: "bg-primary",
	VariantAccent:  "bg-accent",
	VariantSuccess: "bg-success",
	VariantWarning: "bg-warning",
	VariantDanger:  "bg-danger",
}

var sizeClasses = map[Size]string{
	SizeSM: "h-1",
	SizeMD: "h-2",
}

// trackClasses: the full-width muted rail. overflow-hidden clips the
// indeterminate segment as it enters and leaves, and keeps the fill's
// ends inside the rounded track.
func trackClasses(p Props) string {
	size, ok := sizeClasses[p.Size]
	if !ok {
		size = sizeClasses[SizeMD]
	}
	return join("relative w-full overflow-hidden rounded-full bg-muted forced-colors:outline "+size, p.Class)
}

func fillColor(p Props) string {
	if c, ok := variantClasses[p.Variant]; ok {
		return c
	}
	return variantClasses[VariantPrimary]
}

// fillClasses: the determinate fill. The width transition animates value
// changes, including htmx morph swaps that update the style in place.
// forced-colors paints it in the system highlight color.
func fillClasses(p Props) string {
	return "h-full rounded-full transition-[width] duration-300 ease-out motion-reduce:transition-none forced-colors:bg-[Highlight] " + fillColor(p)
}

// indeterminateFillClasses: a 40%-wide segment that loops across the
// track. The animation lives in flint-ui.css (flint-progress-indeterminate)
// along with its reduced-motion fallback.
func indeterminateFillClasses(p Props) string {
	return "flint-progress-indeterminate h-full w-2/5 rounded-full forced-colors:bg-[Highlight] " + fillColor(p)
}

func maxValue(p Props) float64 {
	if p.Max <= 0 {
		return 100
	}
	return p.Max
}

// clamp keeps Value inside 0..Max so a stray server value can't overflow
// the track or report nonsense to assistive tech.
func clamp(p Props) float64 {
	return math.Min(math.Max(p.Value, 0), maxValue(p))
}

func percent(p Props) float64 {
	return clamp(p) / maxValue(p) * 100
}

func valueText(p Props) string {
	if p.ValueText != "" {
		return p.ValueText
	}
	return strconv.Itoa(int(math.Round(percent(p)))) + "%"
}

func fillStyle(p Props) templ.SafeCSS {
	return templ.SafeCSS("width: " + formatNumber(percent(p)) + "%")
}

func formatNumber(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
