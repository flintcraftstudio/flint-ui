package alert

import "strings"

// Variant colors follow the same `bg-{token}/15 text-{token}` recipe as
// Badge so the two components share a palette. Warning is the exception:
// the warning token is a yellow that can't be read against a yellow
// tint, so it uses bg-warning/25 and text-warning-foreground — same
// adjustment Badge makes for the same reason.

var variantClasses = map[Variant]string{
	VariantInfo:    "bg-primary/10 text-primary",
	VariantSuccess: "bg-success/10 text-success",
	VariantWarning: "bg-warning/25 text-warning-foreground",
	VariantDanger:  "bg-danger/10 text-danger",
}

func resolveVariant(v Variant) string {
	if c, ok := variantClasses[v]; ok {
		return c
	}
	return variantClasses[VariantInfo]
}

// rootClasses: flex row with the icon on the left and content on the
// right. `items-start` aligns the icon with the first line of the
// title; the icon itself sits in a `pt-0.5` wrapper (see alert.templ)
// to optically center against the title's cap height.
func rootClasses(p Props) string {
	parts := []string{
		"flex items-start gap-3",
		"rounded-lg",
		"px-4 py-3",
		resolveVariant(p.Variant),
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// titleClasses: bold heading, inherits the alert's variant text color.
// Slightly tighter line-height than Description so title + description
// read as a single stanza rather than two paragraphs.
func titleClasses(p SubProps) string {
	parts := []string{
		"text-sm/5 font-semibold",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// descriptionClasses: body text, inherits the alert's variant color at
// slight opacity reduction for visual hierarchy against the title.
// mt-1 pulls Description close to Title — tighter than the default
// prose rhythm because they're the same block of content.
func descriptionClasses(p SubProps) string {
	parts := []string{
		"mt-1 text-sm/5 text-current/90",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// actionsClasses: optional row under Description. flex-wrap so buttons
// stack gracefully on narrow viewports; gap-3 matches the Modal Actions
// spacing so adjacent Alert + Modal feel like the same design system.
func actionsClasses(p SubProps) string {
	parts := []string{
		"mt-3 flex flex-wrap items-center gap-3",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
