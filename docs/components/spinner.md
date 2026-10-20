# Spinner

An indeterminate loading indicator, from scratch (Catalyst has none). It inherits `currentColor`, so set its color with a `text-*` class on it or an ancestor.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/spinner"
```

## Usage

```go
// Decorative — next to visible "Loading…" copy, or inside a button
@spinner.Spinner(spinner.Props{})

// Announced — role="status" with screen-reader text
@spinner.Spinner(spinner.Props{Size: spinner.SizeLG, Label: "Loading jobs"})
```

## As an htmx indicator

Add the `htmx-indicator` class and point `hx-indicator` at the spinner. htmx's built-in indicator CSS hides it until the request is in flight (`htmx.config.includeIndicatorCSS`, on by default):

```go
@button.Button(button.Props{Outline: true, Attrs: templ.Attributes{
    "hx-get":       "/reports/weekly",
    "hx-target":    "#report",
    "hx-indicator": "#report-spinner",
}}) { Load report }
@spinner.Spinner(spinner.Props{Class: "htmx-indicator", Attrs: templ.Attributes{"id": "report-spinner"}})
```

For a spinner *inside* a button, use `button.Props.Indicator` instead. It swaps the label for a spinner with no extra wiring. See [Button → Loading state](./button.md#loading-state).

## Sizes

| Size              | Dimensions                         | Use for                                   |
| ----------------- | ---------------------------------- | ----------------------------------------- |
| `spinner.SizeSM`  | 16px                               | Inline with small text.                   |
| `spinner.SizeMD`  | 20px on touch, 16px at `sm+`        | Default. Matches button icons.            |
| `spinner.SizeLG`  | 32px                               | Loading a panel or section.               |

## Props

| Field   | Type               | Default  | Notes                                                                 |
| ------- | ------------------ | -------- | --------------------------------------------------------------------- |
| `Size`  | `spinner.Size`     | `SizeMD` |                                                                       |
| `Label` | `string`           | `""`     | Set → wraps the SVG in `role="status"` with `sr-only` text. Empty → `aria-hidden`. |
| `Class` | `string`           | `""`     | Appended to the root (the status wrapper, or the SVG).               |
| `Attrs` | `templ.Attributes` | `nil`    | Spread onto the root.                                                 |

## Accessibility

- Leave `Label` empty when the context already says what's happening (the button label, nearby copy). A labelled spinner inside a labelled button double-announces.
- Screen readers reliably announce changes *inside* an existing live region, not a region that was just inserted. For swapped-in content, announce the result (a toast, or updated text in an `aria-live` container) rather than relying on the spinner.
- Under `prefers-reduced-motion` the spinner keeps turning, slowly (3s per turn). A frozen spinner reads as "stuck".
