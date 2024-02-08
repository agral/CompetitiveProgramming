#include <cassert>
#include <vector>

class Solution {
public:
  int numSquares(int n) {
    // dp / bfs approach:
    std::vector<int> dp(n + 1, n);
    dp[0] = 0; // a corner case; actually 0 == 0 * 0. But it would break bfs.
    dp[1] = 1; // 1 == 1 * 1

    for (int i = 2; i <= n; ++i) {
      for (int square_num = 1; square_num * square_num <= i; ++square_num) {
        // number i can be made by either dp[i] or dp[i-k] + 1; where k is a square number.
        // e.g. 5 can be made by either dp[5], or by dp[1] + 1, when using k = 4 == 2*2.
        dp[i] = std::min(dp[i], dp[i - (square_num * square_num)] + 1);
      }
    }

    return dp[n];
  }
};

int main() {
  Solution s;
  assert(s.numSquares(12) == 3); // 12 == 4 + 4 + 4
  assert(s.numSquares(13) == 2); // 13 == 9 + 4
}
