package fileinput

import (
	"fmt"
	"strings"
)

// wrapperClasses and the frame half of inputClasses mirror input.Input
// (Catalyst's input.tsx), so a file field lines up with text fields in
// the same form. The file: variants restyle the browser's "Choose file"
// button as a muted chip that fits the input's line box.
const wrapperClasses = "relative block w-full " +
	"before:absolute before:inset-px before:rounded-[calc(var(--radius-lg)-1px)] before:bg-input before:shadow-sm " +
	"after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:ring-transparent after:ring-inset sm:focus-within:after:ring-2 sm:focus-within:after:ring-ring " +
	"has-data-disabled:opacity-50 has-data-disabled:before:bg-muted has-data-disabled:before:shadow-none"

const inputClasses = "relative block w-full appearance-none rounded-lg px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] " +
	"text-base/6 text-muted-foreground sm:text-sm/6 " +
	"border border-border data-hover:border-foreground/20 bg-transparent focus:outline-hidden " +
	"data-invalid:border-danger data-invalid:data-hover:border-danger data-disabled:border-border " +
	"file:mr-3 file:rounded-md file:border-0 file:bg-muted file:px-2.5 file:py-0.5 file:text-sm/5 file:font-semibold file:text-foreground " +
	"hover:file:bg-foreground/10 disabled:file:bg-muted"

// dropAreaClasses: the whole area is the <label>, so it's the click
// target. Keyboard focus lands on the sr-only input inside; has-[...]
// draws the ring on the area instead.
const dropAreaClasses = "flex cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed border-border px-6 py-10 text-center " +
	"hover:border-foreground/25 " +
	"has-[input:focus-visible]:outline-2 has-[input:focus-visible]:outline-offset-2 has-[input:focus-visible]:outline-ring " +
	"data-[dragging=true]:border-ring data-[dragging=true]:bg-ring/5 " +
	"data-invalid:border-danger " +
	"data-disabled:cursor-not-allowed data-disabled:opacity-50 data-disabled:hover:border-border " +
	"*:data-[slot=icon]:size-10 *:data-[slot=icon]:text-muted-foreground/60"

const clearClasses = "shrink-0 rounded-sm text-sm/6 font-medium text-foreground underline decoration-foreground/50 hover:decoration-foreground " +
	"focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-ring"

func prompt(p DropzoneProps) string {
	switch {
	case p.Prompt != "":
		return p.Prompt
	case p.Multiple:
		return "Upload files"
	}
	return "Upload a file"
}

// dropzoneData is the Dropzone's Alpine scope.
//
// Browsers apply accept to the picker but not to drops, so drop()
// filters against the same list (extensions, exact MIME types, and
// type/* wildcards) and reports what it skipped. Accepted files are
// written into the real input via DataTransfer, then a bubbling change
// event is dispatched so sync() and any hx-trigger="change" both run.
func dropzoneData(accept string, multiple, disabled bool) string {
	var rules []string
	for _, r := range strings.Split(accept, ",") {
		if r = strings.ToLower(strings.TrimSpace(r)); r != "" {
			rules = append(rules, fmt.Sprintf("%q", r))
		}
	}
	return fmt.Sprintf(`{
	files: [],
	rejected: '',
	dragging: false,
	disabled: %t,
	multiple: %t,
	accept: [%s],
	ok(f) {
		if (!this.accept.length) return true
		const name = f.name.toLowerCase(), type = (f.type || '').toLowerCase()
		return this.accept.some(a => a.startsWith('.') ? name.endsWith(a) : a.endsWith('/*') ? type.startsWith(a.slice(0, -1)) : type === a)
	},
	drop(e) {
		this.dragging = false
		if (this.disabled) return
		let all = [...e.dataTransfer.files], good = all.filter(f => this.ok(f))
		if (!this.multiple) good = good.slice(0, 1)
		const skipped = all.length - good.length
		const dt = new DataTransfer()
		good.forEach(f => dt.items.add(f))
		this.$refs.input.files = dt.files
		this.$refs.input.dispatchEvent(new Event('change', { bubbles: true }))
		this.rejected = skipped ? skipped + (skipped === 1 ? ' file' : ' files') + ' skipped (wrong type or too many)' : ''
	},
	sync() { this.files = [...this.$refs.input.files].map(f => ({ name: f.name, size: f.size })) },
	clear() { this.$refs.input.value = ''; this.rejected = ''; this.sync() },
	size(n) { return n < 1024 ? n + ' B' : n < 1048576 ? Math.round(n / 1024) + ' KB' : (n / 1048576).toFixed(1) + ' MB' },
}`, disabled, multiple, strings.Join(rules, ", "))
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
