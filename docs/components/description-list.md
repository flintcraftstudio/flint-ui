# Description List

Semantic key-value display — native `<dl>` / `<dt>` / `<dd>` with Catalyst's two-column layout. Converted from `catalyst-ui-kit/typescript/description-list.tsx`.

Use for client profiles, invoice metadata, settings summaries — anywhere a "field: value" display helps without the interactivity of a form.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/descriptionlist"
```

(Package name is one word — Go doesn't allow hyphens in package names.)

## Components

| Component                           | Renders   | Notes                                              |
| ----------------------------------- | --------- | -------------------------------------------------- |
| `descriptionlist.DescriptionList`   | `<dl>`    | Outer grid — 1 col on mobile, `min(50%, 20rem) | auto` on sm+. |
| `descriptionlist.Term`              | `<dt>`    | Label side. Muted text, top border separating rows. |
| `descriptionlist.Details`           | `<dd>`    | Value side. Foreground text.                       |

## Quick example

```go
@descriptionlist.DescriptionList(descriptionlist.Props{}) {
    @descriptionlist.Term(descriptionlist.SubProps{}) { Client }
    @descriptionlist.Details(descriptionlist.SubProps{}) { Henderson Ranch }
    @descriptionlist.Term(descriptionlist.SubProps{}) { Contact }
    @descriptionlist.Details(descriptionlist.SubProps{}) { evie@hendersonranch.com }
    @descriptionlist.Term(descriptionlist.SubProps{}) { Balance }
    @descriptionlist.Details(descriptionlist.SubProps{}) { $8,200.00 }
}
```

## Responsive layout

On mobile the list stacks each term over its details in a single column. On sm+ it switches to two columns: the term column caps at `min(50%, 20rem)` so long values wrap naturally in column 2 instead of pushing the layout off-screen.

## Row separators

Each Term carries a top border; the first Term omits its border so the list sits cleanly at the top of a Card or Slide-over panel. Details elements inherit the same border treatment on sm+ for visual consistency.

## Rich value content

`Details` accepts any templ children:

```go
@descriptionlist.Details(descriptionlist.SubProps{}) {
    <div class="flex items-center gap-2">
        <span>Scheduled</span>
        @badge.Badge(badge.Props{Variant: badge.VariantSuccess}) { Confirmed }
    </div>
}
```

## Props reference

### Props (DescriptionList root)

| Field   | Type               | Default | Notes                              |
| ------- | ------------------ | ------- | ---------------------------------- |
| `Class` | `string`           | `""`    | Appended to the `<dl>`.            |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<dl>`.            |

### SubProps (Term / Details)

| Field   | Type               | Default | Notes                              |
| ------- | ------------------ | ------- | ---------------------------------- |
| `Class` | `string`           | `""`    | Appended to the `<dt>` / `<dd>`.  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the element.          |
