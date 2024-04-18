#include <cassert>
#include <vector>

class Solution {
public:
  int islandPerimeter(std::vector<std::vector<int>>& grid) {
    const int HEIGHT = grid.size();
    const int WIDTH = grid[0].size();
    int perimeter = 0;
    for (int h = 0; h < HEIGHT; ++h) {
      for (int w = 0; w < WIDTH; ++w) {
        if (grid[h][w] == LAND) {
          if (h == 0 || grid[h-1][w] == WATER) {
            ++perimeter;
          }
          if (h == HEIGHT - 1 || grid[h+1][w] == WATER) {
            ++perimeter;
          }
          if (w == 0 || grid[h][w-1] == WATER) {
            ++perimeter;
          }
          if (w == WIDTH - 1 || grid[h][w+1] == WATER) {
            ++perimeter;
          }
        }
      }
    }
    return perimeter;
  }

private:
  const int LAND = 1;
  const int WATER = 0;
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> grid = {
      {0, 1, 0, 0},
      {1, 1, 1, 0},
      {0, 1, 0, 0},
      {1, 1, 0, 0},
    };
    assert(s.islandPerimeter(grid) == 16);
  }
  {
    std::vector<std::vector<int>> grid = {
      {1},
    };
    assert(s.islandPerimeter(grid) == 4);
  }
  {
    std::vector<std::vector<int>> grid = {
      {1, 0},
    };
    assert(s.islandPerimeter(grid) == 4);
  }
}
