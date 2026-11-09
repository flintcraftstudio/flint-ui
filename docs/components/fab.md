# Floating action button

A round + button in a phone's thumb zone, just above the tab bar, for the main action of the section on screen (New quote, New client). With two or more actions it opens a short menu above itself. Catalyst has no equivalent. It is hidden from `lg` up, where the tab bar becomes a sidebar: put the same action in the page header there.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/fab"
```

## Components

| Component  | Renders     | Notes |
| ---------- | ----------- | ----- |
| `fab.Fab`  | `<div><a>`  | One action: a round + link. |
| `fab.Menu` | `<details>` | A + button (`<summary>`) that opens its items above it. |
| `fab.Item` | `<a>`       | One of a Menu's actions. Tapping it closes the menu. |

## Example

```go
// One action, opening a sheet with JavaScript and a page without.
@fab.Fab(fab.Props{Label: "New quote", Href: "/quotes/new",
    Attrs: templ.Attributes{"hx-get": "/quotes/new/sheet", "hx-target": "#picker", "hx-swap": "innerHTML"}})

// Two actions.
@fab.Menu(fab.Props{Label: "New"}) {
    @fab.Item(fab.ItemProps{Href: "/clients/new"}) { New client }
    @fab.Item(fab.ItemProps{Href: "/quotes/new"}) { New quote }
}
```

Use a Fab while a section has one action, and a Menu from two: a menu holding one item costs a tap for nothing.

## Placement

- It is fixed at the right, 1rem above flint-ui's tab bar (`bottom: calc(5rem + env(safe-area-inset-bottom))`), at `z-20`. Leave room at the foot of the page so the last row can scroll clear of it (`pb-24` on the page's content is enough).
- Hidden from `lg` up. Render the action in the page header as a button too; the header's copy is the one a wide screen shows.
- A page without the tab bar should not use it: its offset assumes the bar.

## Behaviour

- The Menu is a native `<details>`: it opens, and its links work, without JavaScript.
- With Alpine, `x-data="flintFab"` (registered in flint-ui's script) closes it on a tap outside, on Escape (focus goes back to the button) and on a tap on an item, so an item that opens a sheet doesn't leave the menu open behind it.
- The + turns 45° into an × while the menu is open.
- Don't link an item to an element inside the menu: following a link to an element in a closed `<details>` opens it again.

## Props

### Props

| Field   | Type               | Default | Notes |
| ------- | ------------------ | ------- | ----- |
| `Label` | `string`           | `""`    | The button's accessible name: the action ("New quote") for a Fab, what the menu offers ("New") for a Menu. Required: the button shows only a +. |
| `Href`  | `string`           | `""`    | The Fab's link. A Menu ignores it. |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the Fab's `<a>`, or the Menu's `<details>`. |

### ItemProps

| Field   | Type               | Default | Notes |
| ------- | ------------------ | ------- | ----- |
| `Href`  | `string`           | `""`    | The action's link. |
| `Class` | `string`           | `""`    | Appended to the `<a>`. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<a>` (`hx-get` and friends). |

## Accessibility

- The button is 56px and each item 48px tall, over the 44px tap target; keyboard focus shows the `ring` token's outline.
- `Label` names the icon-only button. A `<summary>` reports itself expanded or collapsed on its own.
- The menu is a disclosure of links, not an ARIA `menu`: Tab moves through the items, Enter follows one.
- The button is the primary hue on primary-foreground, the token pair checked for AA.
