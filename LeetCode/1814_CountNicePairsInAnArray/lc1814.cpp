#include <cassert>
#include <unordered_map>
#include <vector>

class Solution {
public:
  int countNicePairs(std::vector<int>& nums) {
    const int MOD = 1e9 + 7;
    std::unordered_map<int, long> nice_pairs;

    for (int i = 0; i < nums.size(); i++) {
      nice_pairs[nums[i] - reversed(nums[i])] += 1;
    }

    long ans = 0;
    for (const auto& [_, size]: nice_pairs) {
      // A set of size k contains k * (k - 1) / 2 distinct pairs.
      ans = (ans + size * (size - 1) / 2) % MOD;
    }
    return ans;
  }

private:
  int reversed(int num) {
    int ans = 0;
    while (num > 0) {
      ans *= 10;
      ans += (num % 10);
      num /= 10;
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {42, 11, 1, 97};
    assert(s.countNicePairs(nums) == 2);
  }
  {
    std::vector<int> nums = {13, 10, 35, 24, 76};
    assert(s.countNicePairs(nums) == 4);
  }
}
