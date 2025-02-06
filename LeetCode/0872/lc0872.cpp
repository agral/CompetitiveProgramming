#include <cassert>
#include <vector>

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
  bool leafSimilar(TreeNode* root1, TreeNode* root2) {
    std::vector<int> leaves1, leaves2;
    dfs(root1, leaves1);
    dfs(root2, leaves2);
    return leaves1 == leaves2;
  }
private:
  void dfs(TreeNode* node, std::vector<int>& leaves) {
    if (node->left == nullptr && node->right == nullptr) {
      leaves.push_back(node->val);
      return;
    }
    if (node->left != nullptr) {
      dfs(node->left, leaves);
    }
    if (node->right != nullptr) {
      dfs(node->right, leaves);
    }
  }
};

int main() {
  Solution s;
  {
    TreeNode* root1 = makeTree({3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4, -1, -1, -1, -1});
    TreeNode* root2 = makeTree({3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8});
    assert(s.leafSimilar(root1, root2) == true);
  }
  {
    TreeNode* root1 = makeTree({1, 2, 3});
    TreeNode* root2 = makeTree({1, 3, 2});
    assert(s.leafSimilar(root1, root2) == false);
  }
}
