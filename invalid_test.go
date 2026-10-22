package flintui_test

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/flintcraftstudio/flint-ui/components/checkbox"
	"github.com/flintcraftstudio/flint-ui/components/combobox"
	"github.com/flintcraftstudio/flint-ui/components/fileinput"
	"github.com/flintcraftstudio/flint-ui/components/input"
	"github.com/flintcraftstudio/flint-ui/components/radio"
	"github.com/flintcraftstudio/flint-ui/components/selectbox"
	"github.com/flintcraftstudio/flint-ui/components/textarea"
	"github.com/flintcraftstudio/flint-ui/components/toggle"
)

// ariaInvalid matches every aria-invalid attribute, with or without a value.
var ariaInvalid = regexp.MustCompile(`\saria-invalid(="[^"]*")?`)

// TestAriaInvalid: every form control with Invalid set says
// aria-invalid="true", and says nothing when it isn't. A bare aria-invalid
// (as before v1.2.1) is an empty value, which ARIA reads as false, so
// screen readers never heard that a field was invalid.
func TestAriaInvalid(t *testing.T) {
	controls := map[string]func(invalid bool) templ.Component{
		"checkbox":  func(v bool) templ.Component { return checkbox.Checkbox(checkbox.Props{Name: "c", Invalid: v}) },
		"combobox":  func(v bool) templ.Component { return combobox.Combobox(combobox.Props{Name: "c", Invalid: v}) },
		"fileinput": func(v bool) templ.Component { return fileinput.Input(fileinput.Props{Name: "c", Invalid: v}) },
		"dropzone":  func(v bool) templ.Component { return fileinput.Dropzone(fileinput.DropzoneProps{Name: "c", Invalid: v}) },
		"input":     func(v bool) templ.Component { return input.Input(input.Props{Name: "c", Invalid: v}) },
		"radio":     func(v bool) templ.Component { return radio.Radio(radio.Props{Name: "c", Invalid: v}) },
		"select":    func(v bool) templ.Component { return selectbox.Select(selectbox.Props{Name: "c", Invalid: v}) },
		"textarea":  func(v bool) templ.Component { return textarea.Textarea(textarea.Props{Name: "c", Invalid: v}) },
		"switch":    func(v bool) templ.Component { return toggle.Switch(toggle.Props{Name: "c", Invalid: v}) },
	}
	for name, render := range controls {
		for _, invalid := range []bool{true, false} {
			var b strings.Builder
			if err := render(invalid).Render(context.Background(), &b); err != nil {
				t.Fatalf("%s: %v", name, err)
			}
			attrs := ariaInvalid.FindAllString(b.String(), -1)
			want := []string(nil)
			if invalid {
				want = []string{` aria-invalid="true"`}
			}
			if strings.Join(attrs, "|") != strings.Join(want, "|") {
				t.Errorf("%s with Invalid %v: aria-invalid attributes %q, want %q", name, invalid, attrs, want)
			}
		}
	}
}
