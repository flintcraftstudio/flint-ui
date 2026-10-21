// Package progress provides a horizontal progress bar. From-scratch
// design (Catalyst has no progress.tsx).
//
// Two modes:
//   - Determinate: Value out of Max (default 100). The fill width and
//     aria-valuenow track Value.
//   - Indeterminate: a segment loops across the track for work with no
//     known length. aria-valuenow is omitted, which is how assistive
//     tech recognises an indeterminate progressbar.
//
// For server-reported progress, poll with htmx through Attrs and swap the
// bar with hx-swap="outerMorph" so the fill animates to its new width
// instead of jumping (see docs/components/progress.md).
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/progress.md in the flint-ui module.
package progress
