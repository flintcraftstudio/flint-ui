package breadcrumbs

import "strings"

// From-scratch — no Catalyst source to track. Link color uses the
// muted-foreground token so the trail recedes visually; hover flips
// to the full foreground token. Current-page styling adds weight
// (font-medium) so "you are here" reads as the emphasized step.

func ariaLabelOr(label string) string {
	if label != "" {
		return label
	}
	return "Breadcrumb"
}

func rootClasses(p Props) string {
	parts := []string{"text-sm"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// linkClasses: muted until hover. focus:outline-hidden + data-focus:
// (aliased to :focus-visible in flint.css) produce a subtle underline
// so keyboard users see a clear focus indicator without a heavy ring.
func linkClasses(p ItemProps) string {
	parts := []string{
		"text-muted-foreground hover:text-foreground transition-colors",
		"focus:outline-hidden data-focus:text-foreground data-focus:underline data-focus:underline-offset-4",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// currentClasses: full foreground + medium weight. Same text size as
// the links — differentiation is weight + color, not scale.
func currentClasses(p CurrentProps) string {
	parts := []string{"font-medium text-foreground"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
