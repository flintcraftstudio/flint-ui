// Package heading provides page and section headings. Converted from
// catalyst-ui-kit/typescript/heading.tsx.
//
// Two components:
//   - Heading   — page-level title (h1 default, renders slightly larger)
//   - Subheading — section title (h2 default, slightly smaller)
//
// Both accept a Level (1–6) so the HTML element matches the document
// outline even when the visual size is fixed. This is the standard
// "semantic vs. visual" split — a Subheading as an <h2> inside a page
// with a Heading <h1> keeps screen-readers happy.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/heading.md in the flint-ui module.
package heading
