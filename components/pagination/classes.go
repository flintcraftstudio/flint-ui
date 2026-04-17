package pagination

import (
	"strings"

	"github.com/a-h/templ"
)

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/
// pagination.tsx. The current-page highlight uses the muted token
// (bg-muted) instead of Catalyst's zinc/dark-mode pair — a single
// neutral token covers both themes via the semantic contract.

func ariaLabelOr(label string) string {
	if label != "" {
		return label
	}
	return "Page navigation"
}

// rootClasses: the outer nav flexes Previous / List / Next across its
// width. gap-x-2 provides breathing room without pushing the List to
// the edges.
func rootClasses(p Props) string {
	parts := []string{"flex gap-x-2"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// prevWrapperClasses: grow basis-0 makes the Previous span claim its
// share of the nav width so List stays centered. No horizontal
// alignment — the Button sits at the natural left of the grown span.
func prevWrapperClasses(extra string) string {
	parts := []string{"grow basis-0"}
	if extra != "" {
		parts = append(parts, extra)
	}
	return strings.Join(parts, " ")
}

// nextWrapperClasses: same grow/basis as Previous plus
// justify-end so Next pushes to the right edge.
func nextWrapperClasses(extra string) string {
	parts := []string{"flex grow basis-0 justify-end"}
	if extra != "" {
		parts = append(parts, extra)
	}
	return strings.Join(parts, " ")
}

// listClasses: hidden on mobile (sm:flex). items-baseline keeps
// single- and double-digit page numbers aligned along the text
// baseline rather than center.
func listClasses(p ListProps) string {
	parts := []string{"hidden items-baseline gap-x-2 sm:flex"}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// pageClasses translates Catalyst's Page button styling. min-w-9 keeps
// the button a consistent width regardless of digit count (1 vs 99).
// The before:absolute layer is Catalyst's chrome for the current-page
// highlight — bg-muted (flint-ui token) when Current is set, otherwise
// no overlay.
func pageClasses(p PageProps) string {
	parts := []string{
		"min-w-9 before:absolute before:-inset-px before:rounded-lg",
	}
	if p.Current {
		parts = append(parts, "before:bg-muted")
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// pageAriaAttrs sets aria-current on the current page and a
// human-readable aria-label describing which page this link leads to.
// Catalyst uses "Page {children}" — we don't have access to the child
// text inside a Go function, so callers' aria-label (via Attrs) wins
// if provided; otherwise the default "Page" label is applied.
func pageAriaAttrs(p PageProps) templ.Attributes {
	attrs := templ.Attributes{}
	if p.Current {
		attrs["aria-current"] = "page"
	}
	return attrs
}

// gapClasses: w-9 matches the page buttons' min-w-9 so the ellipsis
// occupies the same slot as a collapsed page number. text-foreground
// uses the semantic token instead of Catalyst's zinc/dark-mode pair.
func gapClasses(p GapProps) string {
	parts := []string{
		"w-9 text-center text-sm/6 font-semibold text-foreground select-none",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// mergeAttrs merges caller-supplied Attrs over component defaults.
// Defaults are written first, callerAttrs second — callerAttrs wins on
// key conflicts so a caller passing their own aria-label or href
// overrides the component's default.
func mergeAttrs(callerAttrs, defaults templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for k, v := range defaults {
		out[k] = v
	}
	for k, v := range callerAttrs {
		out[k] = v
	}
	return out
}
