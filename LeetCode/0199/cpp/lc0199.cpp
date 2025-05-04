#include <iostream>
#include <sstream>
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


// Runtime complexity: O(n), where n is the number of nodes in the tree.
// Auxiliary space complexity: O(h), where h is the height of the tree.
class Solution {
public:
    std::vector<int> rightSideView(TreeNode* root) {
        std::vector<int> ans = {};
        dfs(root, ans, 0);
        return ans;
    }
private:
    void dfs(TreeNode* root, std::vector<int>& ans, int level) {
        if (root == nullptr) {
            return;
        }
        if (level == ans.size()) {
            ans.push_back(root->val);
        }
        dfs(root->right, ans, level+1);
        dfs(root->left, ans, level+1);
    }
};

template<typename T>
std::string vectorToString(std::vector<T> vec) {
    std::string separator = "";
    std::stringstream ss;
    ss << "[";
    for (const auto& elem: vec) {
        ss << separator << elem;
        separator = ", ";
    }
    ss << "]";
    return ss.str();
}

int main() {
    struct Testcase {
        TreeNode* root;
        std::vector<int> expected;
    };
    Testcase testcases[] = {
        {makeTree({1, 2, 3, -1, 5, -1, 4}), {1, 3, 4}},
        {makeTree({1, 2, 3, 4, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1}), {1, 3, 4, 5}},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.rightSideView(tc.root);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
