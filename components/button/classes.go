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
const solidClasses = "border-transparent bg-(--btn-border) " +
	"before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) " +
	"before:shadow-sm " +
	"after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] " +
	"after:shadow-[inset_0_1px_--theme(--color-primary-foreground/15%)] " +
	"data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) " +
	"data-disabled:before:shadow-none data-disabled:after:shadow-none"

const outlineClasses = "border-border text-foreground data-active:bg-muted data-hover:bg-muted " +
	"[--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"

const plainClasses = "border-transparent text-foreground data-active:bg-muted data-hover:bg-muted " +
	"[--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"

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

// Classes returns the full className string for the button, combining the
// base, the selected style (outline / plain / solid-with-variant), and any
// caller-supplied extras.
func Classes(p Props) string {
	parts := []string{baseClasses}

	switch {
	case p.Outline:
		parts = append(parts, outlineClasses)
	case p.Plain:
		parts = append(parts, plainClasses)
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
