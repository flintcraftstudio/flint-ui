# Page header

The title block at the top of a dashboard page: optional back link, page title, supporting copy or metadata, and the page's primary actions. It's composed from scratch, since Catalyst ships the pieces (Heading, Text, Button) and lays them out inline on every demo page. Spacing follows the Catalyst demo's orders pages.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/pageheader"
```

## Components

| Component                 | Renders  | Notes                                                        |
| ------------------------- | -------- | ------------------------------------------------------------ |
| `pageheader.PageHeader`   | `<div>`  | Wrapping row. Content left, Actions right (bottom-aligned).  |
| `pageheader.Back`         | `<a>`    | "‹ Parent" link on its own row above the title.             |
| `pageheader.Content`      | `<div>`  | Left column. Full width below `sm`.                          |
| `pageheader.Title`        | `<h1>`   | `heading.Heading` styles. `Level` changes the tag.           |
| `pageheader.Description`  | `<p>`    | Supporting copy under the title.                             |
| `pageheader.Meta`         | `<div>`  | Wrapping row of small facts (date, crew, amount).            |
| `pageheader.Actions`      | `<div>`  | The page's buttons. Wraps under Content below `sm`.          |

## List page

```go
@pageheader.PageHeader(pageheader.Props{}) {
    @pageheader.Content(pageheader.SubProps{}) {
        @pageheader.Title(pageheader.TitleProps{}) { Jobs }
        @pageheader.Description(pageheader.SubProps{}) { Every open job across both crews. }
    }
    @pageheader.Actions(pageheader.SubProps{}) {
        @button.Button(button.Props{Outline: true}) { Export }
        @button.Button(button.Props{Href: "/jobs/new"}) { New job }
    }
}
```

## Detail page

```go
@pageheader.PageHeader(pageheader.Props{}) {
    @pageheader.Back(pageheader.BackProps{Href: "/jobs"}) { Jobs }
    @pageheader.Content(pageheader.SubProps{}) {
        <div class="flex flex-wrap items-center gap-4">
            @pageheader.Title(pageheader.TitleProps{}) { Job #1834 }
            @badge.Badge(badge.Props{Variant: badge.VariantSuccess}) { Scheduled }
        </div>
        @pageheader.Meta(pageheader.SubProps{}) {
            <span><svg data-slot="icon" ...></svg> Apr 22, 2026</span>
            <span><svg data-slot="icon" ...></svg> Crew A</span>
        }
    }
    @pageheader.Actions(pageheader.SubProps{}) {
        @button.Button(button.Props{Outline: true}) { Reschedule }
        @button.Button(button.Props{}) { Send invoice }
    }
}
```

- **Badge beside the title**: wrap Title and the badge in a flex row (above). The badge stays outside the `<h1>`, so it isn't read as part of the title.
- **Meta items**: each child is laid out as an inline icon + text pair. A `data-slot="icon"` SVG is sized to 16px and tinted.
- **Controls under the title**: Content takes any children. Put a search input or filter row under the Title; Actions bottom-align with it.
- **Breadcrumbs instead of Back**: render `breadcrumbs.Breadcrumbs` inside a `<div class="basis-full">` as the first child.

## Props

### Props / SubProps

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Class` | `string`           | `""`    | Appended to the root element. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the root element. |

### TitleProps

| Field   | Type               | Default | Notes                                              |
| ------- | ------------------ | ------- | -------------------------------------------------- |
| `Level` | `heading.Level`    | `1`     | Set when the header sits under another `<h1>`.     |
| `Class` | `string`           | `""`    | Appended to the heading.                           |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the heading.                           |

### BackProps

| Field   | Type               | Default | Notes                         |
| ------- | ------------------ | ------- | ----------------------------- |
| `Href`  | `string`           | —       | Required. The parent page.    |
| `Class` | `string`           | `""`    | Appended to the row wrapper.  |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the `<a>`.        |

## Accessibility

- One `<h1>` per page. Inside a shell, `pageheader.Title` is normally it.
- The Back link's chevron is `aria-hidden`; the link text names the destination.
