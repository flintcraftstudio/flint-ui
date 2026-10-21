// Package descriptionlist provides semantic key-value pair display —
// dl / dt / dd with Catalyst's two-column layout. Converted from
// catalyst-ui-kit/typescript/description-list.tsx.
//
// The layout is single-column on mobile, two-column (label | value)
// on sm+ where the label column caps at min(50%, 20rem). Rows are
// visually separated by a subtle border-top on each row, with the
// first row's border suppressed so the list sits cleanly inside a
// Card body.
//
// Use for client profiles, invoice metadata, settings summaries —
// anywhere a form-like "field: value" display helps without the
// interactivity of a form.
//
// Package name is singular-word (no hyphen) because Go doesn't allow
// hyphens in package names; the exported types keep "Description"
// as the prefix for clarity.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/description-list.md in the flint-ui module.
package descriptionlist
