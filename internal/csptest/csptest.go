// Package csptest renders flint-ui components in headless Chromium under a
// strict Content-Security-Policy and Alpine's CSP build, the way a client
// with no 'unsafe-eval' or inline scripts runs them. A component passes
// when the page logs no console errors and its check script finds what it
// expects in the DOM.
//
// Tests using it are skipped when no chromium binary is on PATH.
package csptest

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/a-h/templ"

	flintui "github.com/flintcraftstudio/flint-ui"
)

// CSP is a client's strict policy: scripts and styles only from the page's
// origin, no inline code, no eval.
const CSP = "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data: blob:; object-src 'none'; base-uri 'self'"

//go:embed testdata/alpine-csp-3.17.4.min.js
var alpineCSP []byte

// Result is what the browser left behind.
type Result struct {
	DOM     string   // the document after scripts ran
	Console []string // console messages at error level, and CSP violations
}

// Run serves body inside a page with flint-ui's script, the given check
// script (plain JavaScript, served as a file) and Alpine's CSP build, loads
// it in headless Chromium, and returns the DOM and any console errors.
func Run(t *testing.T, body templ.Component, check string) Result {
	t.Helper()
	bin, err := exec.LookPath("chromium")
	if err != nil {
		t.Skip("chromium not on PATH")
	}

	mux := http.NewServeMux()
	file := func(path, ctype string, b []byte) {
		mux.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", ctype)
			w.Write(b)
		})
	}
	js := "text/javascript; charset=utf-8"
	file("/flint-ui.js", js, flintui.Script())
	file("/check.js", js, []byte(check))
	file("/alpine.js", js, alpineCSP)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", CSP)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var buf bytes.Buffer
		if err := page(body).Render(r.Context(), &buf); err != nil {
			t.Error(err)
		}
		w.Write(buf.Bytes())
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin,
		"--headless=new", "--no-sandbox", "--disable-gpu", "--disable-extensions",
		"--user-data-dir="+filepath.Join(t.TempDir(), "profile"),
		"--enable-logging=stderr", "--v=0",
		"--virtual-time-budget=3000",
		"--dump-dom", srv.URL+"/")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("chromium: %v\n%s", err, stderr.String())
	}
	return Result{DOM: stdout.String(), Console: consoleErrors(stderr.String())}
}

// page is the minimal document: flint-ui's script and the check script
// before Alpine (deferred scripts run in order), and nothing inline.
func page(body templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if _, err := w.Write([]byte(`<!doctype html><html lang="en"><head><meta charset="utf-8">` +
			`<script src="/flint-ui.js" defer></script><script src="/check.js" defer></script>` +
			`<script src="/alpine.js" defer></script></head><body x-data>`)); err != nil {
			return err
		}
		if err := body.Render(ctx, w); err != nil {
			return err
		}
		_, err := w.Write([]byte(`</body></html>`))
		return err
	})
}

// Chromium writes console messages to stderr as
// [...:INFO:CONSOLE:<level>] "message", source: url (line). Level 1 is
// console.log/info from the page, 2 warning, 3 and up errors; uncaught
// exceptions and CSP violations come through as errors too. Alpine reports
// expression errors with console.error-style messages, so anything
// mentioning Alpine counts regardless of level.
var consoleLine = regexp.MustCompile(`:CONSOLE(?::(\d+))?\] "([^"]*)`)

func consoleErrors(stderr string) []string {
	var out []string
	for _, line := range strings.Split(stderr, "\n") {
		m := consoleLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		msg := m[2]
		if m[1] >= "3" || strings.Contains(msg, "Alpine") || strings.Contains(msg, "Uncaught") ||
			strings.Contains(msg, "Content Security Policy") || strings.Contains(msg, "Refused") {
			out = append(out, msg)
		}
	}
	return out
}
