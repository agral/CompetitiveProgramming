#include <cassert>
#include <utility>
#include <vector>

class Solution {
public:
  int knightDialer(int n) {
    const int MOD = 1'000'000'007;
    std::vector<std::vector<int>> dp(4, std::vector<int>(3, 1));
    dp[3][0] = 0;
    dp[3][2] = 0;

    for (int k = 0; k < n - 1; k++) {
      std::vector<std::vector<int>> next_dp(4, std::vector<int>(3));
      for (int row = 0; row < 4; row++) {
        for (int column = 0; column < 3; column++) {
          if (!isValidKey(row, column)) {
            continue;
          }
          for (const auto& [dy, dx]: jumps) {
            int y = column + dy;
            int x = row + dx;
            if (y < 0 || y > 3 || x < 0 || x > 2) {
              continue;
            }
            if (!isValidKey(y, x)) {
              continue;
            }
            next_dp[row][column] = (dp[y][x] + next_dp[row][column]) % MOD;
          }
        }
      }
      dp = std::move(next_dp);
    }

    int ans = 0;
    for (const auto& row: dp) {
      for (const auto& num: row) {
        ans = (ans + num) % MOD;
      }
    }
    return ans;
  }

private:
  const std::vector<std::pair<int, int>> jumps = {
    {-2, -1}, {-2, 1}, // "up"
    {-1, 2},  {1, 2},  // "right"
    {2, 1},   {2, -1}, // "down"
    {1, -2},  {-1, -2} // "left"
  };

  // A knight can not jump to non-numeric cells, i.e. '#' and '*'.
  [[nodiscard]] inline bool isValidKey(int row, int column) {
    return (row != 3) || (column == 1);
  }
};

int main() {
  Solution s;
  assert(s.knightDialer(1) == 10);
  assert(s.knightDialer(2) == 20);
  assert(s.knightDialer(3131) == 136'006'598);
}
