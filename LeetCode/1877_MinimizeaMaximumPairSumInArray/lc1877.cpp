#include <algorithm>
#include <cassert>
#include <vector>

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
  Solution s;
  {
    std::vector<int> nums = {3, 5, 2, 3};
    assert(s.minPairSum(nums) == 7);
  }
  {
    std::vector<int> nums = {3, 5, 4, 2, 4, 6};
    assert(s.minPairSum(nums) == 8);
  }
}
