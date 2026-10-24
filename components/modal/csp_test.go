package modal_test

import (
	"context"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/flintcraftstudio/flint-ui/components/modal"
	"github.com/flintcraftstudio/flint-ui/internal/csptest"
)

// TestModalUnderCSP: under Alpine's CSP build and a strict CSP, a modal
// opens on its own name (bare, or wrapped as htmx's {value: name}) and not
// on another's, and closes on a close event with no name or its own name
// (not another's), Escape, a click outside the panel and a panel button's
// open = false. Opened, and reopened, it moves focus into the panel.
// Nothing is logged to the console. Before v1.2.6 its state was inline in
// x-data and it never started under the CSP build, and its trap could
// start before x-show had shown the panel, leaving focus outside.
func TestModalUnderCSP(t *testing.T) {
	for _, c := range []struct {
		name  string
		steps []string
		want  string
	}{
		{"by name", []string{
			`state("start"); fire("flint:modal-open", "other")`,
			`state("other"); fire("flint:modal-open", "item")`,
			`state("item"); fire("flint:modal-close", "other")`,
			`state("close-other")`,
		}, "start=closed other=closed item=open close-other=open"},
		{"htmx detail", []string{
			`fire("flint:modal-open", { value: "item" })`,
			`state("open"); fire("flint:modal-close", { value: "item" })`,
			`state("close")`,
		}, "open=open close=closed"},
		{"close all", []string{
			`fire("flint:modal-open", "item")`,
			`fire("flint:modal-close")`,
			`state("close")`,
		}, "close=closed"},
		{"escape", []string{
			`fire("flint:modal-open", "item")`,
			`window.dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }))`,
			`state("escape")`,
		}, "escape=closed"},
		{"outside the panel", []string{
			`fire("flint:modal-open", "item")`,
			`root().querySelector("[role=dialog]").click()`,
			`state("outside")`,
		}, "outside=closed"},
		{"panel button", []string{
			`fire("flint:modal-open", "item")`,
			`document.getElementById("done").click()`,
			`state("done")`,
		}, "done=closed"},
		{"focus", []string{
			`fire("flint:modal-open", "item")`,
			`log.push("focus=" + document.activeElement.id); document.getElementById("done").click()`,
			`fire("flint:modal-open", "item")`,
			`log.push("reopened=" + document.activeElement.id)`,
		}, "focus=inside reopened=inside"},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := runSteps(t, c.steps); got != c.want {
				t.Errorf("states %q, want %q", got, c.want)
			}
		})
	}
}

// runSteps renders a modal named item with two panel buttons, then runs
// each step 200ms apart under CSP and returns what the steps logged.
// Steps may call state(label), fire(event, detail) and root(). It fails on
// any console error.
func runSteps(t *testing.T, steps []string) string {
	t.Helper()
	body := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return modal.Modal(modal.Props{Name: "item"}).Render(templ.WithChildren(ctx, templ.Raw(
			`<button id="inside" type="button">Inside</button><button id="done" type="button" x-on:click="open = false">Done</button>`)), w)
	})
	check := `
document.addEventListener("alpine:initialized", function () {
	var log = [];
	var root = function () { return document.querySelector("[data-flint-name=item]"); };
	var state = function (label) { log.push(label + "=" + (Alpine.$data(root()).open ? "open" : "closed")); };
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
