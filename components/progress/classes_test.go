package progress

import "testing"

func TestValues(t *testing.T) {
	tests := []struct {
		name      string
		p         Props
		wantNow   string
		wantMax   string
		wantText  string
		wantWidth string
	}{
		{"default max", Props{Value: 42}, "42", "100", "42%", "width: 42%"},
		{"custom max", Props{Value: 3, Max: 4}, "3", "4", "75%", "width: 75%"},
		{"rounds text, keeps exact width", Props{Value: 1, Max: 3}, "1", "3", "33%", "width: 33.33333333333333%"},
		{"clamps above max", Props{Value: 140}, "100", "100", "100%", "width: 100%"},
		{"clamps below zero", Props{Value: -5}, "0", "100", "0%", "width: 0%"},
		{"negative max means 100", Props{Value: 50, Max: -1}, "50", "100", "50%", "width: 50%"},
		{"explicit value text", Props{Value: 2, Max: 5, ValueText: "Step 2 of 5"}, "2", "5", "Step 2 of 5", "width: 40%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatNumber(clamp(tt.p)); got != tt.wantNow {
				t.Errorf("aria-valuenow = %q, want %q", got, tt.wantNow)
			}
			if got := formatNumber(maxValue(tt.p)); got != tt.wantMax {
				t.Errorf("aria-valuemax = %q, want %q", got, tt.wantMax)
			}
			if got := valueText(tt.p); got != tt.wantText {
				t.Errorf("aria-valuetext = %q, want %q", got, tt.wantText)
			}
			if got := string(fillStyle(tt.p)); got != tt.wantWidth {
				t.Errorf("style = %q, want %q", got, tt.wantWidth)
			}
		})
	}
}
