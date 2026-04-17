package avatar

import "strings"

// Non-color class strings preserved verbatim from catalyst-ui-kit/
// typescript/avatar.tsx. Catalyst's black/10 + dark white/10 outline
// collapses to a single foreground/10 token; dark-mode inversion
// happens via the semantic contract without a dark: variant.

// rootClasses: inline-grid keeps image + initials stacked in one
// cell. *:col-start-1 *:row-start-1 pins both children to the same
// cell so the img overlays the initials on load, and the initials
// bridge the gap while loading. [--avatar-radius:20%] is Catalyst's
// square-corner radius, applied via rounded-(--avatar-radius).
func rootClasses(p Props) string {
	parts := []string{
		"inline-grid shrink-0 align-middle",
		"[--avatar-radius:20%]",
		"*:col-start-1 *:row-start-1",
		"outline -outline-offset-1 outline-foreground/10",
	}
	if p.Square {
		parts = append(parts, "rounded-(--avatar-radius) *:rounded-(--avatar-radius)")
	} else {
		parts = append(parts, "rounded-full *:rounded-full")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// buttonClasses: relative + inline-grid so the TouchTarget overlay
// can sit absolute-positioned within. Focus ring uses the ring token
// (via outline) so each client's brand accent drives the focus
// color. data-focus is aliased to :focus-visible in flint.css, so
// the ring only appears for keyboard users.
func buttonClasses(p ButtonProps) string {
	parts := []string{
		"relative inline-grid",
		"focus:not-data-focus:outline-hidden",
		"data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring",
	}
	if p.Square {
		parts = append(parts, "rounded-[20%]")
	} else {
		parts = append(parts, "rounded-full")
	}
	if p.Disabled {
		parts = append(parts, "opacity-50 pointer-events-none")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
