package fzf

// maxWidthMargin returns the columns to add to the right margin so that the
// finder, from its left border to its right border, is at most --max-width
// columns wide. The finder stays at its left margin.
func (t *Terminal) maxWidthMargin(screenWidth int, marginInt [4]int, extraMargin [4]int) int {
	if t.maxWidth <= 0 {
		return 0
	}
	// extraMargin holds the border columns, which belong to the finder
	finderWidth := screenWidth - (marginInt[1] - extraMargin[1]) - (marginInt[3] - extraMargin[3])
	return max(0, finderWidth-t.maxWidth)
}
