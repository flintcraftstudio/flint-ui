package fab

import "strings"

// Not from Catalyst. Colors are tokens: the button is the primary hue,
// like a solid button; the menu is a surface card.

// rootClasses place the button at the right, above flint-ui's tab bar
// (4rem tall, plus the home indicator) with a 1rem gap, over the page and
// under the bar's own layer. Hidden from lg up, where the tab bar is a
// sidebar and the action belongs in the page header.
const rootClasses = "fixed right-4 bottom-[calc(5rem+env(safe-area-inset-bottom))] z-20 lg:hidden"

// buttonClasses: 56px round, comfortably over the 44px tap target.
var buttonClasses = strings.Join([]string{
	"flex size-14 cursor-pointer list-none items-center justify-center rounded-full bg-primary text-primary-foreground shadow-lg",
	// Hide the <summary> marker in WebKit.
	"[&::-webkit-details-marker]:hidden",
	// Icon: the + turns to an × while the menu is open.
	"*:data-[slot=icon]:size-6 *:data-[slot=icon]:transition-transform group-open:*:data-[slot=icon]:rotate-45",
	"data-hover:opacity-90",
	"focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring",
}, " ")

// menuClasses: a card standing on the button's right edge, opening upward.
const menuClasses = "absolute right-0 bottom-full mb-3 flex w-max min-w-48 flex-col rounded-xl bg-surface p-1 shadow-lg ring-1 ring-border"

var itemClasses = strings.Join([]string{
	// 48px rows, over the 44px tap target.
	"flex min-h-12 items-center gap-3 rounded-lg px-3 text-base/6 font-medium text-foreground",
	"*:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:text-muted-foreground",
	"data-hover:bg-foreground/5",
	"focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:-outline-offset-2 data-focus:outline-ring",
}, " ")

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
