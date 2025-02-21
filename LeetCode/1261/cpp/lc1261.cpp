#include <unordered_set>
#include <vector>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(): val(0), left(nullptr), right(nullptr) {}
    TreeNode(int v): val(v), left(nullptr), right(nullptr) {}
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

class FindElements {
private:
    std::unordered_set<int> m_values;

public:
    FindElements(TreeNode* root) {
        dfs(root);
    }

    bool find(int target) {
        return m_values.contains(target);
    }
private:
    void dfs(TreeNode* node, int value = 0 /*for root node*/) {
        if (node == nullptr) {
            return;
        }

        m_values.insert(value);
        node->val = value;
        dfs(node->left, 2 * value + 1);
        dfs(node->right, 2 * value + 2);
    }
};

int main() {
    // checked manually. Don't really want to write a test suite for custom binary trees today.
}
