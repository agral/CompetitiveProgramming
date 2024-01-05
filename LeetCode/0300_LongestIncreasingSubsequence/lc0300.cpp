#include <algorithm>
#include <cassert>
#include <vector>

// DP solution with O(n^2) runtime complexity
// and O(n) auxiliary memory requirement.
// Can be done better with a binary search approach, O(nlogn).
class Solution {
public:
  int lengthOfLIS(std::vector<int>& nums) {
    if (nums.empty()) {
      return 0;
    }

    // dp[i] holds length of longest increasing subsequence of nums[0..i].
    std::vector<int> dp(nums.size(), 1);
    for (int i = 1; i < nums.size(); ++i) {
      for (int j = 0; j < i; ++j) {
        if (nums[i] > nums[j]) {
          dp[i] = std::max(dp[i], dp[j]+1);
        }
      }
    }
    return *(std::max_element(dp.begin(), dp.end()));
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {10, 9, 2, 5, 3, 7, 101, 18};
    assert(s.lengthOfLIS(nums) == 4);
  }
  {
    std::vector<int> nums = {0, 1, 0, 3, 2, 3};
    assert(s.lengthOfLIS(nums) == 4);
  }
  {
    std::vector<int> nums = {7, 7, 7, 7, 7, 7, 7};
    assert(s.lengthOfLIS(nums) == 1);
  }
}
