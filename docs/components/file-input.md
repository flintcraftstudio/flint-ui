# File input

File pickers, built from scratch since Catalyst has none:

| Component             | Use for                                                                 |
| --------------------- | ----------------------------------------------------------------------- |
| `fileinput.Input`     | A file field inside a form: avatar, certificate, a single attachment. Lines up with `input.Input`. |
| `fileinput.Dropzone`  | Imports and bulk attachments. A large drag-and-drop area with a selected-file list. |

Both render a real `<input type="file">`, so a plain form submit works with no JS. Dropzone adds Alpine for the drag-over state, the file list, and drop support.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/fileinput"
```

## Uploading with htmx — read this first

**A request only carries files when it's multipart.** Unless the requesting element has `hx-encoding="multipart/form-data"`, or its `<form>` has `enctype="multipart/form-data"`, htmx 4 URL-encodes the body and **the files are silently dropped**.

```go
<form hx-post="/jobs/import" hx-encoding="multipart/form-data" hx-target="#import-result"
      hx-disable="find button[type=submit]">
    @fileinput.Dropzone(fileinput.DropzoneProps{Name: "import", ID: "import", Accept: ".csv,text/csv", Hint: "CSV up to 5 MB", Required: true})
    @button.Button(button.Props{Type: "submit", Indicator: true}) { Import jobs }
</form>
```

- **No upload progress.** htmx 4 sends requests with `fetch`, which doesn't report upload progress (the old `htmx:xhr:progress` event is gone). Show the button's `Indicator` while the request is in flight instead.
- **Errors swap.** htmx 4 swaps 4xx responses, so a handler can return `422` with an error message for `hx-target`.

### Upload as soon as files are chosen

Put the htmx attributes on the input itself via `Attrs`, and trigger on `change`. Dropped files fire `change` too:

```go
@fileinput.Dropzone(fileinput.DropzoneProps{
    Name: "photos", ID: "photos", Accept: "image/*", Multiple: true,
    Attrs: templ.Attributes{
        "hx-post":     "/jobs/1834/photos",
        "hx-trigger":  "change",
        "hx-encoding": "multipart/form-data",
        "hx-target":   "#photo-list",
    },
})
```

## Server side (Go)

```go
func importHandler(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, 5<<20) // hard cap on request size
    if err := r.ParseMultipartForm(1 << 20); err != nil {
        http.Error(w, "upload too large or not multipart", http.StatusUnprocessableEntity)
        return
    }
    f, hdr, err := r.FormFile("import") // or r.MultipartForm.File["photos"] for Multiple
    ...
}
```

`Accept` only filters the picker, and on Dropzone the drop. It's a convenience, not validation. Check the size, sniff the content (`http.DetectContentType`), and never trust `hdr.Filename` as a path.

## Input

```go
@fieldset.Field(fieldset.Props{}) {
    @fieldset.Label(fieldset.LabelProps{For: "avatar"}) { Profile photo }
    @fieldset.Description(fieldset.Props{}) { PNG or JPG, up to 2 MB. }
    @fileinput.Input(fileinput.Props{Name: "avatar", ID: "avatar", Accept: "image/png,image/jpeg"})
}
```

The browser's "Choose file" button is restyled as a muted chip. The "No file chosen" text and the chosen filename come from the browser.

## Dropzone

- **Click or keyboard**: the whole area is a `<label>` for the input, so clicking anywhere opens the picker. The input is visually hidden but focusable: Tab to it and press Space or Enter, and the focus ring is drawn on the area.
- **Drop**: browsers apply `accept` to the picker but not to drops. Dropzone filters dropped files against the same list (extensions like `.csv`, MIME types like `text/csv`, and wildcards like `image/*`), keeps the first file unless `Multiple` is set, and reports how many it skipped. Accepted files are written into the real input.
- **File list**: names and sizes appear under the area, with a Clear link. The list is a polite live region.
- **Required**: works natively. The browser blocks submit until a file is chosen.

## Props

### Props (Input)

| Field      | Type               | Default | Notes                                                    |
| ---------- | ------------------ | ------- | -------------------------------------------------------- |
| `Name`     | `string`           | `""`    | Form field name.                                         |
| `ID`       | `string`           | `""`    | Pair with `fieldset.Label{For}`.                         |
| `Accept`   | `string`           | `""`    | Native accept list: `.csv`, `image/*,.pdf`.              |
| `Multiple` | `bool`             | `false` | Allow several files.                                     |
| `Required` | `bool`             | `false` |                                                          |
| `Disabled` | `bool`             | `false` |                                                          |
| `Invalid`  | `bool`             | `false` | Danger border + `aria-invalid`.                          |
| `Class`    | `string`           | `""`    | Appended to the wrapper.                                 |
| `Attrs`    | `templ.Attributes` | `nil`   | Spread onto the `<input>`. Use for `hx-*`.               |

### DropzoneProps

All of the above, plus:

| Field    | Type     | Default                                   | Notes                                       |
| -------- | -------- | ----------------------------------------- | ------------------------------------------- |
| `ID`     | `string` | —                                         | **Required**: the label targets it.         |
| `Prompt` | `string` | `"Upload a file"` / `"Upload files"`      | Bold call to action, followed by "or drag and drop". |
| `Hint`   | `string` | `""`                                      | Muted line under it: "CSV up to 5 MB".      |

`Class` goes on the Dropzone's root; `Attrs` go on the `<input>`.

## Accessibility

- Both are native file inputs: keyboard and screen-reader behavior come from the browser.
- Dropzone's visible text ("Upload a file or drag and drop", plus the hint) is the input's label, so the hint is announced too.
- Selected files and skipped-file messages are announced through a polite live region.
