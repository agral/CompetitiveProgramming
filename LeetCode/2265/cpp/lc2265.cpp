#include <cassert>
#include <vector>
#include <utility>

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
  int averageOfSubtree(TreeNode* root) {
    m_ans = 0;
    visit(root);
    return m_ans;
  }

  // Returns sum of entire subtree (first) and total count of subtree nodes (second);
  // both information necessary to compute averages further up the binary tree.
  std::pair<int, int> visit(TreeNode* node) {
    if (node == nullptr) {
      return {0, 0};
    };
    auto left = visit(node->left);
    auto right = visit(node->right);
    int sum = left.first + right.first + node->val;
    int count = left.second + right.second + 1;

    // Note: for this task, average is sum / count, rounded down to nearest integer.
    if (node->val == sum / count) {
      m_ans += 1;
    }

    return {sum, count};
  }

private:
  int m_ans;
};

int main() {
  Solution s;
  TreeNode* example1 = makeTree({4, 8, 5, 0, 1, -1, 6});
  assert(s.averageOfSubtree(example1) == 5);

  TreeNode* example2 = makeTree({1});
  assert(s.averageOfSubtree(example2) == 1);
}
