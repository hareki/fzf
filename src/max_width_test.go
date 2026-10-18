package fzf

import "testing"

func TestMaxWidthMargin(t *testing.T) {
	// --margin=0,1 --border=rounded: one border column on the left, the
	// border and scrollbar columns on the right
	marginInt := [4]int{0, 3, 0, 2}
	extraMargin := [4]int{0, 2, 0, 1}
	for _, test := range []struct {
		maxWidth    int
		screenWidth int
		expected    int
	}{
		{0, 200, 0},
		{80, 200, 118},
		{80, 82, 0},
		{80, 60, 0},
	} {
		terminal := Terminal{maxWidth: test.maxWidth}
		if got := terminal.maxWidthMargin(test.screenWidth, marginInt, extraMargin); got != test.expected {
			t.Errorf("maxWidth %d, screenWidth %d: expected %d, got %d", test.maxWidth, test.screenWidth, test.expected, got)
		}
	}
}
