package fieldset

// Non-color class strings preserved verbatim from catalyst-ui-kit/typescript/fieldset.tsx.
// Color classes are translated to semantic tokens (see styles/flint.css);
// dark:* color variants are stripped (dark mode is handled via token redefinition).

const fieldsetBase = "*:data-[slot=text]:mt-1 [&>*+[data-slot=control]]:mt-6"

const legendBase = "text-base/6 font-semibold text-foreground data-disabled:opacity-50 sm:text-sm/6"

const fieldGroupBase = "space-y-8"

// Field uses data-slot sibling selectors to auto-space label/control/
// description/error children. The order of placement in the template
// decides which selectors apply.
const fieldBase = "[&>[data-slot=label]+[data-slot=control]]:mt-3 " +
	"[&>[data-slot=label]+[data-slot=description]]:mt-1 " +
	"[&>[data-slot=description]+[data-slot=control]]:mt-3 " +
	"[&>[data-slot=control]+[data-slot=description]]:mt-3 " +
	"[&>[data-slot=control]+[data-slot=error]]:mt-3 " +
	"*:data-[slot=label]:font-medium"

const labelBase = "text-base/6 text-foreground select-none data-disabled:opacity-50 sm:text-sm/6"

const descriptionBase = "text-base/6 text-muted-foreground data-disabled:opacity-50 sm:text-sm/6"

const errorMessageBase = "text-base/6 text-danger data-disabled:opacity-50 sm:text-sm/6"

func joinClasses(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
