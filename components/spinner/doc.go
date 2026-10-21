// Package spinner provides an indeterminate loading indicator. From-
// scratch design (Catalyst has no spinner.tsx); the SVG is the standard
// Tailwind spinner.
//
// Three uses:
//   - Standalone, with a Label, as a polite status ("Loading jobs…").
//   - As an htmx request indicator: add Class "htmx-indicator" and point
//     hx-indicator at it. htmx's built-in indicator CSS hides it until a
//     request is in flight.
//   - Inside Button via button.Props.Indicator, which renders one for you.
//
// The spinner inherits currentColor, so set its color with a text-*
// class on it or an ancestor.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/spinner.md in the flint-ui module.
package spinner
