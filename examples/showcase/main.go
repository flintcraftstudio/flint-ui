package main

import (
	"log/slog"
	"net/http"
	"os"

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

	mux.Handle("POST /echo", http.HandlerFunc(echoHandler))

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
	_, _ = w.Write([]byte(`<span class="text-emerald-600">pong</span>`))
}
