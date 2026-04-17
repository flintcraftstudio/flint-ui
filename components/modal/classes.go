package modal

import (
	"fmt"
	"strings"
)

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/dialog.tsx
// and alert.tsx. Color/opacity classes are translated to the semantic token
// contract. Catalyst's Headless UI data-closed / data-enter / data-leave
// transition classes are replaced with Alpine x-transition — the modal.templ
// wires enter-start / enter-end / leave-start / leave-end to the panelHidden /
// panelVisible helpers below.

// sizeClasses mirrors Catalyst's sizes map: xs..5xl as sm: max-width constraints.
var sizeClasses = map[Size]string{
	SizeXS:  "sm:max-w-xs",
	SizeSM:  "sm:max-w-sm",
	SizeMD:  "sm:max-w-md",
	SizeLG:  "sm:max-w-lg",
	SizeXL:  "sm:max-w-xl",
	Size2XL: "sm:max-w-2xl",
	Size3XL: "sm:max-w-3xl",
	Size4XL: "sm:max-w-4xl",
	Size5XL: "sm:max-w-5xl",
}

func resolveSize(p Props) string {
	s := p.Size
	if s == "" {
		if p.Alert {
			s = SizeMD
		} else {
			s = SizeLG
		}
	}
	if c, ok := sizeClasses[s]; ok {
		return c
	}
	if p.Alert {
		return sizeClasses[SizeMD]
	}
	return sizeClasses[SizeLG]
}

// backdropClasses covers the full viewport behind the panel. Alert uses a
// lighter wash (15%) than Dialog (25%) — Catalyst's visual distinction
// between an attention-grabbing modal and a softer confirm.
func backdropClasses(p Props) string {
	opacity := "bg-foreground/25"
	if p.Alert {
		opacity = "bg-foreground/15"
	}
	return "fixed inset-0 " + opacity + " focus:outline-0"
}

// scrollClasses: the outer overflow container. Matches Catalyst's
// `fixed inset-0 w-screen overflow-y-auto pt-6 sm:pt-0`.
func scrollClasses(_ Props) string {
	return "fixed inset-0 w-screen overflow-y-auto pt-6 sm:pt-0"
}

// gridClasses: the grid that positions the panel. Dialog pins the panel to
// the bottom on mobile (grid-rows-[1fr_auto]) and centers-ish on sm+
// (1fr_auto_3fr). Alert centers on both — the symmetric 1fr_auto_1fr on
// mobile keeps a small dialog visually anchored.
func gridClasses(p Props) string {
	if p.Alert {
		return "grid min-h-full grid-rows-[1fr_auto_1fr] justify-items-center p-8 sm:grid-rows-[1fr_auto_3fr] sm:p-4"
	}
	return "grid min-h-full grid-rows-[1fr_auto] justify-items-center sm:grid-rows-[1fr_auto_3fr] sm:p-4"
}

// panelClasses: the visible dialog surface. Dialog uses a larger radius on
// the top corners only for the mobile bottom-sheet look, then a uniform
// rounded-2xl on sm+. Alert is rounded-2xl at all breakpoints.
func panelClasses(p Props) string {
	size := resolveSize(p)

	var parts []string
	if p.Alert {
		parts = []string{
			"row-start-2", "w-full",
			size,
			"rounded-2xl", "bg-surface", "p-8", "shadow-lg",
			"ring-1", "ring-border",
			"sm:rounded-2xl", "sm:p-6",
			"forced-colors:outline",
			"will-change-transform",
		}
	} else {
		parts = []string{
			"row-start-2", "w-full", "min-w-0",
			size,
			"rounded-t-3xl", "bg-surface", "p-(--gutter)", "shadow-lg",
			"ring-1", "ring-border",
			"[--gutter:--spacing(8)]",
			"sm:mb-auto", "sm:rounded-2xl",
			"forced-colors:outline",
			"will-change-transform",
		}
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// panelHiddenClasses and panelVisibleClasses drive the Alpine enter/leave
// transitions. Dialog slides up from 12px on mobile and scales-in on sm+;
// Alert just fades + scales on both.
func panelHiddenClasses(p Props) string {
	if p.Alert {
		return "opacity-0 scale-95"
	}
	return "opacity-0 translate-y-12 sm:translate-y-0 sm:scale-95"
}

func panelVisibleClasses() string {
	return "opacity-100 translate-y-0 scale-100"
}

func titleClasses(p SubProps) string {
	var parts []string
	if p.Alert {
		parts = []string{
			"text-center", "text-base/6", "font-semibold", "text-balance",
			"text-foreground",
			"sm:text-left", "sm:text-sm/6", "sm:text-wrap",
		}
	} else {
		parts = []string{
			"text-lg/6", "font-semibold", "text-balance",
			"text-foreground",
			"sm:text-base/6",
		}
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func descriptionClasses(p SubProps) string {
	var parts []string
	if p.Alert {
		parts = []string{
			"mt-2", "text-center", "text-pretty",
			"text-sm/6", "text-muted-foreground",
			"sm:text-left",
		}
	} else {
		parts = []string{
			"mt-2", "text-pretty",
			"text-base/6", "text-muted-foreground",
			"sm:text-sm/6",
		}
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func bodyClasses(p SubProps) string {
	var parts []string
	if p.Alert {
		parts = []string{"mt-4"}
	} else {
		parts = []string{"mt-6"}
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

func actionsClasses(p SubProps) string {
	var parts []string
	if p.Alert {
		parts = []string{
			"mt-6", "flex", "flex-col-reverse", "items-center", "justify-end",
			"gap-3", "*:w-full",
			"sm:mt-4", "sm:flex-row", "sm:*:w-auto",
		}
	} else {
		parts = []string{
			"mt-8", "flex", "flex-col-reverse", "items-center", "justify-end",
			"gap-3", "*:w-full",
			"sm:flex-row", "sm:*:w-auto",
		}
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// alpineData is the inline x-data expression placed on the modal root. It
// carries the modal name so open/close handlers can filter events targeted
// at other modals. The name is validated by safeName — anything that would
// break the JS string literal is stripped so malformed input produces an
// inert modal instead of a syntax error.
func alpineData(name string) string {
	return fmt.Sprintf("{ open: false, name: '%s' }", safeName(name))
}

// openHandler listens for the global open event and matches on name so a
// single dispatch can address one modal on a page with many.
func openHandler(name string) string {
	return fmt.Sprintf("if ($event.detail === '%s') { open = true }", safeName(name))
}

// closeHandler: no detail closes every modal (useful after a successful
// htmx form submit); a detail matching this modal's name closes just this one.
func closeHandler(name string) string {
	return fmt.Sprintf("if (!$event.detail || $event.detail === '%s') { open = false }", safeName(name))
}

// safeName keeps only characters safe to embed in a single-quoted JS string
// literal. Names are expected to be stable identifiers like 'delete-invoice'
// or 'crew:123:edit'; non-matching characters are dropped so a malformed
// name can't break the page.
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
