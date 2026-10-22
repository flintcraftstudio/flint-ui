package textarea

import (
	"strings"
	"testing"
)

// TestText16px: the text never drops below 16px, or iOS Safari zooms into
// the field when it gets focus.
func TestText16px(t *testing.T) {
	if !strings.Contains(textareaBase, "text-base/6") || strings.Contains(textareaBase, "text-sm") || strings.Contains(textareaBase, "text-xs") {
		t.Errorf("textareaBase = %q, want text-base/6 at every width", textareaBase)
	}
}
