# Patterns

How flint-ui components combine into the screens a client admin tool needs. Each recipe is taken from working code in the showcase (`examples/showcase/`), which you can run with `mage Showcase`. Open the per-component docs in [`docs/components/`](./components/) for the full prop surface.

Every recipe assumes **htmx 4**: put `hx-target` / `hx-swap` on the element that makes the request (inheritance is off), and remember that 4xx/5xx responses swap like 200s. See [htmx 4 notes](../README.md#htmx-4).

| I need… | Recipe |
| --- | --- |
| A new admin screen | [Page skeleton](#page-skeleton) |
| A searchable, sortable list with bulk actions | [List page](#list-page) |
| A form that shows server-side errors | [Form with validation](#form-with-validation) |
| Create or edit without leaving the list | [Form in a modal or slide-over](#form-in-a-modal-or-slide-over) |
| "Are you sure?" before deleting | [Confirm a destructive action](#confirm-a-destructive-action) |
| Tell the user it worked | [Feedback toasts from the server](#feedback-toasts-from-the-server) |
| Show that something is happening | [Loading states](#loading-states) |

---

## Page skeleton

Base layout once per app: the shell, the toast container, and (optionally) dark mode. Pages render into it and open with a `pageheader`.

```go
templ Base(title string, nav templ.Component, page templ.Component) {
    <!DOCTYPE html>
    <html lang="en" class="bg-background text-foreground antialiased lg:bg-canvas">
        <head>
            <meta charset="UTF-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1"/>
            <title>{ title }</title>
            <link rel="stylesheet" href="/static/css/app.css"/>
            <script src="/static/js/htmx.min.js" defer></script>
            // Alpine plugins load before Alpine itself. Focus is required
            // (modal, slide-over, command palette, mobile drawer use x-trap);
            // Collapse only if you use accordion.
            <script src="/static/js/alpine-focus.min.js" defer></script>
            <script src="/static/js/alpine-collapse.min.js" defer></script>
            <script src="/static/js/alpine.min.js" defer></script>
        </head>
        <body x-data>
            @layout.Sidebar(layout.SidebarProps{Navbar: mobileNavbar(), Sidebar: nav}) {
                @page
            }
            @toast.Container(toast.ContainerProps{})
        </body>
    </html>
}

templ JobsPage(jobs []Job) {
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
    <div class="mt-8">
        if len(jobs) == 0 {
            @emptystate.EmptyState(emptystate.Props{Bordered: true}) {
                @emptystate.Title(emptystate.TitleProps{Level: 2}) { No jobs yet }
                @emptystate.Description(emptystate.SubProps{}) { New jobs show up here. }
                @emptystate.Actions(emptystate.SubProps{}) {
                    @button.Button(button.Props{Href: "/jobs/new"}) { New job }
                }
            }
        } else {
            @JobsList(jobs)
        }
    </div>
}
```

- `body x-data` gives every Alpine-driven component a common ancestor; modals, slide-overs and toasts listen for window events from it.
- `lg:bg-canvas` on `<html>` makes the content panel read as a card at `lg+`, in light and dark.
- Shell choice: `layout.Sidebar` (left rail), `layout.Stacked` (top navbar), `layout.Auth` (login). Showcase: `appshell.templ`, `/demo/sidebar`.
- Dark mode is opt-in: add `data-theme="system"` to `<html>`. For a toggle, see [Dark mode](../README.md#dark-mode).

## List page

Search, sort, select, bulk actions, and row menus, all server-rendered. The search box sits outside the swapped fragment so it keeps focus; sort links work without JS and hx-get just swaps the table. Showcase: `tables_admin.templ`, `/tables`.

```go
templ JobsList(q JobsQuery) {
    @table.Selectable(table.SelectableProps{}) {
        @input.InputGroup(input.GroupProps{}) {
            @iconSearch()
            @input.Input(input.Props{Type: "search", Name: "q", Value: q.Q, Attrs: templ.Attributes{
                "aria-label": "Search jobs",
                "hx-get":     "/jobs/table",
                "hx-trigger": "input changed delay:300ms, search",
                "hx-target":  "#jobs-table",
                "hx-swap":    "outerHTML",
            }})
        }
        @table.BulkActions(table.BulkActionsProps{}) {
            @button.Button(button.Props{Variant: button.VariantDanger, Indicator: true,
                Attrs: table.BulkAttrs(templ.Attributes{"hx-post": "/jobs/delete", "hx-swap": "none"})}) { Delete }
        }
        @JobsTable(q) // <div id="jobs-table"> … </div>, also what /jobs/table returns
    }
}
```

- **Sorting:** `table.SortHeader` with `Href` (full-page fallback) and `LinkAttrs` (hx-get the fragment). `table.ParseSortDirection` and `SortDirection.Next()` build the URLs.
- **Selection:** `SelectAll` in the header row, `SelectCell{Value: id}` per row. `BulkAttrs` makes the buttons submit the checked ids as `ids` and disables them while nothing is selected.
- **Empty results:** render an `emptystate.EmptyState` inside a single full-width row (`colspan`) so the table chrome stays put.
- **Row menus:** `dropdown.Menu` with `Fixed: true` so it escapes the table's `overflow-x-auto`; trigger it with `dropdown.TriggerAttrs` on a plain button.
- **After a bulk action:** respond `204` with `HX-Trigger` carrying `flint:toast` and `flint:selection-clear` (see [toasts](#feedback-toasts-from-the-server)).

## Form with validation

Validate on the server and re-render the form with errors. htmx 4 swaps a `422` response, so the handler returns the same component with error state and the user's input intact.

```go
templ JobForm(f JobForm, errs map[string]string) {
    <form id="job-form" hx-post="/jobs" hx-target="this" hx-swap="outerHTML" hx-disable="find button">
        @fieldset.FieldGroup(fieldset.Props{}) {
            @fieldset.Field(fieldset.Props{}) {
                @fieldset.Label(fieldset.LabelProps{For: "job-client"}) { Client }
                @input.Input(input.Props{Name: "client", ID: "job-client", Value: f.Client,
                    Invalid: errs["client"] != "",
                    Attrs:   templ.Attributes{"aria-describedby": "job-client-error"}})
                if msg := errs["client"]; msg != "" {
                    @fieldset.ErrorMessage(fieldset.Props{Attrs: templ.Attributes{"id": "job-client-error"}}) { { msg } }
                }
            }
        }
        @button.Button(button.Props{Type: "submit", Indicator: true, Class: "mt-8"}) { Save job }
    </form>
}
```

```go
func createJob(w http.ResponseWriter, r *http.Request) {
    f, errs := parseJobForm(r)
    if len(errs) > 0 {
        w.WriteHeader(http.StatusUnprocessableEntity) // swapped in place by htmx 4
        views.JobForm(f, errs).Render(r.Context(), w)
        return
    }
    // save, then respond with the next view or a toast
}
```

- `Invalid` sets the red border and `aria-invalid="true"`; `aria-describedby` makes screen readers read the error with the field.
- `Indicator` on the submit button shows a spinner while the form is in flight (the `htmx-request` class lands on the `<form>`), and `hx-disable="find button"` blocks double submits.
- The same pattern works for `selectbox`, `textarea`, `checkbox`, `radio` and `toggle`; each has an `Invalid` prop.

## Form in a modal or slide-over

Render the form inside `modal.Modal` (or `slideover.Slideover` for longer forms). On success the server closes it with a header; on a validation failure it returns `422` with the form, which swaps in place and the dialog stays open.

```go
@button.Button(button.Props{Attrs: templ.Attributes{
    "x-on:click": "$dispatch('flint:modal-open', 'new-job')",
}}) { New job }

@modal.Modal(modal.Props{Name: "new-job"}) {
    @modal.Title(modal.SubProps{}) { New job }
    @modal.Body(modal.SubProps{}) {
        @JobForm(JobForm{}, nil)
    }
}
```

```go
// success path
w.Header().Set("HX-Trigger", `{"flint:modal-close":"new-job","flint:toast":{"variant":"success","title":"Job created"}}`)
```

- Slide-over: the same with `flint:slideover-open` / `flint:slideover-close`.
- Close everything with `{"flint:modal-close": true}`. Never send `null`: htmx 4's `htmx.trigger` throws on it.
- A cancel button inside the panel: `x-on:click="open = false"`.

## Confirm a destructive action

Use the modal's `Alert` layout (smaller, centered, lighter backdrop) for yes/no questions. `hx-confirm` (the browser dialog) is fine for internal tools; the Alert modal is the polished version.

```go
@modal.Modal(modal.Props{Name: "delete-job-1834", Alert: true}) {
    @modal.Title(modal.SubProps{Alert: true}) { Delete job #1834? }
    @modal.Description(modal.SubProps{Alert: true}) { This removes it from the crew's schedule. }
    @modal.Actions(modal.SubProps{Alert: true}) {
        @button.Button(button.Props{Plain: true, Attrs: templ.Attributes{"x-on:click": "open = false"}}) { Cancel }
        @button.Button(button.Props{Variant: button.VariantDanger, Indicator: true, Attrs: templ.Attributes{
            "hx-delete": "/jobs/1834",
            "hx-swap":   "none",
        }}) { Delete }
    }
}
```

- Sub-components need `Alert: true` to match the parent.
- The handler responds `204` with `HX-Trigger: {"flint:modal-close": true, "flint:toast": {...}}`, plus whatever refreshes the list.

## Feedback toasts from the server

Render `toast.Container` once in the base layout. Any response can raise a toast through `HX-Trigger`, with no client code:

```go
w.Header().Set("HX-Trigger", headerJSON(map[string]any{
    "flint:toast": map[string]any{"variant": "success", "title": "Archived 3 jobs", "body": "#1834, #1833, #1829"},
}))
w.WriteHeader(http.StatusNoContent)
```

- Payload: `variant` (`info` / `success` / `warning` / `danger`), `title`, optional `body`, optional `duration` in ms (`0` keeps it until dismissed).
- **Escape non-ASCII** in header values (`\uXXXX`): browsers decode headers as Latin-1, so an em dash or accented name arrives garbled. Copy `headerJSON` from `examples/showcase/main.go`.
- From Alpine instead: `$dispatch('flint:toast', {variant: 'success', title: 'Saved'})`.

## Loading states

Pick the smallest indicator that tells the truth:

| Situation | Use |
| --- | --- |
| A button triggered the request | `button.Props{Indicator: true}` (+ `hx-disable="this"` to block repeats) |
| A section is loading | `spinner.Spinner` or an indeterminate `progress.Progress` with `Class: "htmx-indicator"` and `hx-indicator="#id"` on the trigger |
| Lazy content on page load | a placeholder with `hx-get` + `hx-trigger="load"` that swaps itself (showcase `/stats`) |
| A long job with known progress | `progress.Progress` polled with `hx-trigger="every 1s"` + `hx-swap="outerMorph"`; the final response drops the `hx-*` attributes to stop polling (showcase `/progress`) |

File uploads can't report progress through htmx 4 (`fetch` has no upload progress). Use the button indicator.

---

When a pattern here doesn't fit, compose from the per-component docs and keep to the rules in the README: semantic color tokens only, Alpine for local UI state only, `Attrs` for every `hx-*` attribute.
