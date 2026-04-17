# Alert

Inline banner for status messages inside a page. From-scratch design — a peer of Badge, not of Modal.

> **Not to be confused with Modal's alert dialog.** Catalyst's `alert.tsx` is a modal confirmation dialog; that layout ships as [`modal.Props.Alert`](./modal.md). What lives here is the inline banner — the thing at the top of a page saying "Your license expires in 30 days" or "Saved successfully."

## Import

```go
import "github.com/flintcraft/flint-ui/components/alert"
```

## Components

| Component              | Renders                | Notes                                                        |
| ---------------------- | ---------------------- | ------------------------------------------------------------ |
| `alert.Alert`          | `<div role="alert">`   | Outer banner with variant icon + content column.             |
| `alert.Title`          | `<div>`                | Bold heading in the variant's color.                         |
| `alert.Description`    | `<div>`                | Body text in the variant's color at slight opacity reduction. |
| `alert.Actions`        | `<div>`                | Optional row of buttons or links.                            |

## Quick example

```go
@alert.Alert(alert.Props{Variant: alert.VariantWarning}) {
    @alert.Title(alert.SubProps{}) { Montana contractor license renews in 30 days }
    @alert.Description(alert.SubProps{}) {
        Renewal takes roughly a week to process.
    }
    @alert.Actions(alert.SubProps{}) {
        @button.Button(button.Props{}) { Start renewal }
        @button.Button(button.Props{Plain: true}) { Remind me later }
    }
}
```

## Variants

Match Badge's palette so alerts and badges read as the same concept at different scales — a success badge inside a success alert feels intentional.

| Variant                  | Color family                                              |
| ------------------------ | --------------------------------------------------------- |
| `alert.VariantInfo`      | `bg-primary/10 text-primary`. Default if unset.           |
| `alert.VariantSuccess`   | `bg-success/10 text-success`.                             |
| `alert.VariantWarning`   | `bg-warning/25 text-warning-foreground`. The warning token is yellow, so text flips to the foreground pair for readability (same exception as Badge). |
| `alert.VariantDanger`    | `bg-danger/10 text-danger`.                               |

Per-client rebranding flows through the variant tokens automatically — an A-Team alert uses A-Team red on `VariantDanger`, Lo Mo's teal on `VariantInfo`, etc.

## Icon

The variant's icon renders automatically. Same glyph set as Toast so a success alert and a success toast share symbology. No `Icon: false` toggle — if you need a silent alert, override via Class.

## Title only

Drop Description for a compact single-line alert. Good for inline validation messages or narrow card contexts:

```go
@alert.Alert(alert.Props{Variant: alert.VariantSuccess}) {
    @alert.Title(alert.SubProps{}) { All changes saved }
}
```

## Accessibility

- The root carries `role="alert"` so assistive tech announces the content when the element appears in the DOM. This is correct for banners that appear in response to an action (form submit, save, etc.).
- For alerts that are present on first paint (e.g., a license-expiry banner at the top of a dashboard), consider adding `role="status"` instead via `Attrs` — `role="alert"` is disruptive when the banner is always there.
- The decorative icon carries `aria-hidden="true"`; announce the variant via the Title text itself ("Warning: ..." or "Success: ...") if your dashboard's screen-reader audience needs that cue.

## Not dismissable (yet)

V0.1 alerts are sticky. Most inline alerts are informative and persist until the underlying state resolves (license renewed, payment succeeded). If a client needs a dismiss button, add it via Actions with an Alpine `x-on:click="$el.closest('[role=alert]').remove()"` — or ship a proper `Dismissable` prop when the need is real.

## Props reference

### Props (Alert root)

| Field     | Type               | Default        | Notes                                          |
| --------- | ------------------ | -------------- | ---------------------------------------------- |
| `Variant` | `Variant`          | `VariantInfo`  | Picks color family and default icon.           |
| `Class`   | `string`           | `""`           | Appended to the root `<div>`.                  |
| `Attrs`   | `templ.Attributes` | `nil`          | Spread onto the root (use to override `role`). |

### SubProps (Title / Description / Actions)

| Field   | Type               | Default | Notes                                |
| ------- | ------------------ | ------- | ------------------------------------ |
| `Class` | `string`           | `""`    | Appended to the rendered element.    |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the rendered element.    |
