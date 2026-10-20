package emptystate

import "strings"

// Proportions follow the Tailwind UI empty-state pattern; text sizes
// match Catalyst's Text so it sits naturally next to other flint-ui
// copy.

func rootClasses(p Props) string {
	parts := []string{
		"flex flex-col items-center px-6 py-12 text-center",
		"*:data-[slot=icon]:size-12 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:text-muted-foreground/60",
	}
	if p.Bordered {
		parts = append(parts, "rounded-lg border-2 border-dashed border-border")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// titleClasses: mt-2 spaces it from the icon; when there's no icon the
// margin is harmless top padding.
const titleClasses = "mt-2 text-base/6 font-semibold text-foreground sm:text-sm/6"

const descriptionClasses = "mt-1 max-w-sm text-base/6 text-muted-foreground sm:text-sm/6"

const actionsClasses = "mt-6 flex flex-wrap items-center justify-center gap-3"

func resolveLevel(l int) int {
	if l < 2 || l > 6 {
		return 3
	}
	return l
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
