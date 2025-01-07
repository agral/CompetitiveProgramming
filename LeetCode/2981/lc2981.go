package lc2981

// You are given a string s that consists of lowercase English letters.
//
// A string is called special if it is made up of only a single character.
// For example, the string "abc" is not special,
// whereas the strings "ddd", "zz", and "f" are special.
//
// Return the length of the longest special substring of s which occurs
// at least thrice, or -1 if no special substring occurs at least thrice.
//
// A substring is a contiguous non-empty sequence of characters within a string.

func maximumLength(s string) int {
	sz := len(s)
	numLetters := int('z'-'a') + 1
	runningLength := 0
	previousLetter := '!'
	counts := make([][]int, numLetters)
	for i := 0; i < numLetters; i++ {
		counts[i] = make([]int, sz+1)
	}

	for _, ch := range s {
		offset := int(ch - 'a')
		if previousLetter == ch {
			runningLength++
		} else {
			runningLength = 1
			previousLetter = ch
		}
		counts[offset][runningLength]++
	}

	ans := -1
	for offset := 0; offset < numLetters; offset++ {
		count := 0
		for k := sz; k > 0; k-- {
			count += counts[offset][k]
			if count >= 3 {
				ans = max(ans, k)
				continue
			}
		}
	}
	return ans
}
