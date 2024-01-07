#include <cassert>
#include <vector>
#include <unordered_map>

class Solution {
public:
  int numberOfArithmeticSlices(std::vector<int>& nums) {
    const int SIZE = nums.size();

    std::unordered_map<long long, std::vector<int>> valToNumsIndexes{};
    for (int i = 0; i < SIZE; ++i) {
      valToNumsIndexes[nums[i]].push_back(i);
    }

    int ans = 0;
    std::vector<std::vector<int>> dp(SIZE, std::vector<int>(SIZE));
    for (int i = 1; i < SIZE; ++i) {
      for (int j = 0; j < i; ++j) {
        long long target = 2LL * nums[j] - nums[i];
        auto it = valToNumsIndexes.find(target);
        if (it != valToNumsIndexes.end()) {
          for (int idx: it->second) {
            if (idx < j) {
              dp[i][j] += (1 + dp[j][idx]);
            }
          }
        }
        ans += dp[i][j];
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {2, 4, 6, 8, 10};
    assert(s.numberOfArithmeticSlices(nums) == 7);
  }
  {
    std::vector<int> nums = {7, 7, 7, 7, 7};
    assert(s.numberOfArithmeticSlices(nums) == 16);
  }
}
