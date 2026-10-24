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
// panel button's open = false. Without History it never touches history.
// Nothing is logged to the console. Before v1.2.3 its state was inline in
// x-data and it never started under the CSP build.
func TestSlideoverUnderCSP(t *testing.T) {
	for _, c := range []struct {
		name  string
		steps []string
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
		{"no history", []string{
			`fire("flint:slideover-open", "item")`,
			`state("open"); log.push("added=" + (history.length - startLength) + " mark=" + mark())`,
		}, "open=open added=0 mark=none"},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := runSteps(t, slideover.Props{Name: "item"}, "", c.steps); got != c.want {
				t.Errorf("states %q, want %q", got, c.want)
			}
		})
	}
}

// TestSlideoverHistoryUnderCSP: with History, opening pushes one entry
// marked in history.state (keeping what was there), Back closes the panel,
// and every other close removes the entry. Opening and closing three times
// adds one entry, not three: each push replaces the forward entry. A panel
// reopened while its last close's back() is still in flight stays open
// once it lands, and its own close still removes its entry; closed again
// before it lands, it doesn't go back twice. Back onto an entry still
// marked as the panel's (the page pushed one of its own) leaves it open.
// A page loaded onto an entry still marked (a reload while open) has the
// mark removed, so Back closes a reopened panel on the first press.
func TestSlideoverHistoryUnderCSP(t *testing.T) {
	for _, c := range []struct {
		name   string
		before string // JavaScript run ahead of Alpine
		steps  []string
		want   string
	}{
		{"back closes", "", []string{
			`history.replaceState({ page: "kept" }, ""); fire("flint:slideover-open", "item")`,
			`state("open"); log.push("mark=" + mark() + " kept=" + history.state.page); history.back()`,
			`state("back"); log.push("mark=" + mark())`,
		}, "open=open mark=item kept=kept back=closed mark=none"},
		{"done removes the entry", "", []string{
			`fire("flint:slideover-open", "item")`,
			`document.getElementById("done").click()`,
			`state("done"); log.push("mark=" + mark() + " added=" + (history.length - startLength))`,
		}, "done=closed mark=none added=1"},
		{"escape removes the entry", "", []string{
			`fire("flint:slideover-open", "item")`,
			`window.dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }))`,
			`state("escape"); log.push("mark=" + mark())`,
		}, "escape=closed mark=none"},
		{"three times", "", []string{
			`fire("flint:slideover-open", "item")`, `document.getElementById("done").click()`,
			`fire("flint:slideover-open", "item")`, `fire("flint:slideover-close")`,
			`fire("flint:slideover-open", "item")`, `history.back()`,
			`state("end"); log.push("mark=" + mark() + " added=" + (history.length - startLength))`,
		}, "end=closed mark=none added=1"},
		{"reopened before back lands", "", []string{
			// The close's back() lands 300ms late, so the panel is reopened
			// while it is still in flight.
			`var back = history.back; history.back = function () { history.back = back; setTimeout(function () { history.back(); }, 300); };
			fire("flint:slideover-open", "item")`,
			`document.getElementById("done").click()`,
			`fire("flint:slideover-open", "item")`,
			`state("reopened"); log.push("mark=" + mark())`,
			`state("landed"); log.push("mark=" + mark()); document.getElementById("done").click()`,
			`state("done"); log.push("mark=" + mark() + " added=" + (history.length - startLength))`,
		}, "reopened=open mark=item landed=open mark=item done=closed mark=none added=1"},
		{"closed again before back lands", "", []string{
			// The close's back() lands 500ms late, after the second close. A
			// base entry of the page's own, so one back() too many would
			// leave it.
			`history.pushState({ page: "base" }, "");
			var back = history.back; history.back = function () { history.back = back; setTimeout(function () { history.back(); }, 500); };
			fire("flint:slideover-open", "item")`,
			`document.getElementById("done").click()`,
			`fire("flint:slideover-open", "item")`,
			`document.getElementById("done").click()`,
			`state("landed"); log.push("mark=" + mark() + " page=" + (history.state && history.state.page))`,
		}, "landed=closed mark=none page=base"},
		{"back to its own entry", "", []string{
			// The page pushes an entry of its own while the panel is open,
			// keeping the state; Back undoes the page's entry, not the panel.
			`fire("flint:slideover-open", "item")`,
			`history.pushState(Object.assign({}, history.state, { page: 2 }), "")`,
			`history.back()`,
			`state("back"); log.push("mark=" + mark())`,
		}, "back=open mark=item"},
		{"loaded onto a marked entry", `history.replaceState({ flintSlideover: "item", page: "kept" }, "");`, []string{
			// A reload while open left the page on the panel's old entry.
			`log.push("mark=" + mark() + " kept=" + history.state.page); fire("flint:slideover-open", "item")`,
			`history.back()`,
			`state("back"); log.push("mark=" + mark() + " kept=" + history.state.page)`,
		}, "mark=none kept=kept back=closed mark=none kept=kept"},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := runSteps(t, slideover.Props{Name: "item", History: true}, c.before, c.steps); got != c.want {
				t.Errorf("states %q, want %q", got, c.want)
			}
		})
	}
}

// runSteps renders a slide-over with p and two panel buttons, runs before
// (JavaScript) ahead of Alpine, then each step 200ms apart under CSP, and returns what the steps
// logged. Steps may call state(label), fire(event, detail), mark() (the
// history.state mark, or "none"), root(), and read startLength. It fails
// on any console error.
func runSteps(t *testing.T, p slideover.Props, before string, steps []string) string {
	t.Helper()
	body := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return slideover.Slideover(p).Render(templ.WithChildren(ctx, templ.Raw(
			`<button id="inside" type="button">Inside</button><button id="done" type="button" x-on:click="open = false">Done</button>`)), w)
	})
	check := before + `
document.addEventListener("alpine:initialized", function () {
	var log = [];
	var startLength = history.length;
	var root = function () { return document.querySelector("[data-flint-name=item]"); };
	var state = function (label) { log.push(label + "=" + (Alpine.$data(root()).open ? "open" : "closed")); };
	var mark = function () { return (history.state && history.state.flintSlideover) || "none"; };
	var fire = function (name, detail) {
		window.dispatchEvent(new CustomEvent(name, detail === undefined ? {} : { detail: detail }));
	};
	var steps = [` + "function () { " + strings.Join(steps, " },\n\t\tfunction () { ") + " }" + `,
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
	return m[1]
}
