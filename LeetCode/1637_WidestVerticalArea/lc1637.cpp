#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int maxWidthOfVerticalArea(std::vector<std::vector<int>>& points) {
    const int SIZE = points.size();
    std::vector<int> x_coords;
    x_coords.resize(SIZE);
    for (int i = 0; i < SIZE; ++i) {
      x_coords[i] = points[i][0];
      // ignore the y coord completely, it does not matter.
    }

    std::sort(x_coords.begin(), x_coords.end());

    int ans = 0;
    for (int i = 1; i < SIZE; ++i) {
      ans = std::max(ans, x_coords[i] - x_coords[i-1]);
    }
    return ans;
  }
};

int main() {
  Solution s;

  {
    std::vector<std::vector<int>> points = {{8, 7}, {9, 9}, {7, 4}, {9, 7}};
    assert(s.maxWidthOfVerticalArea(points) == 1);
  }
  {
    std::vector<std::vector<int>> points = {{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}};
    assert(s.maxWidthOfVerticalArea(points) == 3);
  }
}
