# Input

Converted from `catalyst-ui-kit/typescript/input.tsx`. Renders a Catalyst-styled `<input>` wrapped in `<span data-slot="control">` so [`Field`](./fieldset.md) can auto-space label, description, control, and error children.

## Import

```go
import "github.com/flintcraft/flint-ui/components/input"
```

## Usage

```go
@input.Input(input.Props{
    ID:          "email",
    Name:        "email",
    Type:        "email",
    Placeholder: "you@example.com",
    Required:    true,
    Autocomplete: "email",
})
```

Wrap with [`fieldset.Field`](./fieldset.md) for label + description + error spacing:

```go
@fieldset.Field(fieldset.Props{}) {
    @fieldset.Label(fieldset.LabelProps{For: "email"}) { Email }
    @fieldset.Description(fieldset.Props{}) { We'll only use this for receipts. }
    @input.Input(input.Props{ID: "email", Type: "email", Name: "email"})
    @fieldset.ErrorMessage(fieldset.Props{}) { Must be a valid address. }
}
```

## Props

| Field          | Type               | Default   | Notes                                                                        |
| -------------- | ------------------ | --------- | ---------------------------------------------------------------------------- |
| `Type`         | `string`           | `"text"`  | `text`, `email`, `number`, `password`, `search`, `tel`, `url`, or a date type (`date`, `datetime-local`, `month`, `time`, `week`). Date types get Catalyst's webkit datetime-edit overrides. |
| `Name`         | `string`           | `""`      | Form field name. Omitted from the rendered element if empty.                 |
| `Value`        | `string`           | `""`      | Initial value. Omitted if empty.                                             |
| `Placeholder`  | `string`           | `""`      | Placeholder text.                                                            |
| `ID`           | `string`           | `""`      | `id` attribute — set this so `fieldset.Label{For: ...}` can wire up.         |
| `Autocomplete` | `string`           | `""`      | `autocomplete` attribute, e.g. `"email"`, `"name"`, `"off"`.                 |
| `Required`     | `bool`             | `false`   | `required` attribute.                                                        |
| `Disabled`     | `bool`             | `false`   | `disabled` attribute. Wrapper uses `has-data-disabled:` to fade itself out.  |
| `Readonly`     | `bool`             | `false`   | `readonly` attribute.                                                        |
| `Invalid`     | `bool`             | `false`   | Adds `data-invalid` + `aria-invalid="true"` — flips to the red border state. |
| `Class`        | `string`           | `""`      | Appended to the wrapper `<span>` (matching Catalyst's `className` behavior). |
| `Attrs`        | `templ.Attributes` | `nil`     | Spread onto the `<input>`. Use for `hx-*`, `data-*`, `aria-*`, `minlength`, `pattern`, etc. |

## InputGroup

`InputGroup` positions one or two SVG icons around a child `Input`, shifting the input padding so the icons don't overlap the text.

```go
@input.InputGroup(input.GroupProps{}) {
    <svg data-slot="icon" ...>…</svg>
    @input.Input(input.Props{Placeholder: "Search orders…"})
}
```

- Leading icon: place the `<svg data-slot="icon">` before `Input`.
- Trailing icon: place it after.
- Both: include two `<svg data-slot="icon">` siblings.

Every icon must carry `data-slot="icon"` — the CSS selectors that position it key off that attribute.

## Accessibility

- Always pair with `fieldset.Label{For: id}` (or surrounding `<label>`) for screen reader users.
- `Invalid` sets `aria-invalid="true"` so AT announces validation state.
- `data-invalid` targets both Headless UI's own attribute *and* the native `:user-invalid` pseudo-class (see `styles/flint.css`), so browser-validated fields light up without extra JS.
- Focus ring is applied via `focus-within:after:ring-2 focus-within:after:ring-blue-500` on the wrapper — visible for both mouse and keyboard users.
