#include <iostream>
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
  TreeNode *recoverFromPreorder(std::string traversal) {
    int idx = 0;
    return rebuildTreeFromPreorderTraversalString(traversal, 0, idx);
  }

private:
  TreeNode *rebuildTreeFromPreorderTraversalString(const std::string &traversal,
                                                   int depth, int &strOffset) {
    int numDashes = 0;
    while (numDashes + strOffset < traversal.size() &&
           traversal[numDashes + strOffset] == '-') {
      ++numDashes;
    }
    if (numDashes != depth) {
      return nullptr;
    }
    strOffset += depth;
    int beginning = strOffset;
    while (strOffset < traversal.size() && std::isdigit(traversal[strOffset])) {
      ++strOffset;
    }
    int value = std::stoi(traversal.substr(beginning, strOffset - beginning));
    return new TreeNode(
        value,
        rebuildTreeFromPreorderTraversalString(traversal, depth + 1, strOffset),
        rebuildTreeFromPreorderTraversalString(traversal, depth + 1,
                                               strOffset));
  }
};

int main() {
  struct Testcase {
    std::string traversalForm;
    TreeNode *expected;
  };
  Testcase testcases[] = {
      {"1-2--3--4-5--6--7", makeTree({1, 2, 5, 3, 4, 6, 7})},
      {"1-2--3---4-5--6---7",
       makeTree({1, 2, 5, 3, -1, 6, -1, 4, -1, -1, -1, 7, -1, -1, -1})},
      {"1-401--349---90--88",
       makeTree({1, 401, -1, 349, 88, -1, -1, 90, -1, -1, -1, -1, -1, -1, -1})},
  };
  Solution s{};
  int numGood = 0, numBad = 0;
  for (const Testcase &tc : testcases) {
    auto actual = s.recoverFromPreorder(tc.traversalForm);
    if (!is_same(actual, tc.expected)) {
      std::cout << "Testcase " << tc.traversalForm << " failed.\n";
      ++numBad;
    } else {
      ++numGood;
    }
  }
  std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " " << numGood << "/"
            << (numBad + numGood) << " testcases passed successfully.\n";
}
