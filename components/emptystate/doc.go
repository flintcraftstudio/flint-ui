// Package emptystate provides the placeholder shown where a list, table,
// or panel has nothing to display yet. From-scratch design (Catalyst has
// no empty state); proportions follow the Tailwind UI empty-state
// pattern.
//
//	@emptystate.EmptyState(emptystate.Props{}) {
//	    <svg data-slot="icon" ...></svg>
//	    @emptystate.Title(emptystate.TitleProps{}) { No jobs yet }
//	    @emptystate.Description(emptystate.SubProps{}) {
//	        Jobs you schedule will show up here.
//	    }
//	    @emptystate.Actions(emptystate.SubProps{}) {
//	        @button.Button(button.Props{Href: "/jobs/new"}) { New job }
//	    }
//	}
//
// An icon (an <svg data-slot="icon">, typically a 24px outline
// Heroicon) placed as a direct child is sized and tinted automatically.
//
// Inside a table, put it in a single full-width cell:
//
//	@table.Row(table.RowProps{}) {
//	    @table.Cell(table.CellProps{Attrs: templ.Attributes{"colspan": "5"}}) {
//	        @emptystate.EmptyState(emptystate.Props{}) { ... }
//	    }
//	}
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/empty-state.md in the flint-ui module.
package emptystate
