package lc1550

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func threeConsecutiveOdds(arr []int) bool {
	counter := 0
	for _, num := range arr {
		if num%2 == 0 {
			counter = 0
		} else {
			counter++
			if counter == 3 {
				return true
			}
		}
	}
	return false
}
