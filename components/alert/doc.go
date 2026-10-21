// Package alert provides inline banner notices — tinted blocks inside a
// page that convey a status to the reader without interrupting their
// flow. From-scratch design.
//
// This is NOT the modal confirmation dialog Catalyst calls "alert". That
// one shipped as modal.Props.Alert and keeps Catalyst's centered-on-
// mobile dialog layout. What lives here is the *inline banner* — the
// thing at the top of a page saying "Your license expires in 30 days"
// or "Saved successfully." It's a peer of Badge, not Modal.
//
// Variants mirror Badge: info / success / warning / danger. Colors
// follow the same `bg-{token}/15 text-{token}` recipe so badges and
// alerts share a palette — a success badge inside a success alert reads
// as the same concept at different scales.
//
// A variant icon is rendered automatically on the left. Callers don't
// need to supply one; if they want a different icon, override via the
// Icon slot (children marked `data-slot="icon"` replace the default).
// This matches Toast's pattern.
//
// Dismissable alerts aren't shipped in v0.1 — most inline alerts are
// informative, not transient. Add if a client needs them; the pattern
// would be an optional Alpine `{ open: true }` + a close button.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/alert.md in the flint-ui module.
package alert
