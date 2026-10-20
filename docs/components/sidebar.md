# Sidebar

Converted from `catalyst-ui-kit/typescript/sidebar.tsx`. Vertical app navigation, used by `layout.Sidebar` (desktop rail and mobile drawer) and as the mobile drawer in `layout.Stacked`.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/sidebar"
```

## Components

| Component         | Renders              | Notes                                                          |
| ----------------- | -------------------- | -------------------------------------------------------------- |
| `sidebar.Sidebar` | `<nav>`              | Full-height flex column.                                       |
| `sidebar.Header`  | `<div>`              | Pinned top region (team switcher, search). Bordered below.     |
| `sidebar.Body`    | `<div>`              | Scrollable middle region holding the nav sections.             |
| `sidebar.Footer`  | `<div>`              | Pinned bottom region (account menu). Bordered above.           |
| `sidebar.Section` | `<div data-slot="section">` | Groups items. Consecutive sections space themselves.    |
| `sidebar.Heading` | `<h3>`               | Section label ("Recent jobs").                                 |
| `sidebar.Item`    | `<a>` or `<button>`  | A nav entry. `Href` → link; empty → button.                    |
| `sidebar.Label`   | `<span>`             | Item text, truncated.                                          |
| `sidebar.Spacer`  | `<div>`              | Inside Body: pushes the following sections to the bottom.      |
| `sidebar.Divider` | `<hr>`               | Rule between sections.                                         |

## Example

```go
@sidebar.Sidebar(sidebar.Props{}) {
    @sidebar.Header(sidebar.SubProps{}) {
        @sidebar.Item(sidebar.ItemProps{Href: "/"}) {
            @avatar.Avatar(avatar.Props{Src: "/static/logo.svg", Square: true})
            @sidebar.Label(sidebar.SubProps{}) { Ridgeline Roofing }
        }
    }
    @sidebar.Body(sidebar.SubProps{}) {
        @sidebar.Section(sidebar.SubProps{}) {
            @sidebar.Item(sidebar.ItemProps{Href: "/jobs", Current: current == "jobs"}) {
                <svg data-slot="icon" ...></svg>
                @sidebar.Label(sidebar.SubProps{}) { Jobs }
            }
            @sidebar.Item(sidebar.ItemProps{Href: "/crews", Current: current == "crews"}) {
                <svg data-slot="icon" ...></svg>
                @sidebar.Label(sidebar.SubProps{}) { Crews }
            }
        }
        @sidebar.Section(sidebar.SubProps{}) {
            @sidebar.Heading(sidebar.SubProps{}) { Recent jobs }
            @sidebar.Item(sidebar.ItemProps{Href: "/jobs/1834"}) { Henderson residence }
        }
        @sidebar.Spacer(sidebar.SubProps{})
        @sidebar.Section(sidebar.SubProps{}) {
            @sidebar.Item(sidebar.ItemProps{Href: "/help"}) { Support }
        }
    }
    @sidebar.Footer(sidebar.SubProps{Class: "max-lg:hidden"}) {
        ...account dropdown, see below
    }
}
```

## Items

- **Icons**: put an `<svg data-slot="icon">` (20px solid Heroicons fit) as the first child. It's sized, tinted `muted-foreground`, and darkens on hover and for the current item.
- **Trailing icon**: an icon as the last child (a chevron) is pushed to the right and drawn smaller.
- **Avatars**: an `avatar.Avatar` child (it carries `data-slot="avatar"`) is sized to match.
- **Current**: `Current: true` draws a bar in the left gutter, darkens the icon, and sets `aria-current="page"` on links. Catalyst animates the bar between items; flint-ui renders it statically per page.

## Account menu in the footer

Use `dropdown.TriggerAttrs` to make an Item the Dropdown trigger, and open the menu upward:

```go
@dropdown.Dropdown(dropdown.Props{Class: "w-full"}) {
    @sidebar.Item(sidebar.ItemProps{Attrs: dropdown.TriggerAttrs(nil)}) {
        @avatar.Avatar(avatar.Props{Initials: "LW", Square: true})
        @sidebar.Label(sidebar.SubProps{}) { Logan }
        <svg data-slot="icon" ...chevron-up></svg>
    }
    @dropdown.Menu(dropdown.MenuProps{Anchor: dropdown.AnchorTopStart, Class: "min-w-64"}) {
        @dropdown.Item(dropdown.ItemProps{Href: "/settings"}) { Settings }
        @dropdown.Item(dropdown.ItemProps{Href: "/logout"}) { Sign out }
    }
}
```

## Props

### Props / SubProps (every component except Item)

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root element. |

### ItemProps

| Field     | Type               | Default | Notes                                                           |
| --------- | ------------------ | ------- | --------------------------------------------------------------- |
| `Href`    | `string`           | `""`    | Set → `<a>`. Empty → `<button type="button">`.                  |
| `Current` | `bool`             | `false` | Marks the current page: bar, dark icon, `aria-current="page"`.  |
| `Class`   | `string`           | `""`    | Appended to the outer `<span>` wrapper.                         |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the `<a>` / `<button>`. Use for `hx-*`, `x-on:*`, `dropdown.TriggerAttrs`. |

## Accessibility

- Renders a `<nav>`. If the page has another nav landmark, add `aria-label` via `Attrs`.
- The current link gets `aria-current="page"`.
- Items show the `ring` token's focus outline for keyboard users, and have a 44px touch target on coarse pointers.
