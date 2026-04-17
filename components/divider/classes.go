package divider

import "strings"

// Catalyst uses `border-zinc-950/5` (soft) and `border-zinc-950/10`
// (regular). flint-ui translates these to token-derived opacities —
// each client's foreground token drives the line color.

func classes(p Props) string {
	parts := []string{"w-full border-t"}
	if p.Soft {
		parts = append(parts, "border-foreground/5")
	} else {
		parts = append(parts, "border-foreground/10")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
