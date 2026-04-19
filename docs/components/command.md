# Command Palette

Cmd+K / Ctrl+K overlay for searching and invoking actions across an app. From-scratch design — Catalyst has no `command.tsx`.

The component sits at the intersection of [Modal](./modal.md) and [Combobox](./combobox.md):

- Modal's plumbing: teleport to `<body>`, backdrop, focus trap (`x-trap.noscroll`), ESC-to-close, event-driven open.
- Combobox's plumbing: filterable text input feeding a list of items, arrow-key navigation, Enter to activate the first match.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/command"
```

## Components

| Component         | Renders                                   | Notes                                                  |
| ----------------- | ----------------------------------------- | ------------------------------------------------------ |
| `command.Palette` | `<template x-teleport="body">`            | Outer overlay. Place once per page (usually in the base layout). |
| `command.Group`   | `<div role="group">` with heading + items | Labeled section. Hides when none of its items match the current query. |
| `command.Item`    | `<button>` or `<a>`                       | Single command. Href → link; no Href → button.         |

## Quick example

```go
@command.Palette(command.Props{}) {
    @command.Group(command.GroupProps{Heading: "Navigation"}) {
        @command.Item(command.ItemProps{Label: "Go to dashboard", Href: "/", Shortcut: "G D"})
        @command.Item(command.ItemProps{Label: "Go to jobs", Href: "/jobs", Shortcut: "G J"})
    }
    @command.Group(command.GroupProps{Heading: "Create"}) {
        @command.Item(command.ItemProps{
            Label:    "New job",
            Shortcut: "⌘ J",
            Attrs:    templ.Attributes{"x-on:click": "$dispatch('flint:modal-open', 'new-job')"},
        })
    }
}
```

Press `Cmd+K` / `Ctrl+K` to open the palette. Type to filter. Enter activates the first visible item. Esc closes.

## Placement

Put the Palette in your **base layout**, not a per-page template. It teleports to `<body>` so its position in source doesn't affect rendering, but the keyboard shortcut listener only fires while the element is in the DOM — layout placement ensures Cmd+K works on every route.

```go
// layout.templ
@command.Palette(command.Props{}) {
    // App-wide commands, sourced from a per-client command catalog
    @command.Group(...) { ... }
}
```

## Opening

Four ways:

| Method                        | How                                           |
| ----------------------------- | --------------------------------------------- |
| Keyboard                      | `Cmd+K` (macOS) or `Ctrl+K` (Linux/Windows)   |
| Dispatch event                | `$dispatch('flint:command-open')`             |
| Trigger button                | `<button x-on:click="$dispatch('flint:command-open')">Commands</button>` |
| Close event                   | `$dispatch('flint:command-close')` (from anywhere) |

The keyboard shortcut is a window-level listener installed by the palette itself — no caller setup required. Disable it via `Props.DisableShortcut` when your app already uses Cmd+K for something else (callers can still trigger via event).

## Keyboard

Requires [`@alpinejs/focus`](https://alpinejs.dev/plugins/focus) — already loaded by the showcase layout.

| Key                           | Behavior                                             |
| ----------------------------- | ---------------------------------------------------- |
| `Cmd+K` / `Ctrl+K`            | Toggle open (unless `DisableShortcut` is set)        |
| Type in input                 | Filter items by substring; empty Groups hide         |
| `ArrowDown` (in input)        | Move focus to first visible item                     |
| `ArrowDown` / `ArrowUp` (in item) | Cycle through visible items; wraps              |
| `Enter` (in input)            | Activate the first visible item                      |
| `Enter` (in item)             | Activate that item (native click)                    |
| `Escape`                      | Close the palette; clear query                       |

## Item types

### Link items (navigation)

```go
@command.Item(command.ItemProps{Label: "Go to jobs", Href: "/jobs"})
```

Renders as `<a href>`. On click, the palette closes and the browser follows the link naturally.

### Button items (actions)

```go
@command.Item(command.ItemProps{
    Label: "Export CSV",
    Attrs: templ.Attributes{
        "hx-post":   "/export",
        "hx-target": "#export-status",
    },
})
```

Renders as `<button type="button">`. Caller wires the action via `x-on:click` or `hx-*`. The palette closes on click; the caller's handler runs via Alpine's handler chain.

### Disabled items

```go
@command.Item(command.ItemProps{
    Label:    "Reset database",
    Shortcut: "admin only",
    Disabled: true,
})
```

Dimmed; click and arrow-focus both skip.

## Shortcuts are visual

`Shortcut` renders a keyboard hint on the right edge of the row — but it doesn't actually bind the shortcut. If you want `G J` to actually take the user to /jobs, add your own global `keydown` listener. The Shortcut prop is purely a discoverability cue.

Separating "shortcut display" from "shortcut binding" keeps the palette component out of the business of global keyboard arbitration. Clients can bind shortcuts however they like (per-app patterns vary).

## Groups hide when empty

Each Group's `x-show` queries its child Items' `data-label` against the current query — the heading stays hidden when nothing below it matches.

Empty-state overall: when the query matches zero items across all groups, a "No results" placeholder renders inside the list. Style via component overrides if you need custom empty text.

## Accessibility

- The panel carries `role="dialog"` + `aria-modal="true"` + `aria-label="Command palette"`.
- `x-trap.noscroll` traps Tab focus inside and locks body scroll — same as Modal.
- Search input has `role="combobox"`, `aria-autocomplete="list"`, `aria-expanded="true"`, `aria-controls` pointing at the listbox.
- Items are native `<button>` or `<a>` with `tabindex="-1"` — programmatically focusable via arrow keys, announced correctly to screen readers.
- Group headings are `aria-hidden="true"`; `role="group"` carries the semantic grouping.

## Required surrounding setup

- **`x-data` on a common ancestor.** Same baseline as Modal / Dropdown / Combobox.
- **`@alpinejs/focus` loaded before Alpine core.**

No additional plugins needed.

## Props reference

### Props (Palette root)

| Field             | Type               | Default                            | Notes                                                           |
| ----------------- | ------------------ | ---------------------------------- | --------------------------------------------------------------- |
| `Placeholder`     | `string`           | `"Type a command or search…"`      | Placeholder in the search input.                                |
| `FooterHint`      | `string`           | `"↑↓ Navigate · ↵ Select · Esc Close"` | Small helper text at the bottom of the palette.             |
| `DisableShortcut` | `bool`             | `false`                            | Turns off the built-in Cmd+K / Ctrl+K window listener.          |
| `Class`           | `string`           | `""`                               | Appended to the panel `<div>`.                                  |
| `Attrs`           | `templ.Attributes` | `nil`                              | Spread onto the panel.                                          |

### GroupProps

| Field     | Type               | Default | Notes                                                 |
| --------- | ------------------ | ------- | ----------------------------------------------------- |
| `Heading` | `string`           | `""`    | Required. Section label shown above the group's items. |
| `Class`   | `string`           | `""`    | Appended to the group `<div>`.                         |
| `Attrs`   | `templ.Attributes` | `nil`   | Spread onto the group.                                 |

### ItemProps

| Field      | Type               | Default | Notes                                                       |
| ---------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Label`    | `string`           | `""`    | Required. Displayed text and filter-match source.            |
| `Shortcut` | `string`           | `""`    | Optional keyboard hint rendered on the right. Visual only.   |
| `Href`     | `string`           | `""`    | Set to render as `<a>`; omit for `<button>`.                |
| `Disabled` | `bool`             | `false` | Dims and blocks activation.                                 |
| `Class`    | `string`           | `""`    | Appended to the item element.                               |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the item. Use for `x-on:click` or `hx-*` wiring. |
