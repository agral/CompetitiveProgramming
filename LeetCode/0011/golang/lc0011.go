package lc0011

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		ans = max(area, ans)
		if height[l] < height[r] {
			l += 1
		} else if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
			r -= 1
		}
	}
	return ans
}
