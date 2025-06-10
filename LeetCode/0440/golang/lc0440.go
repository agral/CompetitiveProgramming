package lc0440

// Runtime complexity: O(n); TLEs at 45th / 69 testcases.
// Auxiliary space: O(1)
func findKthNumberBigOhN(n int, k int) int {
	curr := 1
	for range k - 1 {
		if 10*curr <= n {
			curr *= 10
		} else {
			for curr%10 == 9 || curr == n {
				curr = curr / 10
			}
			curr += 1
		}
	}

	return curr
}

// Runtime complexity: O(log^2(n))
// Aux space: O(1)
func findKthNumber(n int, k int) int {
	curr := 1
	i64 := int64(1)
	k64 := int64(k)
	for i64 < k64 {
		skip := getSkipAheadLength(int64(curr), int64(curr+1), int64(n))
		if i64+skip <= k64 {
			i64 += skip
			curr += 1
		} else {
			i64 += 1
			curr = curr * 10
		}
	}
	return curr
}

func getSkipAheadLength(curr int64, after int64, n int64) int64 {
	ans := int64(0)
	for curr <= n {
		ans += min(n+1, after) - curr
		after = after * 10
		curr = curr * 10
	}
	return ans
}
