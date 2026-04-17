package descriptionlist

import "strings"

// Non-color classes preserved verbatim from catalyst-ui-kit/typescript/
// description-list.tsx. zinc/dark-mode color pairs collapse to
// foreground + muted-foreground via the semantic token contract.

// rootClasses: single-column on mobile, two-column on sm+ where the
// label column is min(50%, --spacing(80)) — roughly 20rem max.
// `text-base/6` on mobile, `text-sm/6` on sm+ for denser desktop
// reading.
func rootClasses(p Props) string {
	parts := []string{
		"grid grid-cols-1 text-base/6",
		"sm:grid-cols-[min(50%,--spacing(80))_auto] sm:text-sm/6",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// termClasses: the <dt> label column. Top border separates rows;
// `first:border-none` removes the border above the first term so the
// list sits cleanly inside a Card body. Muted text so values read
// primary.
func termClasses(p SubProps) string {
	parts := []string{
		"col-start-1 border-t border-foreground/5 pt-3 text-muted-foreground first:border-none",
		"sm:border-t sm:border-foreground/5 sm:py-3",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}

// detailsClasses: the <dd> value column. Foreground text for
// emphasis. On mobile, no top border (term's border covers the
// pair). On sm+, a top border matches the term's, except for the
// second element (first <dd>) where `nth-2:border-none` keeps the
// first row unbroken.
func detailsClasses(p SubProps) string {
	parts := []string{
		"pt-1 pb-3 text-foreground",
		"sm:border-t sm:border-foreground/5 sm:py-3 sm:nth-2:border-none",
	}
	if p.Class != "" {
		parts = append(parts, p.Class)
	}
	return strings.Join(parts, " ")
}
