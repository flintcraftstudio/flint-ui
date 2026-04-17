package selectbox

import "strconv"

// Non-color class strings preserved verbatim from catalyst-ui-kit/typescript/select.tsx.
// Color classes are translated to the semantic token contract in styles/flint.css.
// dark:* color classes are stripped — dark mode is handled via token redefinition
// on :root (see flintcraft-ui-conversion-guide.md § "Dark Mode (Future)").

// wrapperBase is applied to the outer <span data-slot="control">. The `group`
// utility is what the chevron's group-has-data-disabled: selector hooks into.
const wrapperBase = "group relative block w-full " +
	"before:absolute before:inset-px before:rounded-[calc(var(--radius-lg)-1px)] before:bg-input before:shadow-sm " +
	"after:pointer-events-none after:absolute after:inset-0 after:rounded-lg after:ring-transparent after:ring-inset has-data-focus:after:ring-2 has-data-focus:after:ring-ring " +
	"has-data-disabled:opacity-50 has-data-disabled:before:bg-muted has-data-disabled:before:shadow-none"

// selectBaseCommon is the part of the inner <select>'s class string that
// doesn't change between single and multi-select modes.
const selectBaseCommon = "relative block w-full appearance-none rounded-lg py-[calc(--spacing(2.5)-1px)] sm:py-[calc(--spacing(1.5)-1px)] " +
	"[&_optgroup]:font-semibold " +
	"text-base/6 text-foreground placeholder:text-muted-foreground sm:text-sm/6 " +
	"border border-border data-hover:border-foreground/20 " +
	"bg-transparent " +
	"focus:outline-hidden " +
	"data-invalid:border-danger data-invalid:data-hover:border-danger " +
	"data-disabled:border-border data-disabled:opacity-100"

// selectPaddingSingle leaves a wider right-hand gutter for the chevron.
const selectPaddingSingle = "pr-[calc(--spacing(10)-1px)] pl-[calc(--spacing(3.5)-1px)] sm:pr-[calc(--spacing(9)-1px)] sm:pl-[calc(--spacing(3)-1px)]"

// selectPaddingMultiple is symmetric since multi-selects render the native
// list box and don't show a chevron.
const selectPaddingMultiple = "px-[calc(--spacing(3.5)-1px)] sm:px-[calc(--spacing(3)-1px)]"

// chevronWrapperClasses positions the chevron container absolutely over
// the right edge of the wrapper.
const chevronWrapperClasses = "pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2"

// chevronSvgClasses sizes and colors the chevron icon. group-has-data-disabled
// hooks the outer wrapper's disabled state so the icon darkens when the
// <select> is disabled (the wrapper also fades to opacity-50).
const chevronSvgClasses = "size-5 stroke-muted-foreground group-has-data-disabled:stroke-foreground sm:size-4 forced-colors:stroke-[CanvasText]"

func wrapperClasses(p Props) string {
	if p.Class == "" {
		return wrapperBase
	}
	return wrapperBase + " " + p.Class
}

func selectClasses(p Props) string {
	padding := selectPaddingSingle
	if p.Multiple {
		padding = selectPaddingMultiple
	}
	return selectBaseCommon + " " + padding
}

func sizeString(n int) string {
	return strconv.Itoa(n)
}
