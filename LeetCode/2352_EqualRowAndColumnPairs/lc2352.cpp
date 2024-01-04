#include <cassert>
#include <vector>

class Solution {
public:
  // A brute-force solution, O(n^3) runtime. No custom hash functions needed.
  int equalPairs(std::vector<std::vector<int>>& grid) {
    const auto SIZE = grid.size();
    int ans = 0;
    for (int row = 0; row < SIZE; ++row) {
      for (int col = 0; col < SIZE; ++col) {
        int i = 0;
        while (i < SIZE) {
          if (grid[row][i] != grid[i][col]) {
            break;
          }
          ++i;
        }
        if (i == SIZE) { // entire row and col matches
          ++ans;
        }
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> grid = {{3, 2, 1}, {1, 7, 6}, {2, 7, 7}};
    assert(s.equalPairs(grid) == 1);
  }
  {
    std::vector<std::vector<int>> grid = {{3, 1, 2, 2}, {1, 4 , 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}};
    assert(s.equalPairs(grid) == 3);
  }
}
