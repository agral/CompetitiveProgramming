#include <cassert>
#include <vector>

class Solution {
public:
  int numSpecial(std::vector<std::vector<int>>& mat) {
    const auto HEIGHT = mat.size();
    const auto WIDTH = mat[0].size();
    int sole_ones_in_rows = 0;
    int sole_ones_in_columns = 0;
    for (int h = 0; h < HEIGHT; h++) {
      int count = 0;
      for (int w = 0; w < WIDTH; w++) {
        if (mat[h][w] == 1) {
          ++count;
        }
      }
      if (count == 1) {
        ++sole_ones_in_rows;
      }
    }
    for (int w = 0; w < WIDTH; w++) {
      int count = 0;
      for (int h = 0; h < HEIGHT; h++) {
        if (mat[h][w] == 1) {
          ++count;
        }
      }
      if (count == 1) {
        ++sole_ones_in_columns;
      }
    }

    return std::min(sole_ones_in_columns, sole_ones_in_rows);
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> mat = {{1, 0, 0}, {0, 0, 1}, {1, 0, 0}};
    assert(s.numSpecial(mat) == 1);
  }
  {
    std::vector<std::vector<int>> mat = {{1, 0, 0}, {0, 1, 0}, {0, 0, 1}};
    assert(s.numSpecial(mat) == 3);
  }
  {
    std::vector<std::vector<int>> mat = {
      {0,0,0,0,0,1,0,0},
      {0,0,0,0,1,0,0,1},
      {0,0,0,0,1,0,0,0},
      {1,0,0,0,1,0,0,0},
      {0,0,1,1,0,0,0,0}
    };
    assert(s.numSpecial(mat) == 1);
  }
}
