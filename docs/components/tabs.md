# Tabs

From-scratch design (Catalyst has no `tabs.tsx`). Two interaction modes share the same components, props, and visuals. Pick the mode based on **what's behind the content**:

- **`ModeAlpine`** (default) — every panel is server-rendered up front; Alpine's `x-show` toggles which one is visible. No roundtrip. Use when each tab's content is static or already cheap to render.
- **`ModeServer`** — the tablist and the active panel are server-rendered; each Tab swaps the entire tabs section via htmx. Use when content depends on server state that other tabs' actions can mutate (a form in tab A changes what tab B should show) — the server stays the source of truth.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/tabs"
```

## Components

| Component       | Renders                                | Notes                                                                              |
| --------------- | -------------------------------------- | ---------------------------------------------------------------------------------- |
| `tabs.Tabs`     | `<div>` (with `x-data` in ModeAlpine)  | Outer container. Owns active state in ModeAlpine.                                  |
| `tabs.List`     | `<div role="tablist">`                 | Wraps the row of Tabs. Adds Arrow / Home / End handlers in ModeAlpine.             |
| `tabs.Tab`      | `<button role="tab">`                  | Single tab. Sub-components carry `Mode` and `Active` matching the parent.          |
| `tabs.Panel`    | `<div role="tabpanel">`                | One per tab in ModeAlpine. Not used in ModeServer (caller renders active panel).   |

All four sub-components carry `Mode` (and `Active` where relevant) that **must mirror the parent Tabs** — same convention as `modal.SubProps.Alert`. Verbose, but explicit. Real client projects usually wrap this in a per-page helper template (see "Sub-component prop convention" below).

## ModeAlpine — quick example

```go
@tabs.Tabs(tabs.Props{Mode: tabs.ModeAlpine, Active: "overview"}) {
    @tabs.List(tabs.ListProps{Mode: tabs.ModeAlpine}) {
        @tabs.Tab(tabs.TabProps{Name: "overview", Mode: tabs.ModeAlpine, Active: "overview"}) { Overview }
        @tabs.Tab(tabs.TabProps{Name: "schedule", Mode: tabs.ModeAlpine, Active: "overview"}) { Schedule }
        @tabs.Tab(tabs.TabProps{Name: "invoices", Mode: tabs.ModeAlpine, Active: "overview"}) { Invoices }
    }
    @tabs.Panel(tabs.PanelProps{Name: "overview", Active: "overview"}) { ... }
    @tabs.Panel(tabs.PanelProps{Name: "schedule", Active: "overview"}) { ... }
    @tabs.Panel(tabs.PanelProps{Name: "invoices", Active: "overview"}) { ... }
}
```

`Active` is the initial active tab. Alpine takes over after init: clicks set `active = name`, panels' `x-show` re-evaluates, the tab pill updates via `aria-selected`.

### FOUC handling

Two trapdoors handled for you:

1. **Tab pill flash.** Tabs render with static `aria-selected` based on `Active` so the active tab is highlighted on first paint, before Alpine binds. After init, `x-bind:aria-selected` reactively takes over.
2. **All panels visible flash.** Panels whose `Name != Active` render with `x-cloak`, hidden by the `[x-cloak]{display:none!important}` rule already in `styles/flint.css`. After init, Alpine removes `x-cloak` and `x-show` keeps inactive panels hidden via `display:none`. The active panel renders without `x-cloak` so it's visible from the first byte.

This is why `Panel.Active` and `Tab.Active` exist as props — they let the server render correct initial visibility.

## ModeServer — quick example

```go
templ JobTabs(active string) {
    @tabs.Tabs(tabs.Props{
        Mode:   tabs.ModeServer,
        Active: active,
        Attrs:  templ.Attributes{"id": "job-tabs"},
    }) {
        @tabs.List(tabs.ListProps{Mode: tabs.ModeServer}) {
            @tab("overview", active, "Overview")
            @tab("crew",     active, "Crew")
            @tab("billing",  active, "Billing")
        }
        <div class="mt-4">
            @activePanel(active) // server-rendered content for the active tab
        </div>
    }
}

