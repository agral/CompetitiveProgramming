package lc1780

func checkPowersOfThree(n int) bool {
	// Want the ternary representation of a number to not contain any '2' digits.
	// (so if only 0 and 1 is present in the ternary representation, then the number
	// can be represented as a sum of powers of three (implicit: used at most once each)).

	// The ternary representation can be easily checked with e.g. `bc` program:
	// `echo "obase=3;91"|bc` outputs 10101; while `echo "obase=3;21"|bc` outputs 210.
	for n > 1 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}
