package slideover

import (
	"fmt"
	"strings"
)

// Slideover reuses most of Modal's structural vocabulary (backdrop tint,
// panel surface tokens, title/description/body typography) with two
// diffs: the panel is edge-anchored with a height-filling flex column,
// and the enter/leave transitions translate along X instead of fading +
// scaling. Non-color classes are flint-ui originals — there's no
// Catalyst source to track verbatim.

// sizeClasses maps Size to a max-width cap applied at sm+. Below sm the
// panel is w-screen so short viewports don't see a cramped drawer.
var sizeClasses = map[Size]string{
	SizeSM:  "sm:max-w-sm",
	SizeMD:  "sm:max-w-md",
	SizeLG:  "sm:max-w-lg",
	SizeXL:  "sm:max-w-xl",
	Size2XL: "sm:max-w-2xl",
}

func resolveSize(s Size) string {
	if c, ok := sizeClasses[s]; ok {
		return c
	}
	return sizeClasses[SizeLG]
}

func resolveSide(s Side) Side {
	if s == SideLeft {
		return SideLeft
	}
	return SideRight
}

// backdropClasses dims the page behind the panel. Same opacity as
// Modal's dialog variant (25%) so the two components feel like part of
// the same system.
func backdropClasses(_ Props) string {
	return "fixed inset-0 bg-foreground/25 focus:outline-0"
}

// anchorClasses positions the panel-holder flush to the chosen edge,
// full-height. `flex` + `max-w-full` lets the inner panel set its own
// width cap via max-w-* without blowing out on narrow viewports.
func anchorClasses(p Props) string {
	switch resolveSide(p.Side) {
	case SideLeft:
		return "fixed inset-y-0 left-0 flex max-w-full"
	default:
		return "fixed inset-y-0 right-0 flex max-w-full"
	}
}

// panelClasses: the visible drawer surface. flex-col lays out Title /
// Description / Body / Actions vertically; Body's flex-1 + Actions'
// mt-auto handle the "scroll middle, pin buttons" behavior without
// positional tricks.
func panelClasses(p Props) string {
	parts := []string{
		"relative flex h-full w-screen flex-col",
		resolveSize(p.Size),
		"bg-surface", "shadow-xl",
		"ring-1", "ring-border",
		"p-6", "sm:p-8",
		"forced-colors:outline",
		"will-change-transform",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// panelHiddenClasses and panelVisibleClasses drive the Alpine
// enter/leave transitions. Right-side panels slide out of the right
// edge (translate-x-full) when hidden; left-side panels slide out of
// the left edge (-translate-x-full). Visible state is translate-x-0
// regardless of side.
func panelHiddenClasses(p Props) string {
	if resolveSide(p.Side) == SideLeft {
		return "-translate-x-full"
	}
	return "translate-x-full"
}

func panelVisibleClasses() string {
	return "translate-x-0"
}

// titleClasses matches Modal's dialog title typography so a slideover
// that holds form content reads with the same weight as a modal dialog
// of the same purpose.
func titleClasses(p SubProps) string {
	parts := []string{
		"text-lg/6", "font-semibold", "text-balance",
		"text-foreground",
		"sm:text-base/6",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func descriptionClasses(p SubProps) string {
	parts := []string{
		"mt-2", "text-pretty",
		"text-base/6", "text-muted-foreground",
		"sm:text-sm/6",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// bodyClasses: flex-1 makes Body the stretchy child, overflow-y-auto
// puts any scroll inside the body so title + actions stay pinned.
// -mx-* + px-* restores the horizontal padding edge-to-edge, so a
// scrollbar sits flush to the panel edge instead of inside the padded
// column (matches the Tailwind UI slide-over pattern).
func bodyClasses(p SubProps) string {
	parts := []string{
		"mt-6",
		"flex-1", "overflow-y-auto",
		"-mx-6", "px-6",
		"sm:-mx-8", "sm:px-8",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// actionsClasses: mt-auto pins the row to the bottom if Body is short
// or absent. Same responsive column-reverse → row pattern as Modal so
// the primary button is always on top on mobile and right on desktop.
// A top border visually separates the pinned footer from the body.
func actionsClasses(p SubProps) string {
	parts := []string{
		"mt-auto", "pt-6",
		"border-t", "border-border",
		"flex", "flex-col-reverse", "items-center", "justify-end",
		"gap-3", "*:w-full",
		"sm:flex-row", "sm:*:w-auto",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// alpineData is the inline x-data expression placed on the slideover
// root. Mirrors modal.alpineData — name is validated by safeName so a
// malformed name produces an inert slideover instead of a syntax error.
func alpineData(name string) string {
	return fmt.Sprintf("{ open: false, name: '%s' }", safeName(name))
}

// openHandler listens for the global open event and matches on name so
// a single dispatch can address one slideover on a page with many.
func openHandler(name string) string {
	return fmt.Sprintf("if ($event.detail === '%s') { open = true }", safeName(name))
}

// closeHandler: no detail closes every slideover (useful after a
// successful htmx form submit); a detail matching this slideover's name
// closes just this one.
func closeHandler(name string) string {
	return fmt.Sprintf("if (!$event.detail || $event.detail === '%s') { open = false }", safeName(name))
}

// safeName keeps only characters safe to embed in a single-quoted JS
// string literal. Same rule as modal.safeName and tabs.safeName.
func safeName(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '-', r == '_', r == ':', r == '.':
			b.WriteRune(r)
		}
	}
	return b.String()
}
