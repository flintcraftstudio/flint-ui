// Package tooltip provides accessible hint bubbles over a trigger. From-
// scratch design (Catalyst has no tooltip.tsx).
//
// Architecture: the Tooltip component is a wrapper <span> that holds Alpine
// { open: false } state. The caller's trigger goes inside as children; the
// bubble itself is appended after the children as an absolute-positioned
// sibling. Events live on the wrapper — mouseenter/mouseleave for hover,
// focusin/focusout for keyboard (these bubble from any inner focusable,
// so buttons, links, inputs all work without extra wiring), Escape to
// dismiss.
//
// A11y: the bubble renders with role="tooltip" and id = Props.ID. The
// caller MUST set aria-describedby={ID} on the actual focusable trigger
// (typically via button.Props.Attrs) so screen readers announce the
// tooltip when the trigger gains focus. Without that binding the tooltip
// is visual-only.
//
// No teleport — the bubble is absolute-positioned inside the wrapper.
// Ancestor overflow:hidden / transform / filter can clip it, the same
// caveat Dropdown has. For tooltips near the edge of the viewport, flip
// the Anchor.
//
// Content is a plain string (Props.Label). Tooltips are for short hints;
// for rich content use Popover (future). Keeping this to text avoids the
// children-vs-content ambiguity Modal/Dropdown have to manage.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/tooltip.md in the flint-ui module.
package tooltip
