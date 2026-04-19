# Checkbox

Converted from `catalyst-ui-kit/typescript/checkbox.tsx`. Catalyst uses Headless UI's custom `<Checkbox>` control; flint-ui renders a real `<input type="checkbox">` with a styled sibling box + SVG check so form submission, keyboard handling, and focus semantics work without JS. The outer `<label>` carries `data-slot="control"` so `CheckboxField`'s grid layout positions it in the left column.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/checkbox"
```

## Usage

Bare checkbox (table row selection, toolbars):

```go
@checkbox.Checkbox(checkbox.Props{ID: "row-1", Name: "rows", Value: "1"})
```

With label:

```go
@checkbox.CheckboxField(checkbox.GroupProps{}) {
    @checkbox.Checkbox(checkbox.Props{ID: "terms", Name: "terms"})
    @fieldset.Label(fieldset.LabelProps{For: "terms"}) { I agree to the terms }
}
```

With label + description:

```go
@checkbox.CheckboxField(checkbox.GroupProps{}) {
    @checkbox.Checkbox(checkbox.Props{ID: "marketing", Name: "marketing", Checked: true})
    @fieldset.Label(fieldset.LabelProps{For: "marketing"}) { Send me product updates }
    @fieldset.Description(fieldset.Props{}) { About one email a month. }
}
```

Group of checkboxes sharing a name:

```go
@checkbox.CheckboxGroup(checkbox.GroupProps{}) {
    @checkbox.CheckboxField(checkbox.GroupProps{}) {
        @checkbox.Checkbox(checkbox.Props{ID: "n-leads", Name: "notifications", Value: "leads"})
        @fieldset.Label(fieldset.LabelProps{For: "n-leads"}) { New leads }
    }
    @checkbox.CheckboxField(checkbox.GroupProps{}) {
        @checkbox.Checkbox(checkbox.Props{ID: "n-reviews", Name: "notifications", Value: "reviews"})
        @fieldset.Label(fieldset.LabelProps{For: "n-reviews"}) { New reviews }
    }
}
```

Form submits multiple values as `r.Form["notifications"]` on the Go side.

## Props

| Field      | Type               | Default | Notes                                                                         |
| ---------- | ------------------ | ------- | ----------------------------------------------------------------------------- |
| `Name`     | `string`           | `""`    | Form field name. Omitted from the rendered element if empty.                  |
| `ID`       | `string`           | `""`    | `id` attribute — set this so `fieldset.Label{For: ...}` can wire up.          |
| `Value`    | `string`           | `""`    | `value` attribute submitted when checked. Browser default is `"on"` if empty. |
| `Checked`  | `bool`             | `false` | Initial checked state.                                                        |
| `Required` | `bool`             | `false` | `required` attribute.                                                         |
| `Disabled` | `bool`             | `false` | `disabled` attribute. Box fades to opacity-50.                                |
| `Invalid`  | `bool`             | `false` | Adds `data-invalid` + `aria-invalid="true"` — flips to the danger border.     |
| `Class`    | `string`           | `""`    | Appended to the outer `<label>`.                                              |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the `<input>`. Use for `hx-*`, `data-*`, `aria-*`, `x-model`, etc. |

## Layout components

- **`CheckboxField`** — two-column grid that positions the Checkbox in column 1 and `fieldset.Label` / `fieldset.Description` siblings in column 2. Use this whenever a checkbox has a visible label.
- **`CheckboxGroup`** — vertical stack of `CheckboxField`s. Spacing auto-bumps from `space-y-3` to `space-y-6` when any child field contains a description.

Both accept a `GroupProps{Class, Attrs}`.

## Indeterminate state

Not yet exposed. Native `<input type="checkbox">` only supports indeterminate via JS (`input.indeterminate = true`). If you need it, set it via `Attrs` with an `x-init` or `onload` hook, or add Alpine state. A native `Indeterminate` prop may land when a client need forces the decision.

## Accessibility

- The outer `<label>` makes the entire 18×18px (16×16 on sm) box clickable and screen-reader-labelled.
- `Invalid` sets `aria-invalid="true"` and `data-invalid`; the box border turns `border-danger`.
- Focus ring uses `peer-focus-visible:outline-ring` on the box — visible for keyboard users, hidden for mouse.
- `forced-colors:stroke-[HighlightText]` keeps the checkmark visible in Windows high-contrast mode.
