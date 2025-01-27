package lc1462

import (
	"slices"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	testcases := []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		expected      []bool
	}{
		{ // Example 1: Course 1 is a prerequisite of course 0.
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			queries:       [][]int{{0, 1}, {1, 0}},
			expected:      []bool{false, true},
		},
		{ // Example 2: (a degenerated case) two independent courses
			numCourses:    2,
			prerequisites: [][]int{},
			queries:       [][]int{{1, 0}, {0, 1}},
			expected:      []bool{false, false},
		},
		{ // Example 3:
			numCourses:    3,
			prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
			queries:       [][]int{{1, 0}, {1, 2}},
			expected:      []bool{true, true},
		},
		{ // Example 4: (self-made) an indirect dependency
			//  0 --> 1 --> 3
			//        |     |
			//        v     v
			//        2 --> 4
			//
			numCourses:    5,
			prerequisites: [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 4}},
			queries:       [][]int{{0, 2}, {0, 3}, {1, 3}, {1, 4}, {2, 3}, {3, 2}, {0, 4}},
			expected:      []bool{true, true, true, true, false, false, true},
		},
		{ // Testcase #47 - WA #01
			numCourses:    13,
			prerequisites: [][]int{{2, 1}, {2, 7}, {2, 0}, {2, 10}, {2, 11}, {1, 7}, {1, 0}, {1, 9}, {1, 4}, {1, 11}, {7, 3}, {7, 9}, {7, 4}, {7, 11}, {7, 8}, {3, 6}, {3, 12}, {3, 5}, {6, 10}, {6, 8}, {0, 4}, {12, 9}, {12, 8}, {9, 4}, {9, 11}, {9, 8}, {9, 5}, {10, 8}, {4, 8}},
			queries:       [][]int{{12, 11}, {11, 1}, {10, 12}, {9, 10}, {10, 11}, {11, 12}, {2, 7}, {6, 8}, {3, 2}, {9, 5}, {8, 7}, {1, 4}, {3, 12}, {9, 6}, {4, 3}, {11, 4}, {5, 7}, {3, 9}, {3, 1}, {8, 12}, {5, 12}, {0, 8}, {10, 5}, {10, 11}, {12, 11}, {12, 9}, {5, 4}, {11, 5}, {12, 10}, {11, 0}, {6, 10}, {11, 7}, {8, 10}, {2, 1}, {3, 4}, {8, 7}, {11, 6}, {9, 11}, {1, 4}, {10, 8}, {7, 1}, {8, 7}, {9, 7}, {5, 1}, {8, 10}, {11, 8}, {8, 12}, {9, 12}, {12, 11}, {6, 12}, {12, 11}, {6, 10}, {9, 12}, {8, 10}, {8, 11}, {8, 5}, {7, 9}, {12, 11}, {11, 12}, {8, 0}, {12, 11}, {7, 0}, {8, 7}, {5, 11}, {11, 8}, {1, 9}, {4, 10}, {11, 6}, {10, 12}},
			expected:      []bool{true, false, false, false, false, false, true, true, false, true, false, true, true, false, false, false, false, true, false, false, false, true, false, false, true, true, false, false, false, false, true, false, false, true, true, false, false, true, true, true, false, false, false, false, false, false, false, false, true, false, true, true, false, false, false, false, true, true, false, false, true, false, false, false, false, true, false, false, false},
		},
	}

	for i, tc := range testcases {
		actual := checkIfPrerequisite(tc.numCourses, tc.prerequisites, tc.queries)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d failed: want:\n%v, got:\n%v", i+1, tc.expected, actual)
		}
	}
}
