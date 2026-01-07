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

// Runtime complexity: O(n)
// Auxiliary space complexity: O(n)
// Subjective level: hard
// Solved on: 2026-01-07
class Solution {
public:
    int maxProduct(TreeNode *root) {
        constexpr int MOD = 1'000'000'007;
        m_sums = {};
        long total = getSum(root);
        long ans = 0L;
        for (long sum: m_sums) {
            ans = std::max(ans, sum * (total - sum));
        }
        /*
        for (long sum: m_sums) {
            std::cout << sum << " ";
        }
        std::cout << std::endl;
        */
        return ans % MOD;
    }

private:
    long getSum(TreeNode* root) {
        if (root == nullptr) {
            return 0;
        }
        long leftSum = getSum(root->left);
        long rightSum = getSum(root->right);
        long ans = leftSum + rightSum + root->val;
        m_sums.push_back(ans);
        return ans;
    }
    std::vector<long> m_sums;
};

int main() {
    struct Testcase {
        TreeNode *root;
        int expected;
    };
    Testcase testcases[] = {
        {makeTree({1, 2, 3, 4, 5, 6, -1}), 110},
        {makeTree({1, -1, 2, -1, -1, 3, 4, -1, -1, -1, -1, -1, -1, 5, 6}), 90},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.maxProduct(tc.root);
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
