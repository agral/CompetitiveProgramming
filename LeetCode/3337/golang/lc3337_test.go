package lc3337

import (
	"slices"
	"testing"
)

func areMatricesEqual(lhs [][]int, rhs [][]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range len(lhs) {
		if !slices.Equal(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

func Test_lengthAfterTransformation(t *testing.T) {
	testcases := []struct {
		s        string
		t        int
		nums     []int
		expected int
	}{
		{"abcyy", 2, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, 7},
		{"azbk", 1, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 8},
	}

	for i, tc := range testcases {
		actual := lengthAfterTransformations(tc.s, tc.t, tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}

func Test_makeIdentityMatrix(t *testing.T) {
	IDENTITY_MATRIX_ONE := [][]int{{1}}
	IDENTITY_MATRIX_TWO := [][]int{{1, 0}, {0, 1}}
	IDENTITY_MATRIX_FIVE := [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1}}

	created_one := makeIdentityMatrix(1)
	if !areMatricesEqual(IDENTITY_MATRIX_ONE, created_one) {
		t.Errorf("makeIdentityMatrix(%d) failed: want %v, got %v", 1, IDENTITY_MATRIX_ONE, created_one)
	}
	created_two := makeIdentityMatrix(2)
	if !areMatricesEqual(IDENTITY_MATRIX_TWO, created_two) {
		t.Errorf("makeIdentityMatrix(%d) failed: want %v, got %v", 1, IDENTITY_MATRIX_TWO, created_two)
	}
	created_five := makeIdentityMatrix(5)
	if !areMatricesEqual(IDENTITY_MATRIX_FIVE, created_five) {
		t.Errorf("makeIdentityMatrix(%d) failed: want %v, got %v", 1, IDENTITY_MATRIX_FIVE, created_five)
	}
}
