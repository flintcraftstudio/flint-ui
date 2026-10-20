# Layout

Converted from `catalyst-ui-kit/typescript/sidebar-layout.tsx`, `stacked-layout.tsx`, and `auth-layout.tsx`. These are the three application shells: render one as the outermost element of `<body>` and put the page content in its children.

| Shell            | Use for                                                                 |
| ---------------- | ----------------------------------------------------------------------- |
| `layout.Sidebar` | Most admin apps. A fixed 16rem left rail at `lg+`, a hamburger drawer below. |
| `layout.Stacked` | Apps with a handful of top-level sections. A top navbar, a drawer below `lg`. |
| `layout.Auth`    | Login, sign-up, password reset. Centers the children in a raised panel. |

## Import

```go
import (
    "github.com/flintcraftstudio/flint-ui/components/layout"
    "github.com/flintcraftstudio/flint-ui/components/navbar"
    "github.com/flintcraftstudio/flint-ui/components/sidebar"
)
```

## Sidebar shell

Navigation goes in as `templ.Component` props; page content is the children. Build the nav once in a helper that takes the current section, so every page highlights the right item:

```go
templ AppLayout(current string) {
    @layout.Sidebar(layout.SidebarProps{
        Navbar:  appMobileNavbar(),    // shown beside the hamburger below lg
        Sidebar: appSidebar(current),  // the rail at lg+, the drawer below
    }) {
        { children... }
    }
}

templ appSidebar(current string) {
    @sidebar.Sidebar(sidebar.Props{}) {
        @sidebar.Body(sidebar.SubProps{}) {
            @sidebar.Section(sidebar.SubProps{}) {
                @sidebar.Item(sidebar.ItemProps{Href: "/jobs", Current: current == "jobs"}) {
                    @icons.Briefcase()
                    @sidebar.Label(sidebar.SubProps{}) { Jobs }
                }
                ...
            }
        }
    }
}

templ appMobileNavbar() {
    @navbar.Navbar(navbar.Props{}) {
        @navbar.Spacer(navbar.SubProps{})
        @navbar.Section(navbar.SubProps{}) { ...account dropdown }
    }
}
```

Then each page:

```go
templ JobsPage(jobs []Job) {
    @AppLayout("jobs") {
        @pageheader.PageHeader(pageheader.Props{}) { ... }
        ...
    }
}
```

## Stacked shell

`Navbar` is the full top bar at every width. `Sidebar` is the drawer below `lg`, usually the same links laid out vertically. Mark navbar sections that move into the drawer with `Class: "max-lg:hidden"`:

```go
@layout.Stacked(layout.StackedProps{
    Navbar:  appNavbar(current),         // brand, links (max-lg:hidden), spacer, account
    Sidebar: appDrawerSidebar(current),  // the same links, vertical
}) {
    ...page content
}
```

## Auth shell

```go
@layout.Auth(layout.AuthProps{}) {
    <form method="post" action="/login" class="grid w-full max-w-sm grid-cols-1 gap-8">
        ...
    </form>
}
```

The panel is raised (`lg:bg-surface` + ring) only at `lg+`. Give the page a muted background there (e.g. `lg:bg-muted` on `<html>`) so it reads as a card.

## Mobile drawer

Below `lg`, Sidebar and Stacked render a hamburger button that opens the `Sidebar` component in a slide-in drawer. State is Alpine (`mobileNav` on the shell root). It closes on:

- ESC, a backdrop click, or the close button
- **any link click inside it**, so `hx-boost` navigation (which doesn't reload the page) still closes it
- the viewport growing past `lg`, which also releases the scroll lock
- a `flint:sidebar-close` window event, e.g. `HX-Trigger: {"flint:sidebar-close": true}`

Focus is trapped in the drawer and body scroll is locked while it's open (`x-trap.inert.noscroll`).

## Required surrounding setup

- **Alpine Focus plugin, loaded before Alpine core.** Same requirement as Modal; without it the drawer's focus trap is a no-op.
- **The sidebar is rendered twice**, once for the desktop rail and once for the drawer, as in Catalyst. Only one copy is displayed at a time, but don't put fixed `id` attributes inside it.
- **Content width** is capped at `max-w-6xl` (Catalyst's default).

## Props

### SidebarProps / StackedProps

| Field     | Type               | Default | Notes                                                           |
| --------- | ------------------ | ------- | --------------------------------------------------------------- |
| `Navbar`  | `templ.Component`  | `nil`   | Sidebar: the mobile header contents. Stacked: the top bar. `nil` renders nothing. |
| `Sidebar` | `templ.Component`  | `nil`   | Sidebar: the rail + drawer. Stacked: the drawer only.           |
| `Class`   | `string`           | `""`    | Appended to the shell root.                                     |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the shell root (use for `hx-boost`, `hx-target:inherited`, etc.). |

### AuthProps

| Field   | Type               | Default | Notes                      |
| ------- | ------------------ | ------- | -------------------------- |
| `Class` | `string`           | `""`    | Appended to the `<main>`.  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<main>`.  |

## Accessibility

- Content renders in `<main>`; the navigation components render `<nav>`.
- The hamburger and close buttons have `aria-label`s, and the hamburger tracks `aria-expanded`.
- The drawer is `role="dialog"` + `aria-modal="true"`, labelled "Navigation". Siblings are made `inert` while it's open.
