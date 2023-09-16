#include <cassert>
#include <cmath>
#include <climits>
#include <queue>
#include <vector>

#include <iostream>

struct Cell {
  int y;
  int x;
  int cost;
};

class Solution {
public:
  int minimumEffortPath(std::vector<std::vector<int>>& heights) {
    const int HEIGHT = heights.size();
    const int WIDTH = heights[0].size();
    if (HEIGHT == 1 && WIDTH == 1) {
      return 0;
    }
    const int STEPS[] = {0, -1, 0, 1, 0};
    auto compare = [](const Cell& lhs, const Cell& rhs) { return lhs.cost > rhs.cost; };

    std::priority_queue<Cell, std::vector<Cell>, decltype(compare)> minHeap(compare);
    std::vector<std::vector<int>> efforts(HEIGHT, std::vector<int>(WIDTH, INT_MAX));
    std::vector<std::vector<bool>> is_seen(HEIGHT, std::vector<bool>(WIDTH, false));

    minHeap.push({HEIGHT - 1, WIDTH - 1, 0});
    efforts[HEIGHT - 1][WIDTH - 1] = 0;

    while (!minHeap.empty()) {
      const auto [y, x, cost] = minHeap.top();
      minHeap.pop();
      if ((x == 0) && (y == 0)) {
        return cost;
      }
      is_seen[y][x] = true;
      for (int s = 0; s < 4; s++) {
        const int next_y = y + STEPS[s];
        const int next_x = x + STEPS[s+1];
        if ((next_y < 0) || (next_y == HEIGHT) || (next_x < 0) || (next_x == WIDTH)) {
          continue;
        }
        if (is_seen[next_y][next_x]) {
          continue;
        }
        const int newEffort = abs(heights[y][x] - heights[next_y][next_x]);
        const int maxEffort = std::max(efforts[y][x], newEffort);
        if (efforts[next_y][next_x] > maxEffort) {
          efforts[next_y][next_x] = maxEffort;
          minHeap.push({next_y, next_x, maxEffort});
        }
      }
    }

    throw; // this should be unreachable
  }
};

int main() {
  Solution s;
  std::vector<std::vector<int>> example1 = {{1, 2, 2}, {3, 8, 2}, {5, 3, 5}};
  assert(s.minimumEffortPath(example1) == 2);

  std::vector<std::vector<int>> example2 = {{1, 2, 3}, {3, 8, 4}, {5, 3, 5}};
  assert(s.minimumEffortPath(example2) == 1);

  std::vector<std::vector<int>> example3 = {
      {1, 2, 1, 1, 1},
      {1, 2, 1, 2, 1},
      {1, 2, 1, 2, 1},
      {1, 2, 1, 2, 1},
      {1, 1, 1, 2, 1}
  };
  assert(s.minimumEffortPath(example3) == 0);
}
