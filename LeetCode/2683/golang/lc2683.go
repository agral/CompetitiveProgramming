package lc2683

func doesValidArrayExist(derived []int) bool {
	xorsum := 0
	for _, val := range derived {
		xorsum ^= val
	}
	return xorsum == 0
}
