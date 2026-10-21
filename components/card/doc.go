// Package card provides a surface container for grouping related content.
// From-scratch design — Catalyst has no card.tsx.
//
// A Card is structural chrome only: background surface, border, rounded
// corners, and a subtle shadow. Content sits in Header / Body / Footer
// sub-components, each with its own padding. Header carries a bottom
// border and Footer a top border, so composing sections naturally
// produces dividers between them without a separate Divider component.
//
// Typical shapes:
//
//	// Stat tile: body only
//	@card.Card(card.Props{}) {
//	    @card.Body(card.SubProps{}) {
//	        <div class="text-sm text-muted-foreground">Active jobs</div>
//	        <div class="mt-1 text-3xl font-semibold text-foreground">14</div>
//	    }
//	}
//
//	// Report: header + body + footer
//	@card.Card(card.Props{}) {
//	    @card.Header(card.SubProps{}) { ... }
//	    @card.Body(card.SubProps{}) { ... }
//	    @card.Footer(card.SubProps{}) { ... }
//	}
//
// Cards have no variants — padding, typography, and background are set
// once. If a client wants something different (a dashed outline variant,
// a heavier shadow), override per-instance via Class rather than adding
// a variant enum.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/card.md in the flint-ui module.
package card
