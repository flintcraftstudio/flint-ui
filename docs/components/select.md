# Select

Converted from `catalyst-ui-kit/typescript/select.tsx`. Renders a native `<select>` wrapped in `<span data-slot="control">` so [`Field`](./fieldset.md) can auto-space label, description, control, and error children. Options are passed as children — the component keeps the native `<option>` / `<optgroup>` tree intact so server-side rendering and form submission Just Work.

## Import

The Go package name is `selectbox` (not `select` — that's a Go reserved keyword):

```go
import "github.com/flintcraft/flint-ui/components/selectbox"
```

## Usage

```go
@selectbox.Select(selectbox.Props{ID: "region", Name: "region"}) {
    <option value="">Pick a region…</option>
    <option value="us">United States</option>
    <option value="ca">Canada</option>
}
```

Wrap with [`fieldset.Field`](./fieldset.md) for label + description + error spacing:

```go
@fieldset.Field(fieldset.Props{}) {
    @fieldset.Label(fieldset.LabelProps{For: "region"}) { Region }
    @fieldset.Description(fieldset.Props{}) { Where are you based? }
    @selectbox.Select(selectbox.Props{ID: "region", Name: "region"}) {
        <option value="us">US</option>
        <option value="ca">CA</option>
    }
    @fieldset.ErrorMessage(fieldset.Props{}) { Region is required. }
}
```

Mark the selected option natively with the `selected` attribute:

```go
@selectbox.Select(selectbox.Props{Name: "currency"}) {
    <option value="usd" selected>USD</option>
    <option value="cad">CAD</option>
}
```

## Props

| Field          | Type               | Default | Notes                                                                               |
| -------------- | ------------------ | ------- | ----------------------------------------------------------------------------------- |
| `Name`         | `string`           | `""`    | Form field name. Omitted from the rendered element if empty.                        |
| `ID`           | `string`           | `""`    | `id` attribute — set this so `fieldset.Label{For: ...}` can wire up.                |
| `Autocomplete` | `string`           | `""`    | `autocomplete` attribute, e.g. `"country"`, `"off"`.                                |
| `Multiple`     | `bool`             | `false` | Multi-select list box. Drops the chevron and switches to symmetric horizontal padding. |
| `Size`         | `int`              | `0`     | `size` attribute — number of visible rows. Only emitted when > 0.                   |
| `Required`     | `bool`             | `false` | `required` attribute.                                                               |
| `Disabled`     | `bool`             | `false` | `disabled` attribute. Wrapper uses `has-data-disabled:` to fade itself out.         |
| `Invalid`      | `bool`             | `false` | Adds `data-invalid` + `aria-invalid="true"` — flips to the danger border.           |
| `Class`        | `string`           | `""`    | Appended to the wrapper `<span>` (matching Catalyst's `className` behavior).        |
| `Attrs`        | `templ.Attributes` | `nil`   | Spread onto the `<select>`. Use for `hx-*`, `data-*`, `aria-*`, `x-model`, etc.     |

## Multiple

`Multiple: true` turns the control into a native multi-select list box. The chevron is omitted (it only makes sense for dropdowns), and the select's horizontal padding is made symmetric since there's no chevron gutter to leave room for. Pair with `Size` to control how many rows are visible:

```go
@selectbox.Select(selectbox.Props{Name: "tags", Multiple: true, Size: 5}) {
    <option value="new" selected>New lead</option>
    <option value="vip" selected>VIP</option>
    <option value="referral">Referral</option>
}
```

The form submits the list of selected values as `tags` (multi-value form data — `r.Form["tags"]` on the Go side).

## Optgroups

Native `<optgroup>` works out of the box — Catalyst's `[&_optgroup]:font-semibold` selector bolds the labels:

```go
@selectbox.Select(selectbox.Props{Name: "service"}) {
    <optgroup label="Cleaning">
        <option value="standard">Standard clean</option>
        <option value="deep">Deep clean</option>
    </optgroup>
    <optgroup label="Trades">
        <option value="gutter">Gutter install</option>
    </optgroup>
}
```

## Accessibility

- Always pair with `fieldset.Label{For: id}` (or a surrounding `<label>`) for screen reader users.
- `Invalid` sets `aria-invalid="true"` so assistive tech announces validation state.
- `data-invalid` targets both Headless UI's `[data-invalid]` *and* the native `:user-invalid` pseudo-class (see `styles/flint.css`), so browser-validated fields light up without extra JS.
- The chevron icon carries `forced-colors:stroke-[CanvasText]` so it remains visible in Windows high-contrast mode.
- Focus ring is applied via `has-data-focus:after:ring-2 has-data-focus:after:ring-ring` on the wrapper — visible for both mouse and keyboard users.
