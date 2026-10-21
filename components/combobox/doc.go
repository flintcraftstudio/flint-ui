// Package combobox provides a searchable select — a text input with a
// filterable dropdown of options. Converted from catalyst-ui-kit/typescript/
// combobox.tsx; the API is reshaped because Catalyst's Headless UI runtime
// (virtual rendering, async filter hooks, displayValue functions) doesn't
// translate to server-side templ.
//
// Behavior: the caller lists all options up front; Alpine owns a `query`
// string bound to the input and a `selected` value carried in a hidden
// <input type="hidden"> for form submission. Filtering is client-side —
// each Option's `x-show` checks its Label against the query. This covers
// small-to-medium lists (up to a few hundred options) cleanly; a Mode-
// Server variant for longer lists can follow the Tabs pattern when a
// client needs it.
//
// State model:
//
//	{
//	    open: false,             // dropdown visibility
//	    query: '',               // what's in the input (also drives filter)
//	    selected: '',            // chosen option's Value (for form)
//	    selectedLabel: '',       // chosen option's Label (for input display)
//	    matches(label),          // filter predicate used by each Option
//	    select(value, label),    // set selection + close + sync display
//	    ...
//	}
//
// Option values are passed to select() via data-value / data-label
// attributes on each Option element — the Alpine expression stays a
// constant string and templ's HTML-attribute escaping handles special
// characters in values/labels. Same pattern as clipboard.Copy.
//
// Keyboard (requires @alpinejs/focus, already loaded by the showcase):
//
//	ArrowDown from input → open + focus first visible option
//	ArrowDown/Up within options → wrap to next/previous visible option
//	Enter on input → select first visible option
//	Enter on option → select it
//	Escape → close + restore query to selectedLabel
//	Tab → close + release focus
//
// Selected value is sticky — typing a query that doesn't match the
// current selection does NOT clear `selected`; only picking a new
// option changes it. This matches how most comboboxes behave (Slack
// channel picker, Gmail label picker). If a client needs a clearable
// combobox with explicit "reset" behavior, add a Clear button that
// does `selected = ”; selectedLabel = ”; query = ”`.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/combobox.md in the flint-ui module.
package combobox
