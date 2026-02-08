from typing import List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Runtime complexity: O(n*log(n))
# Auxiliary space complexity: O(h), where h is the (maximum) height of the tested binary tree
# Subjective level: medium-
# Solved on: 2026-02-08
class Solution:
    def isBalanced(self, root: TreeNode | None) -> bool:
        if not root:
            return True

        def maxDepth(root: TreeNode | None) -> int:
            if not root:
                return 0
            return (max(maxDepth(root.left), maxDepth(root.right)) + 1)

        return (
            abs(maxDepth(root.left) - maxDepth(root.right)) <= 1 and
            self.isBalanced(root.left) and
            self.isBalanced(root.right)
        )
