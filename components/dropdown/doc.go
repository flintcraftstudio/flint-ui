// Package dropdown provides Catalyst-style menu dropdowns. Converted from
// catalyst-ui-kit/typescript/dropdown.tsx (built on Headless UI's Menu).
//
// State lives in Alpine. The Dropdown root provides { open: false }; the
// trigger Button toggles, the Menu listens for click-outside / ESC / Tab.
// Clicking inside the menu closes it via a single x-on:click on the panel,
// so callers can safely add htmx attrs to Item without their own handler.
//
// Keyboard nav uses the Alpine Focus plugin: ArrowDown on the trigger
// opens the menu and focuses the first item; ArrowUp/ArrowDown inside the
// menu rotate; Home/End jump to first/last; Tab closes; Enter activates
// (native button/link behavior). Load @alpinejs/focus before alpinejs core
// — same requirement as Modal.
//
// The menu is positioned with CSS (absolute inside a `relative inline-block`
// root). An ancestor with overflow other than visible — a table's
// horizontal scroll wrapper, a scrolling panel — clips it. Set
// MenuProps.Fixed in those places: the menu is then positioned against the
// viewport from the trigger's rect when it opens, and closes on any
// ancestor scroll or resize.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/dropdown.md in the flint-ui module.
package dropdown
