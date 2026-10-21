// Package layout provides the three application shells: a sidebar app
// (Sidebar), a top-nav app (Stacked), and a centered auth card (Auth).
// Converted from catalyst-ui-kit/typescript/sidebar-layout.tsx,
// stacked-layout.tsx, and auth-layout.tsx.
//
// Sidebar and Stacked take the navigation as templ.Component props and
// the page content as children:
//
//	@layout.Sidebar(layout.SidebarProps{
//	    Navbar:  appNavbar(),  // mobile header contents (search, avatar)
//	    Sidebar: appSidebar(), // a sidebar.Sidebar
//	}) {
//	    ...page content
//	}
//
// Below lg, both shells collapse the sidebar into a slide-in drawer
// opened by a hamburger button in the header. The drawer is Alpine
// state on the shell root (mobileNav); it closes on ESC, backdrop
// click, the close button, any link click inside it (so hx-boost
// navigation closes it too), resizing past lg, or a
// 'flint:sidebar-close' window event. Focus trap and scroll lock need
// the Alpine Focus plugin, same as Modal.
//
// The Sidebar component is rendered twice (desktop rail + mobile
// drawer), as in Catalyst. Only one copy is ever displayed, but avoid
// fixed id attributes inside it.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/layout.md in the flint-ui module.
package layout
