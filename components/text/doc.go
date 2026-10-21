// Package text provides the body-copy primitives: a paragraph, an inline
// link, strong emphasis, and inline code. Converted from
// catalyst-ui-kit/typescript/text.tsx (Catalyst's Link is a plain <a>, so
// TextLink is folded in here as Link).
//
//	@text.Text(text.Props{}) {
//	    Invoices are sent from{ " " }
//	    @text.Strong(text.Props{}) {
//	        billing@ridgeline.com
//	    }
//	    { " " }— change it in{ " " }
//	    @text.Link(text.LinkProps{Href: "/settings"}) {
//	        settings
//	    }
//	    { "." }
//	}
//
// templ trims whitespace around component calls, so spaces between a word
// and an inline component must be explicit ({ " " }), and inline children
// belong on their own lines — `{ settings }` on one line keeps a trailing
// space inside the element.
//
// Use these instead of hand-rolled text-muted-foreground paragraphs so
// copy stays consistent with the rest of the kit and rebrands with the
// token contract.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/text.md in the flint-ui module.
package text
