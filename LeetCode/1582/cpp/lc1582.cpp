#include <cassert>
#include <vector>

class Solution {
public:
  int numSpecial(std::vector<std::vector<int>>& mat) {
    const auto HEIGHT = mat.size();
    const auto WIDTH = mat[0].size();
    std::vector<int> rowSum(HEIGHT, 0); // holds sum of each row, column-by-column (row stays the same)
    std::vector<int> colSum(WIDTH, 0);
    for (int h = 0; h < HEIGHT; h++) {
      for (int w = 0; w < WIDTH; w++) {
        if (mat[h][w] == 1) {
          ++rowSum[h];
          ++colSum[w];
        }
      }
    }

    int ans = 0;
    for (int h = 0; h < HEIGHT; h++) {
      for (int w = 0; w < WIDTH; w++) {
        if ((mat[h][w] == 1) && (rowSum[h] == 1) && (colSum[w] == 1)) {
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
