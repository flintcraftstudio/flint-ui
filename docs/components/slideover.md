# Slide-over

Side drawers — full-height panels that slide in from the left or right edge. From-scratch design (Catalyst has no `slideover.tsx`). Shares Modal's teleport, event-driven open, and focus-trap plumbing; the only differences are an edge-anchored panel and horizontal slide transitions.

## Import

```go
import "github.com/flintcraft/flint-ui/components/slideover"
```

## Components

| Component                   | Renders                      | Notes                                                          |
| --------------------------- | ---------------------------- | -------------------------------------------------------------- |
| `slideover.Slideover`       | `<template x-teleport="body">` | Outer shell: backdrop, edge-anchored flex container, panel. |
| `slideover.Title`           | `<h2>`                       | Heading at the top of the panel. Shares Modal's dialog typography. |
| `slideover.Description`     | `<p>`                        | Supporting copy under the title.                               |
| `slideover.Body`            | `<div>`                      | `flex-1 overflow-y-auto` — the scrolling middle region.        |
| `slideover.Actions`         | `<div>`                      | Footer row with `mt-auto` + top border. Pins to the bottom.    |

## Quick example

```go
@button.Button(button.Props{
    Attrs: templ.Attributes{
        "x-on:click": "$dispatch('flint:slideover-open', 'job-filters')",
    },
}) { Filter jobs }

@slideover.Slideover(slideover.Props{Name: "job-filters"}) {
    @slideover.Title(slideover.SubProps{}) { Filter jobs }
    @slideover.Description(slideover.SubProps{}) {
        Narrow the list by crew, status, and date range.
    }
    @slideover.Body(slideover.SubProps{}) {
        <form class="space-y-6"> ... </form>
    }
    @slideover.Actions(slideover.SubProps{}) {
        @button.Button(button.Props{
            Plain: true,
            Attrs: templ.Attributes{"x-on:click": "open = false"},
        }) { Cancel }
        @button.Button(button.Props{
            Attrs: templ.Attributes{"x-on:click": "open = false"},
        }) { Apply filters }
    }
}
```

## Opening and closing

Same event API as Modal, with a different event name so the two don't collide:

| Action          | Dispatch                                                      |
| --------------- | ------------------------------------------------------------- |
| Open one        | `$dispatch('flint:slideover-open', 'name')`                   |
| Close one       | `$dispatch('flint:slideover-close', 'name')`                  |
| Close all       | `$dispatch('flint:slideover-close')` (no detail)              |
| Close from panel | `x-on:click="open = false"` on any button inside the panel   |
| ESC             | Bound globally; closes the currently-open slideover           |
| Backdrop click  | Closes                                                        |

Mirror Modal's convention: a successful htmx form submit can dispatch `flint:slideover-close` (no detail) via `HX-Trigger: {"flint:slideover-close":null}` to close whatever drawer drove the form.

## Side

| Side                    | Anchor                       |
| ----------------------- | ---------------------------- |
| `slideover.SideRight`   | Right edge. Default. Standard for detail panes and filter drawers. |
| `slideover.SideLeft`    | Left edge. Common for mobile navigation.                           |

Top and bottom sides aren't supported in v0.1 — add when a client needs a bottom sheet.

## Size

Width cap applied at sm+. Below sm the panel fills the viewport (`w-screen`).

| Size                  | Max-width at sm+ |
| --------------------- | ---------------- |
| `slideover.SizeSM`    | `24rem`          |
| `slideover.SizeMD`    | `28rem`          |
| `slideover.SizeLG`    | `32rem` (default) |
| `slideover.SizeXL`    | `36rem`          |
| `slideover.Size2XL`   | `42rem`          |

Pick the smallest size that fits the content comfortably — a narrow drawer is less visually disruptive when invoked from a crowded dashboard.

## Layout model

The panel is `flex flex-col`:

- `Title` and `Description` sit at the top via natural document flow.
- `Body` has `flex-1 overflow-y-auto` — it grows to fill the middle and scrolls when content overflows.
- `Actions` has `mt-auto` — it pins to the bottom regardless of body length, with a top border separating it from the scrolling area.

This is the one structural difference from Modal. A modal is a compact vertical stack; a slide-over is a full-height drawer where the scrolling middle matters.

### Headerless drawers

If you omit Title and Description, Body fills the full panel. Mobile nav drawers often look cleanest this way — a simple column of links with no chrome.

## Accessibility

- `role="dialog"` + `aria-modal="true"` on the panel container.
- `x-trap.inert.noscroll="open"` on the panel itself: traps Tab focus, marks page siblings `inert` for assistive tech, locks body scroll while the panel is open.
- ESC, backdrop click, and `x-on:click="open = false"` on any panel button all close.
- Forced-colors mode: `forced-colors:outline` preserves the panel boundary.

## Required surrounding setup

Same as Modal and Dropdown:

1. **`x-data` on a common ancestor.** Triggers' `x-on:click="$dispatch(...)"` are ignored without an enclosing Alpine scope. Put `x-data` on `<body>` in your base layout.
2. **Alpine Focus plugin loaded before Alpine core.** `x-trap` is undefined without it and the focus trap silently no-ops.

```html
<script src="https://unpkg.com/@alpinejs/focus@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

## Props reference

### Props (Slideover root)

| Field    | Type               | Default      | Notes                                                         |
| -------- | ------------------ | ------------ | ------------------------------------------------------------- |
| `Name`   | `string`           | `""`         | Required. Stable URL-safe identifier; used for open/close event matching. |
| `Side`   | `Side`             | `SideRight`  | Which edge the panel anchors to.                              |
| `Size`   | `Size`             | `SizeLG`     | Max-width at sm+.                                             |
| `Class`  | `string`           | `""`         | Appended to the panel `<div>`.                                |
| `Attrs`  | `templ.Attributes` | `nil`        | Spread onto the panel.                                        |

### SubProps (Title / Description / Body / Actions)

| Field   | Type               | Default | Notes                              |
| ------- | ------------------ | ------- | ---------------------------------- |
| `Class` | `string`           | `""`    | Appended to the rendered element.  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the rendered element.  |
