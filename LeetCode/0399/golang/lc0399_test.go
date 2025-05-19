package lc0399

import (
	"slices"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	testcases := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{ // # 01
			/*equations=*/ [][]string{{"a", "b"}, {"b", "c"}},
			/*values=*/ []float64{2.0, 3.0},
			/*queries=*/ [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			/*expected=*/ []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{ // # 02
			/*equations=*/ [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			/*values=*/ []float64{1.5, 2.5, 5.0},
			/*queries=*/ [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			/*expected=*/ []float64{3.75, 0.4, 5.0, 0.2},
		},
		{ // # 03
			/*equations=*/ [][]string{{"a", "b"}},
			/*values=*/ []float64{0.5},
			/*queries=*/ [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			/*expected=*/ []float64{0.5, 2.0, -1.0, -1.0},
		},
	}

	for i, tc := range testcases {
		actual := calcEquation(tc.equations, tc.values, tc.queries)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase #%02d failed: want %v, got %v", i+1, tc.expected, actual)
		}
	}
}
