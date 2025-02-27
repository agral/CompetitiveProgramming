package lc0873

func lenLongestFibSubseq(arr []int) int {
	SZ := len(arr)

	//dp is initially an SZ by SZ 2D array filled with values of 2.
	dp := make([][]int, SZ)
	for i := range SZ {
		dp[i] = make([]int, SZ)
		for j := range SZ {
			dp[i][j] = 2
		}
	}

	// numToIndex maps a number from arr to its index in arr.
	numToIndex := make(map[int]int)
	for i := range SZ {
		numToIndex[arr[i]] = i
	}

	ans := 0

	for y := range SZ {
		for x := y + 1; x < SZ; x++ {
			diff := arr[x] - arr[y]
			if diff < arr[y] {
				index, hasKey := numToIndex[diff]
				if hasKey {
					dp[y][x] = dp[index][y] + 1
					ans = max(ans, dp[y][x])
				}
			}
		}
	}

	return ans
}
