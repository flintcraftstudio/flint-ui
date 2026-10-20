# Progress

A horizontal progress bar, from scratch (Catalyst has none). Determinate (`Value` out of `Max`) or indeterminate (a segment that loops across the track for work with no known length). Renders `role="progressbar"`.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/progress"
```

## Usage

```go
// Determinate — 42 of 100
@progress.Progress(progress.Props{Value: 42, Label: "Upload"})

// Custom max and value text
@progress.Progress(progress.Props{Value: 2, Max: 5, Label: "Onboarding", ValueText: "Step 2 of 5"})

// Indeterminate — loops until you replace it
@progress.Progress(progress.Props{Indeterminate: true, Label: "Syncing invoices"})
```

The bar has no built-in text. For a visible label and value, put them in a row above it:

```go
<div class="space-y-2">
    <div class="flex justify-between text-sm/6">
        <span id="export-label" class="font-medium text-foreground">Exporting invoices</span>
        <span class="text-muted-foreground tabular-nums">60%</span>
    </div>
    @progress.Progress(progress.Props{Value: 60, Attrs: templ.Attributes{"aria-labelledby": "export-label"}})
</div>
```

## Server-reported progress (htmx polling)

Poll a wrapper with `hx-trigger="every …"` and swap it with `hx-swap="outerMorph"`. Morphing updates the existing fill's `style` in place, so the width animates to each new value instead of jumping. To stop polling, return the final state **without** the `hx-*` attributes: the morph removes them and htmx clears the interval.

```go
templ ExportJob(job Job) {
    <div
        id="export-job"
        if !job.Done {
            hx-get={ fmt.Sprintf("/jobs/%d/progress", job.ID) }
            hx-trigger="every 1s"
            hx-swap="outerMorph"
        }
    >
        @progress.Progress(progress.Props{
            Value:   float64(job.Processed),
            Max:     float64(job.Total),
            Label:   "Invoice export",
            Variant: jobVariant(job), // e.g. VariantSuccess when done
        })
    </div>
}
```

The handler renders the same component. The showcase's `/progress` page has a working, stateless version (`ProgressJob` in `examples/showcase/templates/progress.templ`).

Upload progress isn't available this way: htmx 4 sends requests with `fetch`, which doesn't report upload progress. For file uploads, use an indeterminate bar as the request indicator (below) or `button.Props.Indicator`.

## As an htmx indicator

An indeterminate bar with the `htmx-indicator` class stays hidden until the request it's pointed at is in flight:

```go
@button.Button(button.Props{Outline: true, Attrs: templ.Attributes{
    "hx-get":       "/reports/weekly",
    "hx-target":    "#report",
    "hx-indicator": "#report-progress",
}}) { Load report }
@progress.Progress(progress.Props{
    Indeterminate: true,
    Size:          progress.SizeSM,
    Label:         "Loading report",
    Class:         "htmx-indicator",
    Attrs:         templ.Attributes{"id": "report-progress"},
})
```

## Variants

| Variant                    | Fill          |
| -------------------------- | ------------- |
| `progress.VariantPrimary`  | `bg-primary` (default) |
| `progress.VariantAccent`   | `bg-accent`   |
| `progress.VariantSuccess`  | `bg-success`  |
| `progress.VariantWarning`  | `bg-warning`  |
| `progress.VariantDanger`   | `bg-danger`   |

The track is `bg-muted` in every variant.

## Sizes

| Size              | Height | Use for                                   |
| ----------------- | ------ | ----------------------------------------- |
| `progress.SizeSM` | 4px    | Table cells, card footers, indicators.    |
| `progress.SizeMD` | 8px    | Default.                                  |

## Props

| Field           | Type               | Default          | Notes |
| --------------- | ------------------ | ---------------- | ----- |
| `Value`         | `float64`          | `0`              | Clamped to `0..Max`. Ignored when `Indeterminate`. |
| `Max`           | `float64`          | `100`            | Zero or negative means 100. |
| `Indeterminate` | `bool`             | `false`          | Loops a segment; omits `aria-valuenow`. |
| `Label`         | `string`           | `""`             | Accessible name (`aria-label`). Set it, or pass `aria-labelledby` in `Attrs`. |
| `ValueText`     | `string`           | rounded percent  | `aria-valuetext`, e.g. "Step 2 of 5". |
| `Variant`       | `progress.Variant` | `VariantPrimary` | Fill color. |
| `Size`          | `progress.Size`    | `SizeMD`         | Track height. |
| `Class`         | `string`           | `""`             | Appended to the track. |
| `Attrs`         | `templ.Attributes` | `nil`            | Spread onto the track (`id`, `aria-labelledby`, `hx-*`). |

## Accessibility

- Always give the bar a name, with `Label` or `aria-labelledby`. Without one it's announced as just "progress bar".
- Indeterminate bars omit `aria-valuenow`, which is how assistive tech tells them apart from determinate ones.
- Screen readers don't reliably announce progressbar value changes. When a long job finishes, announce the result too (a toast, or text in an `aria-live` region).
- Under `prefers-reduced-motion`, the determinate fill jumps instead of animating, and the indeterminate segment fills the track and pulses gently instead of sliding. A frozen bar reads as "stuck".
- In Windows high-contrast mode the fill uses the system `Highlight` color and the track gets an outline.
