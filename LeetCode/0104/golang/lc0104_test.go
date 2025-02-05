package lc0104

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	testcases := []struct {
		head     *TreeNode
		expected int
	}{
		{MakeBinaryTree([]int{3, 9, 20, -10, -10, 15, 7}), 3},
		{MakeBinaryTree([]int{1, -10, 2}), 2},
	}

	for i, tc := range testcases {
		actual := maxDepth(tc.head)
		if actual != tc.expected {
			fmt.Printf("Tree: %s", ToString(tc.head))
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
