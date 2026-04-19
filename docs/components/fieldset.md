# Fieldset helpers

Converted from `catalyst-ui-kit/typescript/fieldset.tsx`. The `fieldset` package is a small set of layout primitives that compose with controls like [`Input`](./input.md) to produce consistently-spaced forms. Each element carries a `data-slot` attribute that other components (notably `Field`) use to auto-space siblings.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/fieldset"
```

## Layout model

```
Fieldset
└── Legend                 ← caption for the whole form section
└── FieldGroup             ← vertical stack of Fields (space-y-8)
    └── Field              ← one label + control + description + error
        ├── Label
        ├── Description
        ├── <control>      ← Input / Select / Textarea / …
        └── ErrorMessage
```

Placement inside `Field` drives spacing. The sibling-selector classes in Catalyst's source pick up these pairs:

| First child slot | Next sibling slot | Spacing applied |
| ---------------- | ----------------- | --------------- |
| `label`          | `control`         | `mt-3`          |
| `label`          | `description`     | `mt-1`          |
| `description`    | `control`         | `mt-3`          |
| `control`        | `description`     | `mt-3`          |
| `control`        | `error`           | `mt-3`          |

As long as children are in a sensible order, Field spaces itself correctly.

## Components

Every component below accepts a `Props{Class, Attrs}` (except `Label`, which also takes `For`). Extra classes are appended after the Catalyst-defined base; `Attrs` is spread onto the rendered element for htmx / data / aria pass-through.

| Component      | Element      | `data-slot`  | Purpose                                      |
| -------------- | ------------ | ------------ | -------------------------------------------- |
| `Fieldset`     | `<fieldset>` | —            | Outer grouping element.                      |
| `Legend`       | `<legend>`   | `legend`     | Caption for the Fieldset.                    |
| `FieldGroup`   | `<div>`      | `control`    | Vertical stack of Fields (`space-y-8`).      |
| `Field`        | `<div>`      | —            | Single row: label + control + description + error. |
| `Label`        | `<label>`    | `label`      | Label text; takes `For` to bind to an input. |
| `Description`  | `<p>`        | `description`| Hint text shown above or below the control.  |
| `ErrorMessage` | `<p>`        | `error`      | Validation error shown below the control.    |

## Usage

```go
@fieldset.Fieldset(fieldset.Props{}) {
    @fieldset.Legend(fieldset.Props{}) { Contact details }
    @fieldset.FieldGroup(fieldset.Props{}) {

        @fieldset.Field(fieldset.Props{}) {
            @fieldset.Label(fieldset.LabelProps{For: "full-name"}) { Full name }
            @input.Input(input.Props{ID: "full-name", Name: "name", Required: true})
        }

        @fieldset.Field(fieldset.Props{}) {
            @fieldset.Label(fieldset.LabelProps{For: "email"}) { Email }
            @fieldset.Description(fieldset.Props{}) { We'll reply here. }
            @input.Input(input.Props{ID: "email", Type: "email", Name: "email"})
            @fieldset.ErrorMessage(fieldset.Props{}) { Invalid address. }
        }

    }
}
```

## Accessibility

- `Label{For: id}` wires up to the input's `id` so screen readers announce the label when focus lands on the control.
- `Fieldset` / `Legend` group related fields semantically — screen readers announce the legend before each field inside the fieldset.
- `ErrorMessage` is a paragraph; pair with `aria-describedby` on the control (via `Input.Attrs`) when you want AT to announce it during validation.
