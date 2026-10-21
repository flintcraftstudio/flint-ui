// Package flintui is the entry point to flint-ui: server-rendered UI
// components for Go + templ + htmx 4 + Alpine.js + Tailwind CSS v4,
// converted from Tailwind's Catalyst UI kit. The components live under
// components/, one package per family; this package holds Script, the
// JavaScript some of them need.
//
// Invoke a component inside a .templ file as
//
//	@button.Button(button.Props{Variant: button.VariantAccent}) { Save }
//
// Every component takes a Props struct with Class and Attrs; Attrs is
// spread onto the root element, which is how hx-*, aria-* and data-*
// attributes reach it.
//
// # Finding things
//
// Run go doc on a package for its overview and API, for example
//
//	go doc github.com/flintcraftstudio/flint-ui/components/table
//
// The full references (usage, every prop, variants, accessibility, htmx
// wiring) are markdown files inside the module. From a client project:
//
//	dir=$(go list -m -f '{{.Dir}}' github.com/flintcraftstudio/flint-ui)
//	$dir/AGENTS.md              start here: catalog, rules, where to look
//	$dir/docs/patterns.md       page recipes: list page, forms, modals, toasts
//	$dir/docs/components/*.md   one reference per component
//
// Client projects set up per the README also have the module linked at
// .flint-ui/, so the same files are at .flint-ui/AGENTS.md and so on.
//
// # Catalog
//
// App shell and page structure:
//
//	layout       Sidebar, Stacked and Auth application shells
//	sidebar      vertical navigation for layout.Sidebar
//	navbar       horizontal navigation for layout.Stacked and the mobile header
//	pageheader   page title block: back link, title, description, actions
//	heading      page and section headings
//	breadcrumbs  navigation trail
//	tabs         tabbed navigation, htmx-driven or Alpine-driven
//	pagination   numbered page navigation
//	divider      horizontal rule
//
// Forms:
//
//	fieldset     Fieldset, FieldGroup, Field, Label, Description, ErrorMessage
//	input        text inputs, and InputGroup for icons
//	textarea     multi-line input
//	selectbox    native select (package name avoids the Go keyword)
//	combobox     searchable single-select
//	checkbox     checkbox, CheckboxField and CheckboxGroup
//	radio        radio, RadioField and RadioGroup
//	toggle       on/off switch (package name avoids "switch")
//	fileinput    file picker and drag-and-drop Dropzone
//	button       solid, outline and plain buttons, with an htmx loading state
//
// Data display:
//
//	table        tables with sorting, row selection and bulk actions
//	stat         KPI tiles
//	descriptionlist  key/value pairs
//	badge        status labels
//	avatar       user image or initials
//	card         surface container
//	text         paragraph, link, strong and code primitives
//	emptystate   placeholder for an empty list or panel
//
// Feedback:
//
//	alert        inline status banner
//	toast        notifications, fired by an Alpine event or an HX-Trigger header
//	spinner      indeterminate loading indicator
//	progress     determinate or looping progress bar
//
// Overlays and disclosure:
//
//	modal        dialog and alert dialog
//	slideover    edge-anchored drawer
//	dropdown     menus, including row-action menus in tables
//	popover      click-activated floating panel
//	tooltip      hover and focus hint
//	command      Cmd+K command palette
//	accordion    disclosure panels
//	clipboard    copy-to-clipboard wrapper
//
// # Rules
//
// Components draw color only from the semantic tokens in
// styles/flint-ui.css; rebrand by overriding those CSS variables, never by
// adding color classes. Alpine holds only local UI state (open/closed,
// selected tab), never application data. htmx attributes go on the element
// that makes the request, because htmx 4 doesn't inherit them.
package flintui
