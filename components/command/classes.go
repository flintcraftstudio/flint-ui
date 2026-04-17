package command

import "strings"

// From-scratch design — no Catalyst source to mirror. Chrome tokens
// match Modal's panel (bg-surface, ring-1 ring-border, shadow-lg)
// and Combobox's listbox (rounded-xl, backdrop-blur) so the palette
// feels like a larger sibling of the combobox dropdown.

func resolvePlaceholder(p string) string {
	if p != "" {
		return p
	}
	return "Type a command or search…"
}

func resolveFooterHint(p string) string {
	if p != "" {
		return p
	}
	return "↑↓ Navigate · ↵ Select · Esc Close"
}

// shortcutHandler wires Cmd+K on mac and Ctrl+K elsewhere to toggle
// the palette. Preventing default stops the browser's focus-URL-bar
// behavior on Ctrl+K in some browsers.
func shortcutHandler() string {
	return "if ($event.key === 'k' && ($event.metaKey || $event.ctrlKey)) { $event.preventDefault(); open ? hide() : show() }"
}

// backdropClasses: full-viewport dim same as Modal's dialog variant.
func backdropClasses() string {
	return "fixed inset-0 bg-foreground/25 focus:outline-0"
}

// panelWrapperClasses positions the panel top-center, reserving
// viewport padding so the backdrop is visible behind it. pt-20 on
// desktop pushes the palette a bit below the top edge — the
// conventional command-palette resting position.
func panelWrapperClasses() string {
	return "fixed inset-0 flex items-start justify-center px-4 pt-20"
}

