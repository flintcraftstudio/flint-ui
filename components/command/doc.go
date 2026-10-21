// Package command provides a Cmd+K command palette — the centered
// overlay modern dashboards use for searching and invoking actions
// across the app. From-scratch design (Catalyst has no command.tsx).
//
// The component sits at the intersection of Modal and Combobox:
//
//   - Modal's plumbing: teleport to <body>, backdrop, x-trap.noscroll
//     focus trap, ESC-to-close, event-driven open/close.
//   - Combobox's plumbing: text input that filters a list of items
//     by substring, arrow-key nav through visible items, Enter to
//     activate the first match.
//
// Opening:
//
//   - Dispatch the global event:
//     $dispatch('flint:command-open')
//   - Or press the built-in shortcut (default Cmd+K on mac, Ctrl+K
//     elsewhere). The palette installs a window-level keydown listener
//     that toggles open/closed. Disable via Props.DisableShortcut if
//     your app uses Cmd+K for something else.
//
// Architecture:
//
//   - Single palette per page. No `Name` prop (unlike Modal/Slideover)
//     — cross-page convention is one palette.
//   - Items are native <button> or <a>. No fake role="option" on a div
//     — native elements handle keyboard / click / screen-reader
//     semantics correctly.
//   - Groups hide when none of their items match the query — keeps
//     empty section headings from cluttering a filtered view.
//
// Keyboard (requires @alpinejs/focus):
//
//	Cmd/Ctrl+K       Toggle open
//	Type in input    Filter items
//	ArrowDown        Move focus from input → first visible item;
//	                 cycle through items (wraps)
//	ArrowUp          Cycle backward through items (wraps)
//	Enter on input   Activate first visible item
//	Enter on item    Activate that item (native click)
//	Escape           Close palette
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/command.md in the flint-ui module.
package command
