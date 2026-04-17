# flint-ui

Server-rendered UI components for FlintCraft Studio client projects. Built with Go + [templ](https://templ.guide) + htmx + Alpine.js + Tailwind CSS v4, converted from the Tailwind [Catalyst UI Kit](https://tailwindcss.com/plus/ui-kit).

> **Status**: early development. See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the architectural plan and component roadmap.

## Install

```sh
go get github.com/flintcraft/flint-ui@latest
```

Import components one family at a time:

```go
import (
    "github.com/flintcraft/flint-ui/components/button"
    "github.com/flintcraft/flint-ui/components/input"
    "github.com/flintcraft/flint-ui/components/fieldset"
)
```

In your `templ` files:

```go
@button.Button(button.Props{Variant: button.VariantAccent}) {
    Save changes
}
```

## Tailwind

Every color in flint-ui draws from the semantic token contract defined in [`styles/flint.css`](./styles/flint.css)'s `@theme` block (`primary`, `accent`, `danger`, `success`, `warning`, `foreground`, `muted`, `border`, `input`, `ring`, plus `-foreground` pairs). Non-color classes (spacing, sizing, layering, transitions) are preserved verbatim from Catalyst.

Client projects must:

1. Use Tailwind CSS v4.
2. Scan flint-ui templates as content sources:
   ```css
   @source "../vendor/github.com/flintcraft/flint-ui/components/**/*.templ";
   ```
3. Include the `@custom-variant` declarations that alias Headless UI's `data-hover`, `data-focus`, `data-active`, `data-disabled`, `data-invalid` state attributes to native pseudo-classes. Copy the block from [`styles/flint.css`](./styles/flint.css).
4. **Rebrand by overriding CSS variables**, not by editing components:
   ```css
   :root {
     --color-primary: #0a0a0a;            /* client brand primary */
     --color-primary-foreground: #ffffff;
     --color-accent: #dc2626;             /* client brand accent */
     --color-accent-foreground: #ffffff;
   }
   ```
   Every other token falls through to flint-ui's defaults.

## Running the reference site

The `examples/showcase/` app is a Go binary that renders every component in every variant — the flint-ui equivalent of [catalyst-demo.tailwindui.com](https://catalyst-demo.tailwindui.com/).

```sh
mage InstallTailwind    # one-time: download the Tailwind v4 standalone CLI
mage Showcase           # templ generate + CSS build + run showcase on :8080
```

For watch mode during development:

```sh
mage WatchCSS           # terminal 1
go run ./examples/showcase   # terminal 2
```

## Layout

```
components/
  button/               one package per component family
    button.templ
    classes.go
  input/                text inputs + InputGroup icon wrapper
  fieldset/             Fieldset, FieldGroup, Field, Label, Description, ErrorMessage
  shared/               shared prop types
styles/
  flint.css             Tailwind v4 entry: @theme tokens + @custom-variant + @source
examples/
  showcase/             reference site, similar to catalyst-demo
    main.go
    templates/          showcase pages, one per component
    static/
docs/
  components/           per-component markdown reference
```

## Principles

Non-negotiable:

- **Server-rendered first.** No React. No client-side state machinery. Components emit complete HTML.
- **htmx-compatible by default.** Every component accepts `Attrs templ.Attributes` for `hx-*` pass-through.
- **Alpine.js only for local UI state.** Dropdown open/closed, modal visibility, tab selection — never business data.
- **Semantic tokens for color; Catalyst structure for everything else.** Components never reference raw Tailwind color classes (`bg-zinc-900`, `bg-red-600`). Spacing, sizing, borders, transitions track the Catalyst source verbatim.
- **Accessibility preserved.** ARIA attributes, focus management, and keyboard behavior track the Catalyst source.

See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the full rules and the semantic token contract.

## Component roadmap

- [x] Button
- [x] Input (+ InputGroup, Fieldset, FieldGroup, Field, Label, Description, ErrorMessage)
- [x] Select (package `selectbox` — `select` is a Go reserved keyword)
- [x] Textarea
- [x] Checkbox
- [x] Badge
- [x] Table, Heading, Card, Alert
- [x] Modal, Dropdown, Tabs, Toast
- [x] Tooltip, Accordion, Slide-over, Copy-to-Clipboard, Popover, Pagination, Breadcrumbs
- [x] Combobox, Command Palette
- [ ] DatePicker

## License

TBD.
