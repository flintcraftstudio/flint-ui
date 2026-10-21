package flintui

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestCatalogComplete keeps the ways an agent discovers a component in
// sync: every component package needs a doc.go package comment, a row in
// the README catalog linking an existing reference in docs/components/,
// and an entry in this package's catalog (doc.go). Every reference must
// belong to a catalog row.
func TestCatalogComplete(t *testing.T) {
	pkgs := componentPackages(t)

	readme := readFile(t, "README.md")
	rowRe := regexp.MustCompile("(?m)^\\| [^|]+ \\| `(\\w+)` \\|.*\\]\\(\\./docs/components/([\\w-]+\\.md)\\) \\|$")
	rows := map[string]string{} // package → doc file
	for _, m := range rowRe.FindAllStringSubmatch(readme, -1) {
		rows[m[1]] = m[2]
	}

	rootDoc := readFile(t, "doc.go")

	for _, pkg := range pkgs {
		t.Run(pkg, func(t *testing.T) {
			if doc := packageDoc(t, filepath.Join("components", pkg, "doc.go")); doc == "" {
				t.Errorf("components/%s/doc.go is missing or has no package comment", pkg)
			} else if !strings.HasPrefix(doc, "Package "+pkg+" ") {
				t.Errorf("components/%s/doc.go: package comment should start %q", pkg, "Package "+pkg+" ")
			}
			file, ok := rows[pkg]
			if !ok {
				t.Errorf("README catalog has no row for `%s`", pkg)
			} else if _, err := os.Stat(filepath.Join("docs", "components", file)); err != nil {
				t.Errorf("README row for `%s` links docs/components/%s, which doesn't exist", pkg, file)
			}
			if !regexp.MustCompile(`(?m)^//\t` + pkg + `\s`).MatchString(rootDoc) {
				t.Errorf("root doc.go catalog has no entry for %s", pkg)
			}
		})
	}

	linked := map[string]bool{}
	for pkg, file := range rows {
		linked[file] = true
		if !contains(pkgs, pkg) {
			t.Errorf("README catalog row `%s` has no components/%s package", pkg, pkg)
		}
	}
	refs, err := filepath.Glob(filepath.Join("docs", "components", "*.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, ref := range refs {
		if !linked[filepath.Base(ref)] {
			t.Errorf("%s isn't linked from any README catalog row", ref)
		}
	}
}

// componentPackages lists components/* directories that hold a component
// (shared holds common types, not a component).
func componentPackages(t *testing.T) []string {
	t.Helper()
	entries, err := os.ReadDir("components")
	if err != nil {
		t.Fatal(err)
	}
	var pkgs []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != "shared" {
			pkgs = append(pkgs, e.Name())
		}
	}
	return pkgs
}

func packageDoc(t *testing.T, path string) string {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments|parser.PackageClauseOnly)
	if err != nil || f.Doc == nil {
		return ""
	}
	return f.Doc.Text()
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func contains(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}
