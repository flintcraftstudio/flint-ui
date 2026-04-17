package popover

import (
	"strings"

	"github.com/a-h/templ"
)

// Anchor and positioning logic mirrors Dropdown's class-by-class. When
// a third consumer (HoverCard, Menu) lands we should factor anchor
// positioning into a shared components/floating/ helper — until then
// the duplication is acceptable because Dropdown's menu-specific classes
// (subgrid, role="menu" styling) don't fit a generic helper anyway.

// rootClasses sets up the relative-positioned wrapper that anchors the
// absolute Panel. inline-block keeps the popover sized to its trigger.
func rootClasses(p Props) string {
	parts := []string{"relative inline-block"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// triggerAttrs is the Alpine wiring layered onto the underlying button.
// Caller-supplied Attrs are merged in last so they win on key conflicts.
// aria-haspopup="dialog" signals "generic popover" to assistive tech —
// as opposed to Dropdown's aria-haspopup="menu".
func triggerAttrs(extra templ.Attributes) templ.Attributes {
	out := templ.Attributes{
		"x-on:click":           "open = !open",
		"x-bind:aria-expanded": "open",
		"aria-haspopup":        "dialog",
	}
	for k, v := range extra {
		out[k] = v
	}
	return out
}

// anchorClasses positions the panel relative to the popover root.
// start pins the panel's left edge to the trigger's left; end pins the
// right edges. top/bottom pick which side of the trigger the panel
// appears on. origin-* sets the scale pivot for the enter/leave
// transition so the panel appears to emerge from the trigger.
var anchorClasses = map[Anchor]string{
	AnchorBottomStart: "left-0 top-full mt-2 origin-top-left",
	AnchorBottomEnd:   "right-0 top-full mt-2 origin-top-right",
	AnchorTopStart:    "left-0 bottom-full mb-2 origin-bottom-left",
	AnchorTopEnd:      "right-0 bottom-full mb-2 origin-bottom-right",
}

func resolveAnchor(a Anchor) string {
	if c, ok := anchorClasses[a]; ok {
		return c
	}
	return anchorClasses[AnchorBottomStart]
}

// panelClasses: chrome + default padding. bg-surface / ring / shadow
// match Dropdown's Menu (and Modal's panel) so popovers, menus, and
// dialogs sit in the same visual family. Default p-4 covers the common
// case of rich inline content; callers needing edge-to-edge content
// (an image, a code block with its own padding) pass Class="p-0".
//
// min-w-[16rem] keeps short-content popovers from looking cramped;
// max-w-xs prevents arbitrary-width explosions when a popover holds a
// long-form description. Callers override via Class for anything
// outside this range.
func panelClasses(p PanelProps) string {
	parts := []string{
		// Position
		"absolute z-50",
		resolveAnchor(p.Anchor),
		// Size
		"min-w-[16rem] max-w-xs",
		// Shape
		"rounded-xl",
		// Padding
		"p-4",
		// Background + chrome
		"bg-surface",
		"shadow-lg ring-1 ring-border",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
