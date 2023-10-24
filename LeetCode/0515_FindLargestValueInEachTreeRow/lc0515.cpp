#include <cassert>
#include <vector>
#include <iostream>

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
    if (i % 2 == 0) {
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
  void visit(TreeNode* node, int lvl=1) {
    if (node == nullptr) {
      return;
    }
    if (m_ans.size() == lvl) {
      // We have entered a new level, for which there is no recorded best value yet.
      // push back current value as best one:
      m_ans.push_back(node->val);
    }
    else {
      // This node is on a level that has already been visited.
      // Update the score if this node has greater value.
      m_ans[lvl] = std::max(m_ans[lvl], node->val);
    }
    // Finally, visit both children nodes, knowing that these are one level higher in tree hierarchy:
    visit(node->left, lvl + 1);
    visit(node->right, lvl + 1);
  }

  std::vector<int> largestValues(TreeNode* root) {
    m_ans = {};
    if (root == nullptr) {
      // a corner case.
      return {};
    }
    m_ans.push_back(root->val);
    visit(root->left);
    visit(root->right);
    return m_ans;
  }

  std::vector<int> m_ans;
};

int main() {
  Solution s;
  assert(s.largestValues(nullptr) == std::vector<int>{});

  TreeNode* example1 = makeTree({1, 3, 2, 5, 3, -1, 9});
  assert(s.largestValues(example1) == std::vector<int>({1, 3, 9}));

  TreeNode* example2 = makeTree({1, 2, 3});
  assert(s.largestValues(example2) == std::vector<int>({1, 3}));
}
