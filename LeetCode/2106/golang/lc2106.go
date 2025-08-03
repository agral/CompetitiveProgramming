package lc2106

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	len_fruits := len(fruits)
	ans := 0
	// maxRight is either startPos or the right-most index that contains a fruit,
	// whichever is greater:
	maxRight := max(startPos, fruits[len_fruits-1][0])
	amounts := make([]int, maxRight+1)
	prefix := make([]int, maxRight+2)
	for _, f := range fruits {
		amounts[f[0]] = f[1]
	}
	prefix[0] = 0
	for i := range amounts {
		prefix[i+1] = prefix[i] + amounts[i]
	}

	calculateHarvest := func(leftSteps int, rightSteps int) int {
		l := max(startPos-leftSteps, 0)
		r := min(startPos+rightSteps, maxRight)
		return prefix[r+1] - prefix[l]
	}

	// Try going right first:
	for numRightSteps := 0; numRightSteps <= min(maxRight-startPos, k); numRightSteps++ {
		numLeftSteps := max(k-2*numRightSteps, 0)
		ans = max(ans, calculateHarvest(numLeftSteps, numRightSteps))
	}

	// Try going left first:
	for numLeftSteps := 0; numLeftSteps <= min(k, startPos); numLeftSteps++ {
		numRightSteps := max(k-2*numLeftSteps, 0)
		ans = max(ans, calculateHarvest(numLeftSteps, numRightSteps))
	}
	return ans
}
