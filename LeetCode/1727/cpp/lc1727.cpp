#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int largestSubmatrix(std::vector<std::vector<int>>& matrix) {
    const int N = matrix[0].size();
    // Since we must rearrange entire columns, a structure is needed to hold number of 1's
    // in a given column. Note that the matrix is not guaranteed to be square (i.e. M can != N).
    std::vector<int> histogram(N);
    int ans = 0;
    for (const std::vector<int>& row: matrix) {
      for (int i = 0; i < N; i++) {
        histogram[i] = (row[i] == 0) ? 0 : histogram[i] + 1;
      }

      std::vector<int> sorted(histogram);
      std::sort(sorted.begin(), sorted.end());
      for (int i = 0; i < N; i++) {
        ans = std::max(ans, sorted[i] * (N - i));
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> matrix = {{0, 0, 1}, {1, 1, 1}, {1, 0, 1}};
    assert(s.largestSubmatrix(matrix) == 4);
  }
  {
    std::vector<std::vector<int>> matrix = {{1, 0, 1, 0, 1}};
    assert(s.largestSubmatrix(matrix) == 3);
  }
  {
    std::vector<std::vector<int>> matrix = {{1, 1, 0}, {1, 0, 1}};
    assert(s.largestSubmatrix(matrix) == 2);
  }
}