// panelClasses: the visible palette surface. w-full + max-w-xl caps
// width; overflow-hidden keeps rounded corners clean since the
// list's scroll area touches the top/bottom edges.
func panelClasses(p Props) string {
	parts := []string{
		"w-full max-w-xl",
		"overflow-hidden",
		"rounded-xl",
		"bg-surface",
		"shadow-xl",
		"ring-1 ring-border",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// searchRowClasses: flex row holding the icon + input. border-b
// separates search from the list below.
func searchRowClasses() string {
	return "flex items-center gap-3 border-b border-border px-4"
}

// inputClasses: borderless text input — the search row's border
// serves as the input's visual boundary. focus:outline-hidden
// suppresses the browser's default focus ring; we don't apply our
// own because the palette itself has the visual focus.
func inputClasses() string {
	return strings.Join([]string{
		"w-full py-3.5",
		"bg-transparent",
		"text-base/6 text-foreground placeholder:text-muted-foreground",
		"border-0 focus:outline-hidden focus:ring-0",
	}, " ")
}

// listClasses: scrollable area for Groups + Items. max-h-96 keeps
// the palette a consistent height regardless of how many commands
// register; overflow-y-auto scrolls within the list when it grows
// beyond that.
func listClasses() string {
	return "max-h-96 overflow-y-auto overscroll-contain py-2"
}

// groupClasses: spacing between groups. First-group negative margin
// removes the extra top gap since it's already adjacent to the
// search border.
func groupClasses(p GroupProps) string {
	parts := []string{"py-1"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// groupHeadingClasses: small uppercase label. aria-hidden on the
// element itself since the role=group carries the semantic grouping
// via the children.
func groupHeadingClasses() string {
	return "px-4 pb-1 pt-2 text-xs font-semibold uppercase tracking-wider text-muted-foreground"
}

// itemClasses: a single command row. Grid-like flex row with the
// label flex-1 + shortcut shrink-0 at the right. data-focus and
// data-hover use the accent token pair so highlight color themes
// per client — same pattern as Dropdown items.
func itemClasses(p ItemProps) string {
	parts := []string{
		"flex w-full items-center gap-3",
		"px-4 py-2",
		"text-sm text-foreground text-left",
		"cursor-default",
		"data-focus:bg-accent data-focus:text-accent-foreground",
		"data-hover:bg-accent data-hover:text-accent-foreground",
		"data-disabled:opacity-50 data-disabled:pointer-events-none",
		"disabled:opacity-50 disabled:pointer-events-none",
		"focus:outline-hidden",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// shortcutClasses: the right-aligned keyboard hint. group-data-focus
// swaps color along with the row so the shortcut doesn't remain
// muted against an accent-colored highlighted row.
func shortcutClasses() string {
	return strings.Join([]string{
		"shrink-0 font-mono text-xs text-muted-foreground",
		"group-data-focus:text-accent-foreground group-data-hover:text-accent-foreground",
	}, " ")
}

// emptyClasses: the "No results" placeholder shown when the query
// matches no items across all groups.
func emptyClasses() string {
	return "px-4 py-8 text-center text-sm text-muted-foreground"
}

// footerClasses: the bottom hint row. border-t separates it from
// the list; muted typography keeps it secondary.
func footerClasses() string {
	return "border-t border-border px-4 py-2 text-xs text-muted-foreground"
}

// alpineData returns the x-data expression. Mirrors combobox's
// pattern — constant string, reads from $refs for DOM queries.
//
// Key methods:
//
//   show() / hide()          open / close the palette (with query reset)
//   matches(label)           filter predicate for items
//   groupHasMatch($el)       true when the group contains any visible item
//   hasAnyMatch()            true when at least one item overall matches
//   focusFirst()             move focus from input to first visible item
//   activateFirst()          simulate a click on first visible item
func alpineData(p Props) string {
	_ = p // Props reserved for future per-instance config (e.g., custom shortcut key)
	return commandAlpineData
}

const commandAlpineData = "{" +
	"open: false," +
	"query: ''," +
	"show() {" +
	"  this.open = true;" +
	"  this.$nextTick(() => { if (this.$refs.input) this.$refs.input.focus() });" +
	"}," +
	"hide() {" +
	"  this.open = false;" +
	"  this.query = '';" +
	"}," +
	"matches(label) {" +
	"  if (this.query === '') return true;" +
	"  return label.toLowerCase().includes(this.query.toLowerCase());" +
	"}," +
	"groupHasMatch(el) {" +
	"  if (this.query === '') return true;" +
	"  const items = el.querySelectorAll('[data-label]');" +
	"  for (const item of items) {" +
	"    if (this.matches(item.dataset.label)) return true;" +
	"  }" +
	"  return false;" +
	"}," +
	"hasAnyMatch() {" +
	"  if (this.query === '') return true;" +
	"  if (!this.$refs.list) return true;" +
	"  const items = this.$refs.list.querySelectorAll('[data-label]');" +
	"  for (const item of items) {" +
	"    if (this.matches(item.dataset.label)) return true;" +
	"  }" +
	"  return false;" +
	"}," +
	"focusFirst() {" +
	"  if (!this.$refs.list) return;" +
	"  const items = this.$refs.list.querySelectorAll('[data-label]:not([data-disabled=true]):not([disabled])');" +
	"  for (const item of items) {" +
	"    if (item.offsetParent !== null) { item.focus(); return; }" +
	"  }" +
	"}," +
	"activateFirst() {" +
	"  if (!this.$refs.list) return;" +
	"  const items = this.$refs.list.querySelectorAll('[data-label]:not([data-disabled=true]):not([disabled])');" +
	"  for (const item of items) {" +
	"    if (item.offsetParent !== null) { item.click(); return; }" +
	"  }" +
	"}" +
	"}"

// matchExpr is the per-item x-show expression. Embeds the Label
// with JS escaping for filter comparison.
func matchExpr(label string) string {
	return "matches('" + escapeJS(label) + "')"
}

// itemClickHandler: on click, close the palette, then let the
// native link/button behavior run (navigation or handler). For a
// disabled item, emit no-op — though the disabled/pointer-events-
// none styling also blocks clicks as a belt-and-suspenders measure.
//
// Note the ordering: hide() runs first (synchronously), then any
// caller-supplied x-on:click from Attrs runs via Alpine's handler
// chain. Navigation via <a href> happens after the current
// microtask, so hide() has already set open=false before the
// browser navigates — the palette doesn't flash on its way out.
func itemClickHandler(disabled bool) string {
	if disabled {
		return ""
	}
	return "hide()"
}

// escapeJS escapes characters that would break a single-quoted
// JS string literal. Same helper as combobox.escapeJS — kept
// local to avoid a cross-component dependency for a three-line
// utility.
func escapeJS(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\'':
			b.WriteString(`\'`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '<':
			b.WriteString(`\u003c`)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
