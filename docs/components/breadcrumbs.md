# Breadcrumbs

Navigation trail — the row of links showing where the user is in a nested hierarchy. From-scratch design (Catalyst has no `breadcrumbs.tsx`).

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/breadcrumbs"
```

## Components

| Component                   | Renders                                   | Notes                                                  |
| --------------------------- | ----------------------------------------- | ------------------------------------------------------ |
| `breadcrumbs.Breadcrumbs`   | `<nav aria-label="Breadcrumb"><ol>`       | Outer wrapper. `<ol>` holds the Items + Current.       |
| `breadcrumbs.Item`          | `<li><a>` + chevron separator             | Linked crumb. Use for every step except the final one. |
| `breadcrumbs.Current`       | `<li><span aria-current="page">`          | Unlinked "you are here" entry. No trailing chevron. Place last. |

## Quick example

```go
@breadcrumbs.Breadcrumbs(breadcrumbs.Props{}) {
    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/"}) { Dashboard }
    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/jobs"}) { Jobs }
    @breadcrumbs.Current(breadcrumbs.CurrentProps{}) { Henderson Ranch }
}
```

## Item and Current

Order matters: Item is always a link with a trailing chevron. Current is always the final entry — no link, no trailing chevron, `aria-current="page"` for assistive tech.

If a caller accidentally places `Item` last, the result is a dangling chevron pointing at nothing. That's a visible error rather than a silent one, so it's usually caught quickly.

## Rich content

Both Item and Current accept arbitrary templ content as children:

```go
@breadcrumbs.Current(breadcrumbs.CurrentProps{}) {
    <span class="flex items-center gap-2">
        Henderson Ranch
        @badge.Badge(badge.Props{Variant: badge.VariantWarning}) { Blocked }
    </span>
}
```

## Wrapping on narrow viewports

The outer `<ol>` is `flex flex-wrap`, so deeply nested trails wrap to multiple lines on phone-sized viewports instead of overflowing horizontally. If your design prefers truncation (e.g., "Dashboard › … › Henderson Ranch"), drop manual Item replacements:

```go
@breadcrumbs.Breadcrumbs(breadcrumbs.Props{}) {
    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/"}) { Dashboard }
    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/clients/42/jobs"}) { … }
    @breadcrumbs.Current(breadcrumbs.CurrentProps{}) { Henderson Ranch }
}
```

## Accessibility

- `<nav aria-label="Breadcrumb">` declares the region semantics (override the label via `Props.AriaLabel` if the page has multiple breadcrumb trails).
- `<ol>` conveys ordered structure.
- Current uses `aria-current="page"`.
- The chevron separator is `aria-hidden="true"` — the nav + ol + aria-current combination already conveys the trail; the chevron is purely visual.

## Styling

- Links: `text-muted-foreground` at rest, `hover:text-foreground`, subtle `data-focus:underline` for keyboard focus visibility.
- Current: `font-medium text-foreground` — differentiated by weight and color, same size as links.
- Chevron: `text-muted-foreground/60` at `size-4`, sits between items with natural gap.

Each client's brand flows through automatically — A-Team Gutters renders its crumbs against A-Team's foreground / muted-foreground pair; Lo Mo renders against theirs.

## Props reference

### Props (Breadcrumbs root)

| Field       | Type               | Default        | Notes                                              |
| ----------- | ------------------ | -------------- | -------------------------------------------------- |
| `AriaLabel` | `string`           | `"Breadcrumb"` | Override for pages with multiple breadcrumb trails. |
| `Class`     | `string`           | `""`           | Appended to the `<nav>`.                           |
| `Attrs`     | `templ.Attributes` | `nil`          | Spread onto the `<nav>`.                           |

### ItemProps

| Field   | Type               | Default | Notes                                                  |
| ------- | ------------------ | ------- | ------------------------------------------------------ |
| `Href`  | `string`           | `""`    | Required. Empty renders a broken link — use Current.   |
| `Class` | `string`           | `""`    | Appended to the anchor.                                |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the anchor. Use for `hx-get` if desired.   |

### CurrentProps

| Field   | Type               | Default | Notes                             |
| ------- | ------------------ | ------- | --------------------------------- |
| `Class` | `string`           | `""`    | Appended to the current span.     |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the current span.     |
