#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<std::vector<int>> divideArray(std::vector<int>& nums, int k) {
    std::ranges::sort(nums);
    std::vector<std::vector<int>> ans;
    ans.reserve(nums.size() / 3); // sunny day scenario. Will not use all that memory if problem is unsolvable.
    for (int i = 0; i < nums.size() - 2; i += 3) {
      if (nums[i + 2] - nums[i] > k) {
        // Can not divide the numbers into this subarray => return empty array.
        return {};
      }
      ans.push_back({nums[i], nums[i + 1], nums[i + 2]});
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {1, 3, 4, 8, 7, 9, 3, 5, 1};
    std::vector<std::vector<int>> expected = {{1, 1, 3}, {3, 4, 5}, {7, 8, 9}};
    assert(s.divideArray(nums, 2) == expected);
  }
  {
    std::vector<int> nums = {1, 3, 3, 2, 7, 3};
    std::vector<std::vector<int>> expected = {};
    assert(s.divideArray(nums, 3) == expected);
  }
}
