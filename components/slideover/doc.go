// Package slideover provides side drawer overlays — full-height panels
// that slide in from the left or right edge of the viewport. From-scratch
// design (Catalyst has no slideover.tsx). The implementation intentionally
// mirrors modal's teleport + event-driven open + focus-trap patterns;
// only the panel's layout and enter/leave transitions differ.
//
// State lives in Alpine. The caller opens a slideover by dispatching a
// window event:
//
//	<button x-on:click="$dispatch('flint:slideover-open', 'job-filters')">
//	    Filters
//	</button>
//
// and closes it with ESC, a backdrop click, any button inside the panel
// with x-on:click="open = false", or by dispatching
// 'flint:slideover-close' (omit the detail to close every slideover on
// the page; pass the name to target one).
//
// The panel teleports to <body> to escape stacking contexts — a parent
// with transform / overflow:hidden / filter won't break positioning.
// Focus trap + body-scroll lock require the Alpine Focus plugin; load
// @alpinejs/focus BEFORE alpinejs core (same requirement as Modal and
// Dropdown).
//
// Layout inside the panel is flex-col:
//   - Title and Description sit at the top (natural flow, shrink-0).
//   - Body is flex-1 overflow-y-auto — fills and scrolls.
//   - Actions has mt-auto so it pins to the bottom of the panel
//     regardless of content length.
//
// This is the one place Slideover's sub-component geometry differs from
// Modal. A modal is a compact vertical stack; a slideover is a full-
// height drawer where the scrollable body needs to know it's the flex
// child.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/slideover.md in the flint-ui module.
package slideover
