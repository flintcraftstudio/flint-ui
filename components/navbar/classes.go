package navbar

import "strings"

// Non-color class strings preserved verbatim from catalyst-ui-kit/
// typescript/navbar.tsx. zinc-950 → foreground, zinc-500 →
// muted-foreground, zinc-950/5 and /10 → foreground/5 and /10. The
// dark: variants collapse into the token contract.

const navbarClasses = "flex flex-1 items-center gap-4 py-2.5"

const dividerClasses = "h-6 w-px bg-foreground/10"

const sectionClasses = "flex items-center gap-3"

const spacerClasses = "-ml-4 flex-1"

// currentIndicatorClasses draws the underline below the current item.
// Catalyst animates it between items with motion's layoutId; here it is
// static.
const currentIndicatorClasses = "absolute inset-x-2 -bottom-2.5 h-0.5 rounded-full bg-foreground"

var itemClasses = strings.Join([]string{
	// Base
	"relative flex min-w-0 items-center gap-3 rounded-lg p-2 text-left text-base/6 font-medium text-foreground sm:text-sm/5",
	// Leading icon/icon-only
	"*:data-[slot=icon]:size-6 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:fill-muted-foreground sm:*:data-[slot=icon]:size-5",
	// Trailing icon (down chevron or similar)
	"*:not-nth-2:last:data-[slot=icon]:ml-auto *:not-nth-2:last:data-[slot=icon]:size-5 sm:*:not-nth-2:last:data-[slot=icon]:size-4",
	// Avatar
	"*:data-[slot=avatar]:-m-0.5 *:data-[slot=avatar]:size-7 *:data-[slot=avatar]:[--avatar-radius:var(--radius-md)] sm:*:data-[slot=avatar]:size-6",
	// Hover
	"data-hover:bg-foreground/5 data-hover:*:data-[slot=icon]:fill-foreground",
	// Active
	"data-active:bg-foreground/5 data-active:*:data-[slot=icon]:fill-foreground",
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
