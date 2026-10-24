# flint-ui for coding agents

flint-ui is the component library for FlintCraft client admin tools: Go + templ + htmx 4 + Alpine.js + Tailwind CSS v4, converted from Tailwind's Catalyst UI kit. If you're building UI in a project that imports `github.com/flintcraftstudio/flint-ui`, use these components instead of writing markup and color classes by hand.

## Where to look, in order

1. **Pick a component** from the catalog: the table in [README.md](./README.md#component-catalog), or `go doc github.com/flintcraftstudio/flint-ui` for the same list grouped by purpose.
2. **Read its reference** in [`docs/components/<name>.md`](./docs/components/) before using it: usage, every prop, variants, accessibility, htmx wiring. `go doc github.com/flintcraftstudio/flint-ui/components/<pkg>` shows the Go API.
3. **Check [docs/patterns.md](./docs/patterns.md)** when the job is a whole screen or flow: page skeleton, list page with search/sort/bulk actions, form with server-side validation, form in a modal, confirm-delete, server-driven toasts, loading states.
4. **Copy from the showcase** when you want working code: `examples/showcase/templates/` has a page per component (`buttons.templ`, `tables_admin.templ`, `appshell.templ`, …) and `examples/showcase/main.go` has the handlers (validation, `HX-Trigger` toasts, polling).

In a client project the module is usually linked at `.flint-ui/` (see the README's client setup); otherwise find it with `go list -m -f '{{.Dir}}' github.com/flintcraftstudio/flint-ui`.

## Rules that save rework

- **Colors come from semantic tokens only** (`bg-surface`, `text-muted-foreground`, `bg-primary`, …, listed in `styles/flint-ui.css`). Never `bg-zinc-900` or hex. Rebrand by overriding the `--color-*` variables in the client stylesheet.
- **Every component takes `Props` with `Class` and `Attrs`.** `Attrs` (`templ.Attributes`) carries `hx-*`, `aria-*` and `data-*` to the root element.
- **htmx 4, not 2:** no attribute inheritance (put `hx-target`/`hx-swap` on the requesting element), 4xx/5xx responses swap (return `422` with the re-rendered form), events are colon-separated (`hx-on::after:request`), `hx-disable` replaces `hx-disabled-elt`. Details in the README's htmx 4 section.
- **Alpine is for local UI state only** (open/closed, selected tab). Modals, slide-overs and toasts are driven by window events (`flint:modal-open`, `flint:toast`, …), which a server can also fire through an `HX-Trigger` header.
- **Pages need** `<body x-data>`, flint-ui's script (`flintui.Script()`, served as a file and loaded before Alpine), Alpine's Focus plugin (and Collapse for accordion), and one `toast.Container` in the base layout.
- **Alpine state lives in `js/flint-ui.js`, never inline.** A component with state or methods registers with `Alpine.data("flintName", …)` there, and its markup says `x-data="flintName"`; markup keeps only simple expressions (calls, assignments, comparisons). Alpine's CSP build parses expressions, never function bodies, so an inline `x-data="{ add() { … } }"` silently fails to start under a strict CSP. Each converted component has a test using `internal/csptest`, which loads it in headless Chromium under a strict CSP and fails on any console error. Converted so far: **toast**, **slideover**, **modal**. The others still have inline state; convert one, with its csptest, before using it under a strict CSP.
- **Dark mode is opt-in** per site with `data-theme` on `<html>`; components need no changes. See the README's Dark mode section.
- **Run `templ generate`** after editing `.templ` files.

## Working on flint-ui itself

- One package per component under `components/<pkg>/`: a `.templ` file, `classes.go` with the class strings, and `doc.go` with the package overview. Non-color classes track the Catalyst source verbatim; colors map to tokens.
- A new component needs: its package (with a `doc.go`), `docs/components/<name>.md`, a row in the README catalog, an entry in the root `doc.go` catalog, and a showcase page plus its nav entry. `go test ./...` fails if the README row, the root catalog, the `doc.go` or the markdown reference is missing.
- `go test ./...` also checks every token pair for WCAG AA in light and dark (`styles/contrast_test.go`).
- `mage Showcase` runs the reference site on `:8080`. Commit generated `_templ.go` files.
- The conversion rules and token contract are in [flintcraft-ui-conversion-guide.md](./flintcraft-ui-conversion-guide.md).
