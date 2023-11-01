#include <cassert>
#include <vector>
#include <unordered_map>
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
  void visit(TreeNode* node) {
    if (node == nullptr) {
      return;
    }

    // Store this node's value in the hashmap:
    if (histogram.find(node->val) == histogram.end()) {
      histogram[node->val] = 1;
    } else {
      histogram[node->val] += 1;
    }

    // visit both children nodes
    visit(node->left);
    visit(node->right);
  }

  std::vector<int> findMode(TreeNode* root) {
    histogram.clear();
    visit(root);

    // Find the mode in the histogram:
    std::vector<int> ans = {};
    int curr_score = -1;
    for (const auto& [k, v]: histogram) {
      if (v > curr_score) {
        // just found a value that is more frequent than currently assumed mode.
        // clear the mode vector, put this new value in. Set curr_score to new mode's count.
        ans.clear();
        ans.push_back(k);
        curr_score = v;
      }
      else if (v == curr_score) {
        // just found a value that has exactly same count as currently assumed mode.
        // Add it to the mode vector, preserve existing value(s).
        ans.push_back(k);
      }
      // Otherwise ignore it.
    }
    return ans;
  }

  std::unordered_map<int, int> histogram;
};

int main() {
  Solution s;

  TreeNode* example1 = makeTree({1, -1, 2, -1, -1, 2, -1});
  assert(s.findMode(example1) == std::vector<int>({2}));

  TreeNode* example2 = makeTree({0});
  assert(s.findMode(example2) == std::vector<int>({0}));
}
