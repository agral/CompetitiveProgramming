package lc1534

// Go's math.Abs() is defined over float64 exclusively.
// That feels too minimalistic; abs over ints is something rather common.
// So let's roll a custom abs function instead:
func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// Runtime complexity: O(n^3). Sadly, this seems optimal, i.e. it can't be optimized further.
// Aux space complexity: O(1).
func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	LEN := len(arr)
	for i := 0; i < LEN; i++ {
		for j := i + 1; j < LEN; j++ {
			for k := j + 1; k < LEN; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans += 1
				}
			}
		}
	}
	return ans
}
