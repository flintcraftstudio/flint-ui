# Text

Converted from `catalyst-ui-kit/typescript/text.tsx`. These are the body-copy primitives: a muted paragraph, an underlined inline link, strong emphasis, and inline code. Catalyst's `Link` is a plain `<a>` (it exists for router integration), so its `TextLink` becomes `text.Link` here.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/text"
```

## Components

| Component     | Renders   | Notes                                                        |
| ------------- | --------- | ------------------------------------------------------------ |
| `text.Text`   | `<p data-slot="text">` | Muted body paragraph.                           |
| `text.Link`   | `<a>`     | Underlined inline link in the foreground color. `Href` required. |
| `text.Strong` | `<strong>` | Medium weight, foreground color — stands out from muted copy. |
| `text.Code`   | `<code>`  | Inline code: IDs, keys, file names.                          |

## Usage

```go
@text.Text(text.Props{}) {
    Invoices are sent from{ " " }
    @text.Strong(text.Props{}) {
        billing@ridgeline.com
    }
    { " " }on the first business day. Change the sender in{ " " }
    @text.Link(text.LinkProps{Href: "/settings/billing"}) {
        billing settings
    }
    { "." }
}
```

### Whitespace in templ

templ trims whitespace around component calls, so inline elements need two habits:

- **Spell out spaces** between a word and an inline component with `{ " " }`. Without it, `sent from` + Strong renders as "frombilling".
- **Put inline children on their own lines.** `@text.Link(...) { settings }` on one line keeps a trailing space *inside* the link, so the underline runs past the word.

## When to use what

- `text.Text` is for running copy: page and section intros, help text, card bodies. Form help text belongs in `fieldset.Description`, which sits correctly in the Field grid.
- `text.Link` is for links inside sentences. For standalone actions ("View all", "New job"), use `button.Button` with `Href`.
- `text.Strong` is for emphasis inside muted copy. Headings are a job for `heading`.

## Props

### Props (Text, Strong, Code)

| Field   | Type               | Default | Notes                    |
| ------- | ------------------ | ------- | ------------------------ |
| `Class` | `string`           | `""`    | Appended to the element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the element. |

### LinkProps

| Field   | Type               | Default | Notes                                   |
| ------- | ------------------ | ------- | --------------------------------------- |
| `Href`  | `string`           | —       | Required.                               |
| `Class` | `string`           | `""`    | Appended to the `<a>`.                  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<a>`. Use for `hx-*`, `target`, `rel`. |

## Accessibility

- Links are underlined, not just colored, so they're distinguishable without color (WCAG 1.4.1).
- Keyboard focus shows the `ring` token's outline.
- Link text should name the destination. Avoid "click here".
