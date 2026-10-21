// Package toast provides ephemeral notifications. From-scratch design;
// Catalyst has no toast.tsx.
//
// Architecture: a single Container is rendered once per page (typically
// in the base layout). Any code on the page can fire a toast by
// dispatching the `flint:toast` window event with a detail payload:
//
//	$dispatch('flint:toast', { variant: 'success', title: 'Saved' })
//
// or via htmx response header (server-driven, no client code):
//
//	HX-Trigger: {"flint:toast":{"variant":"success","title":"Saved"}}
//
// The Container listens for the event, builds a toast from the payload,
// adds it to its reactive list, and schedules auto-dismiss. Hovering
// (or focusing) pauses the timer; mouse-leave / blur resumes. ESC
// dismisses the most recent. The X button dismisses immediately.
//
// Variants — info / success / warning / danger — map to the semantic
// token contract via a colored left border and an icon tinted to the
// matching token (text-primary / text-success / text-warning / text-danger).
//
// Toasts are content-driven, not declarative — there's no per-toast
// templ. Just one Container, dispatched many times. This trades the
// per-modal name pattern for a simpler API: anything that can dispatch
// a custom event can show a toast (htmx HX-Trigger headers, inline
// onclick, third-party scripts, etc.) with no flint-ui binding code.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/toast.md in the flint-ui module.
package toast
