package tabbar

import "strings"

// Not from Catalyst, which has no bottom navigation. Colors are tokens;
// the shape follows the sidebar: a muted icon and label that darken for
// the current item, with the sidebar's hover tint on the rail.

// navClasses: a bar fixed to the bottom below lg, one equal column per
// child (grid-flow-col, so a child hidden below lg takes no column); from
// lg a narrow rail fixed to the left. The bar pads the iPhone home
// indicator, which is non-zero only on a page with viewport-fit=cover.
var navClasses = strings.Join([]string{
	"fixed inset-x-0 bottom-0 z-20 grid auto-cols-fr grid-flow-col border-t border-border bg-surface pb-[env(safe-area-inset-bottom)]",
	"lg:inset-x-auto lg:inset-y-0 lg:left-0 lg:flex lg:w-22 lg:flex-col lg:items-center lg:gap-2 lg:border-t-0 lg:border-r lg:px-2 lg:py-4",
}, " ")

var itemClasses = strings.Join([]string{
	// Base: 64px tall, comfortably over the 44px tap target.
	"flex h-16 min-w-0 flex-col items-center justify-center gap-0.5 text-center text-xs/4 font-medium text-muted-foreground lg:w-full lg:gap-1 lg:rounded-lg",
	// Icon: drawn in currentColor, so it follows the label.
	"*:data-[slot=icon]:size-6 *:data-[slot=icon]:shrink-0",
	// Hover
	"data-hover:text-foreground lg:data-hover:bg-foreground/5",
	// Current
	"data-current:font-semibold data-current:text-foreground lg:data-current:bg-foreground/5",
	// Focus: inset, since the bar's edges are the screen's.
	"focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:-outline-offset-2 data-focus:outline-ring",
}, " ")

// spacerClasses: nothing on the bar; on the rail it pushes the items after
// it to the bottom.
const spacerClasses = "hidden lg:block lg:flex-1"

// ContentClasses keep a page's content clear of the tab bar: room for the
// bar (and the home indicator) at the bottom below lg, and for the rail
// at the left from lg up. Put them on the element wrapping the page; they
// set its padding, so the page's own padding goes on an element inside.
const ContentClasses = "pb-[calc(4rem+env(safe-area-inset-bottom))] lg:pb-0 lg:pl-22"

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
