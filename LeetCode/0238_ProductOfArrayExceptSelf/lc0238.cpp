#include <cassert>
#include <vector>
#include <array>
#include <iostream>

class Solution {
public:
  std::vector<int> productExceptSelf(std::vector<int>& nums) {
    std::vector<int> prefixes, suffixes, ans;
    prefixes.resize(nums.size());
    suffixes.resize(nums.size());
    ans.resize(nums.size());
    prefixes[0] = 1;
    suffixes[nums.size() - 1] = 1;
    for (auto i = 1; i < nums.size(); i++) {
      prefixes[i] = prefixes[i-1] * nums[i-1];
      suffixes[nums.size() - 1 - i] = suffixes[nums.size() - i] * nums[nums.size() - i];
    }
    for (auto i = 0; i < nums.size(); i++) {
      ans[i] = prefixes[i] * suffixes[i];
    }
    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {1, 2, 3, 4};
  std::vector<int> expected_ans1 = {24, 12, 8, 6};
}
