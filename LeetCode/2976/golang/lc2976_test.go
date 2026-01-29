package lc2976

import "testing"

func Test_minimumCost(t *testing.T) {
	testcases := []struct {
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
		expected int64
	}{
		{
			"abcd",                               // source
			"acbe",                               // target
			[]byte{'a', 'b', 'c', 'c', 'e', 'd'}, // original
			[]byte{'b', 'c', 'b', 'e', 'b', 'e'}, // changed
			[]int{2, 5, 5, 1, 2, 20},             //cost
			int64(28),                            //expected
		},
		{
			"aaaa",           // source
			"bbbb",           // target
			[]byte{'a', 'c'}, // original
			[]byte{'c', 'b'}, // changed
			[]int{1, 2},      //cost
			int64(12),        //expected
		},
		{
			"abcd",       // source
			"abce",       // target
			[]byte{'a'},  // original
			[]byte{'e'},  // changed
			[]int{10000}, //cost
			int64(-1),    //expected
		},
		/* copy for new test cases:
		{
			"",       // source
			"",       // target
			[]byte{}, // original
			[]byte{}, // changed
			[]int{},  //cost
			int64(0), //expected
		},
		*/
	}

	for i, tc := range testcases {
		actual := minimumCost(tc.source, tc.target, tc.original, tc.changed, tc.cost)
		if actual != tc.expected {
			t.Errorf("Testcase minimumCost#%02d (%s -> %s) failed: want %d, got %d",
				i+1, tc.source, tc.target, tc.expected, actual)
		}
	}
}
