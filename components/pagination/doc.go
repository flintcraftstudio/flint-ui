// Package pagination provides numbered page navigation for long tables
// and lists. Converted from catalyst-ui-kit/typescript/pagination.tsx.
//
// Server-driven: each page renders as a link (<a href>). Apply hx-*
// attributes via Attrs if you want htmx-driven page swaps instead of
// full navigations — the component doesn't know or care which.
//
// Layout matches Catalyst's: Previous on the left, List of page numbers
// in the middle (hidden on mobile — small viewports rely on Previous /
// Next only), Next on the right. `grow basis-0` on Previous and Next
// keeps them balanced regardless of how many pages render in between.
//
// No Alpine state. No keyboard plugins. Pagination is the simplest
// component in flint-ui after Badge.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/pagination.md in the flint-ui module.
package pagination
