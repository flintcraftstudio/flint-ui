# flint-ui

Server-rendered UI components for FlintCraft Studio client projects. Built with Go + [templ](https://templ.guide) + htmx + Alpine.js + Tailwind CSS v4, converted from the Tailwind [Catalyst UI Kit](https://tailwindcss.com/plus/ui-kit).

> **Status**: early development. See [`flintcraft-ui-conversion-guide.md`](./flintcraft-ui-conversion-guide.md) for the architectural plan and component roadmap.

## For coding agents

flint-ui exists to speed up internal dashboard development. Start at [`AGENTS.md`](./AGENTS.md): it says where to look and the rules that save rework. Then use the component catalog below to pick a component, open its markdown reference for props, variants, and examples, and check [`docs/patterns.md`](./docs/patterns.md) for whole screens (list page, forms with validation, modals, toasts, loading states). From a client project, `go doc github.com/flintcraftstudio/flint-ui` prints the catalog grouped by purpose, and `go doc` on any component package prints its overview.

- **Invocation pattern**: `@<package>.<Func>(<package>.Props{...}) { children }` inside a `.templ` file.
- **Starting a screen**: wrap pages in a shell from [`layout`](./docs/components/layout.md) (with [`sidebar`](./docs/components/sidebar.md) / [`navbar`](./docs/components/navbar.md)), open each page with [`pageheader`](./docs/components/page-header.md), and render [`emptystate`](./docs/components/empty-state.md) when a list comes back empty.
- **Per-component docs**: every component has a reference in [`docs/components/`](./docs/components/) — open it before using the component. Each doc has: import, usage examples, a full props table, variants, and accessibility notes.
- **Semantic tokens**: components never hard-code colors. Rebrand via the CSS variables in [`styles/flint-ui.css`](./styles/flint-ui.css) (see [Tailwind](#tailwind) below).
- **htmx**: targets **htmx 4** (see [htmx 4 notes](#htmx-4)). Every component accepts `Attrs templ.Attributes` for `hx-*` / `data-*` / `aria-*` pass-through on the root element.
- **Alpine**: only for local UI state (open/closed, selected tab). Never application data. Components register their state from `js/flint-ui.js` (`flintui.Script()`; load it before Alpine), so they can run under Alpine's CSP build and a strict Content-Security-Policy. Converted so far: toast, slideover and modal; see AGENTS.md.
- **Showcase**: run `mage Showcase` to see every component live on `:8080`.

## Install

```sh
go get github.com/flintcraftstudio/flint-ui@latest
```

Import components one family at a time:

```go
import (
    "github.com/flintcraftstudio/flint-ui/components/button"
    "github.com/flintcraftstudio/flint-ui/components/input"
    "github.com/flintcraftstudio/flint-ui/components/fieldset"
)
```

In your `templ` files:

```go
@button.Button(button.Props{Variant: button.VariantAccent}) {
    Save changes
}
```

## Component catalog

All components live under `github.com/flintcraftstudio/flint-ui/components/<package>`. `Main templ functions` lists the exported entry points you invoke with `@package.Func(...)`. Click through to the doc for the full prop surface.

| Component | Package | Main templ functions | Purpose | Docs |
| --- | --- | --- | --- | --- |
| Accordion | `accordion` | `Accordion`, `Item`, `Trigger`, `Panel` | Disclosure panels — header reveals/hides paired content. | [accordion.md](./docs/components/accordion.md) |
| Alert | `alert` | `Alert`, `Title`, `Description`, `Actions` | Inline status banner (info / success / warning / danger). | [alert.md](./docs/components/alert.md) |
| Avatar | `avatar` | `Avatar`, `AvatarButton` | User image or initials, optionally wrapped in a button with touch target. | [avatar.md](./docs/components/avatar.md) |
| Badge | `badge` | `Badge` | Compact semantic status label. | [badge.md](./docs/components/badge.md) |
| Breadcrumbs | `breadcrumbs` | `Breadcrumbs`, `Item`, `Current` | Navigation trail for nested hierarchies. | [breadcrumbs.md](./docs/components/breadcrumbs.md) |
| Button | `button` | `Button` | Solid / outline / plain button; renders `<a>` when `Href` is set. `Indicator` shows a spinner during htmx requests. | [button.md](./docs/components/button.md) |
| Card | `card` | `Card`, `Header`, `Body`, `Footer` | Surface container for grouping related content. | [card.md](./docs/components/card.md) |
| Checkbox | `checkbox` | `Checkbox`, `CheckboxGroup`, `CheckboxField` | Native checkbox control with styled sibling box, plus stacked group. | [checkbox.md](./docs/components/checkbox.md) |
| Clipboard | `clipboard` | `Copy` | Alpine wrapper that copies a value to the system clipboard. | [clipboard.md](./docs/components/clipboard.md) |
| Combobox | `combobox` | `Combobox`, `Option` | Searchable single-select — text input with filterable dropdown. | [combobox.md](./docs/components/combobox.md) |
| Command Palette | `command` | `Palette`, `Group`, `Item` | Cmd+K / Ctrl+K overlay for searching and invoking actions. | [command.md](./docs/components/command.md) |
| Description List | `descriptionlist` | `DescriptionList`, `Term`, `Details` | Semantic `<dl>` / `<dt>` / `<dd>` two-column key/value display. | [description-list.md](./docs/components/description-list.md) |
| Divider | `divider` | `Divider` | Semantic `<hr>` with regular or soft variants. | [divider.md](./docs/components/divider.md) |
| Dropdown | `dropdown` | `Dropdown`, `Button`, `Menu`, `Item`, `Header`, `Section`, `Heading`, `Divider`, `Label`, `Description`, `Shortcut` | Alpine-powered menu with click-outside / ESC / Tab close and focus navigation. `TriggerAttrs` makes any button the trigger; `MenuProps.Fixed` escapes overflow containers. | [dropdown.md](./docs/components/dropdown.md) |
| Empty State | `emptystate` | `EmptyState`, `Title`, `Description`, `Actions` | Placeholder for an empty list, table, or panel; optional dashed border. | [empty-state.md](./docs/components/empty-state.md) |
| Fieldset | `fieldset` | `Fieldset`, `Legend`, `FieldGroup`, `Field`, `Label`, `Description`, `ErrorMessage` | Form layout primitives — `Field` auto-spaces label / description / control / error. | [fieldset.md](./docs/components/fieldset.md) |
| File Input | `fileinput` | `Input`, `Dropzone` | Native file picker styled like `Input`, and a drag-and-drop area with accept filtering and a file list. htmx uploads need `hx-encoding="multipart/form-data"`. | [file-input.md](./docs/components/file-input.md) |
| Heading | `heading` | `Heading`, `Subheading` | Page (`<h1>`) and section (`<h2>`) titles; `Level` decouples tag from visual size. | [heading.md](./docs/components/heading.md) |
| Input | `input` | `Input`, `InputGroup` | Text input, plus `InputGroup` for leading/trailing icons. | [input.md](./docs/components/input.md) |
| Layout | `layout` | `Sidebar`, `Stacked`, `Auth` | App shells: fixed sidebar rail, top navbar, or centered auth card. Collapse to a mobile drawer below `lg`. | [layout.md](./docs/components/layout.md) |
| Modal | `modal` | `Modal`, `Title`, `Description`, `Body`, `Actions` | Dialog + alert dialog in one component (`Props.Alert` switches layout). | [modal.md](./docs/components/modal.md) |
| Navbar | `navbar` | `Navbar`, `Section`, `Item`, `Label`, `Spacer`, `Divider` | Horizontal top nav for `layout.Stacked` and the mobile header in `layout.Sidebar`. | [navbar.md](./docs/components/navbar.md) |
| Page Header | `pageheader` | `PageHeader`, `Back`, `Content`, `Title`, `Description`, `Meta`, `Actions` | Page title block: back link, title, description or meta row, and primary actions. | [page-header.md](./docs/components/page-header.md) |
| Pagination | `pagination` | `Pagination`, `Previous`, `Next`, `List`, `Page`, `Gap` | Numbered page navigation for tables and lists. | [pagination.md](./docs/components/pagination.md) |
| Popover | `popover` | `Popover`, `Button`, `Panel` | Click-activated floating panel with arbitrary content. | [popover.md](./docs/components/popover.md) |
| Progress | `progress` | `Progress` | Determinate or looping indeterminate progress bar; polls with htmx `outerMorph` so the fill animates. | [progress.md](./docs/components/progress.md) |
| Radio | `radio` | `Radio`, `RadioGroup`, `RadioField` | Native `<input type="radio">` exclusive-choice control. | [radio.md](./docs/components/radio.md) |
| Select | `selectbox` | `Select` | Native `<select>` styled to match `Input`. Package is `selectbox` because `select` is a Go keyword. | [select.md](./docs/components/select.md) |
| Sidebar | `sidebar` | `Sidebar`, `Header`, `Body`, `Footer`, `Section`, `Heading`, `Item`, `Label`, `Spacer`, `Divider` | Vertical app navigation for `layout.Sidebar`; current-item bar and `aria-current`. | [sidebar.md](./docs/components/sidebar.md) |
| Slide-over | `slideover` | `Slideover`, `Title`, `Description`, `Body`, `Actions` | Edge-anchored side drawer sharing Modal's teleport + focus-trap plumbing. | [slideover.md](./docs/components/slideover.md) |
| Spinner | `spinner` | `Spinner` | Loading indicator; decorative or `role="status"` with a label; works as an `hx-indicator`. | [spinner.md](./docs/components/spinner.md) |
| Stat | `stat` | `Grid`, `Stat` | KPI tiles: label, value, and a change badge colored by sign (`LowerIsBetter` flips it). | [stat.md](./docs/components/stat.md) |
| Switch | `toggle` | `Switch`, `SwitchGroup`, `SwitchField` | On/off toggle — native `<input type="checkbox" role="switch">`. Package is `toggle`. | [switch.md](./docs/components/switch.md) |
| Tab bar | `tabbar` | `TabBar`, `Item`, `Spacer` | App navigation as tabs: a bar fixed to the bottom on phones, a narrow rail from `lg` up. Plain links, no JavaScript. `ContentClasses` keeps the page clear of it. | [tab-bar.md](./docs/components/tab-bar.md) |
| Table | `table` | `Table`, `Head`, `Body`, `Row`, `Header`, `Cell`, `SortHeader`, `Selectable`, `SelectAll`, `SelectCell`, `BulkActions` | Data table with sortable headers, row selection + bulk actions, and row-action menus. Layout flags (`bleed` / `dense` / `grid` / `striped`) go on each sub-component. | [table.md](./docs/components/table.md) |
| Tabs | `tabs` | `Tabs`, `List`, `Tab`, `Panel` | Tabbed navigation with two drivers — set `Mode` to `ModeServer` (htmx) or `ModeAlpine`. | [tabs.md](./docs/components/tabs.md) |
| Text | `text` | `Text`, `Link`, `Strong`, `Code` | Body-copy primitives: muted paragraph, underlined inline link, emphasis, inline code. | [text.md](./docs/components/text.md) |
| Textarea | `textarea` | `Textarea` | Multi-line input; resizable by default, `NonResizable` when layout owns sizing. | [textarea.md](./docs/components/textarea.md) |
| Toast | `toast` | `Container` | Ephemeral notifications. Render `Container` once in layout; fire via `flint:toast` window event (Alpine) or `HX-Trigger` header (htmx). | [toast.md](./docs/components/toast.md) |
| Tooltip | `tooltip` | `Tooltip` | Hover- and focus-activated hint bubble over a trigger. | [tooltip.md](./docs/components/tooltip.md) |

> **DatePicker** is deferred — `<input type="date">` covers most dashboard cases.

## Tailwind

flint-ui requires Tailwind CSS v4. Everything a client stylesheet needs ships in one importable partial, [`styles/flint-ui.css`](./styles/flint-ui.css):

- the semantic token contract as an `@theme` block (`primary`, `accent`, `danger`, `success`, `warning`, `foreground`, `muted`, `canvas`, `border`, `input`, `ring`, `scrim`, plus `-foreground` pairs and derived `-text` colors), with Catalyst-zinc defaults;
- dark values for every token, switched on per site with `data-theme` (see [Dark mode](#dark-mode)), and a matching `dark:` variant;
- `@custom-variant` aliases that map Headless UI's `data-hover`, `data-focus`, `data-active`, `data-disabled` and `data-invalid` state attributes to native pseudo-classes;
- an `@source` for flint-ui's own component templates (resolved relative to the partial, so importing it is enough for Tailwind to scan them);
- the few plain CSS rules the components rely on (`[x-cloak]`, toast animations).

Non-color classes (spacing, sizing, layering, transitions) are preserved verbatim from Catalyst.

### Client setup

1. **Expose the module directory to Tailwind.** Go keeps modules in the module cache, so link it to a stable path in the client project (add `.flint-ui` to `.gitignore`, and rerun after bumping the flint-ui version):
   ```sh
   ln -sfn "$(go list -m -f '{{.Dir}}' github.com/flintcraftstudio/flint-ui)" .flint-ui
   ```
   Projects that use `go mod vendor` can point at `vendor/github.com/flintcraftstudio/flint-ui` instead.
2. **Import the partial** after Tailwind in the client's entry stylesheet:
   ```css
   @import "tailwindcss";
   @import "../.flint-ui/styles/flint-ui.css";

   @source "../templates";            /* the client's own templates */
   ```
3. **Rebrand by overriding CSS variables**, not by editing components. Values are ordinary CSS colors (hex, `rgb()`, `oklch()`), not RGB triplets; opacity modifiers like `bg-primary/50` work automatically:
   ```css
   :root {
     --color-primary: #0a0a0a;            /* client brand primary */
     --color-primary-foreground: #ffffff;
     --color-accent: #dc2626;             /* client brand accent */
     --color-accent-foreground: #ffffff;
   }
   ```
   Every other token falls through to flint-ui's defaults.

No `tailwind.config.js` is needed.

### Point the client's coding agent at flint-ui

Agents working in a client repo don't know flint-ui exists unless the repo says so. Add this to the client's `CLAUDE.md` (or `AGENTS.md`):

```markdown
## UI components

UI is built with flint-ui (github.com/flintcraftstudio/flint-ui), linked at `.flint-ui/`.
Before writing any UI markup, read `.flint-ui/AGENTS.md`, then the component's
reference in `.flint-ui/docs/components/` and the recipes in `.flint-ui/docs/patterns.md`.
Use flint-ui components and semantic color tokens; don't hand-write equivalent markup
or raw Tailwind color classes. Run `templ generate` after editing `.templ` files.
```

If the project doesn't keep the `.flint-ui` link, use the module path from `go list -m -f '{{.Dir}}' github.com/flintcraftstudio/flint-ui` instead.

## Dark mode

flint-ui ships a dark palette (derived from Catalyst's dark mode: zinc-950 page, zinc-900 panels, white/10 borders), but it is **opt-in**. With no `data-theme` attribute a site renders light only, so existing client sites don't change.

**Turn it on** with one attribute on `<html>`:

```html
<html lang="en" data-theme="system">   <!-- follow the OS setting -->
<html lang="en" data-theme="dark">     <!-- always dark -->
<html lang="en" data-theme="light">    <!-- always light (same as no attribute) -->
```

Components need no changes: they draw from the tokens, and the tokens switch. Native controls (scrollbars, date pickers, `<select>` menus) follow too, because flint-ui sets `color-scheme` to match.

**Offer a toggle** by writing the attribute from a small inline script at the top of `<head>`, before the stylesheet, so the page never flashes the wrong theme:

```html
<script>
  (function () {
    var t = null;
    try { t = localStorage.getItem("flint-theme"); } catch (e) {}
    if (t === "light" || t === "dark" || t === "system") {
      document.documentElement.dataset.theme = t;
    }
  })();
</script>
```

The toggle itself sets `document.documentElement.dataset.theme` and saves the choice to `localStorage`. The showcase has a working Light / Dark / System switch (`themeToggle` in `examples/showcase/templates/layout.templ`). To crossfade instead of snapping, wrap the attribute change in `document.startViewTransition()` when it exists and the visitor hasn't asked for reduced motion; the showcase does this with a 300ms fade.

**Set dark brand colors.** In dark mode flint-ui's dark palette replaces every token, including a client's light brand overrides. That's deliberate: a dark-navy primary would otherwise become dark text on a dark page. To bring the brand into dark mode, add dark values after the import, **outside any `@layer`**:

```css
:root[data-theme] {
  @variant dark {
    --color-primary: #93c5fd;            /* brand primary, lightened */
    --color-primary-foreground: #172554;
    --color-accent: #2dd4bf;
    --color-accent-foreground: #042f2e;
  }
}
```

Two rules keep this working. Use the `:root[data-theme]` selector, which beats a plain `:root`. And stay out of `@layer`: flint-ui's dark block is unlayered, so a layered override loses to it. In dark mode, each hue works as a fill under its `-foreground` and as text on a dark surface, so pick lighter shades (roughly Tailwind -300 / -400) with a dark foreground. The `-text` tokens (`accent-text`, `danger-text`, `success-text`, `warning-text`) are mixed from the hue and `foreground`, so they follow your brand colors without an override.

**The `dark:` variant** follows the same rules (`dark:bg-…` applies under `data-theme="dark"`, and under `data-theme="system"` when the OS is dark). Prefer tokens. Reach for `dark:` only when one element needs different treatment in dark mode, as the switch thumb does.

**App shells** should use `lg:bg-canvas` (not `lg:bg-muted`) for the page behind the content panel, e.g. on `<html>` for `layout.Auth`. Canvas follows `muted` in light and drops to `background` in dark, so the panel stays lighter than the page.

**Contrast is tested.** `styles/contrast_test.go` parses `flint-ui.css` and checks every text/background pair the components render for WCAG AA, in both themes. If you change a default token value, run `go test ./styles/`.

## htmx 4

flint-ui targets [htmx 4](https://four.htmx.org/) (`htmx.org@4.0.0`). The differences from htmx 2 that matter when wiring components:

- **Inheritance is explicit.** `hx-target`, `hx-swap`, `hx-confirm` etc. no longer flow down from ancestors. Put them on the element that makes the request (that's what every flint-ui example does), or opt in with the `:inherited` suffix (`hx-target:inherited="#main"`).
- **Error responses swap.** 4xx/5xx bodies are swapped like 200s, so a handler can return `422` with the re-rendered form and its `fieldset.ErrorMessage`s. Only `204` and `304` skip the swap.
- **Event names are colon-separated.** `htmx:afterRequest` → `htmx:after:request`, so `hx-on::after-request` → `hx-on::after:request`. `hx-on` code runs with the event detail in scope; the request context is `ctx`.
- **`hx-disabled-elt` is now `hx-disable`** (and the old `hx-disable` is `hx-ignore`). Use `hx-disable="this"` to block double-submits.
- **`HX-Trigger` (response) is unchanged**, so the `flint:toast` / `flint:modal-*` / `flint:slideover-*` header patterns work as before. `HX-Trigger-After-Swap` and `-After-Settle` are gone.
- **`htmx-request` / `htmx-indicator` classes are unchanged**, so loading indicators still key off them.

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
AGENTS.md               entry point for coding agents
doc.go                  root package: go doc prints the catalog
components/
  button/               one package per component family
    button.templ
    classes.go
    doc.go              package overview for go doc
  input/                text inputs + InputGroup icon wrapper
  fieldset/             Fieldset, FieldGroup, Field, Label, Description, ErrorMessage
  shared/               shared prop types
styles/
  flint-ui.css          importable Tailwind v4 partial: @theme tokens + dark values + @custom-variant + @source
  contrast_test.go      WCAG AA contrast check for every token pair, light and dark
  flint.css             showcase entry: imports tailwindcss + flint-ui.css
examples/
  showcase/             reference site, similar to catalyst-demo
    main.go
    templates/          showcase pages, one per component
    static/
docs/
  components/           per-component markdown reference
  patterns.md           screen recipes composed from components
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
