# flint-ui

Server-rendered UI components for FlintCraft Studio client projects. Built with Go + [templ](https://templ.guide) + htmx + Alpine.js + Tailwind CSS v4, converted from the Tailwind [Catalyst UI Kit](https://tailwindcss.com/plus/ui-kit).

> **Status**: early development. See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the architectural plan and component roadmap.

## For coding agents

flint-ui exists to speed up internal dashboard development. The component catalog below is the entry point — use it to pick the right component, then open its markdown reference for props, variants, and examples.

- **Invocation pattern**: `@<package>.<Func>(<package>.Props{...}) { children }` inside a `.templ` file.
- **Per-component docs**: every component has a reference in [`docs/components/`](./docs/components/) — open it before using the component. Each doc has: import, usage examples, a full props table, variants, and accessibility notes.
- **Semantic tokens**: components never hard-code colors. Rebrand via the CSS variables in [`styles/flint.css`](./styles/flint.css) (see [Tailwind](#tailwind) below).
- **htmx**: every component accepts `Attrs templ.Attributes` for `hx-*` / `data-*` / `aria-*` pass-through on the root element.
- **Alpine**: only for local UI state (open/closed, selected tab). Never application data.
- **Showcase**: run `mage Showcase` to see every component live on `:8080`.

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

## Component catalog

All components live under `github.com/flintcraft/flint-ui/components/<package>`. `Main templ functions` lists the exported entry points you invoke with `@package.Func(...)`. Click through to the doc for the full prop surface.

| Component | Package | Main templ functions | Purpose | Docs |
| --- | --- | --- | --- | --- |
| Accordion | `accordion` | `Accordion`, `Item`, `Trigger`, `Panel` | Disclosure panels — header reveals/hides paired content. | [accordion.md](./docs/components/accordion.md) |
| Alert | `alert` | `Alert`, `Title`, `Description`, `Actions` | Inline status banner (info / success / warning / danger). | [alert.md](./docs/components/alert.md) |
| Avatar | `avatar` | `Avatar`, `AvatarButton` | User image or initials, optionally wrapped in a button with touch target. | [avatar.md](./docs/components/avatar.md) |
| Badge | `badge` | `Badge` | Compact semantic status label. | [badge.md](./docs/components/badge.md) |
| Breadcrumbs | `breadcrumbs` | `Breadcrumbs`, `Item`, `Current` | Navigation trail for nested hierarchies. | [breadcrumbs.md](./docs/components/breadcrumbs.md) |
| Button | `button` | `Button` | Solid / outline / plain button; renders `<a>` when `Href` is set. | [button.md](./docs/components/button.md) |
| Card | `card` | `Card`, `Header`, `Body`, `Footer` | Surface container for grouping related content. | [card.md](./docs/components/card.md) |
| Checkbox | `checkbox` | `Checkbox`, `CheckboxGroup`, `CheckboxField` | Native checkbox control with styled sibling box, plus stacked group. | [checkbox.md](./docs/components/checkbox.md) |
| Clipboard | `clipboard` | `Copy` | Alpine wrapper that copies a value to the system clipboard. | [clipboard.md](./docs/components/clipboard.md) |
| Combobox | `combobox` | `Combobox`, `Option` | Searchable single-select — text input with filterable dropdown. | [combobox.md](./docs/components/combobox.md) |
| Command Palette | `command` | `Palette`, `Group`, `Item` | Cmd+K / Ctrl+K overlay for searching and invoking actions. | [command.md](./docs/components/command.md) |
| Description List | `descriptionlist` | `DescriptionList`, `Term`, `Details` | Semantic `<dl>` / `<dt>` / `<dd>` two-column key/value display. | [description-list.md](./docs/components/description-list.md) |
| Divider | `divider` | `Divider` | Semantic `<hr>` with regular or soft variants. | [divider.md](./docs/components/divider.md) |
| Dropdown | `dropdown` | `Dropdown`, `Button`, `Menu`, `Item`, `Header`, `Section`, `Heading`, `Divider`, `Label`, `Description`, `Shortcut` | Alpine-powered menu with click-outside / ESC / Tab close and focus navigation. | [dropdown.md](./docs/components/dropdown.md) |
| Fieldset | `fieldset` | `Fieldset`, `Legend`, `FieldGroup`, `Field`, `Label`, `Description`, `ErrorMessage` | Form layout primitives — `Field` auto-spaces label / description / control / error. | [fieldset.md](./docs/components/fieldset.md) |
| Heading | `heading` | `Heading`, `Subheading` | Page (`<h1>`) and section (`<h2>`) titles; `Level` decouples tag from visual size. | [heading.md](./docs/components/heading.md) |
| Input | `input` | `Input`, `InputGroup` | Text input, plus `InputGroup` for leading/trailing icons. | [input.md](./docs/components/input.md) |
| Modal | `modal` | `Modal`, `Title`, `Description`, `Body`, `Actions` | Dialog + alert dialog in one component (`Props.Alert` switches layout). | [modal.md](./docs/components/modal.md) |
| Pagination | `pagination` | `Pagination`, `Previous`, `Next`, `List`, `Page`, `Gap` | Numbered page navigation for tables and lists. | [pagination.md](./docs/components/pagination.md) |
| Popover | `popover` | `Popover`, `Button`, `Panel` | Click-activated floating panel with arbitrary content. | [popover.md](./docs/components/popover.md) |
| Radio | `radio` | `Radio`, `RadioGroup`, `RadioField` | Native `<input type="radio">` exclusive-choice control. | [radio.md](./docs/components/radio.md) |
| Select | `selectbox` | `Select` | Native `<select>` styled to match `Input`. Package is `selectbox` because `select` is a Go keyword. | [select.md](./docs/components/select.md) |
| Slide-over | `slideover` | `Slideover`, `Title`, `Description`, `Body`, `Actions` | Edge-anchored side drawer sharing Modal's teleport + focus-trap plumbing. | [slideover.md](./docs/components/slideover.md) |
| Switch | `toggle` | `Switch`, `SwitchGroup`, `SwitchField` | On/off toggle — native `<input type="checkbox" role="switch">`. Package is `toggle`. | [switch.md](./docs/components/switch.md) |
| Table | `table` | `Table`, `Head`, `Body`, `Row`, `Header`, `Cell` | Data table; each sub-component carries the `bleed` / `dense` / `grid` / `striped` flags it needs. | [table.md](./docs/components/table.md) |
| Tabs | `tabs` | `Tabs`, `List`, `Tab`, `Panel` | Tabbed navigation with two drivers — set `Mode` to `ModeServer` (htmx) or `ModeAlpine`. | [tabs.md](./docs/components/tabs.md) |
| Textarea | `textarea` | `Textarea` | Multi-line input; resizable by default, `NonResizable` when layout owns sizing. | [textarea.md](./docs/components/textarea.md) |
| Toast | `toast` | `Container` | Ephemeral notifications. Render `Container` once in layout; fire via `flint:toast` window event (Alpine) or `HX-Trigger` header (htmx). | [toast.md](./docs/components/toast.md) |
| Tooltip | `tooltip` | `Tooltip` | Hover- and focus-activated hint bubble over a trigger. | [tooltip.md](./docs/components/tooltip.md) |

> **DatePicker** is deferred — `<input type="date">` covers most dashboard cases.

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

## License

TBD.
