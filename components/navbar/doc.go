// Package navbar provides the horizontal top navigation bar used by
// layout.Stacked (and as the mobile header in layout.Sidebar).
// Converted from catalyst-ui-kit/typescript/navbar.tsx.
//
// Composition:
//
//	@navbar.Navbar(navbar.Props{}) {
//	    @navbar.Item(navbar.ItemProps{Href: "/"}) { Logo }
//	    @navbar.Divider(navbar.SubProps{Class: "max-lg:hidden"})
//	    @navbar.Section(navbar.SubProps{Class: "max-lg:hidden"}) {
//	        @navbar.Item(navbar.ItemProps{Href: "/jobs", Current: true}) { Jobs }
//	        @navbar.Item(navbar.ItemProps{Href: "/crews"}) { Crews }
//	    }
//	    @navbar.Spacer(navbar.SubProps{})
//	    @navbar.Section(navbar.SubProps{}) { ...search, inbox, avatar dropdown }
//	}
//
// Catalyst animates the current-page underline between items with
// motion's layoutId. flint-ui pages are server-rendered, so the
// underline is static: the server marks the current item and the next
// navigation renders a fresh page.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/navbar.md in the flint-ui module.
package navbar
