package lc1022

import (
	"strconv"
	"testing"
)

func MakeBinaryTree(values []int) *TreeNode {
	// 1 2 3 4 5 6 7 -> 1 | 2 3 | 4 5 6 7
	numNodes := len(values)
	if numNodes == 0 {
		return nil
	}
	nodes := make([]*TreeNode, numNodes)
	head := &TreeNode{values[0], nil, nil}
	nodes[0] = head
	for i := 1; i < len(values); i++ {
		nodes[i] = &TreeNode{values[i], nil, nil}
		parentIdx := int((i - 1) / 2)
		if i%2 == 1 {
			nodes[parentIdx].Left = nodes[i]
		} else {
			nodes[parentIdx].Right = nodes[i]
		}
	}
	return head
}

func ToString(head *TreeNode) string {
	if head == nil {
		return "nil"
	}
	ans := ""
	level := []*TreeNode{}
	level = append(level, head)
	for len(level) > 0 {
		ans += "[ "
		nextLevel := []*TreeNode{}
		for i := 0; i < len(level); i++ {
			ans += strconv.Itoa(level[i].Val) + " "
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		ans += "] "
		level = nextLevel
	}
	return ans
}

func Test_sumRootToLeaf(t *testing.T) {
	testcases := []struct {
		root     *TreeNode
		expected int
	}{
		{MakeBinaryTree([]int{1, 0, 1, 0, 1, 0, 1}), 22},
		{MakeBinaryTree([]int{0}), 0},
	}

	for i, tc := range testcases {
		actual := sumRootToLeaf(tc.root)
		if actual != tc.expected {
			t.Errorf("Testcase sumRootToLeaf#%02d (%v) failed: want %d, got %d",
				i+1, ToString(tc.root), tc.expected, actual)
		}
	}
}
