package input

import (
	"strings"
	"testing"
)

// TestText16px: the text never drops below 16px, or iOS Safari zooms into
// the field when it gets focus.
func TestText16px(t *testing.T) {
	if !strings.Contains(inputBase, "text-base/6") || strings.Contains(inputBase, "text-sm") || strings.Contains(inputBase, "text-xs") {
		t.Errorf("inputBase = %q, want text-base/6 at every width", inputBase)
	}
}
