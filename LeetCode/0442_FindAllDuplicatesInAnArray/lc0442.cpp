#include <cassert>
#include <vector>
#include <cmath>

class Solution {
public:
  std::vector<int> findDuplicates(std::vector<int>& nums) {
    std::vector<int> ans;
    for (const int k: nums) {
      // 1 <= nums[k] <= 10^5, meaning that all the numbers in nums are initially positive.
      // all integers used are in range <1, N> where N is nums' size.
      // For each array entry k, flip sign of an entry at kth-1 index
      // (guaranteed to exist by the above rules).
      // Given that each entry appears either once or twice, push entries for which
      // this value is positive - those are duplicates, seen for the second time.
      const int idx = abs(k) - 1;
      nums[idx] = -nums[idx];
      if (nums[idx] > 0) {
        ans.push_back(idx + 1);
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {4, 3, 2, 7, 8, 2, 3, 1};
    std::vector<int> expected = {2, 3};
    assert(s.findDuplicates(nums) == expected);
  }
  {
    std::vector<int> nums = {1, 1, 2};
    std::vector<int> expected = {1};
    assert(s.findDuplicates(nums) == expected);
  }
  {
    std::vector<int> nums = {1};
    std::vector<int> expected = {};
    assert(s.findDuplicates(nums) == expected);
  }
}
