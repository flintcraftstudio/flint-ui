# Toast

Ephemeral notifications. From-scratch design — Catalyst has no `toast.tsx`.

A single `toast.Container` renders once per page (typically in your base layout). Any code on the page can fire a toast by dispatching the `flint:toast` window event with a payload — including htmx responses via `HX-Trigger`. There's no per-toast templ; toasts are content-driven and assembled at runtime from the dispatched payload.

## Import

```go
import "github.com/flintcraft/flint-ui/components/toast"
```

## Setup — render Container once

In your base layout, after the main content:

```go
templ Layout() {
    <html>
        <body x-data>
            <main>{ children... }</main>
            @toast.Container(toast.ContainerProps{})
        </body>
    </html>
}
```

That's it. Anything on any page can now fire a toast.

The Container is `pointer-events-none fixed inset-0` — it covers the viewport without blocking clicks. Individual toasts re-enable pointer events so the close button and hover-to-pause work.

## Firing a toast

### From Alpine (in-page)

```html
<button x-on:click="$dispatch('flint:toast', {
    variant: 'success',
    title: 'Saved',
    body: 'Job #1834 scheduled for Apr 22.'
})">
    Save
</button>
```

### From htmx (server response)

The killer pattern. Server returns an `HX-Trigger` header naming `flint:toast` with the payload as JSON; htmx dispatches the event client-side.

```go
func handleSave(w http.ResponseWriter, r *http.Request) {
    // ... do work ...
    w.Header().Set("HX-Trigger", `{"flint:toast":{"variant":"success","title":"Saved"}}`)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte(`<span class="text-success">Saved.</span>`))
}
```

No client-side dispatch code. Any htmx-targeting endpoint can show a toast just by setting the header.

### From plain JavaScript

```js
window.dispatchEvent(new CustomEvent('flint:toast', {
    detail: { variant: 'info', title: 'Heads up' }
}));
```

## Payload

| Field      | Type                            | Default | Notes                                                         |
| ---------- | ------------------------------- | ------- | ------------------------------------------------------------- |
| `variant`  | `'info'\|'success'\|'warning'\|'danger'` | `'info'` | Drives left border color and icon. Unknown values fall back to `info`. |
| `title`    | `string`                        | `''`    | Headline text. Required for the toast to be useful.           |
| `body`     | `string`                        | `''`    | Optional supporting line. Hidden when empty.                  |
| `duration` | `number` (ms)                   | `5000`  | Auto-dismiss timeout. Pass `0` to keep the toast until manually dismissed. |

## Behavior

- **Stacking** — multiple toasts append below previous ones (`gap-2` between).
- **Slide transitions** — toasts slide in from the placement-side edge (right edge for right placements, left edge for left placements) and slide back out the same way on dismiss. 300ms enter, 200ms leave.
- **Auto-dismiss** — each toast has its own timer driven by `duration`.
- **Countdown indicator** — a thin variant-tinted progress bar at the bottom drains from full width down to zero over the toast's duration, so the user can see how much time is left before dismiss. Hidden when `duration: 0` (persistent toasts).
- **Pause on interaction** — hovering or focusing inside a toast pauses *that* toast's timer *and* its progress bar (CSS `animation-play-state: paused`); mouse-leave / blur resumes both. So the user can read a long body without it disappearing under their cursor, and the visual countdown agrees with the actual timer.
- **Manual dismiss** — every toast has an X button that removes it immediately.
- **ESC** — dismisses the most recent toast (only the most recent, to avoid wiping unrelated notifications with one keystroke).
- **No queue** — toasts dispatched before Alpine initializes are lost. CustomEvents don't queue by themselves. In practice this only matters for toasts fired from inline scripts that run before `defer`-loaded Alpine; htmx-driven toasts and Alpine-handler toasts are always safe.

## Placement

Default is bottom-right (the conventional dashboard placement). Pick a different corner via `Placement`:

```go
@toast.Container(toast.ContainerProps{Placement: toast.PlacementTopRight})
```

| Constant                       | Position        |
| ------------------------------ | --------------- |
| `toast.PlacementBottomRight`   | Bottom-right (default) |
| `toast.PlacementBottomLeft`    | Bottom-left     |
| `toast.PlacementTopRight`      | Top-right       |
| `toast.PlacementTopLeft`       | Top-left        |

In every placement, new toasts append below older ones in DOM order. For bottom-aligned stacks, this looks like "new toast pushes old ones up". For top-aligned stacks, "new toast appears below existing." Both are conventional; `flex-col-reverse` is intentionally not used in v0.1.

## Variants

Each variant gets a colored left border and an icon tinted to the matching semantic token:

| Variant   | Border / Icon Token | Use for                                  |
| --------- | -------------------- | ---------------------------------------- |
| `info`    | `primary`            | General confirmations, neutral updates.  |
| `success` | `success`            | "It worked." Saves, completions.         |
| `warning` | `warning`            | Conflicts, soft failures, near-misses.   |
| `danger`  | `danger`             | Errors, failed actions.                  |

Themed by the client palette via the same tokens the rest of flint-ui uses.

## Accessibility

- Container is `aria-live="polite"` and `aria-atomic="false"` — assistive tech announces new toasts as they arrive without re-reading the whole stack.
- Each toast has `role="status"` (non-critical announcement). Screen readers will read the title and body when the toast appears.
- The X button has `aria-label="Dismiss notification"`.
- ESC dismisses the most recent toast.
- **`danger` variant is still polite, not assertive.** v0.1 does not escalate to `aria-live="assertive"` for danger toasts. If a failure is truly critical and must interrupt the user, render a Modal instead — toasts are for notification, not gating.

## Required surrounding setup

Same baseline as Modal / Dropdown / Tabs:

1. **`x-data` on a common ancestor** in your base layout (typically `<body x-data>`), so Alpine processes the Container's directives and so trigger buttons can call `$dispatch`.
2. **Alpine and Focus plugin loaded** in the right order (Focus before core). The Container itself doesn't use the Focus plugin, but the rest of flint-ui does, and this is the conventional load order.

```html
<script src="https://unpkg.com/@alpinejs/focus@3.14.1/dist/cdn.min.js" defer></script>
<script src="https://unpkg.com/alpinejs@3.14.1/dist/cdn.min.js" defer></script>
```

## Why no per-toast templ component?

The Modal component is declarative — one `modal.Modal{Name: "x"}` per dispatchable modal — because modals usually have author-controlled, reusable content. Toasts are different: the message almost always comes from a server response or a runtime computation, not from an authored template. Building a `toast.Toast{Name: "x"}` declarative API would require declaring every possible message in advance, which doesn't match how toasts get used.

Single Container + dispatch payload trades a little client-side construction for a much simpler API: anything that can dispatch a custom event can show a toast, with no flint-ui binding code. htmx `HX-Trigger` headers, third-party scripts, inline `onclick` — all "just work."

## Props reference

### ContainerProps

| Field       | Type                | Default              | Notes                                                                   |
| ----------- | ------------------- | -------------------- | ----------------------------------------------------------------------- |
| `Placement` | `Placement`         | `PlacementBottomRight` (zero value) | One of the four corner constants above.                                 |
| `Class`     | `string`            | `""`                 | Appended to the container `<div>`.                                      |
| `Attrs`     | `templ.Attributes`  | `nil`                | Spread onto the container.                                              |
