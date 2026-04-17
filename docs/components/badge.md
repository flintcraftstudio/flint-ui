# Badge

Converted from `catalyst-ui-kit/typescript/badge.tsx`. Catalyst ships 18 color variants (red, orange, amber, yellow, lime, green, emerald, …); flint-ui collapses those to the six semantic variants from the token contract, so badges re-theme per client without component changes.

## Import

```go
import "github.com/flintcraft/flint-ui/components/badge"
```

## Usage

```go
@badge.Badge(badge.Props{Variant: badge.VariantSuccess}) {
    Paid
}
```

The zero-value Props renders a `Muted` (neutral gray) badge:

```go
@badge.Badge(badge.Props{}) { Draft }
```

## Variants

| Variant         | Token pair                           | Typical use                     |
| --------------- | ------------------------------------ | ------------------------------- |
| `VariantMuted`  | `bg-muted` + `text-muted-foreground` | Default / neutral state         |
| `VariantPrimary`| `bg-primary/15` + `text-primary`     | Active / in progress            |
| `VariantAccent` | `bg-accent/15` + `text-accent`       | Featured / highlight            |
| `VariantDanger` | `bg-danger/15` + `text-danger`       | Overdue / failed / critical     |
| `VariantSuccess`| `bg-success/15` + `text-success`     | Complete / paid / passing       |
| `VariantWarning`| `bg-warning/25` + `text-warning-foreground` | Awaiting / at-risk       |

Warning uses `text-warning-foreground` (dark) instead of `text-warning` because the warning token is yellow — yellow-on-yellow fails contrast.

## Props

| Field     | Type               | Default         | Notes                                                          |
| --------- | ------------------ | --------------- | -------------------------------------------------------------- |
| `Variant` | `Variant`          | `VariantMuted`  | One of the six constants above.                                |
| `Class`   | `string`           | `""`            | Appended to the rendered `<span>`.                             |
| `Attrs`   | `templ.Attributes` | `nil`           | Spread onto the `<span>`. Use for `hx-*`, `data-*`, `aria-*`.  |

## Composition

Leading dot (legend-style):

```go
@badge.Badge(badge.Props{Variant: badge.VariantSuccess}) {
    <span class="size-1.5 rounded-full bg-success"></span>
    Paid
}
```

Count badge:

```go
@badge.Badge(badge.Props{Variant: badge.VariantPrimary}) { 12 }
```

## Clickable badges

Not exposed as a component. Wrap a Badge in a Button or `<a>` yourself when needed — keeps Badge focused on the visual chip without duplicating Button's variant system. If several client projects end up needing the same clickable pattern, a `BadgeButton` helper can land in a future version.

## Accessibility

- Badges convey state visually; pair them with readable label text for screen readers.
- `forced-colors:outline` keeps badges visible in Windows high-contrast mode (bg alpha tints disappear there).
