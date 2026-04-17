package toast

import "strings"

// placementClasses positions the Container at one of four corners. The
// container is `pointer-events-none fixed inset-0 flex flex-col` so it
// covers the viewport without intercepting clicks; the placement classes
// just align the toast stack within that area.
//
// Stacks always grow with new toasts after older ones in the array, which
// renders top-to-bottom in DOM order. For top placements that means new
// toasts appear *below* existing ones; for bottom placements they appear
// *below* and visually push the older ones up. Both are conventional —
// no flex-col-reverse trickery in v0.1.
var placementClasses = map[Placement]string{
	PlacementBottomRight: "justify-end items-end",
	PlacementBottomLeft:  "justify-end items-start",
	PlacementTopRight:    "justify-start items-end",
	PlacementTopLeft:     "justify-start items-start",
}

func resolvePlacement(p Placement) string {
	if c, ok := placementClasses[p]; ok {
		return c
	}
	return placementClasses[PlacementBottomRight]
}

// containerClasses builds the wrapper. inset-0 + pointer-events-none lets
// it cover the viewport without blocking the page; padding gives the
// toast stack breathing room from the edges; gap-2 spaces stacked toasts.
func containerClasses(p ContainerProps) string {
	parts := []string{
		"pointer-events-none fixed inset-0 z-50 flex flex-col gap-2 p-4 sm:p-6",
		resolvePlacement(p.Placement),
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// toastClasses builds the per-toast surface. pointer-events-auto re-enables
// interaction inside the otherwise-pass-through container. The colored
// left border is variant-driven via Tailwind's data-[variant=X] modifier
// so the same class string handles all four variants without conditional
// rendering. relative + overflow-hidden so the absolutely-positioned
// progress bar at the bottom respects the rounded corners. group is
// declared so the progress bar can read the toast's data-variant via
// group-data-[variant=X] without duplicating the attribute.
//
// The flint-toast-in-* class triggers the slide-in CSS @keyframes
// animation when the toast is mounted by x-for. Animation rather than
// transition because CSS transitions don't fire reliably on freshly
// inserted elements — the start state is rarely painted before the
// end state, so the change isn't seen. Animations always run.
func toastClasses(p Placement) string {
	return strings.Join([]string{
		"flint-toast-base",
		"group relative overflow-hidden",
		"pointer-events-auto",
		"w-80 max-w-[calc(100vw-2rem)]",
		"rounded-lg bg-surface p-4 shadow-lg",
		"ring-1 ring-border",
		"border-l-4",
		"data-[variant=info]:border-l-primary",
		"data-[variant=success]:border-l-success",
		"data-[variant=warning]:border-l-warning",
		"data-[variant=danger]:border-l-danger",
		enterAnimationClass(p),
	}, " ")
}

// enterAnimationClass picks the right CSS @keyframes animation based on
// placement: right-placed toasts slide in from the right edge, left-placed
// from the left.
func enterAnimationClass(p Placement) string {
	switch p {
	case PlacementBottomLeft, PlacementTopLeft:
		return "flint-toast-in-left"
	default:
		return "flint-toast-in-right"
	}
}

// progressBarClasses styles the countdown bar at the bottom of the toast.
// transform-origin: left + the @keyframes scaleX(1)→scaleX(0) means the
// bar drains from the right edge toward the left, leaving the remaining
// time anchored to the left. Tinted to match the toast's variant via the
// group-data- modifier (toast root carries `group` + `data-variant=X`).
func progressBarClasses() string {
	return strings.Join([]string{
		"absolute bottom-0 left-0 h-1 w-full origin-left",
		"group-data-[variant=info]:bg-primary",
		"group-data-[variant=success]:bg-success",
		"group-data-[variant=warning]:bg-warning",
		"group-data-[variant=danger]:bg-danger",
	}, " ")
}

// leavingBindExpr is the x-bind:class expression that adds the slide-out
// class when t.leaving is true. Toggled by dismiss(); the toast carries
// .flint-toast-base from mount which provides the transition rule, so the
// property change interpolates over 200ms. Right-placed toasts slide out
// the right edge; left-placed mirror to the left.
func leavingBindExpr(p Placement) string {
	cls := "flint-toast-out-right"
	if p == PlacementBottomLeft || p == PlacementTopLeft {
		cls = "flint-toast-out-left"
	}
	return "{ '" + cls + "': t.leaving }"
}

// dismissClasses styles the X button. Negative margin gives it more
// hit area than its visual size suggests; data-focus ring is consistent
// with the rest of flint-ui's interactive elements.
func dismissClasses() string {
	return strings.Join([]string{
		"-m-1 shrink-0 rounded p-1",
		"text-muted-foreground hover:text-foreground",
		"focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring",
	}, " ")
}

// alpineData is the x-data payload on Container. It owns the toasts array,
// schedules per-toast auto-dismiss, and exposes pause/resume for hover and
// focus interactions. Toast objects are plain JS — no flint-ui state lives
// outside this scope, so calling code only needs to dispatch the event.
//
// The script is kept inline (rather than in a global helper) so the
// component is self-contained — clients only need to render Container
// and load Alpine; no separate JS bundle, no script tag ordering trap
// beyond the Modal/Dropdown requirement.
// alpineDataJS — Alpine state for the Container. Notable choice: dismiss()
// runs the leave animation manually instead of using x-transition. It sets
// t.leaving = true (which the toast's x-bind:class picks up to apply the
// .flint-toast-out-* class), then setTimeout(200ms)s the actual array
// splice so the element stays in the DOM long enough for the CSS
// transition to play out. Same pattern as Pines UI's burnToast() —
// x-transition is unreliable on x-for-mounted elements, so we don't use it.
//
// dismiss() is idempotent: a second call on the same toast (e.g. ESC after
// the X click already started the leave) is a no-op via the t.leaving guard.
const alpineDataJS = `{
	toasts: [],
	nextId: 0,
	leaveDuration: 200,
	add(detail) {
		if (!detail) return;
		const allowed = ['info','success','warning','danger'];
		const variant = allowed.includes(detail.variant) ? detail.variant : 'info';
		const t = {
			id: ++this.nextId,
			variant: variant,
			title: String(detail.title || ''),
			body: String(detail.body || ''),
			duration: typeof detail.duration === 'number' ? detail.duration : 5000,
			paused: false,
			leaving: false,
			_timeout: null
		};
		this.toasts.push(t);
		if (t.duration > 0) this._schedule(t);
	},
	_schedule(t) {
		t._timeout = setTimeout(() => this.dismiss(t.id), t.duration);
	},
	pause(t) {
		if (t.leaving) return;
		if (t._timeout) {
			clearTimeout(t._timeout);
			t._timeout = null;
			t.paused = true;
		}
	},
	resume(t) {
		if (t.leaving) return;
		if (t.paused && t.duration > 0) {
			t.paused = false;
			this._schedule(t);
		}
	},
	dismiss(id) {
		const t = this.toasts.find(x => x.id === id);
		if (!t || t.leaving) return;
		if (t._timeout) { clearTimeout(t._timeout); t._timeout = null; }
		t.leaving = true;
		setTimeout(() => {
			this.toasts = this.toasts.filter(x => x.id !== id);
		}, this.leaveDuration);
	}
}`

func alpineData() string {
	return alpineDataJS
}
