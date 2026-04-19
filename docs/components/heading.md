# Heading

Converted from `catalyst-ui-kit/typescript/heading.tsx`. Two components: `Heading` (page title, defaults to `<h1>`) and `Subheading` (section title, defaults to `<h2>`). Both decouple the HTML tag from the visual size — `Level` controls the element, class controls the look.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/heading"
```

## Usage

```go
@heading.Heading(heading.Props{}) { Dashboard }
@heading.Subheading(heading.Props{}) { Outstanding invoices }
```

Override the HTML level when the document outline requires it, without changing the visual weight:

```go
// Still looks like a Heading, but renders as <h2> because the page already has an <h1>.
@heading.Heading(heading.Props{Level: 2}) { Modal title }
```

## Why decouple size from level?

A modal or side panel often has its own title that should register as an `<h2>` in the page outline (the page itself already has an `<h1>`) but should still visually read as a primary heading inside its container. Catalyst's `Heading` solves this by making Level a prop. flint-ui preserves that.

## Props

| Field   | Type               | Default           | Notes                                                    |
| ------- | ------------------ | ----------------- | -------------------------------------------------------- |
| `Level` | `Level` (int 1–6)  | `1` for Heading, `2` for Subheading | Zero or out-of-range values fall back to the default. |
| `Class` | `string`           | `""`              | Appended to the rendered element.                        |
| `Attrs` | `templ.Attributes` | `nil`             | Spread onto the rendered element.                        |

## Accessibility

- Semantic heading elements (`<h1>` … `<h6>`) render as the requested `Level` — screen readers announce them correctly.
- Don't skip levels in the outline (e.g. `<h1>` → `<h4>`). Use `Subheading` with an explicit `Level` when depth matters.
