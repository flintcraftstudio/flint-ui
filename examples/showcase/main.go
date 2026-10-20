package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"

	"github.com/a-h/templ"

	"github.com/flintcraftstudio/flint-ui/examples/showcase/templates"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("examples/showcase/static"))))

	mux.Handle("GET /{$}", page("Overview", "", templates.Index()))
	mux.Handle("GET /buttons", page("Button", "buttons", templates.Buttons()))
	mux.Handle("GET /inputs", page("Input", "inputs", templates.Inputs()))
	mux.Handle("GET /selects", page("Select", "selects", templates.Selects()))
	mux.Handle("GET /textareas", page("Textarea", "textareas", templates.Textareas()))
	mux.Handle("GET /file-inputs", page("File input", "file-inputs", templates.FileInputs()))
	mux.Handle("POST /files/upload", http.HandlerFunc(filesUploadHandler))
	mux.Handle("GET /checkboxes", page("Checkbox", "checkboxes", templates.Checkboxes()))
	mux.Handle("GET /badges", page("Badge", "badges", templates.Badges()))
	mux.Handle("GET /avatars", page("Avatar", "avatars", templates.Avatars()))
	mux.Handle("GET /radios", page("Radio", "radios", templates.Radios()))
	mux.Handle("GET /switches", page("Switch", "switches", templates.Switches()))
	mux.Handle("GET /dividers", page("Divider", "dividers", templates.Dividers()))
	mux.Handle("GET /description-lists", page("Description List", "description-lists", templates.DescriptionLists()))
	mux.Handle("GET /tables", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page("Table", "tables", templates.Tables(templates.ParseJobsQuery(r.URL.Query()))).ServeHTTP(w, r)
	}))
	mux.Handle("GET /tables/jobs", http.HandlerFunc(tablesJobsHandler))
	mux.Handle("POST /tables/jobs/archive", bulkJobsHandler("Archived"))
	mux.Handle("POST /tables/jobs/delete", bulkJobsHandler("Deleted"))
	mux.Handle("GET /headings", page("Heading", "headings", templates.Headings()))
	mux.Handle("GET /texts", page("Text", "texts", templates.Texts()))
	mux.Handle("GET /cards", page("Card", "cards", templates.Cards()))
	mux.Handle("GET /alerts", page("Alert", "alerts", templates.Alerts()))
	mux.Handle("GET /modals", page("Modal", "modals", templates.Modals()))
	mux.Handle("GET /dropdowns", page("Dropdown", "dropdowns", templates.Dropdowns()))
	mux.Handle("GET /tabs", page("Tabs", "tabs", templates.Tabs()))
	mux.Handle("GET /tabs/job", http.HandlerFunc(tabsJobHandler))
	mux.Handle("GET /toasts", page("Toast", "toasts", templates.Toasts()))
	mux.Handle("POST /toasts/echo", http.HandlerFunc(toastsEchoHandler))
	mux.Handle("GET /tooltips", page("Tooltip", "tooltips", templates.Tooltips()))
	mux.Handle("GET /accordions", page("Accordion", "accordions", templates.Accordions()))
	mux.Handle("GET /slideovers", page("Slide-over", "slideovers", templates.Slideovers()))
	mux.Handle("GET /clipboards", page("Clipboard", "clipboards", templates.Clipboards()))
	mux.Handle("GET /popovers", page("Popover", "popovers", templates.Popovers()))
	mux.Handle("GET /paginations", page("Pagination", "paginations", templates.Paginations()))
	mux.Handle("GET /breadcrumbs", page("Breadcrumbs", "breadcrumbs", templates.Breadcrumbs()))
	mux.Handle("GET /comboboxes", page("Combobox", "comboboxes", templates.Comboboxes()))
	mux.Handle("GET /commands", page("Command Palette", "commands", templates.Commands()))

	mux.Handle("GET /layouts", page("Layouts", "layouts", templates.Layouts()))
	mux.Handle("GET /navbars", page("Navbar", "navbars", templates.Navbars()))
	mux.Handle("GET /sidebars", page("Sidebar", "sidebars", templates.Sidebars()))
	mux.Handle("GET /page-headers", page("Page header", "page-headers", templates.PageHeaders()))
	mux.Handle("GET /empty-states", page("Empty state", "empty-states", templates.EmptyStates()))
	mux.Handle("GET /stats", page("Stat", "stats", templates.Stats()))
	mux.Handle("GET /stats/live", http.HandlerFunc(statsLiveHandler))
	mux.Handle("GET /spinners", page("Spinner", "spinners", templates.Spinners()))
	mux.Handle("POST /slow", http.HandlerFunc(slowHandler))

	mux.Handle("GET /demo/sidebar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current, title := templates.DemoNavTitle(r.URL.Query().Get("nav"))
		bare(w, r, title, templates.SidebarDemo(current, title))
	}))
	mux.Handle("GET /demo/stacked", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current, title := templates.DemoNavTitle(r.URL.Query().Get("nav"))
		bare(w, r, title, templates.StackedDemo(current, title))
	}))
	mux.Handle("GET /demo/auth", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bare(w, r, "Sign in", templates.AuthDemo())
	}))

	mux.Handle("GET /tables/detail", http.HandlerFunc(tablesDetailHandler))

	mux.Handle("POST /echo", http.HandlerFunc(echoHandler))
	mux.Handle("POST /inputs/submit", http.HandlerFunc(inputsSubmitHandler))
	mux.Handle("POST /selects/submit", http.HandlerFunc(selectsSubmitHandler))
	mux.Handle("POST /textareas/submit", http.HandlerFunc(textareasSubmitHandler))
	mux.Handle("POST /checkboxes/submit", http.HandlerFunc(checkboxesSubmitHandler))

	slog.Info("flint-ui showcase listening", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}

