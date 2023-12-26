#include <cassert>
#include <vector>

class Solution {
public:
  int numRollsToTarget(int n, int k, int target) {
    const int MOD = 1'000'000'007;
    std::vector<int> dp(target + 1);
    dp[0] = 1;
    for (int diceNo = 0; diceNo < n; diceNo++) {
      std::vector<int> next_dp(target + 1);
      for (int roll = 1; roll <= k; roll++) {
        for (int t = roll; t <= target; t++) {
          next_dp[t] = (next_dp[t] + dp[t - roll]) % MOD;
        }
      }
      dp = std::move(next_dp);
    }
    return dp[target];
  }
};

int main() {
  Solution s;
  assert(s.numRollsToTarget(1, 6, 3) == 1);
  assert(s.numRollsToTarget(2, 6, 7) == 6);
  assert(s.numRollsToTarget(30, 30, 500) == 222'616'187);
}
