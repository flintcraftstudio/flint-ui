// Package popover provides a click-activated floating panel with
// arbitrary content. From-scratch design — Catalyst has no popover.tsx.
//
// Popover vs. Dropdown vs. Tooltip — three nearby components, three
// distinct purposes:
//
//   - Tooltip: hover/focus-activated short text. Decorative.
//   - Dropdown: click-activated menu of focusable items (role="menu",
//     role="menuitem", keyboard nav through items).
//   - Popover: click-activated panel with arbitrary content — forms,
//     rich copy, buttons. Not a menu, not a tooltip.
//
// State lives in Alpine: a `{ open: false }` scope on the root. Click
// the trigger to toggle; click outside or press ESC to close.
//
// Tab is NOT trapped. A popover is a floating panel, not a modal — the
// caller's content usually wants natural focus flow (form fields tab in
// order, last Tab moves past the popover to the next page element and
// leaves the popover open). This also means no @alpinejs/focus
// requirement; Popover is dependency-free aside from Alpine core.
//
// ARIA: the trigger gets aria-haspopup="dialog" and aria-expanded bound
// to state. The Panel has NO default role — role="dialog" implies a
// name via aria-labelledby, and defaulting to a role we can't enforce
// labeling on would be an a11y footgun. Callers with a headed popover
// should add role="dialog" + aria-labelledby via Attrs.
//
// The panel is positioned with CSS `absolute` inside a `relative
// inline-block` root — same pattern as Dropdown, same overflow caveat
// (ancestor overflow:hidden / transform / filter will clip). If a
// client project runs into clipping we'll extract a shared positioning
// helper; not yet.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/popover.md in the flint-ui module.
package popover
