# Button

Converted from `catalyst-ui-kit/typescript/button.tsx`. Solid, outline, and plain styles wired to the semantic token contract in [`styles/flint.css`](../../styles/flint.css). Optionally renders as `<a>` when `Href` is set.

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
| `Variant`  | `Variant`          | `VariantPrimary`  | One of the five semantic variants. Ignored if `Outline` or `Plain`.   |
| `Outline`  | `bool`             | `false`           | Neutral outline (uses `border` + `foreground` + `muted` tokens).      |
| `Plain`    | `bool`             | `false`           | Borderless until hover.                                               |
| `Class`    | `string`           | `""`              | Extra classes appended after the generated className.                 |
| `Href`     | `string`           | `""`              | When set, renders `<a>` with the same styling instead of `<button>`.  |
| `Type`     | `string`           | `"button"`        | Sets `type=` on the `<button>`. Ignored when `Href` is set.           |
| `Disabled` | `bool`             | `false`           | Sets `disabled` on `<button>` or `aria-disabled` on `<a>`.            |
| `Attrs`    | `templ.Attributes` | `nil`             | Passed through to the underlying element. Use for `hx-*`, `data-*`, `aria-*`. |

## Variants

| Variant          | Background token | Text token              | Typical use                 |
| ---------------- | ---------------- | ----------------------- | --------------------------- |
| `VariantPrimary` | `primary`        | `primary-foreground`    | Default action on a screen. |
| `VariantAccent`  | `accent`         | `accent-foreground`     | Secondary brand action.     |
| `VariantDanger`  | `danger`         | `danger-foreground`     | Destructive actions.        |
| `VariantSuccess` | `success`        | `success-foreground`    | Positive confirmations.     |
| `VariantWarning` | `warning`        | `warning-foreground`    | Caution flows.              |

## Style precedence

`Outline` > `Plain` > solid-with-`Variant`. Only one applies at a time; `Variant` is only consulted when both booleans are `false`.

## Theming

Every visible color comes from a CSS variable declared in `styles/flint.css`'s `@theme` block. Client sites rebrand by overriding those variables at `:root`:

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
- Focus outline uses `data-focus:outline-ring`, aliased to `:focus-visible` in `styles/flint.css`.
- A `pointer-fine:hidden` overlay expands the touch target to 44×44px on coarse pointers, matching Catalyst's `<TouchTarget>`.
