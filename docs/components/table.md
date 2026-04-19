# Table

Converted from `catalyst-ui-kit/typescript/table.tsx`. Catalyst uses React context to propagate four layout toggles (`bleed`, `dense`, `grid`, `striped`) from the outer `<Table>` down to every `<TableHeader>` and `<TableCell>`. templ has no equivalent, so flint-ui takes the explicit path: each sub-component carries the flags it renders.

Callers typically factor a project-specific row helper (`@invoiceRow(inv)`) so the repetition lives in one place.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/table"
```

## Components

| Component        | Renders  | Notes                                                  |
| ---------------- | -------- | ------------------------------------------------------ |
| `table.Table`    | `<table>` wrapped in overflow scroll container | Outer wrapper. `Props.Bleed` drops the gutter padding. |
| `table.Head`     | `<thead>` | Wraps the header row. Text renders `text-muted-foreground`. |
| `table.Body`     | `<tbody>` | Plain passthrough.                                     |
| `table.Row`      | `<tr>`    | `Striped: true` adds zebra-striping via `even:bg-muted/60`. |
| `table.Header`   | `<th>`    | Accepts `Bleed`, `Grid`.                               |
| `table.Cell`     | `<td>`    | Accepts `Bleed`, `Dense`, `Grid`, `Striped`.           |

`Striped` on `Row` and `Cell` must agree — striping replaces the per-cell bottom border, so mixing states produces visible glitches.

## Usage

```go
@table.Table(table.Props{}) {
    @table.Head(table.HeadProps{}) {
        <tr>
            @table.Header(table.HeaderProps{}) { Job }
            @table.Header(table.HeaderProps{}) { Client }
            @table.Header(table.HeaderProps{}) { Scheduled }
        </tr>
    }
    @table.Body(table.BodyProps{}) {
        @table.Row(table.RowProps{}) {
            @table.Cell(table.CellProps{}) { Gutter install }
            @table.Cell(table.CellProps{}) { 218 Pine St }
            @table.Cell(table.CellProps{}) { Apr 18 }
        }
    }
}
```

## Clickable rows — use htmx, not href-on-row

Catalyst's React version puts an `<a>` overlay inside the first cell of any row with `href`, plus focus management via `tabIndex`. flint-ui drops this on purpose: it's tightly coupled to Catalyst's `<Link>` and needs JS to manage focus. For server-rendered tables, htmx on the `<tr>` is cleaner:

```go
@table.Row(table.RowProps{
    Class: "cursor-pointer hover:bg-muted/40",
    Attrs: templ.Attributes{
        "hx-get":   "/jobs/123",
        "hx-target": "#job-detail",
    },
}) {
    // cells...
}
```

Works without JS for keyboard users if you also wire a regular link into one of the cells. For pure row-click with keyboard support, wrap the row content in a `<button>` or add `role="button" tabindex="0"` + an Alpine `x-on:keydown.enter` — your call, per-project.

## Gutter variable

Catalyst uses `--gutter` to let the parent page control table edge padding. flint-ui preserves the mechanism verbatim — the outer wrapper uses `-mx-(--gutter)` and inner padding uses `sm:px-(--gutter)`. Set `--gutter` on any ancestor to adjust, or leave it for Tailwind's default `--spacing(2)`.

`Table.Props.Bleed = true` drops the inner gutter padding so the table runs edge-to-edge inside its container — useful for table-inside-card layouts.

## Accessibility

- Semantic `<table>`, `<thead>`, `<tbody>`, `<tr>`, `<th>`, `<td>` elements preserved — screen readers get proper table navigation.
- Add `scope="col"` on `<th>` when column/row headers matter (pass via `HeaderProps.Attrs`).
- For clickable rows, either include an accessible link inside the row or make the row itself focusable (`tabindex="0"` + key handler).

## Props reference

### Props (Table)

| Field   | Type               | Default | Notes                                              |
| ------- | ------------------ | ------- | -------------------------------------------------- |
| `Bleed` | `bool`             | `false` | Drops the gutter padding on the inner wrapper.     |
| `Class` | `string`           | `""`    | Appended to the overflow-scroll wrapper `<div>`.   |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the overflow-scroll wrapper `<div>`.   |

### RowProps

| Field     | Type               | Default | Notes                                          |
| --------- | ------------------ | ------- | ---------------------------------------------- |
| `Striped` | `bool`             | `false` | Adds `even:bg-muted/60` zebra striping.         |
| `Class`   | `string`           | `""`    | Appended to `<tr>`.                            |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto `<tr>`. Use for `hx-*`, `data-*`.  |

### HeaderProps

| Field   | Type               | Default | Notes                                                |
| ------- | ------------------ | ------- | ---------------------------------------------------- |
| `Bleed` | `bool`             | `false` | Must match the parent Table's Bleed.                 |
| `Grid`  | `bool`             | `false` | Adds vertical dividers between cells.                |
| `Class` | `string`           | `""`    | Appended to `<th>`.                                  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto `<th>`. Use for `scope`, `colspan`, etc. |

### CellProps

| Field     | Type               | Default | Notes                                                            |
| --------- | ------------------ | ------- | ---------------------------------------------------------------- |
| `Bleed`   | `bool`             | `false` | Must match the parent Table's Bleed.                             |
| `Dense`   | `bool`             | `false` | `py-2.5` instead of `py-4`.                                      |
| `Grid`    | `bool`             | `false` | Adds vertical dividers between cells.                            |
| `Striped` | `bool`             | `false` | Suppresses per-cell bottom border (striping provides separation). |
| `Class`   | `string`           | `""`    | Appended to `<td>`.                                              |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto `<td>`. Use for `colspan`, `rowspan`, etc.           |
