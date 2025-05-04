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

// Runtime complexity: O(h)
// Auxiliary space complexity: O(h)
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == nullptr || root == p || root == q) {
            return root;
        }
        TreeNode* leftPart = lowestCommonAncestor(root->left, p, q);
        TreeNode* rightPart = lowestCommonAncestor(root->right, p, q);
        if (leftPart != nullptr && rightPart != nullptr) {
            return root;
        }
        if (rightPart == nullptr) {
            return leftPart;
        }
        return rightPart;
    }
};

int main() {
    struct Testcase {
        TreeNode* root;
        TreeNode* p;
        TreeNode* q;
        TreeNode* expected;
    };
    TreeNode *testedTree = makeTree({3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4, -1, -1, -1, -1});
    Testcase testcases[] = {
        {testedTree, /*5=*/testedTree->left, /*1=*/testedTree->right, /*3=*/testedTree},
        {testedTree, /*5=*/testedTree->left, /*4=*/testedTree->left->right->right, /*5=*/testedTree->left},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        TreeNode* actual = s.lowestCommonAncestor(tc.root, tc.p, tc.q);
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.root << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
