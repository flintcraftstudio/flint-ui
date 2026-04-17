package radio

// Structure mirrors components/checkbox/classes.go — same native-input
// + peer-* variant approach, with rounded-full everywhere and a
// centered inner-dot span instead of a checkmark SVG.
//
// Catalyst's 20+ color variants collapse to the primary token.
// Per-client rebranding flows through the token contract; if a client
// needs a danger-colored radio for a destructive-choice form, pass
// Class with a bg-danger override on the peer-checked state.

// labelBase: group-relative wrapper so the absolutely-positioned
// input + circle stack correctly. size matches Catalyst's radio size
// (4.75 on mobile, 4.25 on sm+) — same as Checkbox.
const labelBase = "group relative isolate inline-flex size-[1.125rem] shrink-0 items-center justify-center rounded-full sm:size-4"

// inputClasses: the native <input type="radio">. Absolutely positioned
// full-bleed with opacity-0 so the whole label bounds are clickable.
// `peer` exposes state to sibling spans via peer-checked:,
// peer-focus-visible:, peer-disabled:, peer-data-invalid:.
const inputClasses = "peer absolute inset-0 z-20 m-0 cursor-pointer appearance-none rounded-full opacity-0 disabled:cursor-not-allowed"

// circleClasses: the visible outer ring. Gets the background, border,
// focus ring, and disabled fade. peer-checked swaps bg to primary and
// hides the border so the checked radio reads as a filled disc.
const circleClasses = "absolute inset-0 z-0 rounded-full " +
	"bg-input shadow-sm " +
	"before:absolute before:inset-0 before:rounded-full before:shadow-[inset_0_1px_--theme(--color-white/15%)] " +
	"border border-border " +
	"peer-hover:border-foreground/30 " +
	"peer-checked:bg-primary peer-checked:border-transparent peer-checked:shadow-none " +
	"peer-focus-visible:outline-2 peer-focus-visible:outline-offset-2 peer-focus-visible:outline-ring " +
	"peer-data-invalid:border-danger peer-hover:peer-data-invalid:border-danger " +
	"peer-disabled:opacity-50 peer-disabled:bg-muted peer-disabled:border-border peer-disabled:peer-checked:bg-muted"

// dotClasses: the centered inner dot. Fades in on peer-checked;
// stays z-above the circle so it's visible on the primary fill.
// size is ~1/3 of the radio diameter — size-1.5 on mobile, size-1
// on sm+ matches the scale Catalyst achieves via its border-[4.5px]
// approach.
const dotClasses = "relative z-10 size-1.5 rounded-full bg-primary-foreground opacity-0 peer-checked:opacity-100 peer-disabled:bg-muted-foreground sm:size-1 forced-colors:bg-[HighlightText]"

// groupBase: same rhythm as checkbox — space-y-3 by default, space-y-6
// when any child has a description, and labels in description-bearing
// groups bump to font-medium for visual weight.
const groupBase = "space-y-3 has-data-[slot=description]:space-y-6 has-data-[slot=description]:**:data-[slot=label]:font-medium"

// fieldBase: two-column grid — Radio in col 1, Label in col 2 row 1,
// Description in col 2 row 2. Identical to Checkbox's fieldBase since
// the two components have the same visual footprint.
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
