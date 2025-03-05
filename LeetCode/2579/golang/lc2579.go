package lc2579

func coloredCells(n int) int64 {
	n64 := int64(n)
	return 1 + 2*n64*(n64-1)
}
