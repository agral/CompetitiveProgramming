package lc1524

func numOfSubarrays(arr []int) int {
	// Runtime complexity: O(n)
	// Aux space complexity: O(1)
	const MOD = 1e9 + 7
	dp_even := 0
	dp_odd := 0
	ans := 0
	for _, num := range arr {
		if num%2 == 0 {
			dp_even += 1
		} else {
			dp_even, dp_odd = dp_odd, dp_even+1
		}
		ans = (ans + dp_odd) % MOD
	}
	return ans
}

func numOfSubarrays_firstAttempt(arr []int) int {
	// Runtime complexity: O(n)
	// Aux space complexity: O(n) - two arrays of size len(arr)+1 are allocated.
	const MOD = 1e9 + 7
	SZ := len(arr)

	dp_even := make([]int, SZ+1) // count of subarrays that end in arr[i-1]
	dp_odd := make([]int, SZ+1)  // with an even / odd sum, respectively.
	ans := 0
	for i := range SZ {
		if arr[i]%2 == 0 {
			dp_even[i+1] = dp_even[i] + 1
			dp_odd[i+1] = dp_odd[i]
		} else {
			dp_even[i+1] = dp_odd[i]
			dp_odd[i+1] = dp_even[i] + 1
		}
		ans = (ans + dp_odd[i+1]) % MOD
	}

	return ans
}
