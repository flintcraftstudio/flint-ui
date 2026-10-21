// Package clipboard provides a Copy wrapper that writes a value to the
// system clipboard on demand. Tiny Alpine primitive for API keys, share
// links, wallet addresses, and similar copy-me-please patterns.
//
// Architecture: the wrapper is an inline span with x-data exposing
// { copied, copy() }. Children reference those identifiers via Alpine's
// inherited scope:
//
//	@clipboard.Copy(clipboard.Props{Value: "wlg-live-abc123"}) {
//	    @button.Button(button.Props{
//	        Plain: true,
//	        Attrs: templ.Attributes{"x-on:click": "copy()"},
//	    }) {
//	        <span x-show="!copied">Copy</span>
//	        <span x-show="copied" class="text-success">Copied</span>
//	    }
//	}
//
// The value rides on a data-value attribute; copy() reads it via
// this.$el.dataset.value so the x-data expression itself is a constant
// string. templ handles HTML-attribute escaping automatically, which
// sidesteps every "what if Value contains a quote" concern — the value
// never gets embedded in a JS string literal.
//
// A11y: the component doesn't render a trigger of its own, so there's
// no prescribed a11y pattern. Callers should either (a) add
// aria-label="Copy X" on their trigger when the visible text alone
// doesn't describe the action, or (b) include aria-live="polite" on a
// visible element that reflects the copied state — without it, screen
// reader users get no feedback on click.
//
// Browser requirements: navigator.clipboard.writeText requires a secure
// context (https or localhost). On insecure contexts the Promise
// rejects silently — consider adding an error-path handler via a
// x-on:copy-error listener if this is a concern for your deployment.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/clipboard.md in the flint-ui module.
package clipboard
