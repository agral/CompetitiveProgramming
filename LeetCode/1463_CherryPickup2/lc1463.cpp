#include <cassert>
#include <vector>

class Solution {
public:
  int cherryPickup(std::vector<std::vector<int>>& grid) {
    // bottom-up DP approach. dp[y][x1][x2] denotes maximum cherries collected when
    // robot #1 stands on y, x1 and robot #2 stands on y, x2.
    const int HEIGHT = grid.size();
    const int WIDTH = grid[0].size();
    std::vector<std::vector<std::vector<int>>> dp(HEIGHT + 1,
                std::vector<std::vector<int>>(WIDTH, std::vector<int>(WIDTH)));
    for (int y = HEIGHT - 1; y >= 0; --y) {
      for (int x1 = 0; x1 < WIDTH; ++x1) {
        for (int x2 = 0; x2 < WIDTH; ++x2) {
          const int score_both_here = grid[y][x1] + (grid[y][x2] * (x1 != x2));

          // the following nested for loops generate at most 9 distinct set of moves for both robots;
          // at most three for both first and second one.
          for (int move1 = std::max(0, x1 - 1); move1 < std::min(WIDTH, x1 + 2); ++move1) {
            for (int move2 = std::max(0, x2 - 1); move2 < std::min(WIDTH, x2 + 2); ++move2) {
              dp[y][x1][x2] = std::max(dp[y][x1][x2], score_both_here + dp[y + 1][move1][move2]);
            }
          }
        }
      }
    }

    return dp[0][0][WIDTH - 1];
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> grid = {
      {3, 1, 1},
      {2, 5, 1},
      {1, 5, 5},
      {2, 1, 1}};
    assert(s.cherryPickup(grid) == 24);
  }
  {
    std::vector<std::vector<int>> grid = {
      {1, 0, 0, 0, 0, 0, 1},
      {2, 0, 0, 0, 0, 3, 0},
      {2, 0, 9, 0, 0, 0, 0},
      {0, 3, 0, 5, 4, 0, 0},
      {1, 0, 2, 3, 0, 0, 6}};
    assert(s.cherryPickup(grid) == 28);
  }

};
