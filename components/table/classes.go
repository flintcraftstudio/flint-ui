package table

import "strings"

// All non-color classes preserved verbatim from catalyst-ui-kit/typescript/table.tsx.
// Color classes (zinc, white alphas) map to semantic tokens per the contract
// in flintcraft-ui-conversion-guide.md. The --gutter / --spacing() v4 arbitrary
// properties are kept as-is so consumers can still override via CSS variables.

func outerClasses(p Props) string {
	parts := []string{"-mx-(--gutter)", "overflow-x-auto", "whitespace-nowrap"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func innerClasses(p Props) string {
	parts := []string{"inline-block", "min-w-full", "align-middle"}
	if !p.Bleed {
		parts = append(parts, "sm:px-(--gutter)")
	}
	return strings.Join(parts, " ")
}

func headClasses(p HeadProps) string {
	parts := []string{"text-muted-foreground"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func rowClasses(p RowProps) string {
	parts := []string{}
	if p.Striped {
		parts = append(parts, "even:bg-muted/60")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func headerClasses(p HeaderProps) string {
	parts := []string{
		"border-b", "border-b-border",
		"px-4", "py-2", "font-medium",
		"first:pl-(--gutter,--spacing(2))", "last:pr-(--gutter,--spacing(2))",
	}
	if p.Grid {
		parts = append(parts, "border-l", "border-l-border", "first:border-l-0")
	}
	if !p.Bleed {
		parts = append(parts, "sm:first:pl-1", "sm:last:pr-1")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func cellClasses(p CellProps) string {
	parts := []string{
		"relative",
		"px-4",
		"first:pl-(--gutter,--spacing(2))", "last:pr-(--gutter,--spacing(2))",
	}
	if !p.Striped {
		parts = append(parts, "border-b", "border-border")
	}
	if p.Grid {
		parts = append(parts, "border-l", "border-l-border", "first:border-l-0")
	}
	if p.Dense {
		parts = append(parts, "py-2.5")
	} else {
		parts = append(parts, "py-4")
	}
	if !p.Bleed {
		parts = append(parts, "sm:first:pl-1", "sm:last:pr-1")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// sortLinkClasses: the whole header label is the click target. The
// negative margin + padding gives the focus ring room without shifting
// the label off the column's text edge.
const sortLinkClasses = "group -mx-1 inline-flex items-center gap-1.5 rounded-sm px-1 " +
	"hover:text-foreground focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-ring"

// sortIconClasses: the active column's arrow is full-strength; inactive
// columns show a faint up/down glyph that firms up on hover, so it's
// discoverable without competing with the active one.
func sortIconClasses(d SortDirection) string {
	if d == SortNone {
		return "flex-none text-muted-foreground/50 group-hover:text-muted-foreground"
	}
	return "flex-none text-foreground"
}

func ariaSort(d SortDirection) string {
	switch d {
	case SortAsc:
		return "ascending"
	case SortDesc:
		return "descending"
	}
	return ""
}

// selectColumnClasses shrinks the checkbox column to its content and
// drops the right padding; the next column's left padding supplies the
// gap.
const selectColumnClasses = "w-0 pr-0"

const bulkBarClasses = "mb-4 flex min-h-9 flex-wrap items-center gap-x-4 gap-y-2"

const bulkClearClasses = "rounded-sm text-sm/6 font-medium text-foreground underline decoration-foreground/50 hover:decoration-foreground " +
	"focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-ring"

func selectName(name string) string {
	if name == "" {
		return "ids"
	}
	return name
}

// selectionData is the Selectable scope. count/total are derived from
// the enabled row checkboxes; a MutationObserver re-syncs them whenever
// rows are added or removed (htmx swaps, pagination), so the scope never
// needs to know which htmx event fired. Alpine resolves this.$el to the
// element whose handler is running (e.g. the select-all checkbox), so
// boxes() walks up to the scope root first.
const selectionData = `{
	count: 0,
	total: 0,
	boxes() { return [...this.$el.closest('[data-flint-selection]').querySelectorAll('[data-flint-select]:not(:disabled)')] },
	sync() { const b = this.boxes(); this.total = b.length; this.count = b.filter(x => x.checked).length },
	toggleAll(on) { this.boxes().forEach(x => { x.checked = on }); this.sync() },
	clear() { this.toggleAll(false) },
	init() { this.sync(); new MutationObserver(() => this.sync()).observe(this.$el, { childList: true, subtree: true }) },
}`

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
