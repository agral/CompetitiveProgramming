package lc0104

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
		parent_idx := int((i - 1) / 2)
		if i%2 == 1 {
			nodes[parent_idx].Left = nodes[i]
		} else {
			nodes[parent_idx].Right = nodes[i]
		}
	}
	return head
}

func ToString(head *TreeNode) string {
	ans := ""
	if head == nil {
		return "nil"
	}
	level := []*TreeNode{}
	level = append(level, head)
	for len(level) > 0 {
		ans = ans + "[ "
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
