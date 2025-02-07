package lc3160

func queryResults(limit int, queries [][]int) []int {
	qLength := len(queries)
	ans := make([]int, qLength)
	ballToColor := make(map[int]int)
	colorToCount := make(map[int]int)

	for i, query := range queries {
		ballNumber, color := query[0], query[1]
		oldColor, ok := ballToColor[ballNumber]
		if ok {
			colorToCount[oldColor] -= 1
			if colorToCount[oldColor] == 0 {
				delete(colorToCount, oldColor)
			}
		}
		ballToColor[ballNumber] = color
		colorToCount[color]++
		ans[i] = len(colorToCount)
	}

	return ans
}
