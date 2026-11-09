package fab_test

import (
	"context"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/flintcraftstudio/flint-ui/components/fab"
	"github.com/flintcraftstudio/flint-ui/internal/csptest"
)

// TestMenuUnderCSP: under Alpine's CSP build and a strict CSP, the menu
// opens from its button, and closes on a tap outside, on Escape (focus
// back on the button), and on a tap on an item, whose own click handler
// still runs. A tap inside the menu but not on an item leaves it open.
// Nothing is logged to the console.
func TestMenuUnderCSP(t *testing.T) {
	for _, c := range []struct {
		name  string
		steps []string
		want  string
	}{
		{"opens", []string{
			`state("start"); button().click()`,
			`state("tapped")`,
		}, "start=closed tapped=open"},
		{"outside", []string{
			`button().click()`,
			`document.getElementById("outside").click()`,
			`state("outside")`,
		}, "outside=closed"},
		{"inside", []string{
			`button().click()`,
			`root().querySelector("div").click()`,
			`state("inside")`,
		}, "inside=open"},
		{"escape", []string{
			`button().click()`,
			`root().dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }))`,
			`state("escape"); log.push("focus=" + (document.activeElement === button()))`,
		}, "escape=closed focus=true"},
		{"item", []string{
			`button().click()`,
			`document.getElementById("second").click()`,
			`state("item"); log.push("handled=" + (document.body.dataset.handled || "no"))`,
		}, "item=closed handled=second"},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := runSteps(t, c.steps); got != c.want {
				t.Errorf("states %q, want %q", got, c.want)
			}
		})
	}
}

func runSteps(t *testing.T, steps []string) string {
	t.Helper()
	body := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if _, err := io.WriteString(w, `<button id="outside" type="button">Outside</button>`); err != nil {
			return err
		}
		// Items link to the button outside, so a tap doesn't leave the page.
		// Not to themselves: following a link to an element inside a closed
		// <details> opens it again. The second marks the body, as an hx-get
		// on a real item would start a request.
		items := templ.Raw(`<a href="#outside" id="first">First</a>`)
		menu := fab.Menu(fab.Props{Label: "New"})
		return menu.Render(templ.WithChildren(ctx, templ.Join(items,
			fab.Item(fab.ItemProps{Href: "#outside", Attrs: templ.Attributes{"id": "second", "data-check": "second"}}))), w)
	})
	check := `
document.addEventListener("click", function (e) {
	if (e.target.dataset && e.target.dataset.check) document.body.dataset.handled = e.target.dataset.check;
});
document.addEventListener("alpine:initialized", function () {
	var log = [];
	var root = function () { return document.querySelector("details"); };
	var button = function () { return root().querySelector("summary"); };
	var state = function (label) { log.push(label + "=" + (root().open ? "open" : "closed")); };
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
