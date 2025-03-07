package lc2523

import (
	"math"
)

func primesInRange(start int, end int) []int {
	// Uses the sieve of Eratosthenes' algorithm to generate a list of prime numbers from start to end inclusively.
	isPrime := make([]bool, end+1)

	// Initially mark all the numbers in range from 2 to n as primes:
	for i := 2; i <= end; i++ {
		isPrime[i] = true
	}
	// Mark multiples of primes as not prime:
	for p := 2; p <= int(math.Sqrt(float64(end))); p++ {
		if isPrime[p] { // optimization: don't duplicate work on nums that are already marked as composite.
			for c := p * p; c <= end; c += p {
				// another optimization: c := p * p instead of 2*p; as these multiples will have already
				// been covered by 2, 3, 5 etc.; leaving p*p the first unmarked composite (non-prime)
				// multiple of p. Easy to work it out with pen and paper.
				isPrime[c] = false
			}
		}
	}

	primes := []int{}
	for i := start; i <= end; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func closestPrimes(left int, right int) []int {
	primes := primesInRange(left, right)
	if len(primes) < 2 {
		return []int{-1, -1}
	}
	prime1 := primes[0]
	prime2 := primes[1]
	minDiff := prime2 - prime1
	for i := 2; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDiff {
			prime1, prime2 = primes[i-1], primes[i]
			minDiff = diff
		}
	}
	return []int{prime1, prime2}
}
