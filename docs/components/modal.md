# Modal

Converted from `catalyst-ui-kit/typescript/dialog.tsx` and `alert.tsx`. Both Catalyst primitives share the same backdrop + panel structure; they differ in size defaults, panel padding, title alignment, and transition. flint-ui folds them into one component — set `Props.Alert = true` for the compact alert layout, leave it off for the larger dialog.

## Import

```go
import "github.com/flintcraft/flint-ui/components/modal"
```

## Components

| Component          | Renders  | Notes                                                        |
| ------------------ | -------- | ------------------------------------------------------------ |
| `modal.Modal`      | `<template x-teleport="body">` wrapping the dialog root | Outer container. Owns Alpine state, backdrop, panel, transitions. |
| `modal.Title`      | `<h2>`   | Dialog title. `SubProps.Alert` must match the parent.        |
| `modal.Description`| `<p>`    | Short supporting copy under the title.                       |
| `modal.Body`       | `<div>`  | Main content — forms, longer text.                           |
| `modal.Actions`    | `<div>`  | Footer row of buttons. Stacks reversed on mobile, right-aligned on sm+. |

All sub-components accept `SubProps{Alert bool, Class string, Attrs templ.Attributes}`. The `Alert` flag must match the parent Modal's — the sub-components render different typography and spacing for the two layouts.

## Usage

### Dialog (default)

```go
templ DeleteJobDialog() {
    @modal.Modal(modal.Props{Name: "delete-job"}) {
        @modal.Title(modal.SubProps{}) { Delete this job? }
        @modal.Description(modal.SubProps{}) { This removes it from the crew's schedule. }
        @modal.Body(modal.SubProps{}) {
            <p>Job #1834 — Henderson Ranch — Apr 18.</p>
        }
        @modal.Actions(modal.SubProps{}) {
            @button.Button(button.Props{
                Plain: true,
                Attrs: templ.Attributes{"x-on:click": "open = false"},
            }) { Cancel }
            @button.Button(button.Props{
                Variant: button.VariantDanger,
                Attrs: templ.Attributes{
                    "hx-delete": "/jobs/1834",
                    "hx-target": "#job-list",
                },
            }) { Delete }
        }
    }
}
```

Trigger from anywhere on the page:

```html
<button x-on:click="$dispatch('flint:modal-open', 'delete-job')">
    Delete
</button>
```

### Alert (confirmation)

```go
@modal.Modal(modal.Props{Name: "unsaved-changes", Alert: true}) {
    @modal.Title(modal.SubProps{Alert: true}) { Discard unsaved changes? }
    @modal.Description(modal.SubProps{Alert: true}) { You can't undo this. }
    @modal.Actions(modal.SubProps{Alert: true}) {
        @button.Button(button.Props{
            Plain: true,
            Attrs: templ.Attributes{"x-on:click": "open = false"},
        }) { Keep editing }
        @button.Button(button.Props{
            Variant: button.VariantDanger,
            Attrs:   templ.Attributes{"x-on:click": "open = false"},
        }) { Discard }
    }
}
```

## Opening and closing

The modal owns its own Alpine state (`{ open: false, name: 'your-name' }`) and listens for two window events.

### Required: `x-data` on a common ancestor

Alpine only processes directives (`x-on:click`, etc.) on elements inside an `x-data` scope. The trigger's `x-on:click="$dispatch(...)"` is silently ignored without one. The standard fix is to put an empty `x-data` on `<body>` in your base layout:

```html
<body x-data>
    <!-- every button here can dispatch flint:modal-open -->
</body>
```

This adds no runtime overhead beyond Alpine walking the tree — there's no state to manage — and it means every trigger on the page works without each one needing its own scope.

### Open

Dispatch `flint:modal-open` with the modal's name as the detail:

```html
<button x-on:click="$dispatch('flint:modal-open', 'delete-job')">Delete</button>
```

The name filter means you can have many modals on one page without collisions — only the matching modal opens.

### Close

Any of these close an open modal:

- **ESC key** — bound on `window`.
- **Clicking the backdrop** — empty area around the panel.
- **A button inside the modal** — use `x-on:click="open = false"`. `open` is in the modal's Alpine scope and is accessible from anywhere inside the panel.
- **Global close event** — `$dispatch('flint:modal-close')` closes every open modal. Useful after a successful htmx response (`hx-on::after-request="$dispatch('flint:modal-close')"`).
- **Targeted close** — `$dispatch('flint:modal-close', 'delete-job')` closes only the named modal.

## Focus management

