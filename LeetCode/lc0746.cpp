#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int minCostClimbingStairs(std::vector<int>& cost) {
    cost.push_back(0);
    for (int i = cost.size() - 3; i >= 0; i--) {
      // modify in-place:
      cost[i] = std::min(cost[i+1], cost[i+2]) + cost[i];
    }
    return std::min(cost[0], cost[1]);
  }
};

int main() {
  Solution s;

  std::vector<int> example1 = {10, 15, 20};
  assert(s.minCostClimbingStairs(example1) == 15);

  std::vector<int> example2 = {1, 100, 1, 1, 1, 100, 1, 1, 100, 1};
  assert(s.minCostClimbingStairs(example2) == 6);
}
