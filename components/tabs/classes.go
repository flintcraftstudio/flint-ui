package tabs

import "strings"

// styling is a from-scratch design (Catalyst has no tabs.tsx). The visual
// is the standard dashboard underline-active pattern: muted-foreground
// text, transparent bottom border, primary border + foreground text on
// the active tab. Hover on inactive tabs gets a subtle border-border
// preview. All driven from the semantic token contract.

// tabID and panelID derive stable element ids from the tab name so
// aria-controls / aria-labelledby line up across Tab and Panel without
// callers having to coordinate ids manually.
func tabID(name string) string {
	return "flint-tab-" + safeName(name)
}

func panelID(name string) string {
	return "flint-tabpanel-" + safeName(name)
}

// safeName strips characters that would break a JS string literal or HTML
// id. Tab names are expected to be stable identifiers like 'overview' or
// 'crew:1'; non-matching characters are dropped so a malformed name can't
// break the page (same approach as modal.safeName).
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

// ariaSelected and tabindexAttr return the static attribute values the
// server stamps on first paint. In ModeAlpine these are immediately
// overridden by x-bind once Alpine initializes; in ModeServer they
// remain authoritative.
func ariaSelected(active bool) string {
	if active {
		return "true"
	}
	return "false"
}

func tabindexAttr(active bool) string {
	if active {
		return "0"
	}
	return "-1"
}

// alpineSelectedExpr / alpineTabindexExpr / alpineClickExpr / alpineShowExpr
// are the Alpine expression strings injected via x-bind: / x-on: / x-show.
// They reference `active` from the parent x-data scope.
func alpineSelectedExpr(name string) string {
	return "active === '" + safeName(name) + "' ? 'true' : 'false'"
}

func alpineTabindexExpr(name string) string {
	return "active === '" + safeName(name) + "' ? 0 : -1"
}

func alpineClickExpr(name string) string {
	return "active = '" + safeName(name) + "'"
}

func alpineShowExpr(name string) string {
	return "active === '" + safeName(name) + "'"
}

// moveTabJS is the body of the moveTab(dir) method on the parent x-data.
// Called from List's keydown handlers. Walks the focusable tabs in DOM
// order, skipping disabled ones (disabled buttons are filtered by the
// :not([disabled]) selector), wrapping at the ends. Sets active and
// focuses so screen readers announce the new tab.
const moveTabJS = `,
	moveTab(dir) {
		const tabs = [...this.$root.querySelectorAll('[role=tab]:not([disabled])')];
		const cur = tabs.findIndex(t => t.dataset.tabName === this.active);
		const last = tabs.length - 1;
		let i = dir === 'first' ? 0
			: dir === 'last' ? last
			: tabs.length === 0 ? 0
			: ((cur < 0 ? 0 : cur) + dir + tabs.length) % tabs.length;
		const t = tabs[i];
		if (t) { this.active = t.dataset.tabName; t.focus(); }
	}
}`

// alpineData wires the root x-data: { active, moveTab(dir) }.
func alpineData(active string) string {
	return "{ active: '" + safeName(active) + "'" + moveTabJS
}

func rootClasses(p Props) string {
	parts := []string{"flex flex-col"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func listClasses(p ListProps) string {
	parts := []string{
		"flex items-center gap-1 border-b border-border",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// tabClasses styles a Tab. The active state is driven entirely by
// aria-selected — Tailwind's `aria-selected:` modifier matches the
// "true" value, `aria-[selected=false]:` matches "false". This means the
// same class string serves both modes: ModeAlpine's x-bind:aria-selected
// and ModeServer's static aria-selected both flip the same Tailwind
// variant, so the visual stays consistent regardless of how Active is
// computed.
func tabClasses(p TabProps) string {
	parts := []string{
		// Layout
		"inline-flex items-center justify-center",
		"px-4 py-2 -mb-px",
		// Border
		"border-b-2 border-transparent",
		// Text
		"text-sm font-medium text-muted-foreground",
		// Hover (only when not selected)
		"hover:text-foreground",
		"aria-[selected=false]:hover:border-border",
		// Transition
		"transition-colors",
		// Focus (data-focus is aliased to :focus-visible in flint.css)
		"focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring",
		// Disabled
		"disabled:opacity-50 disabled:pointer-events-none",
		// Active
		"aria-selected:border-primary aria-selected:text-foreground",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func panelClasses(p PanelProps) string {
	parts := []string{"mt-4 focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
