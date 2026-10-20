package badge

import "strings"

// Non-color class strings preserved verbatim from catalyst-ui-kit/typescript/badge.tsx.
// Catalyst's 18 color variants are collapsed to the six semantic variants;
// the tinted-background pattern (bg-{color}/15, text-{color}) is preserved.

const baseClasses = "inline-flex items-center gap-x-1.5 rounded-md px-1.5 py-0.5 text-sm/5 font-medium sm:text-xs/5 forced-colors:outline"

// variantClasses maps each variant to a tinted background + readable text.
// Muted uses the muted token pair (no alpha needed since it's already subtle).
// The colored variants use bg-{token}/15 with text-{token}-text — the same
// recipe Catalyst uses for its tinted badges (a -500 tint under -700 text in
// light, -400 text in dark), routed through semantic tokens. The -text tokens
// pull the hue toward foreground so it stays AA on its own tint in both
// themes. Primary is dark enough (light enough, in dark) to need no -text
// token. Warning keeps a stronger /25 tint because yellow is so light.
var variantClasses = map[Variant]string{
	VariantMuted:   "bg-muted text-muted-foreground",
	VariantPrimary: "bg-primary/15 text-primary",
	VariantAccent:  "bg-accent/15 text-accent-text",
	VariantDanger:  "bg-danger/15 text-danger-text",
	VariantSuccess: "bg-success/15 text-success-text",
	VariantWarning: "bg-warning/25 text-warning-text",
}

// Classes returns the full class string for a badge.
func Classes(p Props) string {
	v := p.Variant
	if v == "" {
		v = VariantMuted
	}
	vc, ok := variantClasses[v]
	if !ok {
		vc = variantClasses[VariantMuted]
	}
	parts := []string{baseClasses, vc}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
