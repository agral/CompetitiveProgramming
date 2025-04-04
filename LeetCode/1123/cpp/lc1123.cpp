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

TreeNode* makeTree(std::vector<int> values) {
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

// Holds a pointer to a TreeNode and its LCA score.
struct LcaTreeNode {
    TreeNode* node;
    int lcaScore;
};

// Runtime complexity: O(h), where h is the height of the tree.
// Auxiliary space complexity: O(h).
class Solution {
public:
    TreeNode* lcaDeepestLeaves(TreeNode* root) {
        return dfs(root).node;
    }
private:
    LcaTreeNode dfs(TreeNode* root) {
        if (root == nullptr) {
            return {nullptr, 0};
        }
        LcaTreeNode lcaLeft = dfs(root->left);
        LcaTreeNode lcaRight = dfs(root->right);
        if (lcaLeft.lcaScore > lcaRight.lcaScore) {
            return {lcaLeft.node, lcaLeft.lcaScore + 1};
        } else if (lcaLeft.lcaScore < lcaRight.lcaScore) {
            return {lcaRight.node, lcaRight.lcaScore + 1};
        } else {
            return {root, lcaRight.lcaScore + 1}; 
        }
    }
};

int main() {
    struct Testcase {
        std::vector<int> tree_representation;
        int expected;
    };
    Testcase testcases[] = {
        {{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4, -1, -1, -1, -1}, 2},
        {{42}, 42},
        {{0, 1, 3, -1, -1, 2, -1}, 2},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        TreeNode* tree = makeTree(tc.tree_representation);
        auto actual = s.lcaDeepestLeaves(tree)->val;
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
