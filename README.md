# flint-ui

Server-rendered UI components for FlintCraft Studio client projects. Built with Go + [templ](https://templ.guide) + htmx + Alpine.js + Tailwind CSS v4, converted from the Tailwind [Catalyst UI Kit](https://tailwindcss.com/plus/ui-kit).

> **Status**: early development. See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the architectural plan and component roadmap.

## Install

```sh
go get github.com/flintcraft/flint-ui@latest
```

Import components one family at a time:

```go
import "github.com/flintcraft/flint-ui/components/button"
```

In your `templ` files:

```go
@button.Button(button.Props{Color: button.ColorIndigo}) {
    Save changes
}
```

## Tailwind

flint-ui preserves Catalyst's Tailwind v4 class strings verbatim. Client projects must:

1. Use Tailwind CSS v4.
2. Scan flint-ui templates as content sources, e.g. in your CSS entry:
   ```css
   @source "../node_modules/github.com/flintcraft/flint-ui/components/**/*.templ";
   ```
3. Include the `@custom-variant` declarations that alias Headless UI's `data-hover`, `data-focus`, `data-active`, `data-disabled` state attributes to native pseudo-classes. Copy the block from [`styles/flint.css`](./styles/flint.css).

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
  shared/               shared prop types (Size, etc.)
styles/
  flint.css             Tailwind v4 entry: @theme + @custom-variant + @source
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
- **Catalyst's classes preserved verbatim.** Upstream updates can be diff-applied without losing custom edits.
- **Accessibility preserved.** ARIA attributes, focus management, and keyboard behavior track the Catalyst source.

See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the full rules.

## Component roadmap

- [x] Button
- [ ] Input
- [ ] Select
- [ ] Textarea
- [ ] Checkbox
- [ ] Badge
- [ ] Table, Card, Alert, Heading
- [ ] Modal, Dropdown, Tabs, Toast
- [ ] DatePicker, Combobox, Pagination, Breadcrumbs

## License

TBD.
