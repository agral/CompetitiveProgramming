package lc2996

// Runtime complexity: O(n), where n is the length of the array. That is, O(50) == O(1).
// Auxiliary space: O(50) == O(1).
// Subjective level: easy.
// Solved on: 2026-08-11

// This question is *really* poorly described. Starting from a clean slate.
func missingIntegerWA(nums []int) int {
	// 1 <= nums.length <= 50
	// 1 <= nums[i] <= 50
	const MAX int = 50

	seen := make([]bool, MAX+1)
	seen[nums[0]] = true
	lsp := nums[0] // longest sequential prefix
	csp := nums[0] // current sequential prefix

	for i := 1; i < len(nums); i++ {
		seen[nums[i]] = true
		if nums[i] == nums[i-1]+1 {
			csp += nums[i]
			lsp = max(lsp, csp)
		} else {
			csp = nums[i]
		}
	}

	for i := lsp; i <= MAX*MAX; i++ { // in theory, there could be a testcase consisting of nums 1 to 50.
		if i > MAX || !seen[i] {
			return i
		}
	}

	return -71826576 // never reachable for valid testcases.
}

func missingInteger(nums []int) int {
	numsSet := map[int]bool{}
	for _, num := range nums {
		numsSet[num] = true
	}

	ans := nums[0]
	for i := 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		ans += nums[i]
	}

	for numsSet[ans] {
		ans += 1
	}

	return ans
}
