package lc0455

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	ans := 0
	for i := 0; i < len(s) && ans < len(g); i++ {
		if g[ans] <= s[i] {
			ans++
		}
	}
	return ans
}
