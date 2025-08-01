package lc0118

import "fmt"

// Runtime complexity: O(n^2)
// Auxiliary space: O(n^2) for storing the answer; O(1) not including that.
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i+1) // in Pascal's triangle, the ith row has i+1 entries
		ans[i][0] = 1
		ans[i][i] = 1
		for k := range i - 1 {
			ans[i][k+1] = ans[i-1][k] + ans[i-1][k+1]
		}
	}
	fmt.Println(ans)
	return ans
}
