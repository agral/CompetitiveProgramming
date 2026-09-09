package lc3871

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy. Conceptually very easy, it's just lots of nested if statements.
// Solved on: 2026-09-09
func countCommas(n int64) int64 {
	// Note: 1 <= n <= 1e15.
	// So there are the following commas count possible:
	//                     1 - 999                 : 0 commas each (count:                   999)
	//                 1_000 - 999999              : 1 commas each (count:               999_000)
	//             1_000_000 - 999_999_999         : 2 commas each (count:           999_000_000)
	//         1_000_000_000 - 999_999_999_999     : 3 commas each (count:       999_000_000_000)
	//     1_000_000_000_000 - 999_999_999_999_999 : 4 commas each (count:   999_000_000_000_000)
	// 1_000_000_000_000_000 (1e15, upper limit)   : 5 commas      (count:                     1)
	//                                                        total count: 1_000_000_000_000_000 == 1e15
	ans := int64(0)
	if n > 999 {
		ans += min(n-999, 999_000)
		if n > 999_999 {
			ans += 2 * min(n-999_999, 999_000_000)
			if n > 999_999_999 {
				ans += 3 * min(n-999_999_999, 999_000_000_000)
				if n > 999_999_999_999 {
					ans += 4 * min(n-999_999_999_999, 999_000_000_000_000)
					if n == 1_000_000_000_000_000 {
						ans += 5 // times one, the 1e15 itself - the maximum allowed number.
					}
				}
			}
		}
	}
	return ans
}
