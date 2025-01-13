package lc3223

func minimumLength(s string) int {
	const NUM_LETTERS int = 'z' - 'a' + 1

	// Calculate frequency composition of letters in s:
	freq := make([]int, NUM_LETTERS)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// For each letter, two instances are deleted as long as there are at least three
	// instances at any time. This means that in the end there will be:
	// * 0, 1, 2 if started with 0, 1, 2
	// * n%2 =1 if started with odd n >= 3, e.g. 3, 5, 7, 9, ...
	// * n%2+2=2 if started with even n >=3, e.g. 4, 6, 8, 10, ...
	ans := 0
	for i := 0; i < NUM_LETTERS; i++ {
		if freq[i] < 3 {
			ans += freq[i]
		} else if freq[i]%2 == 1 {
			ans += 1
		} else {
			ans += 2
		}
	}

	return ans
}
