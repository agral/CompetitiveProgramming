#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int maxFrequency(std::vector<int>& nums, int k) {
    std::sort(nums.begin(), nums.end());
    long range_sum{0L};
    int ans{0};
    for (int b = 0, e = 0; e < nums.size(); e++) {
      range_sum += static_cast<long>(nums[e]);
      while ((static_cast<long>(nums[e]) * (e - b + 1)) - k > range_sum) {
        range_sum -= nums[b++];
      }
      ans = std::max(ans, e - b + 1);
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {1, 2, 4};
    assert(s.maxFrequency(nums, 5) == 3); // can do {4, 4, 4}, so freq=3.
  }
  {
    std::vector<int> nums = {1, 4, 8, 13};
    assert(s.maxFrequency(nums, 5) == 2); // many such solutions, e.g. {4, 4, 8, 13}.
  }
  {
    std::vector<int> nums = {3, 9, 6}; // note: input array can be unsorted.
    assert(s.maxFrequency(nums, 2) == 1); // can't do anything significant. freq=1.
  }
}
