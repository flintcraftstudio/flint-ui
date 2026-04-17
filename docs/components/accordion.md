# Accordion

Disclosure panels — clickable header reveals/hides its paired content. From-scratch design (Catalyst has no `accordion.tsx`).

## Import

```go
import "github.com/flintcraft/flint-ui/components/accordion"
```

## Components

| Component                | Renders                                   | Notes                                                    |
| ------------------------ | ----------------------------------------- | -------------------------------------------------------- |
| `accordion.Accordion`    | `<div>` with `x-data`                     | Owns the Alpine state for the whole group.              |
| `accordion.Item`         | `<div>` with borders                     | One disclosure pair. Wraps a Trigger and a Panel.       |
| `accordion.Trigger`      | `<h3><button aria-expanded aria-controls>` | The clickable header. Chevron auto-appended and rotates on open. |
| `accordion.Panel`        | `<div role="region"><div>…</div></div>`  | Two-level: outer shell x-collapse animates; inner box holds padding/typography. |

## Quick example

```go
@accordion.Accordion(accordion.Props{}) {
    @accordion.Item(accordion.ItemProps{Name: "materials"}) {
        @accordion.Trigger(accordion.TriggerProps{Name: "materials"}) {
            What materials are included?
        }
        @accordion.Panel(accordion.PanelProps{Name: "materials"}) {
            Seamless aluminum, all hangers, end caps, outlets, and downspouts.
        }
    }
    @accordion.Item(accordion.ItemProps{Name: "warranty"}) {
        @accordion.Trigger(accordion.TriggerProps{Name: "warranty"}) {
            Is the work warrantied?
        }
        @accordion.Panel(accordion.PanelProps{Name: "warranty"}) {
            Ten-year workmanship warranty.
        }
    }
}
```

Each `Item` owns a stable `Name`; the matching `Trigger` and `Panel` repeat it. Same naming pattern as Tabs — templ has no React context, so sub-components coordinate via props.

## Multiple vs single

| Type                  | Behavior                                                       |
| --------------------- | -------------------------------------------------------------- |
| `TypeMultiple` (default) | Each panel toggles independently. Several can be open at once. |
| `TypeSingle`          | Opening a panel closes whichever was open before.              |

```go
@accordion.Accordion(accordion.Props{Type: accordion.TypeSingle}) {
    // ...
}
```

Pick `TypeSingle` when the content is reference-style (FAQ, mobile nav) and the reader should focus on one panel at a time. Pick `TypeMultiple` when panels are peers the reader may compare (settings groups, per-feature docs).

## Default open item

```go
@accordion.Accordion(accordion.Props{Default: "overview"}) {
    @accordion.Item(accordion.ItemProps{Name: "overview"}) { ... }
    // ...
}
```

`Default` seeds the initial active array. Only one item can be the Default in v0.1 (multiple-defaults for `TypeMultiple` would need a slice — add if a client needs it).

## Rich trigger content

`Trigger` children is any templ content — drop a Badge, an icon, a metric, whatever fits:

```go
@accordion.Trigger(accordion.TriggerProps{Name: "crew"}) {
    <span class="flex items-center gap-3">
        Crew assignments
        @badge.Badge(badge.Props{Variant: badge.VariantWarning}) { 2 unassigned }
    </span>
}
```

The chevron is appended automatically and rotates on open. Your content goes in column 1 (flex-1); the chevron sits in column 2 (shrink-0).

## Disabled items

```go
@accordion.Trigger(accordion.TriggerProps{Name: "archive", Disabled: true}) {
    Archive (coming soon)
}
```

Renders the button with native `disabled`. The Panel stays closed regardless of the root's active array; opacity + `pointer-events-none` match the Button component's disabled style.

## Heading level

Each Trigger sits inside a heading tag so the accordion fits the surrounding document outline. Default is `<h3>`; override via `TriggerProps.HeadingLevel` (1–6):

```go
@accordion.Trigger(accordion.TriggerProps{Name: "example", HeadingLevel: 2}) {
    // becomes <h2><button>...</button></h2>
}
```

The same convention as the `heading` package — semantic structure tracks the document, not the visual size.

## Collapse animation

