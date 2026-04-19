# Combobox

Searchable select — a text input with a filterable dropdown of options. Reshaped from `catalyst-ui-kit/typescript/combobox.tsx` (Catalyst's Headless UI + virtual rendering approach doesn't translate to server-side templ; the API here is smaller and flat).

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/combobox"
```

## Components

| Component              | Renders                                | Notes                                                     |
| ---------------------- | -------------------------------------- | --------------------------------------------------------- |
| `combobox.Combobox`    | `<div data-slot="control">` with input + listbox | Owns Alpine state; renders the visible text input, the hidden value input, a chevron toggle button, and the floating `<ul role="listbox">`. |
| `combobox.Option`      | `<li role="option">`                   | Single option row. Label drives filter + input display; Description is optional secondary text. |

## Quick example

```go
@fieldset.Field(fieldset.Props{}) {
    @fieldset.Label(fieldset.LabelProps{For: "client-picker"}) { Client }
    @combobox.Combobox(combobox.Props{
        Name:        "client",
        ID:          "client-picker",
        Placeholder: "Search clients…",
    }) {
        @combobox.Option(combobox.OptionProps{Value: "1", Label: "Henderson Ranch"})
        @combobox.Option(combobox.OptionProps{Value: "2", Label: "Western Skies Contracting"})
        @combobox.Option(combobox.OptionProps{Value: "3", Label: "Lo Mo Outfitting"})
    }
}
```

Form submit sends `client=1` (or whatever Value the user picked). The visible text input is `display-only` — not submitted.

## How it works

One Alpine scope on the root holds:

```js
{
    open: false,            // listbox visibility
    query: '',              // text in the input; drives filtering
    selected: '',           // the chosen Option's Value (hidden input binds to this)
    selectedLabel: ''       // the chosen Option's Label (for re-display + match logic)
}
```

Each Option carries `data-value` and `data-label` HTML attributes. Click handlers read them via `$el.dataset.*` so the Alpine expression itself is a constant string — option values and labels can contain any character (apostrophes, backslashes, angle brackets) without JS-escape concerns.

Filter matching: each Option's `x-show` calls `matches(label)`. When `query` is empty OR equals `selectedLabel`, all options show (the user has picked something and is browsing, not searching). When `query` diverges, a case-insensitive substring check filters.

## Keyboard

Requires [`@alpinejs/focus`](https://alpinejs.dev/plugins/focus) — already loaded by the showcase layout; clients must mirror it.

| Key                          | Behavior                                           |
| ---------------------------- | -------------------------------------------------- |
| Type in input                | Open listbox + filter options by substring.        |
| `ArrowDown` (from input)     | Open listbox + focus first visible option.         |
| `ArrowDown` / `ArrowUp` (in options) | Move focus through visible options; wraps. |
| `Enter` (on input)           | Select the first visible option.                   |
| `Enter` (on option)          | Select it.                                         |
| `Escape`                     | Close listbox; restore input to the selected label. |
| `Tab`                        | Close listbox; release focus naturally.            |

## Selected value is sticky

Typing a query that doesn't match the current selection does NOT clear `selected`. The hidden input keeps the last explicitly-picked value until the user picks another.

This matches how most production comboboxes behave (Slack channel picker, Gmail label picker) — the submit value is what the user *chose*, not what's currently in the filter box. If a client needs an explicit "Clear" behavior, add a small reset button:

```go
@button.Button(button.Props{
    Plain: true,
    Attrs: templ.Attributes{
        "x-on:click": "selected = ''; selectedLabel = ''; query = ''",
    },
}) { Clear }
```

## Rich option content

`OptionProps.Description` adds a secondary line under the Label in the dropdown — useful for disambiguating ("Marcus Lee / marcus@flintcraft.studio" vs "Marcus Lee / marcus@hendersonranch.com"):

```go
@combobox.Option(combobox.OptionProps{
    Value:       "marcus",
    Label:       "Marcus Lee",
    Description: "Installer · marcus@flintcraft.studio",
})
```

Description doesn't appear in the input after selection — only the Label does. If you need richer option content (icons, badges), wait for v0.2 or override via the Class prop.

## Disabled options

```go
@combobox.Option(combobox.OptionProps{
    Value: "enterprise",
    Label: "Enterprise",
    Description: "Call for pricing",
    Disabled: true,
})
```

Dimmed visually; filter still matches them (so users see they exist). Click and Enter are no-ops on disabled options; `ArrowDown` / `ArrowUp` skip them.

## Long lists

The listbox caps at `max-h-64` and scrolls internally. Combobox is tuned for client-side filtering of small-to-medium lists — up to a few hundred options renders and filters snappily. Beyond that, consider adding a `ModeServer` variant (following [Tabs](./tabs.md)'s pattern) that delegates filtering to the server over htmx. Not shipped in v0.1 because no client project has needed it yet.

## Accessibility

- The visible input has `role="combobox"`, `aria-autocomplete="list"`, `aria-expanded` bound to Alpine state, and `aria-controls` pointing at the listbox.
- The listbox has `role="listbox"` and a stable `id`.
- Each option has `role="option"`, `aria-selected` bound to whether its Value matches `selected`, and `aria-disabled="true"` when disabled.
- The chevron toggle has `tabindex="-1"` (out of tab order) and `aria-label="Toggle options"` — it's a shortcut for users who prefer clicking over typing, not a separate focus stop.

## Form integration

Two inputs render inside the wrapper:

- The visible text input (no `name`) — user types here; the current value is `query`, which is display-only.
- A hidden input (`type="hidden"`, `name={Props.Name}`) — its value is bound to `selected`, so the form submit carries the chosen option's Value.

If you omit `Name`, the hidden input isn't rendered (useful for pure filter UIs that don't submit).

## Styling

Chrome mirrors the Input component (same `before:inset-px` focus ring, same `data-slot="control"` for Fieldset auto-spacing) so Combobox and Input look like siblings in a form. The listbox mirrors Dropdown's Menu (`bg-surface/90 backdrop-blur-xl`, `shadow-lg`, `ring-1 ring-border`) so popovers, dropdowns, and comboboxes share a visual vocabulary.

## Required surrounding setup

- **`x-data` on a common ancestor.** Same baseline as Modal / Dropdown / Popover — put `x-data` on `<body>` in your base layout.
- **`@alpinejs/focus` loaded before Alpine core** — otherwise `$focus.within(...)` is undefined and keyboard nav silently no-ops.

## Props reference

### Props (Combobox root)

| Field         | Type               | Default                | Notes                                                              |
| ------------- | ------------------ | ---------------------- | ------------------------------------------------------------------ |
| `Name`        | `string`           | `""`                   | Name of the hidden input carrying the selected value. Omit for pure-filter UIs. |
| `ID`          | `string`           | `""`                   | ID on the visible text input — use as the `for` target of a `<label>`. |
| `Placeholder` | `string`           | `""`                   | Placeholder in the visible input.                                  |
| `ListboxID`   | `string`           | `"flint-combobox-{Name}"` | ID on the listbox. Override only when two comboboxes on a page share a Name (rare). |
| `Required`    | `bool`             | `false`                | Marks the visible input required.                                  |
| `Disabled`    | `bool`             | `false`                | Disables the visible input and the chevron toggle.                 |
| `Invalid`     | `bool`             | `false`                | Applies the Input component's invalid border styling.              |
| `Class`       | `string`           | `""`                   | Appended to the outer wrapper.                                     |
| `Attrs`       | `templ.Attributes` | `nil`                  | Spread onto the visible text input.                                |

### OptionProps

| Field         | Type               | Default | Notes                                                       |
| ------------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Value`       | `string`           | `""`    | Required. Form value written to the hidden input on select. |
| `Label`       | `string`           | `""`    | Required. Drives filter match + input display.              |
| `Description` | `string`           | `""`    | Optional secondary line under the label in the dropdown.    |
| `Disabled`    | `bool`             | `false` | Dims the option; click / Enter / arrow-focus all skip it.   |
| `Class`       | `string`           | `""`    | Appended to the option `<li>`.                              |
| `Attrs`       | `templ.Attributes` | `nil`   | Spread onto the option `<li>`.                              |
