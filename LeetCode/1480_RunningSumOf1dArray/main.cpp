#include <cassert>
#include <vector>
#include <iostream>

class Solution {
public:
  std::vector<int> runningSum(std::vector<int>& nums) {
    std::vector<int> ans{};
    ans.resize(nums.size());
    if (nums.size() > 0) {
      ans[0] = nums[0];
      for (std::size_t i = 1; i < nums.size(); i++) {
        ans[i] = ans[i-1] + nums[i];
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> input_1 {1, 2, 3, 4};
  assert(s.runningSum(input_1) == std::vector<int>({1, 3, 6, 10}));

  std::vector<int> input_2 {1, 1, 1, 1, 1};
  assert(s.runningSum(input_2) == std::vector<int>({1, 2, 3, 4, 5}));

  std::vector<int> input_3 {3, 1, 2, 10, 1};
  assert(s.runningSum(input_3) == std::vector<int>({3, 4, 6, 16, 17}));

  std::vector<int> input_4{};
  assert(s.runningSum(input_4) == std::vector<int>({}));

  std::vector<int> input_5{1};
  assert(s.runningSum(input_5) == std::vector<int>({1}));
}
