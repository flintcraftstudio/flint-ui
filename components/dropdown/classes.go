package dropdown

import (
	"strings"

	"github.com/a-h/templ"
)

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/dropdown.tsx.
// Color classes are translated to the semantic token contract. Headless UI's
// data-focus / data-hover / data-disabled attribute selectors map cleanly to
// the @custom-variant aliases declared in styles/flint.css (see flint.css for
// the exact selector list — data-focus also matches :focus-visible, etc.) so
// Catalyst's data-* utility class strings work unchanged here.

// rootClasses sets up the relative-positioned wrapper that anchors the
// absolute Menu. inline-block keeps the dropdown sized to its trigger.
func rootClasses(p Props) string {
	parts := []string{"relative inline-block"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// triggerAttrs is the Alpine wiring layered onto the underlying button.
// Caller-supplied Attrs are merged in last so they win on key conflicts.
// ArrowDown on the trigger opens the menu and focuses the first item —
// matches Headless UI Menu's keyboard behavior.
func triggerAttrs(extra templ.Attributes) templ.Attributes {
	out := templ.Attributes{
		"x-on:click":                "open = !open",
		"x-bind:aria-expanded":      "open",
		"aria-haspopup":             "menu",
		"x-on:keydown.down.prevent": "open = true; $nextTick(() => $refs.menu && $focus.within($refs.menu).first())",
	}
	for k, v := range extra {
		out[k] = v
	}
	return out
}

// anchorClasses positions the menu relative to the dropdown root. start
// pins the menu's left edge to the trigger's left; end pins the right
// edges. top/bottom pick which side of the trigger the menu appears on.
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

// menuClasses translates Catalyst's MenuItems styles. The supports-[grid-
// template-columns:subgrid] block opts the menu into the subgrid layout
// when the browser supports it; items then declare their own column grid
// that aligns across the whole menu.
func menuClasses(p MenuProps) string {
	parts := []string{
		// Position
		"absolute z-50",
		resolveAnchor(p.Anchor),
		// Base
		"isolate w-max min-w-[12rem] rounded-xl p-1",
		// Invisible border that's only visible in forced-colors mode
		"outline outline-transparent focus:outline-hidden",
		// Scroll if menu won't fit
		"overflow-y-auto",
		// Popover background — translucent surface + blur, matches Catalyst
		"bg-surface/90 backdrop-blur-xl",
		// Shadows + ring
		"shadow-lg ring-1 ring-border",
		// Subgrid plumbing for column alignment across items
		"supports-[grid-template-columns:subgrid]:grid supports-[grid-template-columns:subgrid]:grid-cols-[auto_1fr_1.5rem_0.5rem_auto]",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// itemClasses translates Catalyst's MenuItem class string. data-focus and
// data-hover are aliased to :focus-visible / :hover in flint.css, so the
// hover/keyboard-highlighted state uses the accent token — themable per
// client. The grid bits set up the icon | label | description | shortcut
// columns; subgrid kicks in via the parent menu when supported.
func itemClasses(p ItemProps) string {
	parts := []string{
		// Base layout
		"group cursor-default rounded-lg px-3.5 py-2.5 focus:outline-hidden sm:px-3 sm:py-1.5",
		// Text
		"text-left text-base/6 text-foreground sm:text-sm/6 forced-colors:text-[CanvasText]",
		// Highlight (hover or keyboard focus) uses the accent token
		"data-focus:bg-accent data-focus:text-accent-foreground data-hover:bg-accent data-hover:text-accent-foreground",
		// Disabled
		"data-disabled:opacity-50 data-disabled:pointer-events-none",
		// Forced-colors mode
		"forced-color-adjust-none forced-colors:data-focus:bg-[Highlight] forced-colors:data-focus:text-[HighlightText] forced-colors:data-focus:*:data-[slot=icon]:text-[HighlightText]",
		// Item grid: explicit fallback, subgrid when supported
		"col-span-full grid grid-cols-[auto_1fr_1.5rem_0.5rem_auto] items-center supports-[grid-template-columns:subgrid]:grid-cols-subgrid",
		// Icons
		"*:data-[slot=icon]:col-start-1 *:data-[slot=icon]:row-start-1 *:data-[slot=icon]:mr-2.5 *:data-[slot=icon]:-ml-0.5 *:data-[slot=icon]:size-5 sm:*:data-[slot=icon]:mr-2 sm:*:data-[slot=icon]:size-4",
		"*:data-[slot=icon]:text-muted-foreground data-focus:*:data-[slot=icon]:text-accent-foreground data-hover:*:data-[slot=icon]:text-accent-foreground",
		// Avatar
		"*:data-[slot=avatar]:mr-2.5 *:data-[slot=avatar]:-ml-1 *:data-[slot=avatar]:size-6 sm:*:data-[slot=avatar]:mr-2 sm:*:data-[slot=avatar]:size-5",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func headerClasses(p SubProps) string {
	parts := []string{"col-span-5 px-3.5 pt-2.5 pb-1 sm:px-3"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func sectionClasses(p SubProps) string {
	parts := []string{
		"col-span-full supports-[grid-template-columns:subgrid]:grid supports-[grid-template-columns:subgrid]:grid-cols-[auto_1fr_1.5rem_0.5rem_auto]",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func headingClasses(p SubProps) string {
	parts := []string{
		"col-span-full grid grid-cols-[1fr_auto] gap-x-12 px-3.5 pt-2 pb-1 text-sm/5 font-medium text-muted-foreground sm:px-3 sm:text-xs/5",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func dividerClasses(p SubProps) string {
	parts := []string{"col-span-full mx-3.5 my-1 h-px border-0 bg-border sm:mx-3 forced-colors:bg-[CanvasText]"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func labelClasses(p SubProps) string {
	parts := []string{"col-start-2 row-start-1"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func descriptionClasses(p SubProps) string {
	parts := []string{
		"col-span-2 col-start-2 row-start-2 text-sm/5 text-muted-foreground group-data-focus:text-accent-foreground group-data-hover:text-accent-foreground sm:text-xs/5 forced-colors:group-data-focus:text-[HighlightText]",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func shortcutClasses(p ShortcutProps) string {
	parts := []string{"col-start-5 row-start-1 flex justify-self-end"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// shortcutKeyClasses styles a single key inside a Shortcut. Multi-character
// keys after the first get a left pad so "Ctrl+Shift" reads cleanly.
func shortcutKeyClasses(index int, key string) string {
	parts := []string{
		"min-w-[2ch] text-center font-sans text-muted-foreground capitalize group-data-focus:text-accent-foreground group-data-hover:text-accent-foreground forced-colors:group-data-focus:text-[HighlightText]",
	}
	if index > 0 && len(key) > 1 {
		parts = append(parts, "pl-1")
	}
	return strings.Join(parts, " ")
}
