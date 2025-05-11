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

// Runtime complexity: O(h),
// Auxiliary space complexity: O(h),
// where h is the height of the BST; this is equal to O(logn) in balanced BSTs.
class Solution {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr) {
            return nullptr;
        }
        if (root->val == key) {
            // iff the node has only one descendant (left or right), it will take place
            // of the deleted root node:
            if (root->left == nullptr) {
                return root->right;
            }
            if (root->right == nullptr) {
                return root->left;
            }

            // Root node has both descendants; need to keep the tree balanced
            // by swapping root with a minimal node. This operation needs
            // to be done recursively; this simple recursive swap keeps
            // the tree balanced (so it's still a BST, not just any binary tree).
            TreeNode* minimal = getMinimalNode(root->right);
            root->right = deleteNode(root->right, minimal->val);
            minimal->left = root->left;
            minimal->right = root->right;
            root = minimal;
            // note: I'm not bothering with `free()`/`delete`ing the removed root node.
            // Any real implementation needs to take care of that.
        }
        else if (root->val < key) {
            root->right = deleteNode(root->right, key);
        } else { // implied: root->val > key
            root->left = deleteNode(root->left, key);
        }
        return root;
    }
private:
    TreeNode* getMinimalNode(TreeNode* tn) {
        while (tn->left != nullptr) {
            tn = tn->left;
        }
        return tn;
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
