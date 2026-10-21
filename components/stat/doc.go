// Package stat provides KPI tiles for dashboard home pages: a label, a
// headline value, and an optional change badge with context ("+4.5% from
// last week"). Converted from the Stat component in Catalyst's demo app
// (catalyst-ui-kit/demo/.../app/stat.tsx).
//
//	@stat.Grid(stat.GridProps{}) {
//	    @stat.Stat(stat.Props{Label: "Revenue", Value: "$48,200", Change: "+4.5%", Context: "from last month"})
//	    @stat.Stat(stat.Props{Label: "Overdue invoices", Value: "3", Change: "-2", LowerIsBetter: true})
//	}
//
// Grid renders a <dl>; each Stat is a <div> grouping its <dt> label and
// <dd> value, so assistive tech reads label and value as a pair.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/stat.md in the flint-ui module.
package stat
