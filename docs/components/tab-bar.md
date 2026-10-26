# Tab bar

An app's main navigation as tabs: a bar fixed to the bottom of a phone's screen, in the thumb zone, that becomes a narrow rail at the left from `lg` up. Catalyst has no equivalent. Use it instead of a layout shell's drawer when the app has a few top-level sections (three to five) and is used mostly on phones.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/tabbar"
```

## Components

| Component        | Renders | Notes                                                                 |
| ---------------- | ------- | --------------------------------------------------------------------- |
| `tabbar.TabBar`  | `<nav>` | Fixed bar below `lg`, fixed rail from `lg`. One equal column per child. |
| `tabbar.Item`    | `<a>`   | One tab: an icon over a short label.                                  |
| `tabbar.Spacer`  | `<div>` | On the rail, pushes the following items to the bottom. Nothing on the bar. |

`tabbar.ContentClasses` is a class string for the element wrapping the page: it keeps the content clear of the bar below `lg` and of the rail from `lg`. It sets that element's bottom and left padding, so put the page's own padding on an element inside it.

## Example

```go
@tabbar.TabBar(tabbar.Props{Attrs: templ.Attributes{"aria-label": "Main"}}) {
    <span class="mb-3 max-lg:hidden">...brand mark</span>
    @tabbar.Item(tabbar.ItemProps{Href: "/", Current: current == "quotes"}) {
        <svg data-slot="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" aria-hidden="true">...</svg>
        Quotes
    }
    @tabbar.Item(tabbar.ItemProps{Href: "/clients", Current: current == "clients"}) {
        <svg data-slot="icon" ...></svg>
        Clients
    }
    @tabbar.Spacer(tabbar.SubProps{})
    @tabbar.Item(tabbar.ItemProps{Href: "/settings", Current: current == "settings"}) {
        <svg data-slot="icon" ...></svg>
        Settings
    }
}
<main class={ tabbar.ContentClasses }>
    <div class="px-4 py-8">...page content</div>
</main>
```

## Items

- **Icons**: put an `<svg data-slot="icon">` first. It's drawn at 24px and should use `currentColor` (as `stroke` or `fill`), so it follows the label's color.
- **Labels**: one short word. The bar splits the screen's width between the tabs, and a label doesn't wrap well under an icon.
- **Current**: `Current: true` sets `aria-current="page"` and darkens and bolds the label and icon; on the rail it also tints the tab's background. The server marks the current item on each page.
- **Hidden below `lg`**: the bar gives each child a column, but a child with `display: none` takes none. A brand mark with `max-lg:hidden` shows only on the rail, and the Spacer takes no column on the bar.

## Placement

- The bar sits at `z-20`. A page's own fixed bottom bar (a full-width primary action) must sit above it, for example with `bottom-[calc(4rem+env(safe-area-inset-bottom))]` below `lg`.
- Pages that shouldn't show the tabs (an editor, a full-screen form) simply don't render the TabBar, and drop `ContentClasses`.
- The bar pads the iPhone home indicator with `env(safe-area-inset-bottom)`, which is non-zero only when the page's viewport meta has `viewport-fit=cover`.

## Props

### Props / SubProps

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root element. |

### ItemProps

| Field     | Type               | Default | Notes                                               |
| --------- | ------------------ | ------- | --------------------------------------------------- |
| `Href`    | `string`           | `""`    | The tab's link.                                     |
| `Current` | `bool`             | `false` | Marks the current page; sets `aria-current="page"`. |
| `Class`   | `string`           | `""`    | Appended to the `<a>`.                              |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the `<a>`.                              |

## Accessibility

- Renders a `<nav>`; give it an `aria-label` ("Main") via `Attrs`, since a page usually has other nav landmarks.
- The current link gets `aria-current="page"`, and is marked by weight as well as color.
- Each tab is 64px tall, over the 44px tap target, and shows the `ring` token's outline, inset, for keyboard focus.
- Plain links, no JavaScript: it works without it and under a strict CSP.
