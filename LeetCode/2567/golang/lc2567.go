package lc2567

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(6) == O(1)
// Subjective level: medium-
// Solved on: 2026-11-19
func minimizeSum(nums []int) int {
	// Observation: After changing two numbers, the `low` is always 0
	// (as these can be changed to anything, so just change them to some same number).
	// So the answer will always equal `high`, meaning the absolute difference of any two numbers.

	// indices of minimal, next-to-minimal, and next-to-next-to-minimal:
	min1, min2, min3 := math.MaxInt32, math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32 // same, but maxes.

	// Find all the six {min,max}{1-3} indices:
	for _, val := range nums {
		if val < min3 { // found a number that's at least lower than min3
			if val < min2 { // found a number that's at least lower than min2
				if val < min1 { // found a new global minimum (so far!)
					min3 = min2
					min2 = min1
					min1 = val
				} else { // found a number lower than min2, but not lower than min1.
					min3 = min2
					min2 = val
				}
			} else { // found a number lower than min3, but not lower than min2
				min3 = val
			}
		}
		if val > max3 { // found a number that's at least higher than max3
			if val > max2 { // found a number that's at least higher than max2
				if val > max1 { // found a new global maximum, so far
					max3 = max2
					max2 = max1
					max1 = val
				} else { // found a number than's higher than max2, but not higher than max1.
					max3 = max2
					max2 = val
				}
			} else { // found a number that's only higher than max3.
				max3 = val
			}
		}
	}

	// Try removing the two numbers. There are three ways to do it that matter:
	// Variant A) remove min1 and min2 (by making them equal to any of the remaining numbers)
	// Here `high` == max1 - min3.
	ans := max1 - min3

	// Variant B) remove min1 and max1. `high` == max2-min2.
	ans = min(ans, max2-min2)

	// Variant C) remove max1 and max2. `high` == max3-min1 in this case.
	ans = min(ans, max3-min1)

	// Pick the minimum of the above three, it's the final answer - I think.
	return ans
}
