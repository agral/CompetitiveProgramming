#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<std::vector<int>> onesMinusZeros(std::vector<std::vector<int>>& grid) {
    const int HEIGHT = grid.size();
    const int WIDTH = grid[0].size();
    std::vector<int> hScore(HEIGHT, 0);
    std::vector<int> wScore(WIDTH, 0);
    for (int h = 0; h < HEIGHT; ++h) {
      for (int w = 0; w < WIDTH; ++w) {
        hScore[h] += (grid[h][w] == 1 ? 1 : -1);
        wScore[w] += (grid[h][w] == 1 ? 1 : -1);
      }
    }
    std::vector<std::vector<int>> ans(HEIGHT, std::vector<int>(WIDTH, 0));
    for (int h = 0; h < HEIGHT; ++h) {
      for (int w = 0; w < WIDTH; ++w) {
        ans[h][w] = hScore[h] + wScore[w];
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> grid = {{0, 1, 1}, {1, 0, 1}, {0, 0, 1}};
    std::vector<std::vector<int>> expected = {{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}};
    assert(s.onesMinusZeros(grid) == expected);
  }
  {
    std::vector<std::vector<int>> grid = {{1, 1, 1}, {1, 1, 1}};
    std::vector<std::vector<int>> expected = {{5, 5, 5}, {5, 5, 5}};
    assert(s.onesMinusZeros(grid) == expected);
  }
}
