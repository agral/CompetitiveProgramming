package main

import "fmt"

/**
 * @param num the guess, that will be compared to a secret number.
 * @return -1 if num is HIGHER than the secret number
 *          1 if num is LOWER than the secret number
 *          0 if num is EQUAL to the secret number.
 */
var Secret = 0
var NumTries = 0

func guess(num int) int {
	NumTries++
	if num > Secret {
		return -1
	}
	if num < Secret {
		return 1
	}
	return 0
}

// Runtime complexity: O(1)
// Finds the number via binary search. The answer should be found in O(log(n)).
// Auxiliary space: O(1)
func guessNumber(n int) int {
	lowerBound := 1
	upperBound := n

	for lowerBound < upperBound {
		mid := (lowerBound + upperBound) / 2
		result := guess(mid)
		if result == 0 { // Guessed the secret number.
			return mid
		} else if result == -1 { // Secret number is lower than the guess.
			upperBound = mid - 1
		} else { // Secret number is higher than the guess.
			lowerBound = mid + 1
		}
	}
	return lowerBound
}

func main() {
	const N int = 100

	for i := 1; i <= N; i++ {
		Secret = i
		NumTries = 0
		guessNumber(N)
		fmt.Printf("%d/%d found in %d calls.\n", i, N, NumTries)
	}
}
