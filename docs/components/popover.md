# Popover

Click-activated floating panel with arbitrary content — forms, rich info, buttons, whatever fits. From-scratch design (Catalyst has no `popover.tsx`).

## Popover vs. Dropdown vs. Tooltip

Three nearby components with distinct purposes:

| Component                           | Trigger        | Content                              | Role / Nav                                    |
| ----------------------------------- | -------------- | ------------------------------------ | --------------------------------------------- |
| [Tooltip](./tooltip.md)             | Hover / focus  | Plain text hint                      | Decorative; no focusable children             |
| [Dropdown](./dropdown.md)           | Click          | Menu of focusable items              | `role="menu"`, keyboard nav through items     |
| **Popover**                         | Click          | Arbitrary content (forms, text, …)   | No default role; Tab flows naturally          |

Pick Popover when you need a floating panel with *content* that isn't a menu — a "what's this?" explanation, an inline form, a detail card on a badge.

## Import

```go
import "github.com/flintcraft/flint-ui/components/popover"
```

## Components

| Component          | Renders                        | Notes                                                    |
| ------------------ | ------------------------------ | -------------------------------------------------------- |
| `popover.Popover`  | `<div>` with `x-data`          | Outer shell. Owns Alpine state, anchors the panel.       |
| `popover.Button`   | flint-ui Button                | Trigger. Toggles open; sets aria-expanded + aria-haspopup. |
| `popover.Panel`    | `<div>` positioned absolutely  | Floating panel with chrome + default padding.            |

## Quick example

```go
@popover.Popover(popover.Props{}) {
    @popover.Button(popover.ButtonProps{Outline: true}) {
        Add tag
    }
    @popover.Panel(popover.PanelProps{}) {
        <form class="space-y-3">
            @fieldset.Field(fieldset.Props{}) {
                @fieldset.Label(fieldset.LabelProps{For: "tag-name"}) { Name }
                @input.Input(input.Props{Name: "tag", ID: "tag-name"})
            }
            @button.Button(button.Props{
                Attrs: templ.Attributes{"x-on:click": "open = false"},
            }) { Save }
        </form>
    }
}
```

## How it differs from Modal

A popover is **not a modal**:

- No teleport to `<body>` — the panel is absolute-positioned inside the wrapper.
- **No focus trap.** Tab flows naturally through the panel's content and continues past the last focusable element to the next thing on the page. If that's a problem, reach for Modal instead.
- **No body scroll lock.** The page behind stays scrollable.
- No backdrop. Clicking anywhere outside the panel closes it; interactions elsewhere on the page are not blocked.

This matches Radix's and Headless UI's popover semantics. Popovers are meant to be interruptible.

## Anchor positioning

Same four anchors as Dropdown:

| Anchor                       | Panel position                              |
| ---------------------------- | ------------------------------------------- |
| `popover.AnchorBottomStart`  | Below trigger, left edges aligned. Default. |
| `popover.AnchorBottomEnd`    | Below trigger, right edges aligned.         |
| `popover.AnchorTopStart`     | Above trigger, left edges aligned.          |
| `popover.AnchorTopEnd`       | Above trigger, right edges aligned.         |

### Overflow caveat

The panel is *not* teleported — ancestor `overflow:hidden`, `transform`, or `filter` creates a stacking context that clips or reorders it. Same caveat as Dropdown. If a client project runs into clipping, move the popover outside the offending container.

## Closing

Four ways to close a popover:

| Action          | How                                                      |
| --------------- | -------------------------------------------------------- |
| Click outside   | Automatic (`x-on:click.outside` on the panel)            |
| Press ESC       | Automatic (`x-on:keydown.escape.window` on the root)     |
| Explicit button | Add `x-on:click="open = false"` on a button in the panel |
| Toggle trigger  | Click the trigger again                                  |

Clicking *inside* the panel does **not** close it — unlike Dropdown where every menu item click dismisses. Popover content is arbitrary; the caller decides what counts as "done" and wires a Save or Close button.

## Accessibility

- Trigger renders with `aria-haspopup="dialog"` and `aria-expanded` bound to Alpine state.
- Panel renders with **no default role**. `role="dialog"` requires an accessible name via `aria-labelledby` — defaulting to a role we can't enforce labeling on would be an a11y footgun.
- **When to add `role="dialog"`:** the popover has a heading at the top. Wire `Attrs: templ.Attributes{"role": "dialog", "aria-labelledby": "tag-heading"}` and give the heading `id="tag-heading"`.
- **When not to:** the popover holds inline text (definitions, short help). No role needed; screen readers treat the panel as regular content that appears when the trigger is activated.

## Styling

Panel defaults match Dropdown's Menu: `bg-surface`, `rounded-xl`, `shadow-lg`, `ring-1 ring-border`, so popovers and dropdowns sit in the same visual family.

Defaults:

- **Width:** `min-w-[16rem] max-w-xs`. Keeps short content from looking cramped; prevents long content from blowing out. Override per popover via Class (e.g. `Class: "min-w-[20rem]"`).
- **Padding:** `p-4`. For edge-to-edge content (images, code blocks with their own padding), pass `Class: "p-0"`.

## Props reference

### Props (Popover root)

| Field   | Type               | Default | Notes                              |
| ------- | ------------------ | ------- | ---------------------------------- |
| `Class` | `string`           | `""`    | Appended to the root `<div>`.      |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root.              |

### ButtonProps (Trigger)

| Field      | Type               | Default | Notes                                                |
| ---------- | ------------------ | ------- | ---------------------------------------------------- |
| `Variant`  | `button.Variant`   | `""`    | Button variant. Resolves to Primary if empty.        |
| `Outline`  | `bool`             | `false` | Outline button style.                                |
| `Plain`    | `bool`             | `false` | Plain borderless button style.                       |
| `Class`    | `string`           | `""`    | Appended to the underlying button.                   |
| `Disabled` | `bool`             | `false` | Disables the trigger.                                |
| `Attrs`    | `templ.Attributes` | `nil`   | Merged into the button. Caller wins on key conflicts. |

### PanelProps

| Field    | Type               | Default               | Notes                                                  |
| -------- | ------------------ | --------------------- | ------------------------------------------------------ |
| `Anchor` | `Anchor`           | `AnchorBottomStart`   | Where the panel lands relative to the trigger.         |
| `Class`  | `string`           | `""`                  | Appended to the panel `<div>`. Common: `min-w-[20rem]` to widen or `p-0` for edge-to-edge content. |
| `Attrs`  | `templ.Attributes` | `nil`                 | Spread onto the panel. Use for `role="dialog"` + `aria-labelledby` if your content warrants dialog semantics. |

## Required surrounding setup

**`x-data` on a common ancestor.** Triggers' `x-on:click` is ignored without one — same baseline as Modal / Dropdown / Slide-over. Put `x-data` on `<body>` in your base layout.

No Alpine plugins required. Popover is the lightest click-to-open component in flint-ui — no `@alpinejs/focus` (no focus trap), no `@alpinejs/collapse`, no teleport.
