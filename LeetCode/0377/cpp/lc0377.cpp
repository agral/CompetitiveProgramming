#include <cassert>
#include <vector>

class Solution {
public:
  int combinationSum4(std::vector<int>& nums, int target) {
    // In fact, it's a permutation sum, not a combination sum.

    std::vector<unsigned> dp;
    dp.resize(target + 1, 0);
    dp[0] = 1;
    for (int i = 1; i <= target; i++) {
      for (const auto& num: nums) {
        if (num <= i) {
          dp[i] += dp[i - num];
        }
      }
    }
    return dp[target];
  }
};

int main() {
  Solution s;

  std::vector<int> example1 = {1, 2, 3};
  assert(s.combinationSum4(example1, 4) == 7);

  std::vector<int> example2 = {9};
  assert(s.combinationSum4(example2, 3) == 0);
}
