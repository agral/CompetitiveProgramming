package lc2054

import (
	"slices"
)

type Event struct {
	StartTime  int
	IsStarting int // 1 if starting, 0 if stopping. But I want int, not bool there.
	Value      int
}

// Runtime complexity: O(sort)
// Auxiliary space: O(sort) + O(n)
// Subjective level: medium
// Solved on: 2025-12-23
func maxTwoEvents(events [][]int) int {
	// store custom Events: an "event start" and "event ends" for each event in input array.
	ev := make([]Event, 0, 2*len(events))
	for _, event := range events {
		// events[i] == event == {startTime, endTime, value}
		ev = append(ev, Event{event[0], 1, event[2]})
		ev = append(ev, Event{event[1] + 1, 0, event[2]})
	}

	// sort custom Events in ascending order w.r.t. start time.
	// For ties, sort stop events before start events.
	slices.SortFunc(ev, func(lhs Event, rhs Event) int {
		return 2*(lhs.StartTime-rhs.StartTime) + (lhs.IsStarting - rhs.IsStarting)
	})

	maxValuedEvent := 0
	ans := 0

	for _, event := range ev {
		if event.IsStarting == 1 {
			ans = max(ans, event.Value+maxValuedEvent)
		} else {
			maxValuedEvent = max(maxValuedEvent, event.Value)
		}
	}
	return ans
}
