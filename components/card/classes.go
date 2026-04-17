package card

import "strings"

// Non-color classes are flint-ui originals — no Catalyst source to track.
// Surface tokens match Modal and Dropdown (bg-surface + ring-1 +
// ring-border) so cards sit in the same visual family as other panels.

// rootClasses sets up the outer chrome. No padding — sections bring
// their own. `overflow-hidden` keeps the rounded corners clean when a
// Header or Footer hits the edge with its own background.
func rootClasses(p Props) string {
	parts := []string{
		"overflow-hidden", "rounded-xl",
		"bg-surface",
		"ring-1", "ring-border",
		"shadow-sm",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// headerClasses: padded top section with a bottom divider. The pattern
// is title + description inside, but the component doesn't enforce that
// — Header is just a padded bordered box.
func headerClasses(p SubProps) string {
	parts := []string{
		"px-6 py-4",
		"border-b", "border-border",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// bodyClasses: padded middle section, no borders.
func bodyClasses(p SubProps) string {
	parts := []string{
		"px-6 py-5",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// footerClasses: padded bottom section with a top divider. Often an
// actions row — flex + gap inside. Slightly tighter vertical padding
// than Body reads like a "toolbar" rather than content.
func footerClasses(p SubProps) string {
	parts := []string{
		"px-6 py-3",
		"border-t", "border-border",
		"bg-muted/30",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
