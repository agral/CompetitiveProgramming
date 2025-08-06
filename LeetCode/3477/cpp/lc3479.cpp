#include <iostream>
#include <sstream>
#include <vector>

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

class SegmentTree {
public:
    explicit SegmentTree(const std::vector<int>& nums): m_size(nums.size()), m_tree(4 * m_size) {
        buildTree(nums, 0, 0, m_size-1);
    }

    // Returns the first index where tree[i] >= minValue; or -1 if such index is not found.
    int queryFirstIndex(int minValue) {
        return doQueryFirstIndex(0, 0, m_size-1, minValue);
    }

    // Sets the value of `tree[index] = value`
    void update(int index, int value) {
        doUpdate(0, 0, m_size-1, index, value);
    }

private:
    int m_size;
    std::vector<int> m_tree;

    void buildTree(const std::vector<int>& nums, int treeIndex, int low, int high) {
        if (low == high) {
            m_tree[treeIndex] = nums[low];
            return;
        }
        int mid = (low + high) / 2;
        buildTree(nums, 2*treeIndex+1, low, mid);
        buildTree(nums, 2*treeIndex+2, mid+1, high);
        m_tree[treeIndex] = std::max(m_tree[2*treeIndex+1], m_tree[2*treeIndex+2]);
    }

    void doUpdate(int treeIndex, int low, int high, int index, int value) {
        if (low == high) {
            m_tree[treeIndex] = value;
            return;
        }
        int mid = (low + high) / 2;
        if (index <= mid) {
            doUpdate(2*treeIndex+1, low, mid, index, value);
        } else {
            doUpdate(2*treeIndex+2, mid+1, high, index, value);
        }
        m_tree[treeIndex] = std::max(m_tree[2*treeIndex+1], m_tree[2*treeIndex+2]);
    }

    int doQueryFirstIndex(int treeIndex, int low, int high, int minValue) {
        if (m_tree[treeIndex] < minValue) {
            return -1;
        }
        if (low == high) {
            update(low, -1);
            return low;
        }
        int mid = (low + high) / 2;
        int left = m_tree[2*treeIndex+1];
        if (left >= minValue) {
            return doQueryFirstIndex(2*treeIndex+1, low, mid, minValue);
        } else {
            return doQueryFirstIndex(2*treeIndex+2, mid+1, high, minValue);
        }
    }
};

// Runtime complexity: O(nlogn)
// Auxiliary space complexity: O(n)
class Solution {
public:
    int numOfUnplacedFruits(std::vector<int>& fruits, std::vector<int>& baskets) {
        int ans = 0;
        SegmentTree tree(baskets);
        for (const int fruit: fruits) {
            if (tree.queryFirstIndex(fruit) == -1) {
                ans += 1;
            }
        }
        return ans;
    }
};


int main() {
    struct Testcase {
        std::vector<int> fruits;
        std::vector<int> baskets;
        int expected;
    };
    Testcase testcases[] = {
        {std::vector<int>{4, 2, 5}, std::vector<int>{3, 5, 4}, 1},
        {std::vector<int>{3, 6, 1}, std::vector<int>{6, 4, 7}, 0},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.numOfUnplacedFruits(tc.fruits, tc.baskets);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood + numBad + 1) << " ("
                << vectorToString(tc.fruits) << " | " << vectorToString(tc.baskets)
                << ") failed. Got: " << actual << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
