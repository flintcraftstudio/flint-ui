# Button

Converted from `catalyst-ui-kit/typescript/button.tsx`. Solid, outline, and plain styles wired to the semantic token contract in [`styles/flint-ui.css`](../../styles/flint-ui.css). Optionally renders as `<a>` when `Href` is set.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/button"
```

## Usage

```go
@button.Button(button.Props{}) {
    Save changes
}

@button.Button(button.Props{Variant: button.VariantAccent}) {
    Create project
}

@button.Button(button.Props{Outline: true, Href: "/docs"}) {
    Read the docs
}

@button.Button(button.Props{Plain: true, Variant: button.VariantDanger}) {
    Delete draft
}

@button.Button(button.Props{
    Variant: button.VariantDanger,
    Attrs: templ.Attributes{
        "hx-delete":  "/orders/42",
        "hx-confirm": "Delete this order?",
        "hx-target":  "closest tr",
        "hx-swap":    "outerHTML",
    },
}) {
    Delete
}
```

## Props

| Field      | Type               | Default           | Notes                                                                 |
| ---------- | ------------------ | ----------------- | --------------------------------------------------------------------- |
| `Variant`  | `Variant`          | `VariantPrimary` (solid), neutral (outline/plain) | One of the five semantic variants. Solid buttons fill with it; outline and plain buttons take its text color and hover tint. Unset, outline and plain stay neutral. |
| `Outline`  | `bool`             | `false`           | Outline with a `border` token border. Neutral (`foreground` + `muted`) unless `Variant` is set. |
| `Plain`    | `bool`             | `false`           | Borderless until hover. Neutral unless `Variant` is set.             |
| `Class`    | `string`           | `""`              | Extra classes appended after the generated className.                 |
| `Href`     | `string`           | `""`              | When set, renders `<a>` with the same styling instead of `<button>`.  |
| `Type`     | `string`           | `"button"`        | Sets `type=` on the `<button>`. Ignored when `Href` is set.           |
| `Disabled` | `bool`             | `false`           | Sets `disabled` on `<button>` or `aria-disabled` on `<a>`.            |
| `Indicator` | `bool`            | `false`           | Swap the label for a spinner while an htmx request is in flight. See [Loading state](#loading-state). |
| `Attrs`    | `templ.Attributes` | `nil`             | Passed through to the underlying element. Use for `hx-*`, `data-*`, `aria-*`. |

## Loading state

Set `Indicator` and the button shows a spinner in place of its label while htmx has a request in flight — from the button itself (`hx-post` on the button) or from the `<form>` it submits. It keys off htmx's `.htmx-request` class through the `htmx-request:` variant in `flint-ui.css`, so there's no Alpine or JS wiring. The label stays in the layout (just invisible), so the button doesn't change width.

```go
// hx-* on the button
@button.Button(button.Props{Indicator: true, Attrs: templ.Attributes{
    "hx-post":    "/jobs/1834/approve",
    "hx-target":  "#job-status",
    "hx-disable": "this",
}}) { Approve }

// submit button in an htmx form
<form hx-post="/crew/invite" hx-target="#invite-result" hx-disable="find button">
    ...
    @button.Button(button.Props{Type: "submit", Indicator: true}) { Send invite }
</form>
```

- Pair it with `hx-disable` (htmx 4's name for the old `hx-disabled-elt`) to block double submits. The indicator also sets `pointer-events: none` and keeps the disabled button at full opacity while it's busy.
- If the form uses `hx-indicator` to point somewhere else, htmx puts `.htmx-request` there instead and the button won't spin.
- The spinner is decorative. Announce the result (a toast, an updated status) rather than the wait.

## Variants

| Variant          | Solid background | Solid text           | Outline/plain text | Outline/plain hover | Typical use                 |
| ---------------- | ---------------- | -------------------- | ------------------ | ------------------- | --------------------------- |
| `VariantPrimary` | `primary`        | `primary-foreground` | `primary`          | `primary/10`        | Default action on a screen. |
| `VariantAccent`  | `accent`         | `accent-foreground`  | `accent-text`      | `accent/10`         | Secondary brand action.     |
| `VariantDanger`  | `danger`         | `danger-foreground`  | `danger-text`      | `danger/10`         | Destructive actions.        |
| `VariantSuccess` | `success`        | `success-foreground` | `success-text`     | `success/10`        | Positive confirmations.     |
| `VariantWarning` | `warning`        | `warning-foreground` | `warning-text`     | `warning/25`        | Caution flows.              |

Outline and plain buttons use the `-text` tokens (the hue mixed toward `foreground`), the same ones badges and alerts use, because the bare hue fails AA on its own hover tint: `danger` on `danger/10` over `canvas` is 3.78:1. Primary has no `-text` mix and passes as is. The icon takes the same token, at 70% at rest and full on hover, and the outline keeps the neutral `border` token. `styles/contrast_test.go` checks each text token on `background`, `surface` and `canvas`, at rest and over its hover tint, in both themes.

```go
// Red text, no fill: a destructive action that shouldn't shout.
@button.Button(button.Props{Outline: true, Variant: button.VariantDanger}) { Delete item }
```

## Style precedence

`Outline` > `Plain` > solid. Only one style applies at a time, and `Variant` colors whichever it is:

- **Solid** (both booleans `false`): an unset `Variant` means `VariantPrimary`.
- **Outline or plain**: an unset `Variant` means neutral, exactly as before v1.3.0. An explicit `VariantPrimary` tints the button with the primary hue, like any other variant; leave `Variant` unset for a neutral button. An unknown value is also neutral.

## Theming

Every visible color comes from a CSS variable declared in `styles/flint-ui.css`'s `@theme` block. Client sites rebrand by overriding those variables at `:root`:

```css
:root {
  --color-primary: #0a0a0a;           /* A-Team Gutters black */
  --color-primary-foreground: #ffffff;
  --color-accent: #dc2626;            /* A-Team Gutters red */
  --color-accent-foreground: #ffffff;
}
```

`bg-primary`, `text-primary-foreground`, and the custom properties used by Catalyst's layered button structure (`[--btn-bg:var(--color-primary)]`, `[--btn-hover-overlay:var(--color-primary-foreground)]/10`) all update in lockstep — no component changes needed.

## Accessibility

- Renders as a real `<button>` or `<a>` — no role shims, full keyboard support by default.
- `Disabled` maps to the `disabled` attribute on `<button>` and `aria-disabled="true"` + `data-disabled` on `<a>` (anchors can't be truly disabled; the attribute cues AT and our Tailwind selectors).
- Focus outline uses `data-focus:outline-ring`, aliased to `:focus-visible` in `styles/flint-ui.css`.
- A `pointer-fine:hidden` overlay expands the touch target to 44×44px on coarse pointers, matching Catalyst's `<TouchTarget>`.
