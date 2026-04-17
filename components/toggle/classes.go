package toggle

// Approach: the <label> renders the track (its own background fills
// the pill shape); a nested span is the thumb. The input is sr-only
// so it's still in the accessibility tree and form flow, but visual
// styling lives on the label and thumb.
//
// State propagation uses Tailwind's has-* variants. has-checked
// wraps :has(:checked) — the label styles itself based on whether
// its descendant input is checked. has-focus-visible / has-disabled
// / has-data-invalid do the same for the other states. No peer-*
// needed since the input is inside the label, not a sibling of the
// styled elements.

// labelBase is the track + visible shape. w-10 h-6 on mobile,
// w-8 h-5 on sm+ matches Catalyst. p-[3px] reserves padding inside
// the track for the thumb to translate within. items-center aligns
// the thumb vertically. has-checked swaps the track background to
// the primary token — client rebranding flows through automatically.
const labelBase = "group relative inline-flex h-6 w-10 shrink-0 cursor-pointer items-center rounded-full p-[3px] transition " +
	"bg-muted ring-1 ring-inset ring-border " +
	"has-checked:bg-primary has-checked:ring-transparent " +
	"has-focus-visible:outline-2 has-focus-visible:outline-offset-2 has-focus-visible:outline-ring " +
	"has-data-invalid:ring-danger has-disabled:opacity-50 has-disabled:cursor-not-allowed " +
	"sm:h-5 sm:w-8 " +
	"forced-colors:outline forced-colors:has-checked:bg-[Highlight]"

// thumbClasses is the movable circle inside the track. translate-x-0
// at rest, translate-x-4 (sm:-3) when the input is checked — group-
// has-checked targets :has(:checked) on the ancestor `.group`. Size
// leaves a 3px margin inside the track so the thumb sits flush.
const thumbClasses = "pointer-events-none inline-block size-[18px] sm:size-3.5 rounded-full " +
	"bg-surface shadow-sm ring-1 ring-border " +
	"translate-x-0 group-has-checked:translate-x-4 sm:group-has-checked:translate-x-3 " +
	"transition duration-200 ease-in-out " +
	"border border-transparent"

// groupBase: same rhythm as Checkbox and Radio — space-y-3 by default,
// space-y-6 + medium-weight labels when any child has a description.
const groupBase = "space-y-3 has-data-[slot=description]:space-y-6 has-data-[slot=description]:**:data-[slot=label]:font-medium"

// fieldBase: two-column grid with label on the left, switch on the
// right. Opposite of Checkbox/Radio — matches Catalyst's SwitchField.
// Reads naturally as "Email notifications: [toggle]".
const fieldBase = "grid grid-cols-[1fr_auto] gap-x-8 gap-y-1 " +
	"*:data-[slot=control]:col-start-2 *:data-[slot=control]:self-start sm:*:data-[slot=control]:mt-0.5 " +
	"*:data-[slot=label]:col-start-1 *:data-[slot=label]:row-start-1 " +
	"*:data-[slot=description]:col-start-1 *:data-[slot=description]:row-start-2 " +
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
