package lc0403

import "fmt"

func pretty_print(arr [][]bool) {
	for r := 0; r < len(arr); r++ {
		for z := 0; z < len(arr[r]); z++ {
			fmt.Printf("%v\t", arr[r][z])
		}
		fmt.Println()
	}
	fmt.Println()
}

func canCross(stones []int) bool {
	sz := len(stones)

	// dp[i][j] == true if the frog can make a jump of size j from stones[i].
	dp := make([][]bool, sz)
	for i := 0; i < sz; i++ {
		dp[i] = make([]bool, sz+1)
	}
	dp[0][1] = true // initially the frog can jump forward exactly 1 space from the 0th stone.

	for i := 1; i < sz; i++ {
		for j := 0; j < i; j++ {
			k := stones[i] - stones[j]
			if (k <= sz) && (dp[j][k] == true) {
				// ignore jumps where jump-length > len(stones);
				// max jump size for the frog is len(stones) if it used +1 for each jump.
				dp[i][k-1] = true
				dp[i][k] = true
				dp[i][k+1] = true
			}
		}
	}

	// pretty_print(dp)
	for j := 0; j <= sz; j++ {
		if dp[sz-1][j] == true {
			return true
		}
	}

	return false
}
