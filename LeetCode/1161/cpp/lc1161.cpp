#include <algorithm>
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
        if (i % 2 == 1) {
            nodes[parent_idx]->left = nodes[i];
        } else {
            nodes[parent_idx]->right = nodes[i];
        }
    }
    return root;
}

// Runtime complexity: O(n), where n is the total number of nodes in the tree.
// Auxiliary space complexity: O(h), where h is the height of the binary tree.
class Solution {
public:
    int maxLevelSum(TreeNode* root) {
        std::vector<int> levelSums{};
        dfs(root, levelSums, 0);
        std::vector<int>::iterator it = std::max_element(levelSums.begin(), levelSums.end());
        return std::distance(levelSums.begin(), it) + 1;
    }
    void dfs(TreeNode* root, std::vector<int>& levelSums, int level) {
        if (root == nullptr) {
            return;
        }
        if (level == levelSums.size()) {
            levelSums.push_back(0);
        }
        levelSums[level] += root->val;
        dfs(root->left, levelSums, level+1);
        dfs(root->right, levelSums, level+1);
    }
};

int main() {
    struct Testcase {
        TreeNode* root;
        int expected;
    };
    Testcase testcases[] = {
        {makeTree({1, 7, 0, 7, -8, -1, -1}), 2},
        {makeTree({989, -1, 10250, -1, -1, 98693, -89388, -1, -1, -1, -1, -1, -1, -1, -32127}), 2},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.maxLevelSum(tc.root);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
