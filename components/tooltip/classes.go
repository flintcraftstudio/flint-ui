package tooltip

import "strings"

// Positioning is `absolute` inside the wrapper's `relative inline-flex`
// box. Each anchor pairs a side-pin (top/bottom/left/right-full) with a
// centering translate on the perpendicular axis and a small gap toward
// the trigger. `origin-*` sets the scale pivot to the edge closest to
// the trigger so the enter/leave `scale-95 → scale-100` transition looks
// like the bubble is emerging from the trigger instead of from its own
// center.

var anchorClasses = map[Anchor]string{
	AnchorTop:    "bottom-full left-1/2 -translate-x-1/2 mb-1.5 origin-bottom",
	AnchorBottom: "top-full left-1/2 -translate-x-1/2 mt-1.5 origin-top",
	AnchorLeft:   "right-full top-1/2 -translate-y-1/2 mr-1.5 origin-right",
	AnchorRight:  "left-full top-1/2 -translate-y-1/2 ml-1.5 origin-left",
}

func resolveAnchor(a Anchor) string {
	if c, ok := anchorClasses[a]; ok {
		return c
	}
	return anchorClasses[AnchorTop]
}

// contentClasses translates the tooltip bubble. Colors invert the page
// surface so the hint reads against either light or dark themes: the
// `foreground` token is the default body text color, so `bg-foreground`
// + `text-background` is always high-contrast. pointer-events-none keeps
// the tooltip from eating hover/click events meant for the trigger or
// overlapping content.
func contentClasses(p Props) string {
	parts := []string{
		// Position
		"absolute z-50",
		resolveAnchor(p.Anchor),
		// Size/shape — narrow by default; Class can widen
		"w-max max-w-xs rounded-md px-2 py-1",
		// Typography
		"text-xs font-medium whitespace-normal",
		// Colors — inverted contrast against page
		"bg-foreground text-background",
		// Shadow
		"shadow-md",
		// Tooltip is decoration, not interactive
		"pointer-events-none select-none",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
