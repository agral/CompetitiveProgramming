package lc0838

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func pushDominoes(dominoes string) string {
	ans := []rune(dominoes)
	left, right := -1, -1

	for i := 0; i <= len(dominoes); i++ {
		if i == len(dominoes) || dominoes[i] == 'R' {
			if left < right {
				for right < i {
					ans[right] = 'R'
					right++
				}
			}
			right = i
		} else if dominoes[i] == 'L' {
			if left > right || (left == -1 && right == -1) {
				if left == -1 && right == -1 {
					left++
				}
				for left < i {
					ans[left] = 'L'
					left++
				}
			} else {
				l := right + 1
				r := i - 1
				for l < r {
					ans[l] = 'R'
					ans[r] = 'L'
					l++
					r--
				}
			}
			left = i
		}
	}
	return string(ans)
}
