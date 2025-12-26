package lc1298

import "fmt"

type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("Pop() from an empty Queue.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2025-12-27
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	q := Queue{}
	isClosedAndAvailable := make([]bool, n)

	pushBoxes := func(boxes []int) {
		for _, box := range boxes {
			if status[box] != 0 {
				q.Push(box)
			} else {
				isClosedAndAvailable[box] = true
			}
		}
	}

	ans := 0
	pushBoxes(initialBoxes)

	for !q.Empty() {
		// 1. Open the box and add its candies to the answer pile:
		box, _ := q.Pop()
		ans += candies[box]

		// 2. For all the keys in this box, add its corresponding unopened boxes to the queue:
		for _, key := range keys[box] {
			if status[key] == 0 && isClosedAndAvailable[key] {
				q.Push(key)
			}
			status[key] = 1 // in any case, mark this box as opened/processed.
		}

		// 3. Add all the boxes contained in this box to the queue:
		pushBoxes(containedBoxes[box])
	}

	return ans
}
