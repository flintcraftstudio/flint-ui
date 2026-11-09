// Package fab provides a floating action button: a round + button in a
// phone's thumb zone, just above the tab bar, for the section's main
// action. Catalyst has no equivalent. From lg up, where the tab bar is a
// sidebar and the page header has room, it is hidden: put the same action
// in the header there.
//
// One action is a link; it can carry hx-* attributes to open a sheet:
//
//	@fab.Fab(fab.Props{Label: "New quote", Href: "/quotes/new"})
//
// Two or more open a menu above the button:
//
//	@fab.Menu(fab.Props{Label: "New"}) {
//	    @fab.Item(fab.ItemProps{Href: "/clients/new"}) { New client }
//	    @fab.Item(fab.ItemProps{Href: "/quotes/new"}) { New quote }
//	}
//
// The menu is a native <details>, so it opens and its links work without
// JavaScript. With Alpine, x-data="flintFab" closes it on a tap outside,
// on Escape (returning focus to the button) and on a tap on an item, so
// an item that opens a sheet doesn't leave the menu open behind it.
//
// Full reference (props, placement, accessibility): docs/components/fab.md
// in the flint-ui module.
package fab
