#include <cassert>
#include <cmath>
#include <vector>

class Solution {
public:
  int minTimeToVisitAllPoints(std::vector<std::vector<int>>& points) {
    int ans = 0;
    for (int i = 1; i < points.size(); ++i) {
      ans += std::max(std::abs(points[i][0] - points[i-1][0]), std::abs(points[i][1] - points[i-1][1]));
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> points = {{1, 1}, {3, 4}, {-1, 0}};
    assert(s.minTimeToVisitAllPoints(points) == 7);
  }
  {
    std::vector<std::vector<int>> points = {{3, 2}, {-2, 2}};
    assert(s.minTimeToVisitAllPoints(points) == 5);
  }
}
