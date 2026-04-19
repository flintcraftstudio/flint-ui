# Avatar

User-identity avatar: image, initials, or both stacked so initials bridge the image-loading gap and remain visible if the image fails. Converted from `catalyst-ui-kit/typescript/avatar.tsx`.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/avatar"
```

## Components

| Component              | Renders                     | Notes                                                             |
| ---------------------- | --------------------------- | ----------------------------------------------------------------- |
| `avatar.Avatar`        | `<span data-slot="avatar">` | Static display avatar. Stacks image over initials in a 1x1 grid.  |
| `avatar.AvatarButton`  | `<button>` or `<a>`         | Focusable wrapper. Focus ring + expanded touch target for mobile. |

## Quick example

```go
@avatar.Avatar(avatar.Props{
    Src:      "/users/42/photo.jpg",
    Initials: "JS",
    Alt:      "Jaden Sturgis",
    Class:    "size-10",
})
```

Renders a 40×40 round avatar. If the image fails to load, "JS" shows through. If no `Src` is provided, it's an initials-only avatar.

## No Size prop

Avatars don't ship with a Size enum — callers set size via Class. Same pattern as Catalyst, same pattern as `*:data-[slot=avatar]:size-6` rules we already use in Dropdown.

Rationale: avatars show up at wildly different sizes (4×4 inline badges, 32×32 table rows, 128×128 profile heros). Enumerating sizes would either require 6+ variants or leave callers reaching for Class anyway. Skipping the enum keeps the API one prop smaller.

## Initials

```go
@avatar.Avatar(avatar.Props{Initials: "JS", Alt: "Jaden Sturgis", Class: "size-10"})
```

Renders as an SVG text inside the grid cell — scales with the size class, no font-size math needed. Background color is transparent by default; wrap in a colored parent for colored avatars:

```go
<span class="bg-primary text-primary-foreground inline-flex rounded-full">
    @avatar.Avatar(avatar.Props{Initials: "JS", Alt: "Jaden Sturgis", Class: "size-10"})
</span>
```

Per-client branding flows through — A-Team's red avatars, Lo Mo's teal, Western Skies' gold, without touching the component.

## Image with initials fallback

```go
@avatar.Avatar(avatar.Props{
    Src:      "/uploads/avatars/jaden.jpg",
    Initials: "JS",
    Alt:      "Jaden Sturgis",
    Class:    "size-10",
})
```

The grid stacks both children in the same cell. The `<img>` sits on top of the SVG, so:

- **Image loading** → initials visible until the image renders
- **Image loads** → image covers the initials
- **Image fails / returns 404** → initials remain visible (broken img icon is hidden by the rounded clip)

## Square variant

```go
@avatar.Avatar(avatar.Props{Square: true, Initials: "FC", Class: "size-10"})
```

`Square: true` swaps full-round corners for a subtle `rounded-[20%]` (Catalyst's convention). Use for workspace/organization avatars where the square shape signals "entity" rather than "person."

## Stacked groups

Overlapping-avatar rows ("3 people on this thread") come from CSS alone — `-space-x-*` + a ring matching the surrounding background:

```go
<div class="flex items-center -space-x-2">
    <span class="bg-primary text-primary-foreground inline-flex rounded-full ring-2 ring-surface">
        @avatar.Avatar(avatar.Props{Initials: "JS", Alt: "Jaden", Class: "size-8"})
    </span>
    <span class="bg-accent text-accent-foreground inline-flex rounded-full ring-2 ring-surface">
        @avatar.Avatar(avatar.Props{Initials: "ML", Alt: "Marcus", Class: "size-8"})
    </span>
    <span class="bg-muted text-foreground inline-flex rounded-full ring-2 ring-surface">
        @avatar.Avatar(avatar.Props{Initials: "+3", Class: "size-8"})
    </span>
</div>
```

## AvatarButton

```go
@avatar.AvatarButton(avatar.ButtonProps{
    Src:   "/uploads/avatars/jaden.jpg",
    Alt:   "Open profile",
    Href:  "/profile",
    Class: "size-10",
})
```

Renders the avatar inside a focusable element — `<a>` when `Href` is set, otherwise `<button type="button">`. Adds:

- Focus ring (`data-focus:outline-2 data-focus:outline-ring`) — shows on keyboard focus only (aliased to `:focus-visible`)
- Expanded touch target — invisible layer makes the hit area at least 44×44 px on touch devices regardless of the visible avatar size

Wire click handlers via `Attrs: templ.Attributes{"x-on:click": "..."}` for button variants, or `hx-*` for htmx interactions.

## Accessibility

- `Alt` is applied both to the `<img>` and to the SVG's `<title>`, so screen readers announce the avatar whether the image loads or falls back to initials.
- `Alt: ""` marks the avatar decorative (`aria-hidden="true"` on the SVG) — use when the avatar sits next to the user's name as visual flavor.
- `data-slot="avatar"` is stamped on every rendered Avatar so upstream components (Dropdown, eventually Table) can target avatars for sizing / spacing rules without fragile selectors.

## Props reference

### Props (Avatar)

| Field      | Type               | Default | Notes                                                             |
| ---------- | ------------------ | ------- | ----------------------------------------------------------------- |
| `Src`      | `string`           | `""`    | Image URL. Empty renders initials-only.                           |
| `Initials` | `string`           | `""`    | Fallback text. Shown while Src loads or if it fails.              |
| `Alt`      | `string`           | `""`    | Applied to img alt and SVG title. Empty marks avatar decorative.  |
| `Square`   | `bool`             | `false` | Rounded-square corners (~20%) instead of full circle.             |
| `Class`    | `string`           | `""`    | Appended to the root span. Set size here (`size-8`, etc).         |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the root span.                                        |

### ButtonProps (AvatarButton)

| Field       | Type               | Default | Notes                                                              |
| ----------- | ------------------ | ------- | ------------------------------------------------------------------ |
| `Src`       | `string`           | `""`    | Passed through to the inner Avatar.                                |
| `Initials`  | `string`           | `""`    | Passed through to the inner Avatar.                                |
| `Alt`       | `string`           | `""`    | Passed through to the inner Avatar.                                |
| `Square`    | `bool`             | `false` | Passed through to the inner Avatar + the button's border radius.   |
| `Href`      | `string`           | `""`    | Empty renders `<button>`; non-empty renders `<a href>`.            |
| `Disabled`  | `bool`             | `false` | Applies to the button variant only (links can't be natively disabled). |
| `Class`     | `string`           | `""`    | Appended to the button/link. Set size + background here.           |
| `Attrs`     | `templ.Attributes` | `nil`   | Spread onto the button/link.                                       |
