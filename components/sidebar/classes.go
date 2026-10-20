package sidebar

import "strings"

// Non-color class strings preserved verbatim from catalyst-ui-kit/
// typescript/sidebar.tsx. zinc-950 → foreground, zinc-500 →
// muted-foreground, zinc-950/5 → foreground/5. The dark: variants
// collapse into the token contract.

const sidebarClasses = "flex h-full min-h-0 flex-col"

const headerClasses = "flex flex-col border-b border-foreground/5 p-4 [&>[data-slot=section]+[data-slot=section]]:mt-2.5"

const bodyClasses = "flex flex-1 flex-col overflow-y-auto p-4 [&>[data-slot=section]+[data-slot=section]]:mt-8"

const footerClasses = "flex flex-col border-t border-foreground/5 p-4 [&>[data-slot=section]+[data-slot=section]]:mt-2.5"

const sectionClasses = "flex flex-col gap-0.5"

const dividerClasses = "my-4 border-t border-foreground/5 lg:-mx-4"

const spacerClasses = "mt-8 flex-1"

const headingClasses = "mb-1 px-2 text-xs/6 font-medium text-muted-foreground"

// currentIndicatorClasses draws the bar at the left edge of the current
// item, in the sidebar's p-4 gutter. Catalyst animates it between items
// with motion's layoutId; here it is static.
const currentIndicatorClasses = "absolute inset-y-2 -left-4 w-0.5 rounded-full bg-foreground"

var itemClasses = strings.Join([]string{
	// Base
	"flex w-full items-center gap-3 rounded-lg px-2 py-2.5 text-left text-base/6 font-medium text-foreground sm:py-2 sm:text-sm/5",
	// Leading icon/icon-only
	"*:data-[slot=icon]:size-6 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:fill-muted-foreground sm:*:data-[slot=icon]:size-5",
	// Trailing icon (down chevron or similar)
	"*:last:data-[slot=icon]:ml-auto *:last:data-[slot=icon]:size-5 sm:*:last:data-[slot=icon]:size-4",
	// Avatar
	"*:data-[slot=avatar]:-m-0.5 *:data-[slot=avatar]:size-7 sm:*:data-[slot=avatar]:size-6",
	// Hover
	"data-hover:bg-foreground/5 data-hover:*:data-[slot=icon]:fill-foreground",
	// Active
	"data-active:bg-foreground/5 data-active:*:data-[slot=icon]:fill-foreground",
	// Current
	"data-current:*:data-[slot=icon]:fill-foreground",
	// Focus — Catalyst leans on the browser default ring; flint-ui uses
	// the ring token so keyboard focus matches the rest of the kit.
	"focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring",
}, " ")

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
