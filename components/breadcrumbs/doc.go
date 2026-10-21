// Package breadcrumbs provides navigation trails — the row of links at
// the top of nested dashboard pages showing where the user is in the
// hierarchy. From-scratch design (Catalyst has no breadcrumbs.tsx).
//
// Compound API matches flint-ui convention:
//
//	@breadcrumbs.Breadcrumbs(breadcrumbs.Props{}) {
//	    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/"}) { Dashboard }
//	    @breadcrumbs.Item(breadcrumbs.ItemProps{Href: "/jobs"}) { Jobs }
//	    @breadcrumbs.Current(breadcrumbs.CurrentProps{}) { Henderson Ranch }
//	}
//
// Item renders a link followed by a chevron separator; Current renders
// the final unlinked entry with aria-current="page" and no trailing
// separator. Order matters — callers put Current last and only once.
//
// No Alpine, no keyboard wiring. Pure server-rendered navigation.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/breadcrumbs.md in the flint-ui module.
package breadcrumbs
