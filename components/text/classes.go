package text

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/
// text.tsx. zinc-500 → muted-foreground, zinc-950 → foreground,
// zinc-950/N → foreground/N. The dark: variants collapse into the token
// contract.

const textClasses = "text-base/6 text-muted-foreground sm:text-sm/6"

// linkClasses: Catalyst relies on the browser's default focus ring; the
// ring token keeps keyboard focus consistent with the rest of the kit.
const linkClasses = "text-foreground underline decoration-foreground/50 data-hover:decoration-foreground " +
	"rounded-xs focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-ring"

const strongClasses = "font-medium text-foreground"

const codeClasses = "rounded-sm border border-foreground/10 bg-foreground/2.5 px-0.5 text-sm font-medium text-foreground sm:text-[0.8125rem]"

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
