package combobox

import "strings"

// Non-color classes derived from catalyst-ui-kit/typescript/combobox.tsx
// where possible. The Input-chrome portion (before:inset-px background,
// focus ring) mirrors the Input component verbatim so Combobox and
// Input look like siblings in a form. The Listbox portion mirrors
// Dropdown's Menu chrome so the two floating panels share a visual
// vocabulary.

// wrapperClasses: outer <div data-slot="control">. relative +
// before:inset-px matches the Input wrapper so the focus ring renders
// identically on either component.
func wrapperClasses(p Props) string {
	parts := []string{
		"relative block w-full",
		"before:absolute before:inset-px before:rounded-[calc(var(--radius-lg)-1px)] before:bg-input before:shadow-sm",
		"after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:ring-transparent after:ring-inset sm:focus-within:after:ring-2 sm:focus-within:after:ring-ring",
		"has-data-disabled:opacity-50 has-data-disabled:before:bg-muted has-data-disabled:before:shadow-none",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// inputClasses: the visible text input. Mirrors Input's inputBase
// plus a right-padding bump (pr-10) to make room for the chevron
// button.
func inputClasses(p Props) string {
	parts := []string{
		"relative block w-full appearance-none rounded-lg",
		"pl-[calc(--spacing(3.5)-1px)] pr-[calc(--spacing(10)-1px)]",
		"py-[calc(--spacing(2.5)-1px)]",
		"sm:pl-[calc(--spacing(3)-1px)] sm:pr-[calc(--spacing(9)-1px)]",
		"sm:py-[calc(--spacing(1.5)-1px)]",
		"text-base/6 text-foreground placeholder:text-muted-foreground sm:text-sm/6",
		"border border-border data-hover:border-foreground/20",
		"bg-transparent",
		"focus:outline-hidden",
		"data-invalid:border-danger data-invalid:data-hover:border-danger",
		"data-disabled:border-border",
	}
	return strings.Join(parts, " ")
}

// chevronButtonClasses: the trailing toggle button inside the input's
// right padding gutter. absolute-positioned inside the wrapper; tabindex
// -1 (set on the template) keeps it out of the tab order — the input
// itself handles keyboard toggling.
func chevronButtonClasses() string {
	return strings.Join([]string{
		"absolute inset-y-0 right-0 flex items-center px-2",
		"group",
		"disabled:cursor-not-allowed",
	}, " ")
}

// listboxClasses: the floating options panel. Chrome matches
// Dropdown's Menu (bg-surface/90 with backdrop-blur, shadow-lg, ring).
// mt-2 gap to the input, max-h-64 so long lists scroll within the
// panel rather than pushing the viewport.
func listboxClasses() string {
	return strings.Join([]string{
		"absolute z-50 mt-2 w-full",
		"max-h-64 overflow-y-auto",
		"rounded-xl p-1",
		"bg-surface/90 backdrop-blur-xl",
		"shadow-lg ring-1 ring-border",
		"outline outline-transparent focus:outline-hidden",
		"origin-top",
	}, " ")
}

// optionClasses: a single listbox row. Grid layout so Label/Description
// stack cleanly and the trailing check icon sits in its own column.
// data-focus highlights (keyboard nav) + hover swap to the accent
// token pair — themable per client, same pattern as Dropdown items.
func optionClasses(p OptionProps) string {
	parts := []string{
		"group/option flex w-full cursor-default items-start gap-x-2",
		"rounded-lg py-2.5 pl-3.5 pr-2 sm:py-1.5 sm:pl-3 sm:pr-2",
		"text-base/6 text-foreground sm:text-sm/6",
		"focus:outline-hidden",
		"data-focus:bg-accent data-focus:text-accent-foreground",
		"data-hover:bg-accent data-hover:text-accent-foreground",
		"data-disabled:opacity-50 data-disabled:pointer-events-none",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// descriptionClasses: secondary line under the option's Label. Muted
// at rest; flips to the accent-foreground pair when the parent option
// is focused/hovered so the whole row reads as one highlighted block.
func descriptionClasses() string {
	return strings.Join([]string{
		"text-xs text-muted-foreground",
		"group-data-focus/option:text-accent-foreground",
		"group-data-hover/option:text-accent-foreground",
		"truncate",
	}, " ")
}

// emptyClasses: the "No results" row shown when the query matches no
// option. rounded-lg so it looks consistent with options but
// presentation-role so assistive tech doesn't count it.
func emptyClasses() string {
	return "px-3.5 py-2.5 sm:px-3 sm:py-1.5 text-sm text-muted-foreground"
}

// resolveListboxID derives the listbox id from Name (the canonical
// default) unless ListboxID is explicitly set. Falls back to a fixed
// string when Name is absent — multiple no-Name comboboxes on one
// page collide; callers rendering that scenario must supply
// ListboxID explicitly.
func resolveListboxID(p Props) string {
	if p.ListboxID != "" {
		return p.ListboxID
	}
	if p.Name != "" {
		return "flint-combobox-" + p.Name
	}
	return "flint-combobox"
}

// alpineData returns the x-data expression — a JS object literal
// with all mutable state (query, selected value/label, open) and
// methods (matches, select, toggle, close, focusFirst,
// selectFirstVisible, hasVisibleOptions). The expression is a
// constant string; option values travel via data-value / data-label
// attributes on each option element, read through $el.dataset inside
// click handlers. That keeps this function callsite-independent and
// sidesteps JS-string-escaping concerns in user-supplied values.
func alpineData() string {
	return comboboxAlpineData
}

const comboboxAlpineData = "{" +
	"open: false," +
	"query: ''," +
	"selected: ''," +
	"selectedLabel: ''," +
	"matches(label) {" +
	"  const q = this.query.toLowerCase();" +
	"  if (q === '' || q === this.selectedLabel.toLowerCase()) return true;" +
	"  return label.toLowerCase().includes(q);" +
	"}," +
	"hasVisibleOptions() {" +
	"  const q = this.query.toLowerCase();" +
	"  if (q === '' || q === this.selectedLabel.toLowerCase()) return true;" +
	"  const options = this.$refs.listbox.querySelectorAll('[role=option]');" +
	"  for (const o of options) {" +
	"    if (o.dataset.label && o.dataset.label.toLowerCase().includes(q)) return true;" +
	"  }" +
	"  return false;" +
	"}," +
	"toggle() {" +
	"  this.open = !this.open;" +
	"  if (this.open) this.$nextTick(() => this.focusFirst());" +
	"}," +
	"openAndFocusFirst() {" +
	"  this.open = true;" +
	"  this.$nextTick(() => this.focusFirst());" +
	"}," +
	"focusFirst() {" +
	"  const options = this.$refs.listbox.querySelectorAll('[role=option]:not([data-disabled=true])');" +
	"  for (const o of options) {" +
	"    if (o.offsetParent !== null) { o.focus(); return; }" +
	"  }" +
	"}," +
	"selectFirstVisible() {" +
	"  const options = this.$refs.listbox.querySelectorAll('[role=option]:not([data-disabled=true])');" +
	"  for (const o of options) {" +
	"    if (o.offsetParent !== null) { o.click(); return; }" +
	"  }" +
	"}," +
	"select(value, label) {" +
	"  this.selected = value;" +
	"  this.selectedLabel = label;" +
	"  this.query = label;" +
	"  this.open = false;" +
	"}," +
	"close() {" +
	"  this.open = false;" +
	"  this.query = this.selectedLabel;" +
	"}" +
	"}"

// alpineMatchExpr: x-show expression for an option. Embeds the
// option's Label so Alpine can compare against the current query.
// The label IS embedded in the JS here (not passed via dataset) because
// x-show runs in a reactive context and repeatedly reading dataset is
// slightly more work per pass; the label is already safely string-
// escaped for JS below.
//
// Escaping: single-quote replacement protects against caller labels
// containing apostrophes ("D'Angelo's Pizza") — the alternative of
// reading from dataset would also work but costs a reactive recompute
// whenever Alpine re-evaluates the option.
func alpineMatchExpr(label string) string {
	return "matches('" + escapeJS(label) + "')"
}

// alpineSelectedExpr: used in aria-selected on the option and in x-show
// on the trailing check icon. Returns a boolean comparison.
func alpineSelectedExpr(value string) string {
	return "selected === '" + escapeJS(value) + "'"
}

// alpineSelectExpr: click / enter handler. Reads value+label from
// the option's dataset — sidesteps re-escaping at the call site. For
// disabled options, emit a no-op so aria-disabled visuals are the
// only signal of disabled state.
func alpineSelectExpr(disabled bool) string {
	if disabled {
		return ""
	}
	return "select($el.dataset.value, $el.dataset.label)"
}

// escapeJS escapes characters that would break a single-quoted JS
// string literal. Mirrors Tabs' and Modal's safeName helpers but as
// an escape rather than a restrict — combobox labels can legitimately
// contain any character, while safeName strips.
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
			// Defensive: prevent </script> from terminating a surrounding
			// inline script block if the label is ever embedded in one.
			b.WriteString(`\u003c`)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

