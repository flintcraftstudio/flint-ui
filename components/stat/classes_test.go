package stat

import (
	"testing"

	"github.com/flintcraftstudio/flint-ui/components/badge"
)

func TestChangeVariant(t *testing.T) {
	tests := []struct {
		name string
		p    Props
		want badge.Variant
	}{
		{"plus is good", Props{Change: "+4.5%"}, badge.VariantSuccess},
		{"hyphen minus is bad", Props{Change: "-0.5%"}, badge.VariantDanger},
		{"typographic minus is bad", Props{Change: "−2"}, badge.VariantDanger},
		{"leading space ignored", Props{Change: " +3"}, badge.VariantSuccess},
		{"unsigned is flat", Props{Change: "0%"}, badge.VariantMuted},
		{"lower is better flips up", Props{Change: "+2", LowerIsBetter: true}, badge.VariantDanger},
		{"lower is better flips down", Props{Change: "-2", LowerIsBetter: true}, badge.VariantSuccess},
		{"explicit trend wins over sign", Props{Change: "+1", Trend: TrendDown}, badge.VariantDanger},
		{"explicit flat", Props{Change: "+1", Trend: TrendFlat}, badge.VariantMuted},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeVariant(tt.p); got != tt.want {
				t.Errorf("changeVariant(%+v) = %q, want %q", tt.p, got, tt.want)
			}
		})
	}
}
