# Pagination

Numbered page navigation for long tables and lists. Converted from `catalyst-ui-kit/typescript/pagination.tsx`.

## Import

```go
import "github.com/flintcraft/flint-ui/components/pagination"
```

## Components

| Component                | Renders                          | Notes                                                   |
| ------------------------ | -------------------------------- | ------------------------------------------------------- |
| `pagination.Pagination`  | `<nav aria-label="…">`           | Outer wrapper with a flex row for Previous / List / Next. |
| `pagination.Previous`    | `<span>` + flint-ui Button       | Left-aligned back-one-page link. Disabled when Href is empty. |
| `pagination.List`        | `<span>` (hidden on mobile)      | Wraps the numbered Pages between Previous and Next.     |
| `pagination.Page`        | flint-ui Button                  | Single numbered link. `Current: true` highlights it.    |
| `pagination.Gap`         | `<span aria-hidden>`             | Ellipsis standing in for collapsed page ranges.         |
| `pagination.Next`        | `<span>` + flint-ui Button       | Right-aligned forward-one-page link.                    |

## Quick example

```go
@pagination.Pagination(pagination.Props{}) {
    @pagination.Previous(pagination.LinkProps{Href: "?page=2"}) { Previous }
    @pagination.List(pagination.ListProps{}) {
        @pagination.Page(pagination.PageProps{Href: "?page=1"}) { 1 }
        @pagination.Page(pagination.PageProps{Href: "?page=2"}) { 2 }
        @pagination.Page(pagination.PageProps{Href: "?page=3", Current: true}) { 3 }
        @pagination.Page(pagination.PageProps{Href: "?page=4"}) { 4 }
        @pagination.Page(pagination.PageProps{Href: "?page=5"}) { 5 }
    }
    @pagination.Next(pagination.LinkProps{Href: "?page=4"}) { Next }
}
```

## Server-driven by default

Each Page/Previous/Next renders as a native `<a href>`. Full page navigation works out of the box with zero JS.

### htmx-driven variant

To swap a page fragment without a full reload, drop `hx-*` attrs via `Attrs`:

```go
@pagination.Page(pagination.PageProps{
    Href: "#",
    Attrs: templ.Attributes{
        "hx-get":    "/jobs?page=2",
        "hx-target": "#jobs-table",
        "hx-swap":   "outerHTML",
    },
}) { 2 }
```

Href is still required (falsy Href triggers the disabled state); use `#` or the intended deep-link URL. htmx intercepts the click before the browser follows the link.

## Disabled edges

Catalyst's trick: when `Href == ""`, the component renders a disabled button instead of a link. This lets the caller pass the same `Previous` / `Next` components on every page without conditional swaps:

```go
// Server logic computes previous/next hrefs — empty when at the edge.
var prev, next string
if page > 1 { prev = fmt.Sprintf("?page=%d", page-1) }
if page < totalPages { next = fmt.Sprintf("?page=%d", page+1) }

@pagination.Previous(pagination.LinkProps{Href: prev}) { Previous }
@pagination.Next(pagination.LinkProps{Href: next}) { Next }
```

On the first page `prev == ""` and Previous auto-disables; symmetric on the last page for Next.

## Mobile behavior

`List` hides on viewports below `sm`. Small screens show only Previous / Next, which is the standard mobile pattern — swiping to the right page number with your thumb is impractical on phones. If your mobile design genuinely needs page numbers, pass `ListProps.Class: "flex"` to override.

## Accessibility

- `<nav>` with `aria-label="Page navigation"` (override via `Props.AriaLabel` when a page has multiple paginations).
- Current page gets `aria-current="page"`.
- Previous / Next have `aria-label="Previous page"` / `"Next page"` so icon-only renderings are still announced.
- `Gap` is `aria-hidden` — the collapsed range isn't meaningful to a screen reader; the explicit page numbers before and after convey what they need.

## Props reference

### Props (Pagination root)

| Field       | Type               | Default             | Notes                                         |
| ----------- | ------------------ | ------------------- | --------------------------------------------- |
| `AriaLabel` | `string`           | `"Page navigation"` | Override for pages with multiple paginations. |
| `Class`     | `string`           | `""`                | Appended to the `<nav>`.                      |
| `Attrs`     | `templ.Attributes` | `nil`               | Spread onto the `<nav>`.                      |

### LinkProps (Previous / Next)

| Field   | Type               | Default | Notes                                            |
| ------- | ------------------ | ------- | ------------------------------------------------ |
| `Href`  | `string`           | `""`    | Empty renders a disabled button.                 |
| `Class` | `string`           | `""`    | Appended to the wrapper span.                    |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the underlying button. Use for `hx-*`. |

### ListProps

| Field   | Type               | Default | Notes                           |
| ------- | ------------------ | ------- | ------------------------------- |
| `Class` | `string`           | `""`    | Appended to the list wrapper.   |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the list wrapper.   |

### PageProps

| Field     | Type               | Default | Notes                                                         |
| --------- | ------------------ | ------- | ------------------------------------------------------------- |
| `Href`    | `string`           | `""`    | Required. Use `#` for htmx-intercepted links.                 |
| `Current` | `bool`             | `false` | Adds `before:bg-muted` highlight + `aria-current="page"`.     |
| `Class`   | `string`           | `""`    | Appended to the button class.                                 |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the button. Use for `hx-*` or custom aria-label. |

### GapProps

| Field   | Type               | Default | Notes                          |
| ------- | ------------------ | ------- | ------------------------------ |
| `Class` | `string`           | `""`    | Appended to the gap span.      |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the gap span.      |
