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

bool is_same(TreeNode *tree1, TreeNode *tree2) {
    if (tree1 == nullptr) {
        return tree2 == nullptr;
    }
    if (tree2 == nullptr) {
        return false;
    }
    if (tree1->val != tree2->val) {
        return false;
    }
    return is_same(tree1->left, tree2->left) &&
           is_same(tree1->right, tree2->right);
}

// Runtime complexity:
// Auxiliary space complexity:
class Solution {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        TreeNode* ans = {};
        // TODO 
        return ans;
    }
};

int main() {
    TreeNode* testedTree = makeTree({5, 3, 6, 2, 4, -1, 7});
    struct Testcase {
        TreeNode* root;
        int key;
        TreeNode* expected;
    };
    Testcase testcases[] = {
        {makeTree({5, 3, 6, 2, 4, -1, 7}), 3, makeTree({5, 4, 6, 2, -1, -1, 7})},
        {makeTree({5, 3, 6, 2, 4, -1, 7}), 0, makeTree({5, 3, 6, 2, 4, -1, 7})},
        {nullptr, 0, nullptr},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.deleteNode(tc.root, tc.key);
        if (!is_same(actual, tc.expected)) {
            std::cout << "Testcase #" << (numGood + numBad + 1) << " failed.\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
