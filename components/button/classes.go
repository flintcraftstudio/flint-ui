package button

import "strings"

// Non-color class strings are preserved verbatim from catalyst-ui-kit/typescript/button.tsx.
// Color classes are translated to the semantic token contract defined in
// styles/flint.css (see flintcraft-ui-conversion-guide.md § "Semantic Color Tokens").

const baseClasses = "relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold " +
	"px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 " +
	"focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring " +
	"data-disabled:opacity-50 " +
	"*:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText]"

// solidClasses implement Catalyst's two-layer "optical border + hover overlay"
// button look. The per-variant map below sets --btn-bg / --btn-border /
// --btn-hover-overlay / --btn-icon, which this structural template references.
// The after: layer's 1px top highlight is white for every variant, as in
// Catalyst: it's a lighting effect, not a brand color. Deriving it from
// primary-foreground drew a dark line across light buttons in dark mode.
const solidClasses = "border-transparent bg-(--btn-border) " +
	"before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) " +
	"before:shadow-sm " +
	"after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] " +
	"after:shadow-[inset_0_1px_--theme(--color-white/15%)] " +
	"data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) " +
	"data-disabled:before:shadow-none data-disabled:after:shadow-none"

// outlineClasses and plainClasses are the structure of the two unfilled
// styles; their color comes from neutralTintClasses or tintClasses.
const outlineClasses = "border-border"

const plainClasses = "border-transparent"

// neutralTintClasses color an outline or plain button with no Variant.
const neutralTintClasses = "text-foreground data-active:bg-muted data-hover:bg-muted " +
	"[--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"

// tintClasses color an outline or plain button with an explicit Variant:
// the label in the hue's text token, the icon in the same token (muted at
// rest, full on hover), and a light tint of the hue behind it on hover.
// The text tokens and tint strengths are the ones badges and alerts use,
// so styles/contrast_test.go covers them on their own hover tint. Primary
// has no -text mix; the primary hue itself passes on its /10 tint.
var tintClasses = map[Variant]string{
	VariantPrimary: "text-primary data-active:bg-primary/10 data-hover:bg-primary/10 " +
		"[--btn-icon:var(--color-primary)]/70 data-active:[--btn-icon:var(--color-primary)] data-hover:[--btn-icon:var(--color-primary)]",

	VariantAccent: "text-accent-text data-active:bg-accent/10 data-hover:bg-accent/10 " +
		"[--btn-icon:var(--color-accent-text)]/70 data-active:[--btn-icon:var(--color-accent-text)] data-hover:[--btn-icon:var(--color-accent-text)]",

	VariantDanger: "text-danger-text data-active:bg-danger/10 data-hover:bg-danger/10 " +
		"[--btn-icon:var(--color-danger-text)]/70 data-active:[--btn-icon:var(--color-danger-text)] data-hover:[--btn-icon:var(--color-danger-text)]",

	VariantSuccess: "text-success-text data-active:bg-success/10 data-hover:bg-success/10 " +
		"[--btn-icon:var(--color-success-text)]/70 data-active:[--btn-icon:var(--color-success-text)] data-hover:[--btn-icon:var(--color-success-text)]",

	VariantWarning: "text-warning-text data-active:bg-warning/25 data-hover:bg-warning/25 " +
		"[--btn-icon:var(--color-warning-text)]/70 data-active:[--btn-icon:var(--color-warning-text)] data-hover:[--btn-icon:var(--color-warning-text)]",
}

