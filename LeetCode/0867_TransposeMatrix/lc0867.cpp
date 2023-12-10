#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<std::vector<int>> transpose(std::vector<std::vector<int>>& matrix) {
    const int HEIGHT = matrix.size();
    const int WIDTH = matrix[0].size();
    std::vector<std::vector<int>> transposed(WIDTH, std::vector<int>(HEIGHT, 0));
    for (int h = 0; h < HEIGHT; h++) {
      for (int w = 0; w < WIDTH; w++) {
        transposed[w][h] = matrix[h][w];
      }
    }
    return transposed;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> matrix = {{1, 2, 3}, {4, 5 ,6}, {7, 8, 9}};
    std::vector<std::vector<int>> transposed = {{1, 4, 7}, {2, 5, 8}, {3, 6, 9}};
    assert(s.transpose(matrix) == transposed);
  }
  {
    std::vector<std::vector<int>> matrix = {{1, 2, 3}, {4, 5 ,6}};
    std::vector<std::vector<int>> transposed = {{1, 4}, {2, 5}, {3, 6}};
    assert(s.transpose(matrix) == transposed);
  }
}
