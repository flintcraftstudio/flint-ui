# Tooltip

Hover- and focus-activated hint bubble over a trigger. From-scratch design (Catalyst has no `tooltip.tsx`).

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/tooltip"
```

## Quick example

```go
@tooltip.Tooltip(tooltip.Props{ID: "tip-save", Label: "Save draft"}) {
    @button.Button(button.Props{Attrs: templ.Attributes{"aria-describedby": "tip-save"}}) {
        Save
    }
}
```

The component renders a wrapper `<span>` with the trigger inside. The hint bubble is appended as an absolute-positioned sibling, shown on `mouseenter` / `focusin`, hidden on `mouseleave` / `focusout` / `Escape`.

## How it works

The wrapper holds the Alpine scope `{ open: false }`. Events on the wrapper:

- `mouseenter` → `open = true`
- `mouseleave` → `open = false`
- `focusin` → `open = true` (bubbles up from any inner focusable)
- `focusout` → `open = false`
- `keydown.escape.stop` → `open = false`

`focusin` / `focusout` bubble from descendant elements, so any focusable trigger works — buttons, links, inputs, custom `tabindex=0` elements.

## Accessibility

**Caller must add `aria-describedby={Props.ID}`** on the actual focusable element. Without it, the tooltip is visual-only — screen readers won't announce the hint on focus. The tooltip `<span>` carries `role="tooltip"` and `id={Props.ID}` automatically.

```go
@tooltip.Tooltip(tooltip.Props{ID: "tip-archive", Label: "Archive job"}) {
    @button.Button(button.Props{
        Plain: true,
        Attrs: templ.Attributes{
            "aria-describedby": "tip-archive",  // required for a11y
            "aria-label":       "Archive",      // also needed for icon-only buttons
        },
    }) { @iconArchive() }
}
```

The bubble itself is `pointer-events-none` so it can't intercept clicks meant for the trigger or any overlapping content.

## Anchor positioning

| Anchor                   | Bubble position                    |
| ------------------------ | ---------------------------------- |
| `tooltip.AnchorTop`      | Above trigger, centered. Default.  |
| `tooltip.AnchorBottom`   | Below trigger, centered.           |
| `tooltip.AnchorLeft`     | Left of trigger, centered vertically. |
| `tooltip.AnchorRight`    | Right of trigger, centered vertically. |

No start/end variants — tooltips are narrow enough that perpendicular-axis centering is the right default. For edge-of-viewport triggers, flip to the opposite anchor.

### Overflow caveat

The bubble is *not* teleported to `<body>`. Ancestor `overflow:hidden`, `transform`, or `filter` will clip it — same caveat Dropdown has. If you hit clipping, move the tooltip outside the offending container or remove the overflow constraint.

## Styling

The bubble uses `bg-foreground text-background` for inverted contrast — dark on light themes, light on dark themes automatically via the token contract. Default width is `w-max max-w-xs`; override via `Props.Class`:

```go
@tooltip.Tooltip(tooltip.Props{ID: "tip-wide", Label: "...", Class: "max-w-md"}) { ... }
```

## Content

`Props.Label` is a plain string. For rich content (links, markup, keyboard shortcut chips) reach for a Popover (future component). Keeping Tooltip to text avoids the children-vs-content ambiguity Modal and Dropdown have to manage and keeps the bubble narrow enough that center-anchoring works.

## Keyboard

| Key      | Behavior                                   |
| -------- | ------------------------------------------ |
| `Tab`    | Focus the trigger. Bubble appears.         |
| `Shift+Tab` | Focus leaves the trigger. Bubble hides. |
| `Esc`    | Hide the bubble even while trigger is focused. |

## Required surrounding setup

- **`x-data` on a common ancestor** — the wrapper's Alpine directives need an enclosing scope. The showcase puts `x-data` on `<body>`; clients must do the same in their base layout.

No Alpine plugins required. Tooltip does not use `@alpinejs/focus` (unlike Modal and Dropdown).

## Props reference

### Props

| Field    | Type               | Default       | Notes                                                             |
| -------- | ------------------ | ------------- | ----------------------------------------------------------------- |
| `Label`  | `string`           | `""`          | The plain-text hint shown in the bubble.                          |
| `ID`     | `string`           | `""`          | Required for screen-reader linkage. Caller sets `aria-describedby={ID}` on the trigger. Use URL/id-safe chars; unique per page. |
| `Anchor` | `Anchor`           | `AnchorTop`   | Which side of the trigger the bubble sits on.                     |
| `Class`  | `string`           | `""`          | Appended to the bubble `<span>`. Common: `max-w-md` to widen.     |
| `Attrs`  | `templ.Attributes` | `nil`         | Spread onto the bubble. Use for `data-*` hooks, extra `x-on:*`, etc. |
