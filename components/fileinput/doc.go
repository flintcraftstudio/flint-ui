// Package fileinput provides file pickers: Input, a native file input
// styled to match input.Input, and Dropzone, a drag-and-drop area for
// imports and attachments. From-scratch design (Catalyst has no file
// input); Dropzone follows the Tailwind UI upload pattern.
//
// Both render a real <input type="file">, so a plain form submit works
// with no JS. Dropzone adds Alpine for the drag-over state, the selected
// file list, and dropping files onto it.
//
// htmx: a request only carries files when it's multipart. Put
// hx-encoding="multipart/form-data" on the requesting element (or
// enctype on the <form>); otherwise htmx 4 URL-encodes the body and the
// files are silently dropped. htmx 4 uploads with fetch, which reports no
// upload progress — show button.Props.Indicator while it's in flight.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/file-input.md in the flint-ui module.
package fileinput
