#include <cassert>
#include <vector> // required by makeTree (testing)

struct TreeNode {
  TreeNode(): val{}, left(nullptr), right(nullptr) {}
  TreeNode(int x): val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode* left, TreeNode* right): val(x), left(left), right(right) {}
  int val;
  TreeNode* left;
  TreeNode* right;
};

TreeNode* makeTree(std::vector<int> values) {
  if (values.empty()) {
    return nullptr;
  }

  std::vector<TreeNode*> nodes(values.size(), nullptr);
  TreeNode* root = new TreeNode(values[0]);
  nodes[0] = root;
  for (int i = 1; i < values.size(); i++) {
    if (values[i] == -1) {
      continue; // -1 is used to denote null nodes.
    }
    TreeNode* curr = new TreeNode(values[i]);
    nodes[i] = curr;
    //     0       Link parent to this child.
    //    / \      Left-side children have odd indices, while
    //   1   2     right-side children have even indices.
    //  / \ / \    The parent's index is always floor((n-1)/2).
    // 3  45   6
    int parent_idx = (i-1) / 2;
    if (i % 2 == 1) {
      nodes[parent_idx]->left = nodes[i];
    }
    else {
      nodes[parent_idx]->right = nodes[i];
    }
  }
  return root;
}

class Solution {
public:
  int rangeSumBST(TreeNode* root, int low, int high) {
    if (root == nullptr) {
      return 0;
    }
    if (root->val == 5) {
    }
    if (root->val < low) {
      // don't include this node, ignore the left subtree, continue checking the right subtree.
      return rangeSumBST(root->right, low, high);
    }
    if (root->val > high) {
      // don't include this node, continue checking the left subtree only, ignore the right subtree.
      return rangeSumBST(root->left, low, high);
    }
    // A default case - root's value is in desired range.
    // Include this node's value, continue checking in both subtrees.
    return root->val + rangeSumBST(root->left, low, high) + rangeSumBST(root->right, low, high);
  }
};

int main() {
  Solution s;
  {
    TreeNode* root = makeTree({10, 5, 15, 3, 7, -1, 18});
    auto v = s.rangeSumBST(root, 7, 15);
    assert(s.rangeSumBST(root, 7, 15) == 32);
  }
  /*
  {
    TreeNode* root = makeTree({10, 5, 15, 3, 7, 13, 18, 1, -1, 6, -1, -1, -1, -1, -1});
    assert(s.rangeSumBST(root, 7, 15) == 23);
  }
  */
}
