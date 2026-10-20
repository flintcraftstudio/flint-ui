# Stat

KPI tiles for dashboard home pages: a label, a headline value, and an optional change badge with context ("+4.5% from last month"). Converted from the `Stat` in Catalyst's demo app (`demo/.../app/stat.tsx`) and its dashboard grid.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/stat"
```

## Components

| Component   | Renders  | Notes                                                       |
| ----------- | -------- | ----------------------------------------------------------- |
| `stat.Grid` | `<dl>`   | One column on mobile, two at `sm`, four at `xl`.            |
| `stat.Stat` | `<div>` with `<dt>` + `<dd>`s | One KPI. Must sit inside `Grid` (or another `<dl>`). |

## Usage

```go
@stat.Grid(stat.GridProps{}) {
    @stat.Stat(stat.Props{Label: "Revenue", Value: "$48,200", Change: "+4.5%", Context: "from last month"})
    @stat.Stat(stat.Props{Label: "Average job value", Value: "$6,890", Change: "-0.5%", Context: "from last month"})
    @stat.Stat(stat.Props{Label: "Jobs completed", Value: "27", Change: "+4", Context: "from last month"})
    @stat.Stat(stat.Props{Label: "Crew utilization", Value: "86%"}) // no change row
}
```

`Value` and `Change` are preformatted strings. Format numbers on the server, where currency, rounding and locale are known.

## Change colors

The badge color follows the sign at the start of `Change`:

| Change                      | Trend | Badge (default) | Badge with `LowerIsBetter` |
| --------------------------- | ----- | --------------- | -------------------------- |
| `+4.5%`                     | up    | success         | danger                     |
| `-0.5%` or `−0.5%` (U+2212) | down  | danger          | success                    |
| anything else (`0`, `n/a`)  | flat  | muted           | muted                      |

- **`LowerIsBetter`** is for metrics where a rise is bad: overdue invoices, cancellations, response time.
- **`Trend`** overrides the inference when `Change` has no sign, e.g. `Change: "4 more", Trend: stat.TrendUp`.

## Loading with htmx

When a number is expensive to compute, render the page without it and load the grid afterwards:

```go
<div hx-get="/dashboard/stats" hx-trigger="load" hx-swap="outerHTML">
    @spinner.Spinner(spinner.Props{Size: spinner.SizeLG, Label: "Loading stats"})
</div>
```

The endpoint returns a `stat.Grid`. Add `every 60s` to `hx-trigger` to keep a wall-display dashboard fresh.

## Props

### GridProps

| Field   | Type               | Default | Notes                                              |
| ------- | ------------------ | ------- | -------------------------------------------------- |
| `Class` | `string`           | `""`    | Appended to the `<dl>`. Change the column count here, e.g. `xl:grid-cols-3`. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<dl>`.                            |

### Props

| Field           | Type               | Default     | Notes                                                  |
| --------------- | ------------------ | ----------- | ------------------------------------------------------ |
| `Label`         | `string`           | `""`        | What's measured. Renders as the `<dt>`.                |
| `Value`         | `string`           | `""`        | Headline number, preformatted.                         |
| `Change`        | `string`           | `""`        | Change badge text. Empty omits the change row.         |
| `Trend`         | `stat.Trend`       | `TrendAuto` | `TrendUp` / `TrendDown` / `TrendFlat` override the sign inference. |
| `LowerIsBetter` | `bool`             | `false`     | Swap the up/down badge colors.                         |
| `Context`       | `string`           | `""`        | Muted text after the badge ("from last month").        |
| `Class`         | `string`           | `""`        | Appended to the stat's `<div>`.                        |
| `Attrs`         | `templ.Attributes` | `nil`       | Spread onto the stat's `<div>`.                        |

## Accessibility

- `Grid` is a description list and each stat is a `<dt>`/`<dd>` group, so screen readers announce the label with its value.
- The change is plain text with a sign ("+4.5%"), so the direction doesn't depend on the badge color.
- Catalyst draws the rule above each stat with an `<hr>`. A `<dl>` group may contain only `<dt>`/`<dd>`, so here the rule is the group's top border.
