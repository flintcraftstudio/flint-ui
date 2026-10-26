// Package tabbar provides an app's main navigation as tabs: a bar fixed
// to the bottom of a phone's screen, in the thumb zone, that becomes a
// narrow rail at the left from lg up. Catalyst has no equivalent; use it
// in place of the layout shells' drawer when an app has a few top-level
// sections.
//
// Composition:
//
//	@tabbar.TabBar(tabbar.Props{Attrs: templ.Attributes{"aria-label": "Main"}}) {
//	    @tabbar.Item(tabbar.ItemProps{Href: "/", Current: true}) {
//	        <svg data-slot="icon" ...></svg>
//	        Quotes
//	    }
//	    @tabbar.Item(tabbar.ItemProps{Href: "/clients"}) { ... Clients }
//	    @tabbar.Spacer(tabbar.SubProps{})
//	    @tabbar.Item(tabbar.ItemProps{Href: "/settings"}) { ... Settings }
//	}
//	<main class={ tabbar.ContentClasses }>
//	    <div class="px-4 py-8">...page content</div>
//	</main>
//
// The bar gives each child an equal column, so a child hidden below lg
// (a brand mark with max-lg:hidden, or the Spacer) takes no room there.
// Items are plain links: it needs no JavaScript and runs under a strict
// CSP as it is.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/tab-bar.md in the flint-ui module.
package tabbar