The smooth height transition requires the [Alpine Collapse plugin](https://alpinejs.dev/plugins/collapse). Load it before Alpine core, same pattern as `@alpinejs/focus`:

```html
<script src="https://unpkg.com/@alpinejs/collapse@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

Without the plugin, Alpine silently ignores the `x-collapse` directive and panels snap open/closed via plain `x-show`. Progressive enhancement — same as Modal/Dropdown's relationship with `@alpinejs/focus`.

The Panel is a two-level structure: an outer shell the plugin animates (no padding or border on it — padding on a collapsing element causes a jumpy look) and an inner content box holding the padding, text, and caller-supplied `Class`. Duration is explicitly 200ms on both Panel (`x-collapse.duration.200ms`) and Trigger chevron (`transition-transform duration-200 ease-in-out`) so the two motions read as one.

## Keyboard

| Key               | Behavior                            |
| ----------------- | ----------------------------------- |
| `Tab`             | Focus the next trigger.             |
| `Shift+Tab`       | Focus the previous trigger.         |
| `Space` / `Enter` | Toggle the focused trigger (native button). |

No arrow-key navigation in v0.1 — the W3C accordion pattern suggests it but WCAG doesn't require it. Add via `@accordion.Accordion` Attrs if your dashboard needs it; Tabs' `moveTab` helper is a reasonable template.

## Accessibility

- Each Trigger is a native `<button>` with `aria-expanded` bound to its Alpine state and `aria-controls` pointing at the Panel.
- Each Panel carries `role="region"`, `id={panelID}`, and `aria-labelledby={triggerID}` so screen readers announce the Trigger's label when the region gains focus.
- The Trigger sits inside a heading tag (`<h3>` default, overridable per-Item) so the accordion preserves document outline.
- Disabled triggers use native `disabled`; focus skips them and clicks are blocked.
- `x-cloak` on each Panel prevents a FOUC while Alpine initializes.

## Why Alpine-only

Accordion has no ModeServer variant. Which panels are expanded is local, ephemeral UI state — there's no correctness reason for the server to own it. If you need server-driven reveal of content (e.g., "only show the approvals panel when pending approvals exist"), reach for [Tabs](./tabs.md) in ModeServer instead.

## Styling

| Element   | Default classes                                                                  | Override via             |
| --------- | -------------------------------------------------------------------------------- | ------------------------ |
| Root      | `flex flex-col`                                                                  | `Props.Class`            |
| Item      | `border-b border-border first:border-t`                                          | `ItemProps.Class`        |
| Trigger   | `flex w-full items-center justify-between px-4 py-3 text-base font-medium hover:bg-muted` | `TriggerProps.Class`    |
| Panel     | `px-4 pb-4 pt-1 text-sm text-muted-foreground`                                   | `PanelProps.Class`       |

When the accordion sits inside a card that already has borders, pass `ItemProps.Class: "border-t-0 first:border-t-0"` or similar to avoid a double border at the edges.

## Required surrounding setup

- **`x-data` on a common ancestor.** The accordion root's `x-data` needs to be inside a live Alpine scope. The showcase puts `x-data` on `<body>`; clients must do the same in their base layout. (Same requirement as Modal/Dropdown/Toast.)
- **`@alpinejs/collapse` loaded before Alpine core** if you want smooth height animations. Optional.

## Props reference

### Props (Accordion root)

| Field      | Type               | Default        | Notes                                                        |
| ---------- | ------------------ | -------------- | ------------------------------------------------------------ |
| `Type`     | `Type`             | `TypeMultiple` | `TypeSingle` closes the previous panel when a new one opens. |
| `Default`  | `string`           | `""`           | Name of the Item that starts open. Empty means nothing open. |
| `Class`    | `string`           | `""`           | Appended to the root `<div>`.                                |
| `Attrs`    | `templ.Attributes` | `nil`          | Spread onto the root.                                        |

### ItemProps

| Field      | Type               | Default | Notes                                                 |
| ---------- | ------------------ | ------- | ----------------------------------------------------- |
| `Name`     | `string`           | `""`    | Required. Stable identifier; URL/id-safe characters.  |
| `Class`    | `string`           | `""`    | Appended to the Item `<div>`.                         |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the Item.                                 |

### TriggerProps

| Field          | Type               | Default | Notes                                                       |
| -------------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Name`         | `string`           | `""`    | Required. Must mirror the parent Item's Name.               |
| `HeadingLevel` | `int`              | `3`     | 1–6. Picks the `<h1>`–`<h6>` tag wrapping the button.       |
| `Disabled`     | `bool`             | `false` | Native disabled on the button; panel stays closed.          |
| `Class`        | `string`           | `""`    | Appended to the button.                                     |
| `Attrs`        | `templ.Attributes` | `nil`   | Spread onto the button.                                     |

### PanelProps

| Field      | Type               | Default | Notes                                                 |
| ---------- | ------------------ | ------- | ----------------------------------------------------- |
| `Name`     | `string`           | `""`    | Required. Must mirror the parent Item's Name.         |
| `Class`    | `string`           | `""`    | Appended to the **inner** content box (padding, typography, background). |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the **outer** shell (id, role, aria-labelledby already present). |
