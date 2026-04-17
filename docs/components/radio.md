# Radio

Exclusive-choice form control — pick one of a set. Native `<input type="radio">` under the hood. Converted from `catalyst-ui-kit/typescript/radio.tsx`.

## Import

```go
import "github.com/flintcraft/flint-ui/components/radio"
```

## Components

| Component          | Renders                          | Notes                                                 |
| ------------------ | -------------------------------- | ----------------------------------------------------- |
| `radio.Radio`      | `<label data-slot="control">` wrapping `<input type="radio">` | Single styled radio. |
| `radio.RadioGroup` | `<div role="radiogroup">`        | Vertical stack. Auto-bumps spacing when children have descriptions. |
| `radio.RadioField` | `<div>` with a 2-col grid        | Radio in col 1, Label / Description in col 2.         |

## Quick example

```go
@fieldset.Fieldset(fieldset.Props{}) {
    @fieldset.Legend(fieldset.Props{}) { Billing plan }
    @radio.RadioGroup(radio.GroupProps{}) {
        @radio.RadioField(radio.GroupProps{}) {
            @radio.Radio(radio.Props{ID: "plan-starter", Name: "plan", Value: "starter", Checked: true})
            @fieldset.Label(fieldset.LabelProps{For: "plan-starter"}) { Starter }
        }
        @radio.RadioField(radio.GroupProps{}) {
            @radio.Radio(radio.Props{ID: "plan-pro", Name: "plan", Value: "pro"})
            @fieldset.Label(fieldset.LabelProps{For: "plan-pro"}) { Pro }
        }
    }
}
```

All Radios in a group share `Name`. Browsers handle exclusive selection natively — zero JS.

## Exclusive selection via shared `Name`

Unlike React's Headless UI which wraps state in a RadioGroup provider, flint-ui relies on the HTML spec: `<input type="radio">` elements with the same `name` attribute form an exclusive set. The browser handles checked/unchecked toggling when a user picks one.

## Pairing with Fieldset

Radio is typed to compose with `fieldset.Label` and `fieldset.Description` — same pattern as Checkbox. The RadioField grid positions Radio in column 1 and Label/Description in column 2 via `data-slot` selectors.

## Styling

Radio mirrors Checkbox's visual weight — same size, same focus ring, same primary-token checked fill. The indicator is a small centered dot that fades in on checked. Per-client rebranding flows through: A-Team's red radios, Lo Mo's teal, etc.

## Props reference

### Props (Radio)

| Field      | Type               | Default | Notes                                                  |
| ---------- | ------------------ | ------- | ------------------------------------------------------ |
| `Name`     | `string`           | `""`    | Shared name across radios in an exclusive-choice group. |
| `ID`       | `string`           | `""`    | `id` on the input — use as the `for` target of a `<label>`. |
| `Value`    | `string`           | `""`    | Submitted value when this radio is the checked one.    |
| `Checked`  | `bool`             | `false` | Initial checked state.                                 |
| `Required` | `bool`             | `false` | Marks the input required.                              |
| `Disabled` | `bool`             | `false` | Dims the radio and blocks interaction.                 |
| `Invalid`  | `bool`             | `false` | Adds data-invalid / aria-invalid for validation error styling. |
| `Class`    | `string`           | `""`    | Appended to the outer label.                           |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the inner input. Use for `hx-*`, custom aria. |

### GroupProps (RadioGroup / RadioField)

| Field   | Type               | Default | Notes                           |
| ------- | ------------------ | ------- | ------------------------------- |
| `Class` | `string`           | `""`    | Appended to the wrapper.        |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the wrapper.        |
