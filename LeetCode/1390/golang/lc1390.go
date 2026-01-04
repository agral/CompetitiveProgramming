package lc1390

// Runtime complexity: O(len(nums) * sqrt(num))
// Auxiliary space: O(1)
// Subjective level: medium. Requires familiarity with some number theory concepts
// to solve this in better time than bruteforcing.
// Solved on: 2026-01-04
func sumFourDivisors(nums []int) int {
	const NO_DIVISORS_YET = -42
	ans := 0
	for _, num := range nums {
		// Every number `num` has two divisors, `1` and `num`. Even prime numbers do.
		// So `num` has exactly four divisors, if it has one extra pair of divisors
		// other than `1` and `num`. If it has more than one extra pair of such divisors,
		// it will have more than 4 divisors in total.
		// But note that the above approach does not work for square numbers - (e.g. 9 == 1, 3, 9),
		// as these will always have odd count of divisors, but the above algo
		// results in the root number (here 3) counted twice (as in: 1, 3, 3, 9).
		// So in the end it is sufficient to check for divisors that are greater than one,
		// but less than the square root of `num`. If there's exactly one such divisor `div`,
		// `num` has four divisors in total, in ascending order: `1`, `div`, `num`/`div` and `num`.
		div := NO_DIVISORS_YET
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				if div == NO_DIVISORS_YET {
					div = i
				} else {
					// this number has more than four divisors in total. Disregard it.
					div = NO_DIVISORS_YET
					break
				}
			}
		}

		// note: have to take the squares' problem into account
		// That's fixed by the 2nd conditional, which eliminates squares.
		if div != NO_DIVISORS_YET && div*div < num {
			ans += 1 + div + (num / div) + num
		}
	}
	return ans
}
