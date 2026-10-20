# Empty state

The placeholder for a list, table, or panel that has nothing to show yet. From-scratch design (Catalyst has no empty state); proportions follow the Tailwind UI empty-state pattern.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/emptystate"
```

## Components

| Component                  | Renders        | Notes                                              |
| -------------------------- | -------------- | -------------------------------------------------- |
| `emptystate.EmptyState`    | `<div>`        | Centered column. `Bordered` adds a dashed outline. |
| `emptystate.Title`         | `<h3>`         | One-line summary. `Level` changes the tag.         |
| `emptystate.Description`   | `<p>`          | Why it's empty, or what to do next.                |
| `emptystate.Actions`       | `<div>`        | The call to action, usually one Button.            |

## Standalone

```go
@emptystate.EmptyState(emptystate.Props{Bordered: true}) {
    <svg data-slot="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" ...></svg>
    @emptystate.Title(emptystate.TitleProps{}) { No jobs yet }
    @emptystate.Description(emptystate.SubProps{}) { Jobs you schedule will show up here. }
    @emptystate.Actions(emptystate.SubProps{}) {
        @button.Button(button.Props{Href: "/jobs/new"}) { New job }
    }
}
```

A direct-child `<svg data-slot="icon">` is sized to 48px and tinted a soft `muted-foreground`. 24px outline Heroicons fit best.

## Inside a table

Put it in a single full-width cell. The table already frames it, so leave `Bordered` off:

```go
@table.Body(table.BodyProps{}) {
    if len(invoices) == 0 {
        @table.Row(table.RowProps{}) {
            @table.Cell(table.CellProps{Attrs: templ.Attributes{"colspan": "4"}}) {
                @emptystate.EmptyState(emptystate.Props{Class: "py-8"}) {
                    @emptystate.Title(emptystate.TitleProps{}) { No invoices match these filters }
                    @emptystate.Actions(emptystate.SubProps{}) {
                        @button.Button(button.Props{Outline: true, Href: "/invoices"}) { Clear filters }
                    }
                }
            }
        }
    }
    for _, inv := range invoices { ... }
}
```

With htmx-driven filtering, the server returns this row whenever the filtered result is empty. There's no client-side state to manage.

## Props

### Props

| Field      | Type               | Default | Notes                                                     |
| ---------- | ------------------ | ------- | --------------------------------------------------------- |
| `Bordered` | `bool`             | `false` | Dashed outline, for an empty state standing alone on a page. |
| `Class`    | `string`           | `""`    | Appended to the root. Use `py-*` to tighten inside tables. |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the root.                                     |

### TitleProps

| Field   | Type               | Default | Notes                                    |
| ------- | ------------------ | ------- | ---------------------------------------- |
| `Level` | `int` (2–6)        | `3`     | Heading element. Out-of-range → `<h3>`.  |
| `Class` | `string`           | `""`    | Appended to the heading.                 |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the heading.                 |

### SubProps (Description, Actions)

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the element.      |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the element.      |

## Accessibility

- Pick a `Level` that fits the outline. Under a page `<h1>` with section `<h2>`s, the default `<h3>` is right.
- The icon is decorative. Mark it `aria-hidden="true"`.
