package lc1980

import "testing"

func isValidAnswer(candidate string, nums []string) bool {
	if len(candidate) != len(nums[0]) {
		return false
	}
	for _, num := range nums {
		if num == candidate {
			return false
		}
	}
	return true
}

func Test_findDifferentBinaryString(t *testing.T) {
	testcases := []struct {
		nums []string
	}{
		{[]string{"01", "10"}},
		{[]string{"00", "01"}},
		{[]string{"111", "011", "001"}},
	}

	for i, tc := range testcases {
		actual := findDifferentBinaryString(tc.nums)
		if !isValidAnswer(actual, tc.nums) {
			t.Errorf("Testcase %02d failed: got %s which does not match %v", i+1, actual, tc.nums)
		}
	}
}
