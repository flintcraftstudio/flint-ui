package accordion

import "strings"

// triggerID and panelID derive stable element ids from the Item name.
// Matches the modal/tabs pattern — callers don't have to coordinate ids
// manually, and the button's aria-controls lines up with the panel's id.

func triggerID(name string) string {
	return "flint-acc-trigger-" + safeName(name)
}

func panelID(name string) string {
	return "flint-acc-panel-" + safeName(name)
}

// safeName strips characters that would break a JS string literal or an
// HTML id. Accordion names are expected to be stable identifiers like
// "pricing" or "crew:scheduling"; anything else is dropped so a malformed
// name can't break the page. Same approach as tabs.safeName.
func safeName(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '-', r == '_', r == ':', r == '.':
			b.WriteRune(r)
		}
	}
	return b.String()
}

func resolveHeadingLevel(l int) int {
	if l >= 1 && l <= 6 {
		return l
	}
	return 3
}

// alpineData wires the root x-data: `active` (array of open Item names),
// `multi` (TypeMultiple flag baked in at render time), `toggle(name)`,
// and `isOpen(name)`. Centralizing the state on the root is what lets
// TypeSingle close the previously-open item when a new one opens.
//
// `Default` is seeded into active on first paint; use the empty string
// to start with nothing open. For TypeMultiple with multiple defaults,
// callers can x-init their own state via Props.Attrs.
func alpineData(p Props) string {
	multi := "true"
	if p.Type == TypeSingle {
		multi = "false"
	}
	initial := "[]"
	if p.Default != "" {
		initial = "['" + safeName(p.Default) + "']"
	}
	return "{ active: " + initial + ", multi: " + multi + ", toggle(n) { if (this.active.includes(n)) { this.active = this.active.filter(x => x !== n) } else { this.active = this.multi ? [...this.active, n] : [n] } }, isOpen(n) { return this.active.includes(n) } }"
}

// alpineIsOpenExpr / alpineToggleExpr / alpineChevronExpr are injected
// via x-bind: / x-on: / x-bind:class. They reference the root x-data
// state (active, multi, toggle, isOpen) through Alpine's inherited
// scope — same mechanism Tabs' sub-components use.

func alpineIsOpenExpr(name string) string {
	return "isOpen('" + safeName(name) + "')"
}

func alpineToggleExpr(name string) string {
	return "toggle('" + safeName(name) + "')"
}

// alpineChevronExpr returns the object-form x-bind:class expression:
// when the item is open, add `rotate-180`; otherwise no extra classes.
// Pairs with the base `transition-transform` on the svg so the rotation
// animates.
func alpineChevronExpr(name string) string {
	return "{ 'rotate-180': isOpen('" + safeName(name) + "') }"
}

func rootClasses(p Props) string {
	parts := []string{"flex flex-col"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// itemClasses draws a thin border between items. First item gets a top
// border too via `first:border-t`, so the stack reads as a clean set
// regardless of surrounding layout. Override via Class if the accordion
// sits inside a card that already provides its own borders.
func itemClasses(p ItemProps) string {
	parts := []string{
		"border-b border-border first:border-t",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// triggerClasses: full-width button, left-aligned heading with a
// right-aligned chevron. Hover uses the muted token for a subtle
// background shift; focus uses the ring token via the standard
// data-focus variant (aliased to :focus-visible in flint.css).
// Disabled kills the hover and dims via opacity, matching Button.
func triggerClasses(p TriggerProps) string {
	parts := []string{
		// Layout
		"flex w-full items-center justify-between",
		"px-4 py-3",
		// Text
		"text-left text-base font-medium text-foreground",
		// Hover
		"hover:bg-muted",
		// Focus (data-focus aliased to :focus-visible in flint.css)
		"focus:outline-hidden data-focus:bg-muted",
		// Disabled
		"disabled:opacity-50 disabled:pointer-events-none",
		// Transition
		"transition-colors",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// panelClasses: matches the trigger's horizontal padding but uses a
// tighter vertical rhythm. The default is body copy on the page surface;
// Class overrides for denser content or different backgrounds.
func panelClasses(p PanelProps) string {
	parts := []string{
		"px-4 pb-4 pt-1",
		"text-sm text-muted-foreground",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
