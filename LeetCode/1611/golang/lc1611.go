package lc1611

/*
Note from agral: once again this one is a daily question on 08 Nov '25.
I've solved it in Golang this time; and copied the original note from C++
solution (dated November 2023) below.
*/

// *** Note from agral:
// After doing it, I *hate* this, and it's a daily question for 30 Nov '23.
// There's no way this tests anything in a candidate,
// except having grinded a lot of programming challenges before attempting this one.

// Crucial observation (pen & paper, but this may take some time to spot):
// solving 2^k -> 0 requires (2^(k+1))-1 operations.
// No. | Op | Num | Notes
// 0   |    | 100 | 2^2, will need 2^3-1 == 7 operations
// 1   | 1  | 101 |
// 2   | 2  | 111 |
// 3   | 1  | 110 |
// 4   | 2  | 010 | this is 2^1, it will require 3 ops
// 5   | 1  | 011 |
// 6   | 2  | 001 | this is 2^0, requires 1 ops
// 7   | 1  | 110 |

// So, for n = 1XXXX, the approach is somewhat recursive:
// - if the 2nd bit is 0, then the "extra" cost of flipping 2nd bit to 1 has to be added,
// - if the 2nd bit is 1, there's no extra cost.

// Runtime complexity: O(log(n))
// Auxiliary space: O(log(n))
// Subjective level: Hard, but for all the wrong reasons. Don't bother. Lots of better challenges out there.
func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	maxPowerOfTwoLessThanN := 1
	for 2*maxPowerOfTwoLessThanN <= n {
		maxPowerOfTwoLessThanN *= 2
	}
	var halfMaxPowerOfTwoLessThanN int = maxPowerOfTwoLessThanN / 2

	return maxPowerOfTwoLessThanN +
		minimumOneBitOperations(n^(maxPowerOfTwoLessThanN|halfMaxPowerOfTwoLessThanN))
}
