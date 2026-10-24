package slideover_test

import (
	"context"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/flintcraftstudio/flint-ui/components/slideover"
	"github.com/flintcraftstudio/flint-ui/internal/csptest"
)

// TestSlideoverUnderCSP: under Alpine's CSP build and a strict CSP, a
// slide-over opens on its own name (bare, or wrapped as htmx's
// {value: name}) and not on another's, and closes on a close event with
// no name or its own name (not another's), Escape, a backdrop click and a
// panel button's open = false. Nothing is logged to the console. Before
// v1.2.3 its state was inline in x-data and it never started under the
// CSP build.
//
// Each case is its own page with at most one open and one close:
// headless Chromium's virtual time runs almost no animation frames, so a
// transition never finishes, and reopening during a stalled leave makes
// Alpine reject the cancelled transition's promise. In a real browser
// the same sequences log nothing.
func TestSlideoverUnderCSP(t *testing.T) {
	body := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return slideover.Slideover(slideover.Props{Name: "item"}).Render(templ.WithChildren(ctx, templ.Raw(
			`<button id="inside" type="button">Inside</button><button id="done" type="button" x-on:click="open = false">Done</button>`)), w)
	})
	for _, c := range []struct {
		name  string
		steps []string // JavaScript, run 200ms apart; each may call state(label)
		want  string
	}{
		{"by name", []string{
			`state("start"); fire("flint:slideover-open", "other")`,
			`state("other"); fire("flint:slideover-open", "item")`,
			`state("item"); fire("flint:slideover-close", "other")`,
			`state("close-other")`,
		}, "start=closed other=closed item=open close-other=open"},
		{"htmx detail", []string{
			`fire("flint:slideover-open", { value: "item" })`,
			`state("open"); fire("flint:slideover-close", { value: "item" })`,
			`state("close")`,
		}, "open=open close=closed"},
		{"close all", []string{
			`fire("flint:slideover-open", "item")`,
			`fire("flint:slideover-close")`,
			`state("close")`,
		}, "close=closed"},
		{"escape", []string{
			`fire("flint:slideover-open", "item")`,
			`window.dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }))`,
			`state("escape")`,
		}, "escape=closed"},
		{"backdrop", []string{
			`fire("flint:slideover-open", "item")`,
			`root().querySelector("[aria-hidden=true]").click()`,
			`state("backdrop")`,
		}, "backdrop=closed"},
		{"panel button", []string{
			`fire("flint:slideover-open", "item")`,
			`document.getElementById("done").click()`,
			`state("done")`,
		}, "done=closed"},
	} {
		t.Run(c.name, func(t *testing.T) {
			check := `
document.addEventListener("alpine:initialized", function () {
	var log = [];
	var root = function () { return document.querySelector("[data-flint-name=item]"); };
	var state = function (label) { log.push(label + "=" + (Alpine.$data(root()).open ? "open" : "closed")); };
	var fire = function (name, detail) {
		window.dispatchEvent(new CustomEvent(name, detail === undefined ? {} : { detail: detail }));
	};
	var steps = [` + "function () { " + strings.Join(c.steps, " },\n\t\tfunction () { ") + " }" + `,
		function () {
			var pre = document.createElement("pre");
			pre.id = "result";
			pre.textContent = log.join(" ");
			document.body.appendChild(pre);
		}];
	var i = 0;
	var next = function () { steps[i++](); if (i < steps.length) setTimeout(next, 200); };
	setTimeout(next, 100);
});`
			res := csptest.Run(t, body, check)
			if len(res.Console) > 0 {
				t.Errorf("console errors under CSP:\n%s", strings.Join(res.Console, "\n"))
			}
			m := regexp.MustCompile(`<pre id="result">([^<]*)</pre>`).FindStringSubmatch(res.DOM)
			if m == nil {
				t.Fatal("the check didn't finish")
			}
			if m[1] != c.want {
				t.Errorf("states %q, want %q", m[1], c.want)
			}
		})
	}
}
