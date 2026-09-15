package tui

import (
	"testing"

	"github.com/zzusec/restore-session/internal/tui/text"
)

// TestFormatBytes pins the list's size column: every output is exactly
// sizeWidth cells so the unit suffixes line up across rows, and the boundaries
// between B / KB / MB / GB land where a glance at the list would expect.
func TestFormatBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
		want string
	}{
		{"empty session shows a dash", 0, "     —"},
		{"negative is treated as empty", -1, "     —"},
		{"a few bytes", 12, "   12B"},
		{"just under a kilobyte", 1023, " 1023B"},
		{"one kilobyte", 1 << 10, "   1KB"},
		{"tens of kilobytes", 45<<10 + 200, "  45KB"},
		{"hundreds of kilobytes", 500 << 10, " 500KB"},
		{"one megabyte exactly drops the decimal", 1 << 20, "   1MB"},
		{"a few megabytes with a decimal", 1234 << 10, " 1.2MB"},
		{"twenty megabytes has no decimal", 20 << 20, "  20MB"},
		{"near a gigabyte stays in megabytes without overflowing", 1023<<20 - 1, "1023MB"},
		{"one gigabyte exactly drops the decimal", 1 << 30, "   1GB"},
		{"a couple of gigabytes with a decimal", int64(2.5 * float64(1<<30)), " 2.5GB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := formatBytes(test.in)
			if got != test.want {
				t.Errorf("formatBytes(%d) = %q, want %q", test.in, got, test.want)
			}
			if w := text.Width(got); w != sizeWidth {
				t.Errorf("formatBytes(%d) is %d cells wide, want fixed %d", test.in, w, sizeWidth)
			}
		})
	}
}
