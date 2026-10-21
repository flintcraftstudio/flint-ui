// Package pageheader provides the title block at the top of a dashboard
// page: an optional back link, the page title, supporting copy or
// metadata, and the page's primary actions. From-scratch composition
// (Catalyst ships the pieces — Heading, Text, Button — and its demo app
// lays them out inline on every page); spacing follows the Catalyst
// demo's list and detail pages.
//
// List page:
//
//	@pageheader.PageHeader(pageheader.Props{}) {
//	    @pageheader.Content(pageheader.SubProps{}) {
//	        @pageheader.Title(pageheader.TitleProps{}) { Jobs }
//	        @pageheader.Description(pageheader.SubProps{}) { Every job across both crews. }
//	    }
//	    @pageheader.Actions(pageheader.SubProps{}) {
//	        @button.Button(button.Props{Href: "/jobs/new"}) { New job }
//	    }
//	}
//
// Detail page: add Back above Content, and Meta under Title. To put a
// Badge next to the title, wrap Title and the badge in
// <div class="flex flex-wrap items-center gap-4">.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/page-header.md in the flint-ui module.
package pageheader
