package templates

// NavItem represents a single entry in the showcase sidebar.
type NavItem struct {
	Slug  string
	Title string
	Phase string
}

// NavSections groups components by the conversion phases in
// flintcraft-ui-conversion-guide.md. Add entries as components land.
var NavSections = []struct {
	Heading string
	Items   []NavItem
}{
	{
		Heading: "Foundation",
		Items: []NavItem{
			{Slug: "buttons", Title: "Button", Phase: "1"},
		},
	},
}