func page(title, current string, body templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := templates.Layout(title, current, body).Render(r.Context(), w); err != nil {
			slog.Error("render error", "path", r.URL.Path, "err", err)
		}
	})
}

// bare renders an app-shell demo as a full page without the showcase chrome.
func bare(w http.ResponseWriter, r *http.Request, title string, body templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.Bare(title, body).Render(r.Context(), w); err != nil {
		slog.Error("render error", "path", r.URL.Path, "err", err)
	}
}

// statsLiveHandler serves the lazy-loaded stat grid after a delay, so the
// loading spinner is visible.
func statsLiveHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(1500 * time.Millisecond):
	case <-r.Context().Done():
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.StatsLive().Render(r.Context(), w); err != nil {
		slog.Error("render error", "path", r.URL.Path, "err", err)
	}
}

// filesUploadHandler lists the files a multipart upload carried. It reads
// only the part headers and sizes; nothing is stored.
func filesUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 20<<20)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := r.ParseMultipartForm(8 << 20); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`<span class="text-danger">No files received — is the request multipart? (` + html.EscapeString(err.Error()) + `)</span>`))
		return
	}
	var b strings.Builder
	for field, headers := range r.MultipartForm.File {
		for _, h := range headers {
			fmt.Fprintf(&b, `<li><span class="font-medium">%s</span> <span class="text-muted-foreground">%s · %d bytes · field “%s”</span></li>`,
				html.EscapeString(h.Filename), html.EscapeString(h.Header.Get("Content-Type")), h.Size, html.EscapeString(field))
		}
	}
	if b.Len() == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`<span class="text-danger">Choose a file first.</span>`))
		return
	}
	time.Sleep(600 * time.Millisecond) // long enough to see the button indicator
	_, _ = w.Write([]byte(`<p class="text-success">Received:</p><ul class="mt-1 space-y-1">` + b.String() + `</ul>`))
}

// slowHandler sleeps long enough to see loading states (spinners,
// button.Props.Indicator) before responding.
func slowHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(1500 * time.Millisecond):
	case <-r.Context().Done():
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<span class="text-success">done</span>`))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<span class="text-success">pong</span>`))
}

func inputsSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<span class="text-danger">form parse error</span>`))
		return
	}
	name := html.EscapeString(r.FormValue("name"))
	email := html.EscapeString(r.FormValue("email"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(
		`<span class="text-success">received: ` + name + ` &lt;` + email + `&gt;</span>`,
	))
}

func selectsSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<span class="text-danger">form parse error</span>`))
		return
	}
	region := html.EscapeString(r.FormValue("region"))
	currency := html.EscapeString(r.FormValue("currency"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(
		`<span class="text-success">saved: ` + region + ` / ` + currency + `</span>`,
	))
}

func textareasSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<span class="text-danger">form parse error</span>`))
		return
	}
	subject := html.EscapeString(r.FormValue("subject"))
	body := r.FormValue("body")
	chars := len(body)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(
		`<span class="text-success">sent: ` + subject + ` (` + strconv.Itoa(chars) + ` chars)</span>`,
	))
}

// tablesJobsHandler re-renders the admin-list table fragment for sort and
// search swaps.
func tablesJobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.JobsTable(templates.ParseJobsQuery(r.URL.Query())).Render(r.Context(), w); err != nil {
		slog.Error("render error", "path", r.URL.Path, "err", err)
	}
}

// headerJSON marshals v for an HX-Trigger header. Browsers decode header
// values as Latin-1, so any non-ASCII rune (an em dash, an accented name)
// must travel as a JSON \uXXXX escape or it arrives as mojibake.
func headerJSON(v any) string {
	b, _ := json.Marshal(v)
	var out strings.Builder
	for _, r := range string(b) {
		if r < 0x80 {
			out.WriteRune(r)
			continue
		}
		for _, u := range utf16.Encode([]rune{r}) {
			fmt.Fprintf(&out, "\\u%04x", u)
		}
	}
	return out.String()
}

// bulkJobsHandler answers a bulk (or single-row) action with a toast and
// a selection-clear event. The demo data is static, so nothing changes.
func bulkJobsHandler(verb string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}
		ids := r.Form["ids"]
		label := "job"
		if len(ids) != 1 {
			label = "jobs"
		}
		trigger := headerJSON(map[string]any{
			"flint:toast": map[string]string{
				"variant": "success",
				"title":   verb + " " + strconv.Itoa(len(ids)) + " " + label,
				"body":    "#" + strings.Join(ids, ", #") + " — demo only, nothing changed.",
			},
			"flint:selection-clear": true,
		})
		w.Header().Set("HX-Trigger", trigger)
		w.WriteHeader(http.StatusNoContent)
	})
}

func tablesDetailHandler(w http.ResponseWriter, r *http.Request) {
	name := html.EscapeString(r.URL.Query().Get("name"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if name == "" {
		_, _ = w.Write([]byte(`<span class="text-muted-foreground">no lead selected</span>`))
		return
	}
	_, _ = w.Write([]byte(
		`<span class="text-success">loaded detail for: ` + name + `</span>`,
	))
}

// toastsEchoHandler demonstrates HX-Trigger-driven toasts. The server
// returns a JSON header naming the flint:toast event with the payload
// inline; htmx dispatches it client-side and the page-level toast
// Container picks it up. The response body itself is just a small
// confirmation that gets swapped into #toast-result.
func toastsEchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Trigger", `{"flint:toast":{"variant":"success","title":"Saved on the server","body":"HX-Trigger fired \u2014 no client-side dispatch code."}}`)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<span class="text-success">server responded — toast fired via HX-Trigger</span>`))
}

// tabsJobHandler re-renders the server-mode tabs section for the htmx
// showcase. Each tab click hits this with ?tab=X; the handler returns
// the whole JobTabsServer fragment so htmx can outerHTML-swap it into
// the same #server-tabs-root container.
func tabsJobHandler(w http.ResponseWriter, r *http.Request) {
	tab := r.URL.Query().Get("tab")
	switch tab {
	case "overview", "crew", "billing":
	default:
		tab = "overview"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.JobTabsServer(tab).Render(r.Context(), w); err != nil {
		slog.Error("render error", "path", r.URL.Path, "err", err)
	}
}

func checkboxesSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<span class="text-danger">form parse error</span>`))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	sections := r.Form["sections"]
	if len(sections) == 0 {
		_, _ = w.Write([]byte(`<span class="text-muted-foreground">nothing selected</span>`))
		return
	}
	joined := ""
	for i, s := range sections {
		if i > 0 {
			joined += ", "
		}
		joined += html.EscapeString(s)
	}
	_, _ = w.Write([]byte(
		`<span class="text-success">saved: ` + joined + `</span>`,
	))
}
