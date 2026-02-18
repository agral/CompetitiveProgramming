package lc0693

import (
	"fmt"
	"strconv"
	"strings"
)

// Runtime complexity: O(1) (we're binary searching in an array of 34 numbers for a match)
// Auxiliary space: O(34) == O(1)
// Subjective level: easy
// Solved on: 2026-02-18
func hasAlternatingBits(n int) bool {
	ALTERNATING := []int{0, 1, 2, 5, 10, 21, 42, 85, 170, 341, 682, 1365, 2730, 5461, 10922, 21845, 43690, 87381, 174762, 349525, 699050, 1398101, 2796202, 5592405, 11184810, 22369621, 44739242, 89478485, 178956970, 357913941, 715827882, 1431655765, 2863311530, 5726623061}
	l, r := 0, 33
	for l <= r {
		mid := (l + r) / 2
		if n == ALTERNATING[mid] {
			return true
		} else if n < ALTERNATING[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func GenerateAlternatingBitNumbers() {
	// 0 -> 1 (01) -> 10 -> 101 -> 1010 -> 10101 -> ... -> 1010...101 (31 bits).
	ans := []string{}
	num := 0
	for num <= (1 << 32) { // the last result will be 33 bits long anyway. Use all.
		num *= 2
		ans = append(ans, strconv.Itoa(num))
		num *= 2
		num += 1
		ans = append(ans, strconv.Itoa(num))
	}

	fmt.Println(strings.Join(ans, ", "))
}
