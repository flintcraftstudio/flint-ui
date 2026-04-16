# Button

Converted from `catalyst-ui-kit/typescript/button.tsx`. Solid, outline, and plain styles across Catalyst's 20-color palette. Optionally renders as `<a>` when `Href` is set.

## Import

```go
import "github.com/flintcraft/flint-ui/components/button"
```

## Usage

```go
@button.Button(button.Props{}) {
    Save changes
}

@button.Button(button.Props{Color: button.ColorIndigo}) {
    Create project
}

@button.Button(button.Props{Outline: true, Href: "/docs"}) {
    Read the docs
}

@button.Button(button.Props{
    Color: button.ColorRed,
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

| Field      | Type               | Default        | Notes                                                                        |
| ---------- | ------------------ | -------------- | ---------------------------------------------------------------------------- |
| `Color`    | `Color`            | `ColorDarkZinc` | One of 20 solid palettes. Ignored if `Outline` or `Plain` is set.            |
| `Outline`  | `bool`             | `false`        | Neutral outline style; takes precedence over `Color`.                        |
| `Plain`    | `bool`             | `false`        | Borderless style; takes precedence over `Color`.                             |
| `Class`    | `string`           | `""`           | Extra classes appended after the generated className.                        |
| `Href`     | `string`           | `""`           | When set, renders `<a>` with the same styling instead of `<button>`.         |
| `Type`     | `string`           | `"button"`     | Sets `type=` on the `<button>`. Ignored when `Href` is set.                  |
| `Disabled` | `bool`             | `false`        | Sets `disabled` on `<button>` or `aria-disabled` on `<a>`.                   |
| `Attrs`    | `templ.Attributes` | `nil`          | Passed through to the underlying element. Use for `hx-*`, `data-*`, `aria-*`. |

## Style precedence

`Outline` > `Plain` > solid-with-`Color`. Only one applies at a time; `Color` is only consulted when both booleans are `false`.

## Colors

```go
button.ColorDarkZinc   button.ColorLight     button.ColorDarkWhite
button.ColorDark       button.ColorWhite     button.ColorZinc
button.ColorIndigo     button.ColorCyan      button.ColorRed
button.ColorOrange     button.ColorAmber     button.ColorYellow
button.ColorLime       button.ColorGreen     button.ColorEmerald
button.ColorTeal       button.ColorSky       button.ColorBlue
button.ColorViolet     button.ColorPurple    button.ColorFuchsia
button.ColorPink       button.ColorRose
```

## Accessibility

- Renders as a real `<button>` or `<a>` — no role shims, full keyboard support by default.
- `Disabled` maps to the `disabled` attribute on `<button>` and `aria-disabled="true"` + `data-disabled` on `<a>` (anchors can't be truly disabled; the attribute cues AT and our Tailwind selectors).
- Focus outline uses `data-focus:outline-blue-500`, mapped to `:focus-visible` in `styles/flint.css`.
- A `pointer-fine:hidden` overlay expands the touch target to 44×44px on coarse pointers, matching Catalyst's `<TouchTarget>`.

## Relation to Catalyst

Catalyst relies on Headless UI's `<Button>` wrapper to set `data-hover`, `data-focus`, `data-active`, and `data-disabled` attributes based on React state. Since we render HTML on the server without Headless UI, `styles/flint.css` uses Tailwind v4 `@custom-variant` declarations to alias those `data-*` selectors onto the native `:hover`, `:focus-visible`, `:active`, and `:disabled` pseudo-classes. This lets us preserve Catalyst's class strings verbatim.
