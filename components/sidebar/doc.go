// Package sidebar provides the vertical app navigation used by
// layout.Sidebar (desktop rail + mobile drawer) and layout.Stacked
// (mobile drawer only). Converted from catalyst-ui-kit/typescript/
// sidebar.tsx.
//
// Composition:
//
//	@sidebar.Sidebar(sidebar.Props{}) {
//	    @sidebar.Header(sidebar.SubProps{}) { ...team switcher }
//	    @sidebar.Body(sidebar.SubProps{}) {
//	        @sidebar.Section(sidebar.SubProps{}) {
//	            @sidebar.Item(sidebar.ItemProps{Href: "/", Current: true}) {
//	                <svg data-slot="icon" ...></svg>
//	                @sidebar.Label(sidebar.SubProps{}) { Home }
//	            }
//	        }
//	        @sidebar.Spacer(sidebar.SubProps{})
//	        @sidebar.Section(sidebar.SubProps{}) { ...support, changelog }
//	    }
//	    @sidebar.Footer(sidebar.SubProps{}) { ...account dropdown }
//	}
//
// Catalyst animates the current-page bar between items with motion's
// layoutId. flint-ui pages are server-rendered, so the bar is static:
// the server marks the current item and the next navigation renders a
// fresh page.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/sidebar.md in the flint-ui module.
package sidebar
