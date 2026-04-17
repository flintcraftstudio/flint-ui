package textarea

import "strconv"

// Non-color class strings preserved verbatim from catalyst-ui-kit/typescript/textarea.tsx.
// Color classes are translated to the semantic token contract in styles/flint.css.
// dark:* color classes are stripped — dark mode is handled via token redefinition
// on :root (see flintcraft-ui-conversion-guide.md § "Dark Mode (Future)").

// wrapperBase is applied to the outer <span data-slot="control">. Uses
// focus-within (not has-data-focus) because <textarea> is a native element
// with no Headless UI data-focus attribute.
const wrapperBase = "relative block w-full " +
	"before:absolute before:inset-px before:rounded-[calc(var(--radius-lg)-1px)] before:bg-input before:shadow-sm " +
	"after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:ring-transparent after:ring-inset sm:focus-within:after:ring-2 sm:focus-within:after:ring-ring " +
	"has-data-disabled:opacity-50 has-data-disabled:before:bg-muted has-data-disabled:before:shadow-none"

// textareaBase is applied to the <textarea> element inside the wrapper.
const textareaBase = "relative block h-full w-full appearance-none rounded-lg px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] " +
	"text-base/6 text-foreground placeholder:text-muted-foreground sm:text-sm/6 " +
	"border border-border data-hover:border-foreground/20 " +
	"bg-transparent " +
	"focus:outline-hidden " +
	"data-invalid:border-danger data-invalid:data-hover:border-danger " +
	"data-disabled:border-border"

func wrapperClasses(p Props) string {
	if p.Class == "" {
		return wrapperBase
	}
	return wrapperBase + " " + p.Class
}

func textareaClasses(p Props) string {
	resize := "resize-y"
	if p.NonResizable {
		resize = "resize-none"
	}
	return textareaBase + " " + resize
}

func rowsString(n int) string {
	return strconv.Itoa(n)
}
