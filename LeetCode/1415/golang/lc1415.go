package lc1415

import "fmt"

type StringQueue []string

func (q *StringQueue) Enqueue(val string) {
	*q = append(*q, val)
}

func (q *StringQueue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "error", fmt.Errorf("Dequeue() from an empty Queue.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

// intentionally no bounds check; use only on confirmed Len() > 0.
func (q *StringQueue) UnsafeFront() string {
	return (*q)[0]
}

func (q *StringQueue) Len() int {
	return len(*q)
}

// Runtime complexity: O(3^n)
// Auxiliary space: O(3^n) - for the queued items
// (there are exactly 3^n happy strings of length n)
// Subjective level: medium
// Solved on: 2026-03-14
func getHappyString(n int, k int) string {
	letterToNext := map[byte]string{
		'a': "bc",
		'b': "ac",
		'c': "ab",
	}
	q := StringQueue{"a", "b", "c"}
	for len(q.UnsafeFront()) < n {
		base, _ := q.Dequeue()
		lastBaseChar := base[len(base)-1]
		for _, ch := range letterToNext[lastBaseChar] {
			q.Enqueue(base + string(ch))
		}
	}

	if q.Len() < k {
		return ""
	}

	// note: starting from 1, so that (k-1) entries are removed.
	// Then the k-th entry is the one on top.
	for i := 1; i < k; i++ {
		_, _ = q.Dequeue()
	}
	return q.UnsafeFront()
}
