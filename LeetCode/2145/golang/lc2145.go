package lc2145

func numberOfArrays(differences []int, lower int, upper int) int {
	// Int can overflow, have to use int64s just to be sure.
	current, low, high := int64(0), int64(0), int64(0)
	for _, difference := range differences {
		current += int64(difference)
		low = min(low, current)
		high = max(high, current)
	}
	return max(upper-lower-int(high-low)+1, 0)
}
