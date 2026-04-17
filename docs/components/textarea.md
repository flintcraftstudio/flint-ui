# Textarea

Converted from `catalyst-ui-kit/typescript/textarea.tsx`. Renders a Catalyst-styled `<textarea>` wrapped in `<span data-slot="control">` so [`Field`](./fieldset.md) can auto-space label, description, control, and error children. Vertically resizable by default; flip `NonResizable` when the surrounding layout owns sizing.

## Import

```go
import "github.com/flintcraft/flint-ui/components/textarea"
```

## Usage

```go
@textarea.Textarea(textarea.Props{
    ID:          "brief",
    Name:        "brief",
    Rows:        5,
    Placeholder: "Project brief…",
    Required:    true,
})
```

Wrap with [`fieldset.Field`](./fieldset.md) for label + description + error spacing:

```go
@fieldset.Field(fieldset.Props{}) {
    @fieldset.Label(fieldset.LabelProps{For: "brief"}) { Project brief }
    @fieldset.Description(fieldset.Props{}) { Site, scope, and gotchas. }
    @textarea.Textarea(textarea.Props{ID: "brief", Name: "brief", Rows: 5})
    @fieldset.ErrorMessage(fieldset.Props{}) { Brief is required. }
}
```

## Props

| Field          | Type               | Default | Notes                                                                              |
| -------------- | ------------------ | ------- | ---------------------------------------------------------------------------------- |
| `Name`         | `string`           | `""`    | Form field name. Omitted from the rendered element if empty.                       |
| `ID`           | `string`           | `""`    | `id` attribute — set this so `fieldset.Label{For: ...}` can wire up.               |
| `Placeholder`  | `string`           | `""`    | Placeholder text.                                                                  |
| `Value`        | `string`           | `""`    | Initial text content (rendered between the tags, not as a `value` attribute).      |
| `Autocomplete` | `string`           | `""`    | `autocomplete` attribute, e.g. `"street-address"`, `"off"`.                        |
| `Rows`         | `int`              | `0`     | `rows` attribute — visible row count. Only emitted when > 0 (browser default: 2).  |
| `NonResizable` | `bool`             | `false` | Turns off the vertical resize handle (`resize-none` instead of `resize-y`).        |
| `Required`     | `bool`             | `false` | `required` attribute.                                                              |
| `Disabled`     | `bool`             | `false` | `disabled` attribute. Wrapper uses `has-data-disabled:` to fade itself out.        |
| `Readonly`     | `bool`             | `false` | `readonly` attribute.                                                              |
| `Invalid`      | `bool`             | `false` | Adds `data-invalid` + `aria-invalid="true"` — flips to the danger border.          |
| `Class`        | `string`           | `""`    | Appended to the wrapper `<span>` (matching Catalyst's `className` behavior).       |
| `Attrs`        | `templ.Attributes` | `nil`   | Spread onto the `<textarea>`. Use for `hx-*`, `data-*`, `aria-*`, `maxlength`, `x-model`, etc. |

## Resizable

Catalyst's default is vertically resizable (`resize-y`). flint-ui keeps the same default — the common-case zero-value Props needs no extra plumbing. Set `NonResizable: true` for fixed-height textareas (inline editors, modal forms where the surrounding layout controls sizing):

```go
@textarea.Textarea(textarea.Props{Rows: 3, NonResizable: true})
```

Horizontal resizing isn't exposed — Catalyst only ships `resize-y` vs `resize-none`.

## Accessibility

- Always pair with `fieldset.Label{For: id}` (or a surrounding `<label>`) for screen reader users.
- `Invalid` sets `aria-invalid="true"` so assistive tech announces validation state.
- `data-invalid` targets both Headless UI's `[data-invalid]` *and* the native `:user-invalid` pseudo-class (see `styles/flint.css`), so browser-validated fields light up without extra JS.
- Focus ring is applied via `sm:focus-within:after:ring-2 sm:focus-within:after:ring-ring` on the wrapper — visible for both mouse and keyboard users.
