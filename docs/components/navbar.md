# Navbar

Converted from `catalyst-ui-kit/typescript/navbar.tsx`. Horizontal top navigation, used across the top of `layout.Stacked` and as the mobile header in `layout.Sidebar`.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/navbar"
```

## Components

| Component        | Renders             | Notes                                                     |
| ---------------- | ------------------- | --------------------------------------------------------- |
| `navbar.Navbar`  | `<nav>`             | Flex row.                                                 |
| `navbar.Section` | `<div>`             | Groups items in a row.                                    |
| `navbar.Item`    | `<a>` or `<button>` | A nav entry. `Href` → link; empty → button.               |
| `navbar.Label`   | `<span>`            | Item text, truncated.                                     |
| `navbar.Spacer`  | `<div>`             | Pushes the following sections to the right edge.          |
| `navbar.Divider` | `<div>`             | Short vertical rule between groups. Decorative.           |

## Example

```go
@navbar.Navbar(navbar.Props{}) {
    @navbar.Item(navbar.ItemProps{Href: "/", Class: "max-lg:hidden"}) {
        @avatar.Avatar(avatar.Props{Src: "/static/logo.svg", Square: true})
        @navbar.Label(navbar.SubProps{}) { Ridgeline Roofing }
    }
    @navbar.Divider(navbar.SubProps{Class: "max-lg:hidden"})
    @navbar.Section(navbar.SubProps{Class: "max-lg:hidden"}) {
        @navbar.Item(navbar.ItemProps{Href: "/jobs", Current: current == "jobs"}) { Jobs }
        @navbar.Item(navbar.ItemProps{Href: "/crews", Current: current == "crews"}) { Crews }
    }
    @navbar.Spacer(navbar.SubProps{})
    @navbar.Section(navbar.SubProps{}) {
        @navbar.Item(navbar.ItemProps{Href: "/search", Attrs: templ.Attributes{"aria-label": "Search"}}) {
            <svg data-slot="icon" ...></svg>
        }
        @dropdown.Dropdown(dropdown.Props{}) {
            @navbar.Item(navbar.ItemProps{Attrs: dropdown.TriggerAttrs(templ.Attributes{"aria-label": "Account menu"})}) {
                @avatar.Avatar(avatar.Props{Initials: "LW", Square: true})
            }
            @dropdown.Menu(dropdown.MenuProps{Anchor: dropdown.AnchorBottomEnd}) { ... }
        }
    }
}
```

In `layout.Stacked`, sections marked `max-lg:hidden` disappear below `lg`; put the same links in the shell's `Sidebar` so they're reachable from the drawer.

## Items

- **Icons**: an `<svg data-slot="icon">` child is sized and tinted, and darkens on hover. Give icon-only items an `aria-label`.
- **Avatars**: an `avatar.Avatar` child is sized with slightly rounded corners.
- **Current**: `Current: true` draws an underline below the item and sets `aria-current="page"` on links. It's rendered statically per page; Catalyst animates it between items.

## Props

### Props / SubProps

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root element. |

### ItemProps

| Field     | Type               | Default | Notes                                                     |
| --------- | ------------------ | ------- | --------------------------------------------------------- |
| `Href`    | `string`           | `""`    | Set → `<a>`. Empty → `<button type="button">`.            |
| `Current` | `bool`             | `false` | Underline + `aria-current="page"`.                        |
| `Class`   | `string`           | `""`    | Appended to the outer `<span>` wrapper.                   |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the `<a>` / `<button>`.                       |

## Accessibility

- Renders a `<nav>`. Add `aria-label` via `Attrs` when the page has more than one.
- Icon-only items need an `aria-label`.
- Items show the `ring` token's focus outline for keyboard users, and have a 44px touch target on coarse pointers.
