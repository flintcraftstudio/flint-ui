package checkbox

// Non-color class strings preserved from catalyst-ui-kit/typescript/checkbox.tsx
// where possible; structural changes were required because Catalyst uses
// Headless UI's <Checkbox> component (a styled button with data-* state),
// while flint-ui uses a real <input type="checkbox"> with peer-* variants.
// Color classes are translated to the semantic token contract in styles/flint.css.

// labelBase is the outer <label>. `group relative isolate` sets up the
// stacking context for the absolutely-positioned input + box siblings.
const labelBase = "group relative isolate inline-flex size-[1.125rem] shrink-0 items-center justify-center rounded-[0.3125rem] sm:size-4"

// inputClasses is the native <input type="checkbox">. Absolutely positioned
// on top of the box + svg so the label's full bounds are clickable. `peer`
// lets the sibling box/svg react via peer-checked:, peer-focus-visible:, etc.
const inputClasses = "peer absolute inset-0 z-20 m-0 cursor-pointer appearance-none rounded-[0.3125rem] opacity-0 disabled:cursor-not-allowed"

// boxClasses is the visible box span. Sits under the input and SVG and
// handles background, border, focus ring, disabled fade, invalid border.
const boxClasses = "absolute inset-0 z-0 rounded-[0.3125rem] " +
	"bg-input shadow-sm " +
	"before:absolute before:inset-0 before:rounded-[calc(0.3125rem-1px)] before:shadow-[inset_0_1px_--theme(--color-white/15%)] " +
	"border border-border " +
	"peer-hover:border-foreground/30 " +
	"peer-checked:bg-primary peer-checked:border-transparent peer-checked:shadow-none " +
	"peer-focus-visible:outline-2 peer-focus-visible:outline-offset-2 peer-focus-visible:outline-ring " +
	"peer-data-invalid:border-danger peer-hover:peer-data-invalid:border-danger " +
	"peer-disabled:opacity-50 peer-disabled:bg-muted peer-disabled:border-border peer-disabled:peer-checked:bg-muted"

// checkSvgClasses is the checkmark. Sits above the box (z-10) and fades in
// when the input is checked. Forced-colors keeps it visible in Windows HC mode.
const checkSvgClasses = "relative z-10 size-4 stroke-primary-foreground opacity-0 peer-checked:opacity-100 peer-disabled:stroke-muted-foreground sm:size-3.5 forced-colors:stroke-[HighlightText]"

// groupBase stacks CheckboxFields. space-y-6 when any child has a description,
// space-y-3 otherwise — matches Catalyst's CheckboxGroup behavior.
const groupBase = "space-y-3 has-data-[slot=description]:space-y-6 has-data-[slot=description]:**:data-[slot=label]:font-medium"

// fieldBase positions the Checkbox (data-slot=control, col 1) alongside Label
// (data-slot=label, col 2 row 1) and Description (data-slot=description, col 2 row 2).
const fieldBase = "grid grid-cols-[1.125rem_1fr] gap-x-4 gap-y-1 sm:grid-cols-[1rem_1fr] " +
	"*:data-[slot=control]:col-start-1 *:data-[slot=control]:row-start-1 *:data-[slot=control]:mt-0.75 sm:*:data-[slot=control]:mt-1 " +
	"*:data-[slot=label]:col-start-2 *:data-[slot=label]:row-start-1 " +
	"*:data-[slot=description]:col-start-2 *:data-[slot=description]:row-start-2 " +
	"has-data-[slot=description]:**:data-[slot=label]:font-medium"

func labelClasses(p Props) string {
	if p.Class == "" {
		return labelBase
	}
	return labelBase + " " + p.Class
}

func joinClasses(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