templ tab(name, active, label string) {
    @tabs.Tab(tabs.TabProps{
        Name: name, Mode: tabs.ModeServer, Active: active,
        Attrs: templ.Attributes{
            "hx-get":    "/jobs/1834/tabs?tab=" + name,
            "hx-target": "#job-tabs",
            "hx-swap":   "outerHTML",
        },
    }) { { label } }
}
```

Server handler:

```go
func handleJobTabs(w http.ResponseWriter, r *http.Request) {
    tab := r.URL.Query().Get("tab")
    // ... validate, fetch fresh data ...
    templates.JobTabs(tab).Render(r.Context(), w)
}
```

Each tab click hits the server, which re-renders the entire tabs section with the new `Active`. htmx `outerHTML`-swaps it into the same `#job-tabs` container. The new HTML's static `aria-selected` highlights the right tab. No client state to manage.

### Why swap the whole section, not just the panel?

It's the simplest correct behavior. Swapping only the panel leaves the active-pill state stale until the next page load (or requires a client-side update via `hx-on::after-request`). Swapping the whole section makes the server the source of truth for both content *and* selection — one render, one swap, consistent.

If you need partial swaps (e.g. very large tablists), you can do it manually — wire `hx-target` to a panel-only container and write your own `hx-on::after-request` that updates the tab pill. The component doesn't fight you, but the default pattern is whole-section swap.

## Sub-component prop convention

Every sub-component carries `Mode` and (where the static initial state matters) `Active`. They must match the parent. This is the same convention `modal.SubProps.Alert` uses.

Two reasons:

1. **No implicit context.** templ has no React-style context, so a child component can't read the parent's props at runtime. Passing them explicitly is the only correct option without a runtime hack.
2. **Server-rendered correctness.** `Tab.Active` and `Panel.Active` aren't just decoration — they drive the static `aria-selected` and `x-cloak` attrs that prevent FOUC. The component literally needs them.

The verbosity is real. The standard mitigation is a tiny per-page wrapper:

```go
templ jobTab(name, label string) {
    @tabs.Tab(tabs.TabProps{Name: name, Mode: tabs.ModeAlpine, Active: g_active}) { { label } }
}
```

…or pass `active` through as a function arg. Each project gets to pick the shape that fits.

## Keyboard navigation

ModeAlpine wires the [WAI-ARIA Tabs Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/tabs/) keyboard model:

| Key                       | Behavior                                     |
| ------------------------- | -------------------------------------------- |
| `ArrowRight`              | Focus and activate the next non-disabled tab. Wraps. |
| `ArrowLeft`               | Focus and activate the previous non-disabled tab. Wraps. |
| `Home`                    | Focus and activate the first non-disabled tab. |
| `End`                     | Focus and activate the last non-disabled tab. |
| `Tab` (browser)           | Move focus into / out of the tablist (roving tabindex puts only the active tab in the tab order). |
| `Enter` / `Space`         | Native button activation. Same effect as clicking. |

These are wired in `tabs.List` via Alpine `x-on:keydown` directives that call a `moveTab(dir)` method on the parent x-data. Disabled tabs (native `disabled` attribute) are skipped.

ModeServer **does not** include arrow-key roving in v0.1. Native `Tab` key moves focus between tab buttons; `Enter` activates the focused tab (which fires its `hx-get`). Adding arrow-roving to ModeServer is a future change — it's harder there because each move would trigger a server roundtrip; we want to think through debouncing first.

## Required surrounding setup

Same baseline as Modal and Dropdown:

1. **`x-data` on a common ancestor** in your base layout (typically `<body x-data>`), so Alpine processes the directives.
2. **Alpine and (for moving focus) the Focus plugin loaded in the right order**:

