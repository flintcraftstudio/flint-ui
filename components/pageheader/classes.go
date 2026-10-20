package pageheader

// Spacing follows the Catalyst demo's orders list and order detail
// pages. Text sizes match Catalyst's Text (text-base/6 sm:text-sm/6) and
// its back link (text-sm/6 text-zinc-500 → muted-foreground).

const rootClasses = "flex flex-wrap items-end justify-between gap-4"

// backRowClasses: basis-full forces Back onto its own line in the
// wrapping row; the negative margin tightens the gap above Title.
const backRowClasses = "basis-full -mb-1 lg:mb-2"

const backLinkClasses = "inline-flex items-center gap-2 rounded-sm text-sm/6 text-muted-foreground " +
	"hover:text-foreground focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-ring"

const contentClasses = "min-w-0 max-sm:w-full sm:flex-1"

const descriptionClasses = "mt-2 max-w-2xl text-base/6 text-muted-foreground sm:text-sm/6"

// metaClasses: each child is an inline icon + text pair.
const metaClasses = "mt-2.5 flex flex-wrap items-center gap-x-6 gap-y-2 text-base/6 text-muted-foreground sm:text-sm/6 " +
	"*:inline-flex *:items-center *:gap-2 " +
	"[&_[data-slot=icon]]:size-4 [&_[data-slot=icon]]:shrink-0 [&_[data-slot=icon]]:fill-muted-foreground"

const actionsClasses = "flex flex-wrap items-center gap-4"

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
