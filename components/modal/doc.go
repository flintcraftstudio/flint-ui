// Package modal provides Catalyst-style dialog and alert overlays. Converted
// from catalyst-ui-kit/typescript/dialog.tsx and alert.tsx — both share the
// same backdrop + panel structure with Alert rendering a smaller, centered
// panel. A single component handles both: set Props.Alert = true for the
// compact layout.
//
// State lives in Alpine. The caller opens the modal by dispatching a window
// event:
//
//	<button x-on:click="$dispatch('flint:modal-open', 'delete-invoice')">
//	    Delete
//	</button>
//
// and closes it with ESC, a backdrop click, any button inside the modal with
// x-on:click="open = false", or by dispatching 'flint:modal-close' (omit the
// detail to close every modal; pass the name to target one).
//
// The modal teleports to <body> to escape stacking contexts. This means a
// parent with transform / overflow:hidden / filter won't break positioning.
//
// Focus management requires the Alpine Focus plugin. The panel uses
// x-trap.inert.noscroll to trap Tab, mark siblings inert, and lock body
// scroll while open — matching Headless UI's behavior in the Catalyst
// original. Load @alpinejs/focus BEFORE alpinejs core in your page.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/modal.md in the flint-ui module.
package modal
