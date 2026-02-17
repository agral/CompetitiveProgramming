package lc0401

import (
	"fmt"
	"strconv"
	"strings"
)

// Runtime complexity: O(12*60) == O(1)
// Auxiliary space: O(12 + 60 + O(ans)) == O(1), given that len(ans) < 720 == O(1).
// Subjective level: easy+
// Solved on: 2026-02-17
func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	//print_popcount_table(12, "HOURS")
	//print_popcount_table(60, "MINUTES")
	POPCOUNT_HOURS := [12]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3}
	POPCOUNT_MINUTES := [60]int{
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3,
		2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4,
		2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5,
	}

	for h := range 12 {
		for m := range 60 {
			if POPCOUNT_HOURS[h]+POPCOUNT_MINUTES[m] == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return ans
}

// Used to create the popcount tables:
func popcount(num int) int {
	ans := 0
	for num > 0 {
		if num&1 == 1 {
			ans += 1
		}
		num >>= 1
	}
	return ans
}
func print_popcount_table(count int, tag string) { // prints out ready-to-use code
	table := make([]string, count)
	for i := range count {
		table[i] = strconv.Itoa(popcount(i))
	}
	fmt.Printf("POPCOUNT_%s := [%d]int{%s}\n", tag, count, strings.Join(table, ", "))
}
