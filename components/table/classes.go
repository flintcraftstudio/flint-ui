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
