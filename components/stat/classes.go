package stat

import (
	"strings"

	"github.com/flintcraftstudio/flint-ui/components/badge"
)

// Non-color classes preserved verbatim from Catalyst's demo stat.tsx and
// its dashboard grid. Badge colors lime / pink map to success / danger;
// text-zinc-500 maps to muted-foreground.

const gridClasses = "grid gap-8 sm:grid-cols-2 xl:grid-cols-4"

// statClasses reproduces Catalyst's <Divider> (border-zinc-950/10).
const statClasses = "border-t border-foreground/10"

const labelClasses = "mt-6 text-lg/6 font-medium text-foreground sm:text-sm/6"

const valueClasses = "mt-3 text-3xl/8 font-semibold text-foreground sm:text-2xl/8"

const changeClasses = "mt-3 text-sm/6 sm:text-xs/6"

// resolveTrend returns the explicit Trend, or infers it from the sign of
// Change. U+2212 (the typographic minus) counts as down.
func resolveTrend(p Props) Trend {
	if p.Trend != TrendAuto {
		return p.Trend
	}
	c := strings.TrimSpace(p.Change)
	switch {
	case strings.HasPrefix(c, "+"):
		return TrendUp
	case strings.HasPrefix(c, "-"), strings.HasPrefix(c, "−"):
		return TrendDown
	}
	return TrendFlat
}

// changeVariant colors the badge: good changes are success, bad are
// danger, flat is muted. LowerIsBetter swaps good and bad.
func changeVariant(p Props) badge.Variant {
	good, bad := badge.VariantSuccess, badge.VariantDanger
	if p.LowerIsBetter {
		good, bad = bad, good
	}
	switch resolveTrend(p) {
	case TrendUp:
		return good
	case TrendDown:
		return bad
	}
	return badge.VariantMuted
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
