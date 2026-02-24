package lc1022

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime complexity: O(n), where n is the number of the elements in the binary tree.
// Auxiliary space: O(h), where h is the height of the binary tree; generally O(log2(n)).
// Subjective level: medium
// Solved on: 2026-02-24
func sumRootToLeaf(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, accumulated int)
	dfs = func(root *TreeNode, accumulated int) {
		if root == nil {
			return
		}
		accumulated = accumulated*2 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += accumulated
			return
		}
		dfs(root.Left, accumulated)
		dfs(root.Right, accumulated)
	}

	dfs(root, 0)
	return ans
}
