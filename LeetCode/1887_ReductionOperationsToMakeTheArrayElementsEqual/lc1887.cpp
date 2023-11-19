#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int reductionOperations(std::vector<int>& nums) {
    std::sort(nums.begin(), nums.end());
    int ans = 0;
    int count_of_different_numbers = 0;
    for (int i = 1; i < nums.size(); i++) {
      if (nums[i] > nums[i-1]) {
        ++count_of_different_numbers;
      }
      ans += count_of_different_numbers;
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {5, 1, 3};
    assert(s.reductionOperations(nums) == 3);
  }
  {
    std::vector<int> nums = {1, 1, 1};
    assert(s.reductionOperations(nums) == 0);
  }
  {
    std::vector<int> nums = {1, 1, 2, 2, 3};
    assert(s.reductionOperations(nums) == 4);
  }
}
