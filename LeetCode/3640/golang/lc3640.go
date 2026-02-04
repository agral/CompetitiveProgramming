package lc3640

import (
	"math"
)

// Look ma, Go has enums! Kinda...
const (
	PHASE_INITIAL = iota
	PHASE_UP_FIRST
	PHASE_DOWN
	PHASE_UP_SECOND
)

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: hard-. Also *poorly* described, so maybe a `hard` is in order.
// lots of corner cases to consider.

// Not Solved on: 2026-02-04
// WAs on the last provided testcase #14. Don't want to work on this any more.
func maxSumTrionic(nums []int) int64 {
	n := len(nums)

	ans := int64(math.MinInt64 / 2)
	phase := PHASE_INITIAL
	sumUpFirst := int64(0)
	sumDown := int64(0)
	sumUpSecond := int64(0)
	sumSecondWhenFirst := int64(0)

	for i := 1; i < n; i++ {
		switch phase {
		case PHASE_INITIAL:
			// in contrast to the solution for IsTrionic (LC3637), here the slice
			// may start from a nonincreasing phase - and still contain a valid
			// trionic subslice. It is actually guaranteed, that a valid
			// trionic subslice exists. Remain in PHASE_INITIAL until
			// first increasing sequence is reached.
			if nums[i] > nums[i-1] {
				sumUpFirst = int64(nums[i-1]) + int64(nums[i])
				phase = PHASE_UP_FIRST
				//fmt.Printf("INITIAL -> UP_FIRST at i=%d, sumUpFirst=%d\n", i, sumUpFirst)
			}

		case PHASE_UP_FIRST:
			if nums[i] == nums[i-1] {
				// If a number is repeated, the sequence is no longer increasing nor decreasing.
				// Switch to PHASE_INITIAL, forget the accumulated sum, start over.
				sumUpFirst = int64(0)
				phase = PHASE_INITIAL
			} else if nums[i] > nums[i-1] {
				// keep adding to the sum as long as the numbers form an increasing subsequence.
				sumUpFirst += int64(nums[i])
				// However! If *starting* here results in bigger sumUpFirst than whatever's accumulated,
				// just start here. It will result in a better (bigger) answer.
				// Worked it out myself without getting a WA, so proud.
				sumUpFirst = max(sumUpFirst, int64(nums[i]))
			} else {
				// once the first number that forms a decreasing subsequence is spotted,
				// switch to PHASE_DOWN and set this phase's sum as this first nonconforming number.
				sumDown = int64(nums[i])
				phase = PHASE_DOWN
				//fmt.Printf("UP_FIRST -> DOWN at i=%d, sumUpFirst=%d, sumDown=%d\n", i, sumUpFirst, sumDown)
			}

		case PHASE_DOWN:
			if nums[i] == nums[i-1] {
				// If a number is repeated, the sequence is once again no longer increasing nor decreasing.
				// Switch to PHASE_INITIAL, forget the accumulated sumFirstUp and sumDown, start over.
				sumUpFirst = int64(0)
				sumDown = int64(0)
				phase = PHASE_INITIAL
			} else if nums[i] < nums[i-1] {
				// again, keep adding to the sum as long as the numbers form a decreasing subsequence.
				sumDown += int64(nums[i])
			} else {
				// once the first number that forms an increasing subsequence is spotted,
				// switch to PHASE_UP_SECOND and set this phase's sum as this first nonconforming number.
				// also remove the last number of the down-sequence (i-1 now) from its sum,
				// include it in this sum. This makes it consistent when substituting
				// sumUpFirst with sumUpSecond in the PHASE_UP_SECOND step.
				sumDown -= int64(nums[i-1])
				sumUpSecond = int64(nums[i-1]) + int64(nums[i])

				// but remember that when the second sequence becomes the first, it can start
				// from any index - the second can not, but it can terminate whenever.
				// so when changing from second to first, use the correct starting point.
				// Done by keeping the best score it would have if it could have been
				// a subarray of second starting at any index.
				sumSecondWhenFirst = max(sumUpSecond, int64(nums[i]))
				phase = PHASE_UP_SECOND
				ans = max(ans, sumUpFirst+sumDown+sumUpSecond)
				//fmt.Printf("DOWN -> UP_SECOND at i=%d, sumUpFirst=%d, sumDown=%d, sumUpSecond=%d\n", i, sumUpFirst, sumDown, sumUpSecond)
			}

		case PHASE_UP_SECOND:
			if nums[i] == nums[i-1] {
				// If a number is repeated, the sequence is no longer increasing nor decreasing.
				// Switch to PHASE_INITIAL, forget the accumulated sums, start over.
				sumUpFirst = int64(0)
				sumDown = int64(0)
				sumUpSecond = int64(0)
				phase = PHASE_INITIAL
			} else if nums[i] < nums[i-1] {
				// When a decreasing-subsequence-forming number is spotted,
				// treat the sumUpSecond accumulated so far as the sum of first up-phase,
				// and switch to PHASE_DOWN. Thus it's possible to keep scanning for best
				// trionic subarray in O(n), with no need to repeat scanning.
				sumUpFirst = sumSecondWhenFirst
				sumUpSecond = int64(0)
				sumSecondWhenFirst = int64(0)
				sumDown = int64(nums[i])
				phase = PHASE_DOWN
				//fmt.Printf("UP_SECOND->DOWN at i=%d, sumUpFirst=%d, sumDown=%d, sumUpSecond=%d\n", i, sumUpFirst, sumDown, sumUpSecond)
			} else {
				// Keep adding to the sum as long as the numbers form an increasing subsequence.
				// Keep the score of the best trionic subarray seen so far as the answer.
				sumUpSecond += int64(nums[i])
				sumSecondWhenFirst += int64(nums[i])
				sumSecondWhenFirst = max(sumSecondWhenFirst, int64(nums[i]))
				ans = max(ans, sumUpFirst+sumDown+sumUpSecond)
			}
		}
	}

	//fmt.Printf("sumUpFirst: %d, sumDown: %d, sumUpSecond: %d, ans: %d\n", sumUpFirst, sumDown, sumUpSecond, ans)
	return ans
}