The panel uses `x-trap.inert.noscroll="open"` (from the [Alpine Focus plugin](https://alpinejs.dev/plugins/focus)):

- **Trap**: Tab cycles inside the panel only.
- **Inert**: siblings of the modal are marked `inert`, so screen readers skip them and mouse clicks outside the panel are intercepted by the backdrop.
- **Noscroll**: `<html>` gets `overflow: hidden` while the modal is open — body scroll lock.

**You must load `@alpinejs/focus` before `alpinejs` core** or the trap directive is a no-op:

```html
<script src="https://unpkg.com/@alpinejs/focus@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

When the modal closes, focus returns to the element that was focused when it opened — this comes free with `x-trap`.

## htmx patterns

### Open a modal after a server response

Any htmx-triggered response can open a modal by dispatching the event — either via `hx-on` on the trigger or via `HX-Trigger` response header:

```
HX-Trigger: {"flint:modal-open": "subscription-required"}
```

### Close a modal after a successful submit

Put `hx-on::after-request="$dispatch('flint:modal-close')"` on the form or submit button — closes every open modal. Useful when you don't know (or don't want to hard-code) which modal is open.

### Render the modal body via htmx

Put `hx-get` on the trigger button and `hx-target` on a `<div>` inside `modal.Body`. The open event fires immediately; the body populates as the response arrives.

## Teleport and stacking contexts

The component wraps everything in `<template x-teleport="body">`, so Alpine moves the rendered modal to the end of `<body>` at init time. This is what makes the modal escape parents with `transform`, `filter`, `overflow: hidden`, or `will-change` — all of which create stacking contexts that would otherwise trap `fixed inset-0` children.

Consequences:

- **The modal is no longer in its authored DOM position.** `display: contents` sibling selectors, `:has()` from ancestors, and CSS counters don't reach it. Style it with its own classes.
- **htmx swaps containing the modal still work.** When the `<template>` is removed by a swap, Alpine tears down the teleported element automatically.
- **Multiple modals on the same page are independent.** Each has its own Alpine scope; they don't share state.

## Sizes

| Size      | `sm:max-w-*` | Default for |
| --------- | ------------ | ----------- |
| `SizeXS`  | `xs`         | —           |
| `SizeSM`  | `sm`         | —           |
| `SizeMD`  | `md`         | Alert       |
| `SizeLG`  | `lg`         | Dialog      |
| `SizeXL`  | `xl`         | —           |
| `Size2XL` | `2xl`        | —           |
| `Size3XL` | `3xl`        | —           |
| `Size4XL` | `4xl`        | —           |
| `Size5XL` | `5xl`        | —           |

Below the `sm` breakpoint the panel fills the viewport regardless of size.

## Accessibility

- Panel renders with `role="dialog"` and `aria-modal="true"` — screen readers announce it as a modal dialog.
- `x-trap.inert` marks siblings outside the modal as `inert` — screen readers skip them while open.
- Focus is trapped inside the panel and restored to the trigger when closed (Alpine Focus plugin).
- `aria-hidden="true"` on the backdrop keeps it out of the accessibility tree.
- If your title isn't obvious from context, wire it to `aria-labelledby`: pass an `id` via `SubProps.Attrs` on the Title and add `aria-labelledby` to `Props.Attrs` on the Modal.

## Props reference

### Props (Modal)

| Field   | Type               | Default | Notes                                                                                                                                                |
| ------- | ------------------ | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Name`  | `string`           | —       | Required. Stable identifier used to target open/close events. Restricted to `[A-Za-z0-9_\-:.]`; unsafe characters are stripped to protect the JS literal. |
| `Size`  | `Size`             | `""`    | Panel max-width on sm+. Zero-value resolves to LG (Dialog) or MD (Alert).                                                                            |
| `Alert` | `bool`             | `false` | Compact alert-dialog layout — centered on mobile, smaller padding, fade-only transition.                                                             |
| `Class` | `string`           | `""`    | Appended to the panel `<div>`.                                                                                                                       |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the panel `<div>`. Use for `aria-labelledby`, `aria-describedby`, `id`, extra `data-*`, etc.                                             |

### SubProps (Title / Description / Body / Actions)

| Field   | Type               | Default | Notes                                                                                      |
| ------- | ------------------ | ------- | ------------------------------------------------------------------------------------------ |
| `Alert` | `bool`             | `false` | Must match the parent Modal's `Alert`. Controls typography and spacing for the sub-element. |
| `Class` | `string`           | `""`    | Appended to the rendered element.                                                          |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the rendered element.                                                          |
