# Dropdown

Converted from `catalyst-ui-kit/typescript/dropdown.tsx` (Catalyst's wrapper around Headless UI's `Menu`). flint-ui swaps the Headless UI runtime for Alpine: a `{ open: false }` scope on the root, click-outside / ESC / Tab to close, and the [Alpine Focus plugin](https://alpinejs.dev/plugins/focus) for keyboard navigation.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/dropdown"
```

## Components

| Component               | Renders                              | Notes                                                      |
| ----------------------- | ------------------------------------ | ---------------------------------------------------------- |
| `dropdown.Dropdown`     | `<div>` with `x-data`                | Outer shell. Owns Alpine state, anchors the menu.          |
| `dropdown.Button`       | flint-ui Button                      | Trigger. Toggles open, sets aria-expanded, ArrowDown opens + focuses first item. |
| `dropdown.Menu`         | `<div role="menu">`                  | Floating panel. Positioned via `Anchor`. Click anywhere inside closes. |
| `dropdown.Item`         | `<button role="menuitem">` or `<a>`  | Single row. Set `Href` for a link, leave it empty for a button. |
| `dropdown.Header`       | `<div>`                              | Non-interactive top row (e.g. signed-in user info).        |
| `dropdown.Section`      | `<div role="group">`                 | Groups Items under a Heading.                              |
| `dropdown.Heading`      | `<h3>`                               | Section title.                                             |
| `dropdown.Divider`      | `<hr role="separator">`              | Visual break.                                              |
| `dropdown.Label`        | `<span data-slot="label">`           | Item's primary text (column 2 in the subgrid).             |
| `dropdown.Description`  | `<span data-slot="description">`     | Item's secondary line (column 2, row 2).                   |
| `dropdown.Shortcut`     | `<kbd>` of `<kbd>`s                  | Keyboard hint at the right edge of an Item.                |

## Quick example

```go
@dropdown.Dropdown(dropdown.Props{}) {
    @dropdown.Button(dropdown.ButtonProps{Outline: true}) {
        Actions
        @chevronDown()
    }
    @dropdown.Menu(dropdown.MenuProps{}) {
        @dropdown.Item(dropdown.ItemProps{Href: "/jobs/1834/edit"}) { Edit job }
        @dropdown.Item(dropdown.ItemProps{Attrs: templ.Attributes{
            "hx-post":   "/jobs/1834/duplicate",
            "hx-target": "#job-list",
        }}) { Duplicate }
        @dropdown.Divider(dropdown.SubProps{})
        @dropdown.Item(dropdown.ItemProps{Class: "text-danger"}) {
            Delete
        }
    }
}
```

## Anchor positioning

The menu is `absolute`-positioned inside the Dropdown root (`relative inline-block`). Pick the anchor based on where the trigger sits in the layout — `end` alignment keeps menus from overflowing the viewport when the trigger is near the right edge.

| Anchor                        | Menu position                              |
| ----------------------------- | ------------------------------------------ |
| `dropdown.AnchorBottomStart`  | Below trigger, left edges aligned. Default. |
| `dropdown.AnchorBottomEnd`    | Below trigger, right edges aligned.        |
| `dropdown.AnchorTopStart`     | Above trigger, left edges aligned.         |
| `dropdown.AnchorTopEnd`       | Above trigger, right edges aligned.        |

### Overflow caveat

The menu is *not* teleported to `<body>` (unlike Modal). Ancestor `overflow:hidden`, `transform`, or `filter` will create a stacking context that clips or reorders the menu. If you hit that, either:

- Move the dropdown outside the offending container, or
- Remove the overflow constraint on that ancestor.

A future version may add an Alpine Anchor + Teleport pattern to escape these. For now, plan ancestors accordingly.

## Items: button vs link

`Item` switches on `Href`:

- `Href == ""` — renders `<button type="button">`. Use for client-side actions (htmx posts, Alpine handlers).
- `Href != ""` — renders `<a href="...">`. Use for navigation.

Both render with `role="menuitem"`, both are focusable, both close the menu on click (the click bubbles up to the Menu's `x-on:click`).

### htmx pass-through

Add `hx-*` via `Attrs`:

```go
@dropdown.Item(dropdown.ItemProps{Attrs: templ.Attributes{
    "hx-post":   "/api/sign-out",
    "hx-target": "body",
    "hx-swap":   "outerHTML",
}}) { Sign out }
```

The menu closes via the bubbled click; htmx fires its request. Both happen.

### Disabled

`Disabled: true` sets the native `disabled` on a button or `aria-disabled="true"` on a link. The `data-disabled:opacity-50 data-disabled:pointer-events-none` classes (mapped to `:disabled`/`[aria-disabled="true"]` via the `data-disabled` custom variant in `flint.css`) handle both visuals and click suppression.

## Items with icon, description, shortcut

The menu uses CSS subgrid (with a fallback grid) to align icon | label | description | shortcut columns across all items. Icons must carry `data-slot="icon"` so the shared item styles position and tint them.

```go
@dropdown.Item(dropdown.ItemProps{Href: "#"}) {
    @userPlusIcon() // <svg data-slot="icon" ...>
    @dropdown.Label(dropdown.SubProps{}) { Add crew member }
    @dropdown.Description(dropdown.SubProps{}) { Send an invite by email }
    @dropdown.Shortcut(dropdown.ShortcutProps{Keys: []string{"⌘", "N"}})
}
```

Icons are tinted with `text-muted-foreground` by default and re-tinted to `text-accent-foreground` when the item is highlighted (hover or keyboard focus). If you use an `<img>` avatar instead, mark it `data-slot="avatar"` for the slightly larger sizing.

### Keys array

`Shortcut.Keys` is rendered one `<kbd>` per entry. Multi-character entries after the first get a small left pad so combinations like `["Ctrl", "Shift", "K"]` read cleanly.

## Sections, headings, header

```go
@dropdown.Menu(dropdown.MenuProps{Anchor: dropdown.AnchorBottomEnd}) {
    @dropdown.Header(dropdown.SubProps{}) {
        <div class="text-sm font-medium text-foreground">Jaden Sturgis</div>
        <div class="text-xs text-muted-foreground">jaden@flintcraft.studio</div>
    }
    @dropdown.Divider(dropdown.SubProps{})
    @dropdown.Section(dropdown.SubProps{}) {
        @dropdown.Heading(dropdown.SubProps{}) { Account }
        @dropdown.Item(dropdown.ItemProps{Href: "/profile"}) {
            @dropdown.Label(dropdown.SubProps{}) { Profile }
        }
    }
    @dropdown.Divider(dropdown.SubProps{})
    @dropdown.Item(dropdown.ItemProps{Attrs: templ.Attributes{
        "hx-post": "/sign-out",
    }}) {
        @dropdown.Label(dropdown.SubProps{}) { Sign out }
    }
}
```

`Section` keeps the subgrid running through it, so headings and items in different sections still column-align.

## Keyboard

Driven by the [Alpine Focus plugin](https://alpinejs.dev/plugins/focus). **Load `@alpinejs/focus` before `alpinejs` core** — same requirement as Modal. The showcase layout already does this; clients must mirror it in their base layout:

```html
<script src="https://unpkg.com/@alpinejs/focus@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

| Key                       | Behavior                                       |
| ------------------------- | ---------------------------------------------- |
| `Enter` / `Space` on trigger | Open menu (native button click).            |
| `ArrowDown` on trigger    | Open menu, focus first item.                   |
| `ArrowDown` / `ArrowUp` in menu | Move focus, wrapping at ends.            |
| `Home` / `End` in menu    | First / last item.                             |
| `Enter` on item           | Activate (native button click or link follow). |
| `Tab` in menu             | Close.                                         |
| `Esc`                     | Close.                                         |

`x-on:keydown.escape.window` is bound on the root, so ESC closes from anywhere on the page (including focus inside an htmx-loaded form).

## Required surrounding setup

Same baseline as Modal:

1. **`x-data` on a common ancestor.** Triggers' `x-on:click` is silently ignored without one. The standard fix is `<body x-data>` in your base layout.
2. **Alpine Focus plugin loaded before Alpine core.** Otherwise `$focus.within(...)` is undefined and keyboard nav silently no-ops.

## Trigger styles

`dropdown.Button` forwards `Variant` / `Outline` / `Plain` / `Disabled` / `Class` to the underlying `button.Button`. The Alpine attrs are layered on top; caller-supplied `Attrs` win on conflict so you can override the click handler if you need a different toggle behavior.

```go
// Default — primary solid button
@dropdown.Button(dropdown.ButtonProps{}) { Actions }

// Subtle, fits in dense table rows
@dropdown.Button(dropdown.ButtonProps{Plain: true}) { More }

// Strong call-to-action
@dropdown.Button(dropdown.ButtonProps{Variant: button.VariantAccent}) { New }
```

## Highlight color

Hovered or keyboard-focused items use `bg-accent` / `text-accent-foreground`. This means the highlight inherits each client's brand accent — Lo Mo's teal, A-Team's red, Western Skies' gold. Override per-item via `Class` if you need a different color (e.g. the Delete item using `data-focus:bg-danger` in the showcase).

## Accessibility

- Trigger renders with `aria-haspopup="menu"` and `aria-expanded` bound to Alpine state.
- Menu renders with `role="menu"`; items with `role="menuitem"`; sections with `role="group"`; dividers with `role="separator"`.
- Disabled items use native `disabled` (button) or `aria-disabled` (link). Both map to the `data-disabled` variant for visuals.
- Forced-colors mode (Windows high-contrast): item highlight falls back to `Highlight` / `HighlightText`; divider falls back to `CanvasText`. Preserved verbatim from Catalyst.
- Focus is *not* trapped (menus shouldn't trap — Tab closes and exits).

## Props reference

### Props (Dropdown root)

| Field   | Type               | Default | Notes                                                         |
| ------- | ------------------ | ------- | ------------------------------------------------------------- |
| `Class` | `string`           | `""`    | Appended to the root `<div>`.                                 |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root. Use for `data-*`, extra Alpine refs, etc. |

### ButtonProps (Trigger)

| Field      | Type               | Default | Notes                                                       |
| ---------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Variant`  | `button.Variant`   | `""`    | Button variant. Resolves to `Primary` if empty.             |
| `Outline`  | `bool`             | `false` | Outline button style.                                       |
| `Plain`    | `bool`             | `false` | Plain (borderless) button style.                            |
| `Class`    | `string`           | `""`    | Appended to the underlying button's class.                  |
| `Disabled` | `bool`             | `false` | Disables the trigger.                                       |
| `Attrs`    | `templ.Attributes` | `nil`   | Merged into the button. Caller wins on key conflicts.       |

### MenuProps (Panel)

| Field    | Type               | Default               | Notes                                                |
| -------- | ------------------ | --------------------- | ---------------------------------------------------- |
| `Anchor` | `Anchor`           | `AnchorBottomStart`   | Where the menu lands relative to the trigger.        |
| `Class`  | `string`           | `""`                  | Appended to the menu `<div>`. Common: `min-w-[18rem]` to widen for description rows. |
| `Attrs`  | `templ.Attributes` | `nil`                 | Spread onto the menu.                                |

### ItemProps

| Field      | Type               | Default | Notes                                                      |
| ---------- | ------------------ | ------- | ---------------------------------------------------------- |
| `Href`     | `string`           | `""`    | If non-empty, renders as `<a>`. Otherwise as `<button>`.   |
| `Disabled` | `bool`             | `false` | Native `disabled` on button, `aria-disabled` on link.      |
| `Class`    | `string`           | `""`    | Appended to the item.                                      |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the item. Use for `hx-*`, `x-on:click`, etc.   |

### SubProps (Header / Section / Heading / Divider / Label / Description)

| Field   | Type               | Default | Notes                            |
| ------- | ------------------ | ------- | -------------------------------- |
| `Class` | `string`           | `""`    | Appended to the rendered element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the rendered element. |

### ShortcutProps

| Field   | Type               | Default | Notes                                                       |
| ------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Keys`  | `[]string`         | `nil`   | Each entry renders as a separate `<kbd>`. e.g. `["⌘", "S"]`. |
| `Class` | `string`           | `""`    | Appended to the outer `<kbd>`.                              |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the outer `<kbd>`.                              |
