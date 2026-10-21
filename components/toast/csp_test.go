package toast_test

import (
	"strings"
	"testing"

	"github.com/flintcraftstudio/flint-ui/components/toast"
	"github.com/flintcraftstudio/flint-ui/internal/csptest"
)

// TestContainerUnderCSP: the toast container starts under Alpine's CSP
// build and a strict CSP, renders dispatched toasts (title and body), keeps
// a duration-0 toast, and logs nothing to the console. Before v1.2.0 the
// container's state was inline in x-data and never started under the CSP
// build, so every toast was silently lost.
func TestContainerUnderCSP(t *testing.T) {
	check := `
document.addEventListener("alpine:initialized", function () {
	window.dispatchEvent(new CustomEvent("flint:toast", {
		detail: { variant: "danger", title: "Lasting notice", body: "Stays put", duration: 0 }
	}));
	window.dispatchEvent(new CustomEvent("flint:toast", {
		detail: { variant: "success", title: "Timed notice", duration: 60000 }
	}));
});`
	res := csptest.Run(t, toast.Container(toast.ContainerProps{}), check)
	if len(res.Console) > 0 {
		t.Errorf("console errors under CSP:\n%s", strings.Join(res.Console, "\n"))
	}
	for _, want := range []string{"Lasting notice", "Stays put", "Timed notice"} {
		if !strings.Contains(res.DOM, ">"+want+"<") {
			t.Errorf("toast %q wasn't rendered", want)
		}
	}
}
