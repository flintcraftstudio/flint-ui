# Card

Surface container for grouping related content. From-scratch design — Catalyst has no `card.tsx`.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/card"
```

## Components

| Component          | Renders | Notes                                               |
| ------------------ | ------- | --------------------------------------------------- |
| `card.Card`        | `<div>` | Outer shell. Chrome only (bg, border, radius, shadow); no padding. |
| `card.Header`      | `<div>` | Padded top section with a bottom border.            |
| `card.Body`        | `<div>` | Padded middle section with no borders.              |
| `card.Footer`      | `<div>` | Padded bottom section with a top border, lightly tinted bg. |

## Quick example

```go
@card.Card(card.Props{}) {
    @card.Header(card.SubProps{}) {
        <h3 class="text-base font-semibold text-foreground">Invoice #1028</h3>
        <p class="mt-1 text-sm text-muted-foreground">Due Apr 28</p>
    }
    @card.Body(card.SubProps{}) {
        <p>Line items...</p>
    }
    @card.Footer(card.SubProps{}) {
        @button.Button(button.Props{}) { Mark paid }
    }
}
```

## Why Card has no padding

Sections (`Header` / `Body` / `Footer`) bring their own padding. Keeping the root chrome-only means the section borders can go edge-to-edge cleanly, and a Body-only card has the same outer rhythm as a multi-section card.

If you need a padding-only card (no sections), drop content directly inside `Body`:

```go
@card.Card(card.Props{}) {
    @card.Body(card.SubProps{}) {
        <p>Anything here.</p>
    }
}
```

## Dividers come for free

Header's `border-b` and Footer's `border-t` mean that composing `Header + Body` or `Body + Footer` produces a divider without a separate component. If you want a divider mid-body (e.g., between list sections), use a plain `<hr class="border-border" />` inside Body.

## No variants

Card has a single styling on purpose. If a client needs a different look — a dashed outline, a darker surface, a heavier shadow — pass those classes via `Class`. Adding a `Variant` enum would fragment the surface vocabulary with no clear payoff for an internal component library.

## Props reference

### Props (Card root)

| Field   | Type               | Default | Notes                            |
| ------- | ------------------ | ------- | -------------------------------- |
| `Class` | `string`           | `""`    | Appended to the root `<div>`.    |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root.            |

### SubProps (Header / Body / Footer)

| Field   | Type               | Default | Notes                              |
| ------- | ------------------ | ------- | ---------------------------------- |
| `Class` | `string`           | `""`    | Appended to the rendered section.  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the rendered section.  |
