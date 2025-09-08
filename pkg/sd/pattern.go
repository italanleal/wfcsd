package sd

// Pattern represents a combination of exactly two items with associated positive row indices
type Pattern struct {
	Items []Item // subgroup items
	Index []int  // indices of rows where this pattern occurs (positive cases)
	Freq  int    // absolute frequency = len(Index)
}
