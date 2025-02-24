#include <iostream>
#include <unordered_map>
#include <vector>

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int v) : val(v), left(nullptr), right(nullptr) {}
  TreeNode(int v, TreeNode *left, TreeNode *right)
      : val(v), left(left), right(right) {}
};

TreeNode *makeTree(std::vector<int> values) {
  if (values.empty()) {
    return nullptr;
  }

  std::vector<TreeNode *> nodes(values.size(), nullptr);
  TreeNode *root = new TreeNode(values[0]);
  nodes[0] = root;
  for (int i = 1; i < values.size(); i++) {
    if (values[i] == -1) {
      continue; // -1 is used to denote null nodes.
    }
    TreeNode *curr = new TreeNode(values[i]);
    nodes[i] = curr;
    //     0       Link parent to this child.
    //    / \      Left-side children have odd indices, while
    //   1   2     right-side children have even indices.
    //  / \ / \    The parent's index is always floor((n-1)/2).
    // 3  45   6
    int parent_idx = (i - 1) / 2;
    if (i % 2 == 0) {
      nodes[parent_idx]->left = nodes[i];
    } else {
      nodes[parent_idx]->right = nodes[i];
    }
  }
  return root;
}

bool is_same(TreeNode *tree1, TreeNode *tree2) {
  if (tree1 == nullptr) {
    return tree2 == nullptr;
  }
  if (tree2 == nullptr) {
    return false;
  }
  return is_same(tree1->left, tree2->left) &&
         is_same(tree1->right, tree2->right);
}

class Solution {
public:
  TreeNode *constructFromPrePost(std::vector<int> &preorder,
                                 std::vector<int> &postorder) {
    m_postorderToIndex = {};
    for (int i = 0; i < postorder.size(); i++) {
      m_postorderToIndex[postorder[i]] = i;
    }

    return reconstructTree(preorder, 0, preorder.size() - 1, postorder, 0,
                           postorder.size() - 1);
  }

private:
  TreeNode *reconstructTree(std::vector<int> &preorder, int preStartIdx,
                            int preEndIdx, std::vector<int> &postorder,
                            int postStartIdx, int postEndIdx) {
    if (preStartIdx > preEndIdx) {
      return nullptr;
    }
    if (preStartIdx == preEndIdx) {
      return new TreeNode(preorder[preStartIdx]);
    }
    int nodeVal = preorder[preStartIdx];
    int nodeLeftVal = preorder[preStartIdx + 1];
    int nodeLeftPostorderIdx = m_postorderToIndex[nodeLeftVal];
    int leftLen = nodeLeftPostorderIdx + 1 - postStartIdx;

    TreeNode *node = new TreeNode(nodeVal);
    node->left =
        reconstructTree(preorder, preStartIdx + 1, preStartIdx + leftLen,
                        postorder, postStartIdx, nodeLeftPostorderIdx);
    node->right =
        reconstructTree(preorder, preStartIdx + leftLen + 1, preEndIdx,
                        postorder, nodeLeftPostorderIdx + 1, postEndIdx - 1);
    return node;
  }

  std::unordered_map<int, int> m_postorderToIndex;
};

int main() {
  struct Testcase {
    std::vector<int> preorder;
    std::vector<int> postorder;
    TreeNode *expected;
  };
  Testcase testcases[] = {
      {{1, 2, 4, 5, 3, 6, 7},
       {4, 5, 2, 6, 7, 3, 1},
       makeTree({1, 2, 5, 3, 4, 6, 7})},

      {{1}, {1}, makeTree({1})},
  };
  Solution s{};
  int numGood = 0, numBad = 0;
  for (Testcase &tc : testcases) {
    auto actual = s.constructFromPrePost(tc.preorder, tc.postorder);
    if (!is_same(actual, tc.expected)) {
      std::cout << "Testcase failed.\n";
      ++numBad;
    } else {
      ++numGood;
    }
  }
  std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " " << numGood << "/"
            << (numBad + numGood) << " testcases passed successfully.\n";
}
