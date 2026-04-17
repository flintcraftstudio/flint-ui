# Divider

Semantic horizontal rule — `<hr role="presentation">` with regular or soft variants. Converted from `catalyst-ui-kit/typescript/divider.tsx`.

## Import

```go
import "github.com/flintcraft/flint-ui/components/divider"
```

## Quick example

```go
@divider.Divider(divider.Props{})          // regular
@divider.Divider(divider.Props{Soft: true}) // softer
```

## Variants

| Variant   | Class                  | Use                                                       |
| --------- | ---------------------- | --------------------------------------------------------- |
| Default   | `border-foreground/10` | Section boundaries inside a page or card.                 |
| Soft      | `border-foreground/5`  | Subtle divisions inside an already-bordered container.    |

## Why `role="presentation"`

A native `<hr>` without role is announced as a "separator" by screen readers. That's often fine, but a separator is supposed to separate groups of related content — when the divider is purely visual (inside a Card body between rows, for instance), `role="presentation"` tells assistive tech to skip it. If you want the semantic separator announcement, pass `Attrs: templ.Attributes{"role": "separator"}` to override.

## Props reference

| Field   | Type               | Default | Notes                                    |
| ------- | ------------------ | ------- | ---------------------------------------- |
| `Soft`  | `bool`             | `false` | Use the lighter opacity variant.          |
| `Class` | `string`           | `""`    | Appended to the `<hr>` — add margins here (e.g., `my-4`). |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<hr>`.                  |
