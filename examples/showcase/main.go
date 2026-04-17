package main

import (
	"html"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"

	"github.com/flintcraft/flint-ui/examples/showcase/templates"
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
	mux.Handle("GET /checkboxes", page("Checkbox", "checkboxes", templates.Checkboxes()))
	mux.Handle("GET /badges", page("Badge", "badges", templates.Badges()))
	mux.Handle("GET /tables", page("Table", "tables", templates.Tables()))
	mux.Handle("GET /headings", page("Heading", "headings", templates.Headings()))
	mux.Handle("GET /modals", page("Modal", "modals", templates.Modals()))
	mux.Handle("GET /dropdowns", page("Dropdown", "dropdowns", templates.Dropdowns()))
	mux.Handle("GET /tabs", page("Tabs", "tabs", templates.Tabs()))
	mux.Handle("GET /tabs/job", http.HandlerFunc(tabsJobHandler))
	mux.Handle("GET /toasts", page("Toast", "toasts", templates.Toasts()))
	mux.Handle("POST /toasts/echo", http.HandlerFunc(toastsEchoHandler))

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
	w.Header().Set("HX-Trigger", `{"flint:toast":{"variant":"success","title":"Saved on the server","body":"HX-Trigger fired — no client-side dispatch code."}}`)
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
