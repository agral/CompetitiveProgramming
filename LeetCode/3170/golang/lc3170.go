package lc3170

import "strings"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func clearStars(s string) string {
	const NUM_LETTERS = 'z' - 'a' + 1
	const BLANK = '_'
	ans := []byte(s)
	letters := make([][]int, NUM_LETTERS)
	/*
		for i := range NUM_LETTERS {
			letters[i] = make([]int, 0)
		}
	*/
	for i, letter := range s {
		if letter == '*' {
			ans[i] = BLANK
			j := 0
			for len(letters[j]) == 0 {
				j++
			}

			ans[letters[j][len(letters[j])-1]] = BLANK
			letters[j] = letters[j][:len(letters[j])-1]
		} else {
			letters[letter-'a'] = append(letters[letter-'a'], i)
		}
	}
	sans := string(ans)
	return strings.ReplaceAll(sans, string(BLANK), "")
}
