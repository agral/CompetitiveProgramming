from typing import List
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+
# Solved on: 2026-02-09
class Solution:
    def balanceBST(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        values = []

        def traverse_inorder(root: Optional[TreeNode]) -> None:
            if not root:
                return
            traverse_inorder(root.left)
            values.append(root.val)
            traverse_inorder(root.right)

        traverse_inorder(root)

        def build_balanced_binary_tree(left: int, right: int) -> Optional[TreeNode]:
            if left > right:
                return None
            mid = (left + right) // 2
            return TreeNode(values[mid],
                            build_balanced_binary_tree(left, mid-1),
                            build_balanced_binary_tree(mid+1, right))

        return build_balanced_binary_tree(0, len(values) - 1)
