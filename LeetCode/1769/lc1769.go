package lc1769

func minOperations(boxes string) []int {
	sz := len(boxes)
	ans := make([]int, sz)

	for i, moves, count := 0, 0, 0; i < sz; i++ {
		ans[i] = moves
		count += int(boxes[i] - '0')
		moves += count
	}

	for i, moves, count := sz-1, 0, 0; i >= 0; i-- {
		ans[i] += moves
		count += int(boxes[i] - '0')
		moves += count
	}

	return ans
}
