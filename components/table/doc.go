// Package table provides Catalyst-style tables. Converted from
// catalyst-ui-kit/typescript/table.tsx.
//
// Catalyst uses React context to propagate the four layout toggles
// (bleed, dense, grid, striped) from <Table> down to every <TableHeader>
// and <TableCell>. templ has no analog, so flint-ui takes the boring but
// honest path: each sub-component accepts the flags it cares about.
// Callers typically set the flags once in a per-project render helper,
// so the duplication collapses at the use site.
//
// The one Catalyst feature intentionally dropped: href-on-row (absolute
// <Link> overlay inside the first cell). It's tightly coupled to
// Catalyst's <Link> + browser focus management, and for flint-ui's
// server-rendered, htmx-first world it's cleaner to put htmx attrs
// directly on the <tr> via RowProps.Attrs — `hx-get="/jobs/123"`,
// `hx-target="#detail"`, `hx-push-url="true"` — which works without JS
// focus hacks.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/table.md in the flint-ui module.
package table
