#include <algorithm>
#include <cassert>
#include <string>
#include <vector>
#include <iostream>

class TreeNode {
public:
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x = 0, TreeNode* left = nullptr, TreeNode *right = nullptr)
  : val(x), left(left), right(right) {
  }
};

TreeNode* makeTree(std::vector<int> values) {
  const int SZ = values.size();
  // number of values has to be 2^n - 1, where n is the tree depth.
  // E.g. a tree of depth 3 has 2^3 - 1 == 7 nodes.
  assert((SZ & SZ+1) == 0);
  std::vector<TreeNode*> nodes;
  for (int i = 0; i < SZ; ++i) {
    TreeNode* node = (values[i] == -1 ? nullptr : new TreeNode(values[i]));
    if (i > 0) {
      if (i % 2 == 0) {
        if (nodes[(i-1) / 2] != nullptr) {
          nodes[(i-1) / 2]->right = node;
        }
      } else {
        if (nodes[(i-1) / 2] != nullptr) {
          nodes[(i-1) / 2]->left = node;
        }
      }
    }
    nodes.push_back(node);
  }
  return SZ == 0 ? nullptr : nodes[0];
}

#include <deque>
void printTree(TreeNode* root) {
  if (root == nullptr) {
    return;
  }
  std::deque<TreeNode*> row{root};
  std::deque<TreeNode*> next_row{};
  for (int k = 0; k < 3; ++k) { //ugly hack
    for (const TreeNode* n: row) {
      std::cout << (n == nullptr ? "*" : std::to_string(n->val)) << "\t";
      if (n == nullptr) {
        next_row.push_back(nullptr);
        next_row.push_back(nullptr);
      } else {
        next_row.push_back(n->left);
        next_row.push_back(n->right);
      }
    }
    std::cout << "\n";
    row = next_row;
    next_row = {};
  }
}

class Solution {
public:
  std::string smallestFromLeaf(TreeNode* root) {
    std::string ans;
    dfs(root, "", ans);
    return ans;
  }

private:
  void dfs(TreeNode* root, std::string path, std::string& ans) {
    if (nullptr == root) {
      return;
    }
    path.push_back(static_cast<char>(root->val + 'a'));

    if (nullptr == root->left && nullptr == root->right) {
      std::reverse(path.begin(), path.end());
      if ((ans == "") || (ans > path)) {
        ans = path;
      }
      // Note: two calls to reverse() to undo changes are still faster than copying the string + one reverse().
      std::reverse(path.begin(), path.end());
    }

    dfs(root->left, path, ans);
    dfs(root->right, path, ans);
    path.pop_back();
  }
};

int main() {
  Solution s;
  {
    TreeNode* tree = makeTree({0, 1, 2, 3, 4, 3, 4});
    assert(s.smallestFromLeaf(tree) == "dba");
  }
  {
    TreeNode* tree = makeTree({25, 1, 3, 1, 3, 0, 2});
    assert(s.smallestFromLeaf(tree) == "adz");
  }
  {
    TreeNode* tree = makeTree({
                               2, 
                       2,             1, 
                  -1,     1,      0,     -1, 
                -1, -1, 0, -1, -1, -1, -1, -1});
    assert(s.smallestFromLeaf(tree) == "abc");
  }
}
