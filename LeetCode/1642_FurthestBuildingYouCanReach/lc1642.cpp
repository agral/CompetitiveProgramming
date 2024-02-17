#include <cassert>
#include <queue>
#include <vector>

class Solution {
public:
  int furthestBuilding(std::vector<int>& heights, int bricks, int ladders) {
    std::priority_queue<int, std::vector<int>, std::greater<int>> min_heap;
    for (int h = 0; h < heights.size() - 1; ++h) {
      const int delta = heights[h + 1] - heights[h];
      if (delta <= 0) {
        // jumping down is free, requiring no ladders nor bricks.
        continue;
      }
      min_heap.push(delta);
      // cheat: initially use ladders, but then as we run out of ladders,
      // substitute cheapest ladder with bricks. This is greedy and fast.
      if (min_heap.size() > ladders) {
        bricks -= min_heap.top();
        min_heap.pop();
      }
      if (bricks < 0) {
        return h;
      }
    }
    return heights.size() - 1;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> heights = {4, 2, 7, 6, 9, 14, 12};
    assert(s.furthestBuilding(heights, /*bricks=*/5, /*ladders=*/1) == 4);
  }
  {
    std::vector<int> heights = {4, 12, 2, 7, 3, 18, 20, 3, 19};
    assert(s.furthestBuilding(heights, /*bricks=*/10, /*ladders=*/2) == 7);
  }
  {
    std::vector<int> heights = {14, 3, 19, 3};
    assert(s.furthestBuilding(heights, /*bricks=*/17, /*ladders=*/0) == 3);
  }
}
