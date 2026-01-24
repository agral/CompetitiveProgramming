#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

// Runtime complexity: O(sort)
// Auxiliary space complexity: O(sort)
// Subjective level: easy+
// Solved on: 2023-11-17
// Additionally, minor fixes applied on 2026-01-24. Solution not changed.
class Solution {
public:
  int minPairSum(std::vector<int>& nums) {
    std::sort(nums.begin(), nums.end());
    int ans = -1;
    for (int b = 0, e = nums.size() - 1; b < e; b++, e--) {
      int pairsum = nums[b] + nums[e];
      if (pairsum > ans) {
        ans = pairsum;
      }
    }
    return ans;
  }
};

int main() {
    struct Testcase {
        std::vector<int> nums;
        int expected;
    };
    Testcase testcases[] = {
        {std::vector<int>{3, 5, 2, 3}, 7},
        {std::vector<int>{3, 5, 4, 2, 4, 6}, 8},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.minPairSum(tc.nums);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood+numBad+1) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
