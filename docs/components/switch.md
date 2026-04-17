# Switch

Accessible on/off toggle for settings. Native `<input type="checkbox" role="switch">` under the hood. Converted from `catalyst-ui-kit/typescript/switch.tsx`.

> **Package name**: `toggle` — `switch` is a Go reserved keyword so the package can't use it. The exported type is `Switch` so calls read naturally: `@toggle.Switch(toggle.Props{...})`.

## Import

```go
import "github.com/flintcraft/flint-ui/components/toggle"
```

## Components

| Component              | Renders                                      | Notes                                                 |
| ---------------------- | -------------------------------------------- | ----------------------------------------------------- |
| `toggle.Switch`        | `<label data-slot="control">` wrapping an sr-only checkbox | Single styled switch. |
| `toggle.SwitchGroup`   | `<div>`                                      | Vertical stack. Auto-bumps spacing when children have descriptions. |
| `toggle.SwitchField`   | `<div>` with a 2-col grid                    | Label in col 1, Switch in col 2 (reversed from Checkbox/Radio). |

## Quick example

```go
@toggle.SwitchField(toggle.GroupProps{}) {
    @fieldset.Label(fieldset.LabelProps{For: "notif-email"}) { Email notifications }
    @fieldset.Description(fieldset.Props{}) {
        Daily summary of new leads and outstanding jobs.
    }
    @toggle.Switch(toggle.Props{ID: "notif-email", Name: "notif_email", Value: "1", Checked: true})
}
```

## How it works

A `<label>` wraps an `sr-only <input type="checkbox" role="switch">`. The label renders the track (via its own background) and a thumb span. State propagation uses Tailwind's `has-checked:` and `has-focus-visible:` variants — the label styles itself based on whether its descendant input is checked or focused.

The form submits `notif_email=1` when on, nothing when off — standard checkbox form semantics.

Keyboard:

- **Tab** focuses the (sr-only) input inside the label
- **Space** toggles
- **has-focus-visible** on the label shows the ring when the input is focused

The `role="switch"` + `aria-checked` combo makes screen readers announce "switch, on/off" rather than "checkbox, checked/unchecked."

## SwitchField column order

SwitchField uses `grid-cols-[1fr_auto]` — label on the left, switch pushed to the right. This is the opposite of CheckboxField and RadioField, which place the control first.

Rationale: on/off switches read naturally as "setting name: [toggle]" — iPhone settings style. Form fields with checkboxes or radios read the other direction ("[x] Remember me").

## Why not a native checkbox?

You could ship a checkbox and call it a switch, but screen readers would announce it as a checkbox. Adding `role="switch"` + matching visual design lets assistive tech correctly describe the affordance.

## Props reference

### Props (Switch)

| Field      | Type               | Default | Notes                                                  |
| ---------- | ------------------ | ------- | ------------------------------------------------------ |
| `Name`     | `string`           | `""`    | Form field name. The switch submits `name=value` when checked. |
| `ID`       | `string`           | `""`    | `id` on the input — use as the `for` target of a `<label>`. |
| `Value`    | `string`           | `""`    | Submitted value when the switch is on.                 |
| `Checked`  | `bool`             | `false` | Initial state.                                         |
| `Required` | `bool`             | `false` | Marks the input required.                              |
| `Disabled` | `bool`             | `false` | Dims the switch and blocks interaction.                |
| `Invalid`  | `bool`             | `false` | Adds data-invalid / aria-invalid for validation styling. |
| `Class`    | `string`           | `""`    | Appended to the outer label.                           |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the inner input. Use for `hx-*`, custom aria. |

### GroupProps (SwitchGroup / SwitchField)

| Field   | Type               | Default | Notes                           |
| ------- | ------------------ | ------- | ------------------------------- |
| `Class` | `string`           | `""`    | Appended to the wrapper.        |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the wrapper.        |
