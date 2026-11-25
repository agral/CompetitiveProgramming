package lc3934

// Runtime complexity: O(n^2):
//   - O(n) for checking windows of increasing sizes, from 1 up to L-1 elements in the array
//   - each of these has to slide a window of size k over the entire array. That's another O(n).
//   - each of the operations done in each slide of the window is O(1).
//
// Auxiliary space: O(n):
//   - I've converted nums to int64; this required allocating L int64s in another array;
//   - hashToCount is a map that will store (L-winSize) hashes; that's O(n) too.
//
// Subjective level: hard.
// Solved on: 2026-11-25
func smallestUniqueSubarray(nums []int) int {
	// constants for computing the rolling hash (base=23, modulus=1e9+7):
	const BASE = int64(23)
	const MOD int64 = 1_000_000_007

	L := len(nums)

	// precompute BASE^i (modulo MOD); for all the possible values of 1 <= i <= L:
	POWERS := make([]int64, L+1)
	POWERS[0] = 1
	for i := 1; i <= L; i++ {
		POWERS[i] = (POWERS[i-1] * BASE) % MOD
	}

	// convert nums to int64:
	nums64 := make([]int64, L)
	for i, num := range nums {
		nums64[i] = int64(num)
	}

	windowContainsUniqueSubarray := func(winSize int) bool {
		hashToCount := map[int64]int{} // maps each window hash to its number of occurrences.

		// compute the entire hash of the first window ( [0, winSize) )
		hash := int64(0)
		for i := 0; i < winSize; i++ {
			hash = (hash*BASE + nums64[i]) % MOD
		}
		hashToCount[hash] = 1 // OK, the hash of the first window is now counted.

		// now move the window step-by-step to the right,
		// until end of array is reached for current window size:
		for offset := 1; offset <= L-winSize; offset++ {
			// remove from current hash the value contributed by the leftmost element (leaving):
			leaving := POWERS[winSize-1] * nums64[offset-1] % MOD
			hash = (hash - leaving) % MOD

			// shift the remaining elements one power up:
			hash = (hash * BASE) % MOD

			// add the new element from shifting right:
			hash = (hash + nums64[winSize-1+offset]) % MOD

			// store the just calculated hash in the map:
			hashToCount[hash] += 1 // default value of an int in a map is 0; no need to check for existence.
		}

		// Finally, check if there's any unique hash (count == 1) in the hashToCount map:
		for _, count := range hashToCount {
			if count == 1 {
				return true
			}
		}

		return false
	}

	for i := 1; i < L; i++ { // try increasing sizes of the window, from 1 to L-1.
		if windowContainsUniqueSubarray(i) {
			return i
		}
	}

	return L // if none of the subarrays were unique, the entire array itself is definitely unique.
}
