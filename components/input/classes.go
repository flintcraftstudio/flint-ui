package input

import "strings"

// Non-color class strings preserved verbatim from catalyst-ui-kit/typescript/input.tsx.
// Color classes are translated to the semantic token contract in styles/flint.css.
// dark:* color classes are stripped — dark mode is handled via token redefinition
// on :root (see flintcraft-ui-conversion-guide.md § "Dark Mode (Future)").

// dateTypes are the <input type> values that need Catalyst's webkit
// datetime-edit overrides to render consistently across browsers.
var dateTypes = map[string]struct{}{
	"date":           {},
	"datetime-local": {},
	"month":          {},
	"time":           {},
	"week":           {},
}

const dateTypeClasses = "[&::-webkit-datetime-edit-fields-wrapper]:p-0 " +
	"[&::-webkit-date-and-time-value]:min-h-[1.5em] " +
	"[&::-webkit-datetime-edit]:inline-flex " +
	"[&::-webkit-datetime-edit]:p-0 " +
	"[&::-webkit-datetime-edit-year-field]:p-0 " +
	"[&::-webkit-datetime-edit-month-field]:p-0 " +
	"[&::-webkit-datetime-edit-day-field]:p-0 " +
	"[&::-webkit-datetime-edit-hour-field]:p-0 " +
	"[&::-webkit-datetime-edit-minute-field]:p-0 " +
	"[&::-webkit-datetime-edit-second-field]:p-0 " +
	"[&::-webkit-datetime-edit-millisecond-field]:p-0 " +
	"[&::-webkit-datetime-edit-meridiem-field]:p-0"

// wrapperBase is applied to the outer <span data-slot="control">.
const wrapperBase = "relative block w-full " +
	"before:absolute before:inset-px before:rounded-[calc(var(--radius-lg)-1px)] before:bg-input before:shadow-sm " +
	"after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:ring-transparent after:ring-inset sm:focus-within:after:ring-2 sm:focus-within:after:ring-ring " +
	"has-data-disabled:opacity-50 has-data-disabled:before:bg-muted has-data-disabled:before:shadow-none"

// inputBase is applied to the <input> element inside the wrapper.
const inputBase = "relative block w-full appearance-none rounded-lg px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] " +
	"text-base/6 text-foreground placeholder:text-muted-foreground sm:text-sm/6 " +
	"border border-border data-hover:border-foreground/20 " +
	"bg-transparent " +
	"focus:outline-hidden " +
	"data-invalid:border-danger data-invalid:data-hover:border-danger " +
	"data-disabled:border-border"

// inputGroupBase wraps an <input> and one or two <svg data-slot="icon">
// siblings, repositioning the input padding to make room for the icons.
const inputGroupBase = "relative isolate block " +
	"has-[[data-slot=icon]:first-child]:[&_input]:pl-10 has-[[data-slot=icon]:last-child]:[&_input]:pr-10 sm:has-[[data-slot=icon]:first-child]:[&_input]:pl-8 sm:has-[[data-slot=icon]:last-child]:[&_input]:pr-8 " +
	"*:data-[slot=icon]:pointer-events-none *:data-[slot=icon]:absolute *:data-[slot=icon]:top-3 *:data-[slot=icon]:z-10 *:data-[slot=icon]:size-5 sm:*:data-[slot=icon]:top-2.5 sm:*:data-[slot=icon]:size-4 " +
	"[&>[data-slot=icon]:first-child]:left-3 sm:[&>[data-slot=icon]:first-child]:left-2.5 [&>[data-slot=icon]:last-child]:right-3 sm:[&>[data-slot=icon]:last-child]:right-2.5 " +
	"*:data-[slot=icon]:text-muted-foreground"

func wrapperClasses(p Props) string {
	if p.Class == "" {
		return wrapperBase
	}
	return wrapperBase + " " + p.Class
}

func inputClasses(p Props) string {
	parts := []string{inputBase}
	if _, ok := dateTypes[p.Type]; ok {
		parts = append(parts, dateTypeClasses)
	}
	return strings.Join(parts, " ")
}

func resolveType(t string) string {
	if t == "" {
		return "text"
	}
	return t
}

func joinClasses(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
