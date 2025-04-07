#include <iostream>
#include <numeric>
#include <sstream>
#include <vector>

// Assuming that:
//  * n is the length of nums (a.k.a. the count of the elements in nums),
//  * s is the sum of all the numbers in nums;
// Runtime complexity: O(ns)
// Auxiliary space complexity: O(s)
class Solution {
public:
    bool canPartition(const std::vector<int>& nums) {
        const int sum = std::accumulate(nums.begin(), nums.end(), 0);
        if (sum & 1) {
            // Can't partition numbers into two sets of equal sum, if their sum is odd.
            return false;
        }

        const int target = sum / 2;
        // dp[i] is true iff i can be formed by nums seen so far.
        std::vector<bool> dp(target + 1);
        dp[0] = true;
        for (const int num: nums) {
            for (int i = target; i >= num; i--) {
                dp[i] = dp[i] || dp[i - num];
            }
        }

        return dp[target];
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
        std::vector<int> nums;
        bool expected;
    };
    Testcase testcases[] = {
        {{1, 5, 11, 5}, true},
        {{1, 2, 3, 5}, false},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        bool actual = s.canPartition(tc.nums);
        if (actual != tc.expected) {
            std::cout << "Testcase " << vectorToString(tc.nums) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
