package lc0077

// Runtime complexity: O(C(n, k))
// Auxiliary space: O(n*k)
// Subjective level: hard-
// Solved on: 2026-01-26

// Note: Even the total count of combinations of k elements from set of n is: (n!) / ((k!)*(n-k)!);
// that's too much to compute using ints for the upper constraint case of n=20.

func combine(n int, k int) [][]int {
	ans := [][]int{}
	path := []int{}

	var dfs func(nextNum int)
	dfs = func(nextNum int) {
		if len(path) == k {
			var pathClone []int // have to work around Go's list shallow cloning.
			pathClone = append(pathClone, path...)
			ans = append(ans, pathClone)
			return
		}
		for i := nextNum; i <= n; i++ {
			path = append(path, i) // puts i to the path slice
			dfs(i + 1)
			path = path[:len(path)-1] // pops the last element from path slice
		}
	}

	dfs(1)

	return ans
}
