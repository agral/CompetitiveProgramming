package lc0898

// Runtime complexity: O(n * log(max_elem(arr)))
// Auxiliary space: O(n)
func subarrayBitwiseORs(arr []int) int {
	vec := []int{}
	left := 0
	for _, a := range arr {
		right := len(vec)
		vec = append(vec, a)
		for i := left; i < right; i++ {
			ored := vec[i] | a
			if vec[len(vec)-1] != ored {
				vec = append(vec, ored)
			}
		}
		left = right
	}
	set := map[int]bool{}
	for _, a := range vec {
		set[a] = true
	}
	return len(set)
}
