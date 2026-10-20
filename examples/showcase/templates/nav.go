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
			{Slug: "inputs", Title: "Input", Phase: "1"},
			{Slug: "selects", Title: "Select", Phase: "1"},
			{Slug: "textareas", Title: "Textarea", Phase: "1"},
			{Slug: "checkboxes", Title: "Checkbox", Phase: "1"},
			{Slug: "radios", Title: "Radio", Phase: "1"},
			{Slug: "switches", Title: "Switch", Phase: "1"},
			{Slug: "badges", Title: "Badge", Phase: "1"},
			{Slug: "avatars", Title: "Avatar", Phase: "1"},
		},
	},
	{
		Heading: "App shell",
		Items: []NavItem{
			{Slug: "layouts", Title: "Layouts", Phase: "5"},
			{Slug: "sidebars", Title: "Sidebar", Phase: "5"},
			{Slug: "navbars", Title: "Navbar", Phase: "5"},
			{Slug: "page-headers", Title: "Page header", Phase: "5"},
			{Slug: "empty-states", Title: "Empty state", Phase: "5"},
			{Slug: "spinners", Title: "Spinner", Phase: "5"},
		},
	},
	{
		Heading: "Layout",
		Items: []NavItem{
			{Slug: "tables", Title: "Table", Phase: "2"},
			{Slug: "headings", Title: "Heading", Phase: "2"},
			{Slug: "texts", Title: "Text", Phase: "2"},
			{Slug: "cards", Title: "Card", Phase: "2"},
			{Slug: "alerts", Title: "Alert", Phase: "2"},
			{Slug: "dividers", Title: "Divider", Phase: "2"},
			{Slug: "description-lists", Title: "Description List", Phase: "2"},
		},
	},
	{
		Heading: "Interactive",
		Items: []NavItem{
			{Slug: "modals", Title: "Modal", Phase: "3"},
			{Slug: "dropdowns", Title: "Dropdown", Phase: "3"},
			{Slug: "tabs", Title: "Tabs", Phase: "3"},
			{Slug: "toasts", Title: "Toast", Phase: "3"},
		},
	},
	{
		Heading: "Utility",
		Items: []NavItem{
			{Slug: "tooltips", Title: "Tooltip", Phase: "4"},
			{Slug: "accordions", Title: "Accordion", Phase: "4"},
			{Slug: "slideovers", Title: "Slide-over", Phase: "4"},
			{Slug: "clipboards", Title: "Copy-to-Clipboard", Phase: "4"},
			{Slug: "popovers", Title: "Popover", Phase: "4"},
			{Slug: "paginations", Title: "Pagination", Phase: "4"},
			{Slug: "breadcrumbs", Title: "Breadcrumbs", Phase: "4"},
			{Slug: "comboboxes", Title: "Combobox", Phase: "4"},
			{Slug: "commands", Title: "Command Palette", Phase: "4"},
		},
	},
}