```html
<script src="https://unpkg.com/@alpinejs/focus@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

Tabs uses native `element.focus()` rather than the Focus plugin's `$focus.next()`, but the rest of flint-ui assumes the Focus plugin is present, so loading it costs nothing extra.

## ARIA

- `tabs.Tabs` root: no role (just a `<div>`).
- `tabs.List`: `role="tablist"`.
- `tabs.Tab`: `role="tab"`, `aria-selected`, `aria-controls={panel-id}`, roving `tabindex` (`0` for active, `-1` for others).
- `tabs.Panel`: `role="tabpanel"`, `aria-labelledby={tab-id}`, `tabindex="0"` (panel itself is focusable so screen readers can read the content).
- IDs are auto-derived from `Name`: `flint-tab-X` and `flint-tabpanel-X`. Stable per tab; safe characters only (`[A-Za-z0-9_-:.]`).

## Visual style

Underline-active — the standard dashboard pattern. Inactive tabs are `text-muted-foreground`; the active one gets `border-b-2 border-primary text-foreground`. Hover on inactive tabs previews with a `border-border` underline.

The active state is driven entirely by `aria-selected`, so the *same* class string serves both modes — ModeAlpine's `x-bind:aria-selected` and ModeServer's static `aria-selected` both flip Tailwind's `aria-selected:` modifier.

Override per-tab via `Class` on `tabs.TabProps` if you need a different look for one tab. A `Variant` (pill style, etc.) can be added later if a client project needs it — phase-3 v0.1 ships underline only.

## Props reference

### Props (Tabs root)

| Field   | Type               | Default | Notes                                                                                                                |
| ------- | ------------------ | ------- | -------------------------------------------------------------------------------------------------------------------- |
| `Mode`  | `Mode`             | `""` (= ModeAlpine) | `ModeAlpine` for client-side switching, `ModeServer` for htmx-swapped tabs.                                          |
| `Active`| `string`           | —       | Required. ModeAlpine: initial active tab. ModeServer: currently-active tab (drives static `aria-selected`).          |
| `Class` | `string`           | `""`    | Appended to the root `<div>`.                                                                                        |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root. Use for `id` (so htmx can target the section), `data-*`, etc.                                  |

### ListProps

| Field   | Type               | Default | Notes                                                                            |
| ------- | ------------------ | ------- | -------------------------------------------------------------------------------- |
| `Mode`  | `Mode`             | `""`    | Must mirror parent. ModeAlpine adds Arrow / Home / End keyboard handlers.        |
| `Class` | `string`           | `""`    | Appended to the tablist `<div>`.                                                 |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the tablist.                                                         |

### TabProps

| Field      | Type               | Default | Notes                                                                                          |
| ---------- | ------------------ | ------- | ---------------------------------------------------------------------------------------------- |
| `Name`     | `string`           | —       | Required. Stable identifier; matches `Panel.Name`. Restricted to `[A-Za-z0-9_-:.]`.            |
| `Mode`     | `Mode`             | `""`    | Must mirror parent.                                                                            |
| `Active`   | `string`           | `""`    | Parent's Active. Drives static `aria-selected` / `tabindex` for FOUC-free first paint.         |
| `Disabled` | `bool`             | `false` | Native `disabled`. Skipped by keyboard nav.                                                    |
| `Class`    | `string`           | `""`    | Appended to the tab.                                                                           |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the tab. ModeServer: add `hx-*` here for the swap.                                 |

### PanelProps (ModeAlpine only)

| Field    | Type               | Default | Notes                                                                                |
| -------- | ------------------ | ------- | ------------------------------------------------------------------------------------ |
| `Name`   | `string`           | —       | Required. Matches a `Tab.Name`.                                                      |
| `Active` | `string`           | `""`    | Parent's Active. Panels whose Name != Active render with `x-cloak` (FOUC prevention). |
| `Class`  | `string`           | `""`    | Appended to the panel.                                                               |
| `Attrs`  | `templ.Attributes` | `nil`   | Spread onto the panel.                                                               |
