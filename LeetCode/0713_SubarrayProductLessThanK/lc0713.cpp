#include <cassert>
#include <vector>

class Solution {
public:
  int numSubarrayProductLessThanK(std::vector<int>& nums, int k) {
    if (k <= 1) {
      return 0;
    }

    int result = 1;
    int ans = 0;
    for (int left = 0, right = 0; right < nums.size(); ++right) {
      result *= nums[right];
      while (result >= k) {
        result /= nums[left];
        left += 1;
      }
      ans += right - left + 1;
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {10, 5, 2, 6};
    assert(s.numSubarrayProductLessThanK(nums, 100) == 8);
  }
  {
    std::vector<int> nums = {1, 2, 3};
    assert(s.numSubarrayProductLessThanK(nums, 0) == 0);
  }
}
