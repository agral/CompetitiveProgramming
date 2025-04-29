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

// Complexity analysis assumes:
// best case: a perfectly balanced (level) tree of n nodes;
// worst case: a maximally degenerated tree - in effect, a linked list of n nodes.
//
// Runtime complexity: O(n*logn) [best case] to O(n^2) [worst case].
// Auxiliary space complexity: O(logn) [best case] to O(n) [worst case].
class Solution {
public:
    int pathSum(TreeNode* root, int targetSum) {
        if (root == nullptr) {
            return 0;
        }
        return dfs(root, targetSum) + pathSum(root->left, targetSum)
            + pathSum(root->right, targetSum);
    }
private:
    int dfs(TreeNode* node, long long targetSum) {
        if (node == nullptr) {
            return 0;
        }
        return (node->val == targetSum) +
            dfs(node->left, targetSum - node->val) +
            dfs(node->right, targetSum - node->val);
    }
};

int main() {
    struct Testcase {
        TreeNode* root;
        int targetSum;
        int expected;
    };
    Testcase testcases[] = {
        {makeTree({10, 5, -3, 3, 2, -1, 11, 3, -1, -1, 1, -1, -1, -1, -1}), 8, 3},
        {makeTree({5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1, -1, -1}), 22, 3},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.pathSum(tc.root, tc.targetSum);
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
