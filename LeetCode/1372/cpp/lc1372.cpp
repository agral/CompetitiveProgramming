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
        if (i % 2 == 0) {
            nodes[parent_idx]->left = nodes[i];
        } else {
            nodes[parent_idx]->right = nodes[i];
        }
    }
    return root;
}

struct TreeStats {
    int maxLeft;
    int maxRight;
    int maxSubtree;
};

// Runtime complexity: O(n), where n is the total number of nodes in the tree.
// Auxiliary space complexity: O(h), where h denotes the height of the tree.
class Solution {
public:
    int longestZigZag(TreeNode* root) {
        return dfs(root).maxSubtree;
    }

private:
    TreeStats dfs(TreeNode* node) {
        if (node == nullptr) {
            return {-1, -1, -1};
        }
        TreeStats leftStats = dfs(node->left);
        TreeStats rightStats = dfs(node->right);
        int maxLeft = leftStats.maxRight + 1;
        int maxRight = rightStats.maxLeft + 1;
        int maxSubtree = std::max({maxLeft, maxRight,
            leftStats.maxSubtree, rightStats.maxSubtree});
        return {maxLeft, maxRight, maxSubtree};
    }
};

int main() {
    struct Testcase {
        TreeNode* root;
        int expected;
    };
    // I should rewrite the makeTree() at some future point to better account for sparse trees.
    // Writing these testcases was not that fun.
    Testcase testcases[] = {
        { makeTree({                                       +1,                                              // 1
                                    -1,                                             +1,                     // 2
                        -1,                     -1,                     +1,                    +1,          // 4
                 -1,         -1,         -1,         -1,         -1,         -1,         +1,         +1,    // 8
              -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   +1,   -1,   -1, // 16
            -1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,+1,-1,-1,-1,-1,// 32
        }), 3},
        { makeTree({                                       +1,                                              // 1
                                    +1,                                             +1,                     // 2
                        -1,                     +1,                     -1,                    -1,          // 4
                 -1,         -1,         +1,         +1,         -1,         -1,         -1,         -1,    // 8
              -1,   -1,   -1,   -1,   -1,   +1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1,   -1, // 16
        }), 4},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.longestZigZag(tc.root);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << actual << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
