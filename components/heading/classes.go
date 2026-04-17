package heading

import "strings"

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/heading.tsx.
// Catalyst's text-zinc-950 dark:text-white maps to text-foreground.

const (
	baseHeading    = "text-2xl/8 font-semibold text-foreground sm:text-xl/8"
	baseSubheading = "text-base/7 font-semibold text-foreground sm:text-sm/6"
)

func headingClasses(p Props) string {
	return join(baseHeading, p.Class)
}

func subheadingClasses(p Props) string {
	return join(baseSubheading, p.Class)
}

func resolveLevel(l Level, fallback int) int {
	if l < 1 || l > 6 {
		return fallback
	}
	return int(l)
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return strings.Join([]string{base, extra}, " ")
}
