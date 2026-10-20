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
| `table.SortHeader` | `<th aria-sort>` + `<a>` | Sortable column header. See [Sortable columns](#sortable-columns). |
| `table.Selectable` | `<div>` with `x-data` | Row-selection scope around BulkActions + Table. |
| `table.SelectAll`  | `<th>` + checkbox | Select-all header: checked / indeterminate / toggles every row. |
| `table.SelectCell` | `<td>` + checkbox | A row's selection checkbox; submits `Value` under `Name` (default `ids`). |
| `table.BulkActions` | `<div>` | Toolbar: live "N selected" count, Clear, and your action buttons. |

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

## Sortable columns

Sorting is server-side. The URL carries the sort, and each header links to the column's *next* sort, so it works without JS and can be bookmarked. `SortHeader` renders the link, the direction arrow, and `aria-sort`:

```go
// sortQuery returns the query string that applies col's next sort.
func (q JobsQuery) sortQuery(col string) string {
    dir := table.SortNone
    if q.Sort == col {
        dir = q.Dir
    }
    return url.Values{"sort": {col}, "dir": {string(dir.Next())}, "q": {q.Q}}.Encode()
}

@table.SortHeader(table.SortHeaderProps{
    Href:      "/jobs?" + q.sortQuery("client"),
    Direction: q.direction("client"), // SortNone unless this column is the active sort
    LinkAttrs: templ.Attributes{      // optional: swap just the table
        "hx-get":    "/jobs/table?" + q.sortQuery("client"),
        "hx-target": "#jobs-table",
        "hx-swap":   "outerHTML",
    },
}) { Client }
```

- `table.ParseSortDirection(r.URL.Query().Get("dir"))` reads the direction safely; anything unexpected is `SortNone`.
- `SortDirection.Next()`: unsorted or descending → ascending, ascending → descending.
- Inactive columns show a faint up/down glyph that firms up on hover. The active column shows a solid arrow.

## Row selection and bulk actions

Wrap the toolbar and the table in `Selectable`, add a `SelectAll` header and a `SelectCell` per row, and wire action buttons with `BulkAttrs`:

```go
@table.Selectable(table.SelectableProps{}) {
    @table.BulkActions(table.BulkActionsProps{}) {
        @button.Button(button.Props{Outline: true, Indicator: true, Attrs: table.BulkAttrs(templ.Attributes{
            "hx-post": "/jobs/archive",
            "hx-swap": "none",
        })}) { Archive }
    }
    @table.Table(table.Props{}) {
        @table.Head(table.HeadProps{}) {
            @table.Row(table.RowProps{}) {
                @table.SelectAll(table.SelectAllProps{Label: "Select all jobs"})
                @table.Header(table.HeaderProps{}) { Job }
                ...
            }
        }
        @table.Body(table.BodyProps{}) {
            for _, j := range jobs {
                @table.Row(table.RowProps{}) {
                    @table.SelectCell(table.SelectCellProps{Value: j.ID, Label: "Select job #" + j.ID})
                    ...
                }
            }
        }
    }
}
```

- **Selection is form state.** Row checkboxes are real inputs named `ids`. `BulkAttrs` sets `hx-include="closest [data-flint-selection]"`, so the server receives `ids=1834&ids=1829`. Read them with `r.Form["ids"]`. Any other named inputs inside the scope (a search box, for example) are included too.
- **Buttons disable themselves.** `BulkAttrs` also binds `disabled` while nothing is selected. `BulkActions` is always rendered, so checking the first row doesn't shift the table down.
- **Swaps stay in sync.** The count is recomputed whenever rows are added or removed (sorting, searching, pagination), so a swapped-in table starts with the right count.
- **Clear from the server.** After a successful bulk action, respond with `HX-Trigger: {"flint:selection-clear": true}`, plus a toast if you like. A `204` is enough when there's nothing to swap.
- **Locked rows**: set `SelectCellProps.Disabled`. Select-all and the count skip disabled checkboxes.

## Row actions

Put a Dropdown in the last cell and set `Fixed` on its menu. The table's horizontal scroll wrapper would otherwise clip the menu on the bottom rows:

```go
@table.Cell(table.CellProps{Class: "text-right"}) {
    @dropdown.Dropdown(dropdown.Props{}) {
        @button.Button(button.Props{Plain: true, Attrs: dropdown.TriggerAttrs(templ.Attributes{
            "aria-label": "Actions for job #" + j.ID,
        })}) {
            <svg data-slot="icon" ...ellipsis-horizontal></svg>
        }
        @dropdown.Menu(dropdown.MenuProps{Anchor: dropdown.AnchorBottomEnd, Fixed: true}) {
            @dropdown.Item(dropdown.ItemProps{Href: "/jobs/" + j.ID}) { View }
            @dropdown.Item(dropdown.ItemProps{Attrs: templ.Attributes{
                "hx-post": "/jobs/archive",
                "hx-vals": `{"ids": "` + j.ID + `"}`,
                "hx-swap": "none",
            }}) { Archive }
        }
    }
}
```

Give the actions column a visually hidden header: `@table.Header(table.HeaderProps{}) { <span class="sr-only">Actions</span> }`.

## Empty results

When a query returns nothing, render one full-width row holding an [Empty State](./empty-state.md):

```go
if len(jobs) == 0 {
    @table.Row(table.RowProps{}) {
        @table.Cell(table.CellProps{Attrs: templ.Attributes{"colspan": "7"}}) {
            @emptystate.EmptyState(emptystate.Props{Class: "py-8"}) {
                @emptystate.Title(emptystate.TitleProps{}) { No jobs match your search }
            }
        }
    }
}
```

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

### SortHeaderProps

| Field       | Type                  | Default    | Notes                                                        |
| ----------- | --------------------- | ---------- | ------------------------------------------------------------ |
| `Href`      | `string`              | —          | Required. URL applying this column's next sort.              |
| `Direction` | `table.SortDirection` | `SortNone` | This column's current sort: `SortAsc`, `SortDesc`, or `SortNone`. |
| `Bleed`, `Grid` | `bool`            | `false`    | Same as HeaderProps.                                         |
| `Class`     | `string`              | `""`       | Appended to the `<th>`.                                      |
| `Attrs`     | `templ.Attributes`    | `nil`      | Spread onto the `<th>`.                                      |
| `LinkAttrs` | `templ.Attributes`    | `nil`      | Spread onto the `<a>`. Use for `hx-get` / `hx-target` / `hx-swap`. |

### SelectAllProps

| Field   | Type               | Default | Notes                                              |
| ------- | ------------------ | ------- | -------------------------------------------------- |
| `Label` | `string`           | `""`    | Screen-reader name of the checkbox ("Select all jobs"). |
| `Bleed` | `bool`             | `false` | Same as HeaderProps.                               |
| `Class` | `string`           | `""`    | Appended to the `<th>`.                            |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<th>`.                            |

### SelectCellProps

| Field     | Type               | Default | Notes                                                   |
| --------- | ------------------ | ------- | ------------------------------------------------------- |
| `Name`    | `string`           | `"ids"` | Form field the checked rows submit under.               |
| `Value`   | `string`           | `""`    | This row's ID.                                          |
| `Label`   | `string`           | `""`    | Screen-reader name ("Select job #1834").                |
| `Checked` | `bool`             | `false` | Pre-check the row (e.g. preserving selection across a server render). |
| `Disabled` | `bool`            | `false` | Disable the checkbox; skipped by select-all and the count. |
| `Bleed`, `Dense`, `Striped` | `bool` | `false` | Same as CellProps.                           |
| `Class`   | `string`           | `""`    | Appended to the `<td>`.                                 |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the `<td>`.                                 |

### SelectableProps / BulkActionsProps

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root element. |
