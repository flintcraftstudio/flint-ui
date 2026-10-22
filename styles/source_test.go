package styles

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestSourcesCoverClasses: the partial tells Tailwind to scan every kind of
// file that holds component classes. Most components keep their classes in
// classes.go; before v1.2.2 only .templ files were scanned, so a client
// importing the partial got almost none of them and unstyled components.
func TestSourcesCoverClasses(t *testing.T) {
	css, err := os.ReadFile("flint-ui.css")
	if err != nil {
		t.Fatal(err)
	}
	exts := map[string]bool{}
	files, err := filepath.Glob("../components/*/*")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") || strings.HasSuffix(f, "doc.go") {
			continue
		}
		if ext := filepath.Ext(f); ext == ".go" || ext == ".templ" {
			exts[ext] = true
		}
	}
	for ext := range exts {
		if want := `@source "../components/**/*` + ext + `";`; !strings.Contains(string(css), want) {
			t.Errorf("flint-ui.css lacks %s: classes in component %s files would never be built", want, ext)
		}
	}
}
