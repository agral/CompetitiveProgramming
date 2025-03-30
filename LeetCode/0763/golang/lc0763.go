package lc0763

func partitionLabels(s string) []int {
	ans := []int{}
	letterToLastPos := make(map[rune]int)
	for i, ch := range s {
		letterToLastPos[ch] = i
	}
	l, r := 0, 0 // boundaries of the currently considered partition
	for i, ch := range s {
		r = max(r, letterToLastPos[ch])
		if r == i {
			ans = append(ans, r-l+1)
			l = r + 1
		}
	}

	return ans
}