// variantClasses sets the CSS custom properties that solidClasses reads.
// Every variant has the same structure: foreground text, --btn-bg from the
// semantic token, --btn-border matching, a 10%-alpha hover overlay in the
// foreground color, and a slightly-muted icon tint.
var variantClasses = map[Variant]string{
	VariantPrimary: "text-primary-foreground [--btn-bg:var(--color-primary)] [--btn-border:var(--color-primary)] [--btn-hover-overlay:var(--color-primary-foreground)]/10 " +
		"[--btn-icon:var(--color-primary-foreground)]/70 data-active:[--btn-icon:var(--color-primary-foreground)] data-hover:[--btn-icon:var(--color-primary-foreground)]",

	VariantAccent: "text-accent-foreground [--btn-bg:var(--color-accent)] [--btn-border:var(--color-accent)] [--btn-hover-overlay:var(--color-accent-foreground)]/10 " +
		"[--btn-icon:var(--color-accent-foreground)]/70 data-active:[--btn-icon:var(--color-accent-foreground)] data-hover:[--btn-icon:var(--color-accent-foreground)]",

	VariantDanger: "text-danger-foreground [--btn-bg:var(--color-danger)] [--btn-border:var(--color-danger)] [--btn-hover-overlay:var(--color-danger-foreground)]/10 " +
		"[--btn-icon:var(--color-danger-foreground)]/70 data-active:[--btn-icon:var(--color-danger-foreground)] data-hover:[--btn-icon:var(--color-danger-foreground)]",

	VariantSuccess: "text-success-foreground [--btn-bg:var(--color-success)] [--btn-border:var(--color-success)] [--btn-hover-overlay:var(--color-success-foreground)]/10 " +
		"[--btn-icon:var(--color-success-foreground)]/70 data-active:[--btn-icon:var(--color-success-foreground)] data-hover:[--btn-icon:var(--color-success-foreground)]",

	VariantWarning: "text-warning-foreground [--btn-bg:var(--color-warning)] [--btn-border:var(--color-warning)] [--btn-hover-overlay:var(--color-warning-foreground)]/25 " +
		"[--btn-icon:var(--color-warning-foreground)]/70 data-active:[--btn-icon:var(--color-warning-foreground)] data-hover:[--btn-icon:var(--color-warning-foreground)]",
}

// indicatorClasses hide the label (text and element children alike) and
// reveal the spinner slot while an htmx request is in flight. The label
// stays in the layout so the button keeps its width. pointer-events-none
// swallows repeat clicks; hx-disable is still the real double-submit guard,
// and since it sets disabled, opacity-100 keeps a busy button from looking
// unavailable.
const indicatorClasses = "htmx-request:pointer-events-none htmx-request:text-transparent htmx-request:opacity-100 " +
	"htmx-request:*:not-data-[slot=spinner]:invisible htmx-request:*:data-[slot=spinner]:flex"

// indicatorSlotClasses center the spinner over the (hidden) label. It
// takes the variant's icon tint, since the button's own text color is
// transparent during the request.
const indicatorSlotClasses = "absolute inset-0 hidden items-center justify-center text-(--btn-icon)"

// Classes returns the full className string for the button, combining the
// base, the selected style (outline / plain / solid-with-variant), and any
// caller-supplied extras. Outline and plain are neutral unless Variant is
// set, which tints them; solid treats an unset Variant as primary.
func Classes(p Props) string {
	parts := []string{baseClasses}

	switch {
	case p.Outline || p.Plain:
		if p.Outline {
			parts = append(parts, outlineClasses)
		} else {
			parts = append(parts, plainClasses)
		}
		// Unlike solid, an unset (or unknown) Variant means neutral here.
		if tc, ok := tintClasses[p.Variant]; ok {
			parts = append(parts, tc)
		} else {
			parts = append(parts, neutralTintClasses)
		}
	default:
		v := p.Variant
		if v == "" {
			v = VariantPrimary
		}
		if vc, ok := variantClasses[v]; ok {
			parts = append(parts, solidClasses, vc)
		} else {
			parts = append(parts, solidClasses, variantClasses[VariantPrimary])
		}
	}

	if p.Indicator {
		parts = append(parts, indicatorClasses)
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// buttonElClasses adds Catalyst's cursor-default that normally comes from
// Headless UI's <Button> wrapper — applied only to <button>, not <a>.
func buttonElClasses(p Props) string {
	return Classes(p) + " cursor-default"
}

func resolveType(t string) string {
	if t == "" {
		return "button"
	}
	return t
}
