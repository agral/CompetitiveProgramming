#include <cassert>
#include <queue>
#include <set>
#include <utility>
#include <vector>

class Solution {
public:
  int shortestPathLength(std::vector<std::vector<int>>& graph) {
    int ans = 0;
    std::queue<std::vector<int>> queue;
    std::set<std::pair<int, int>> set;
    for (int i = 0; i < graph.size(); i++) {
      int mask = 1 << i;
      queue.push({0, i, mask});
      set.insert({i, mask});
    }

    while (!queue.empty()) {
      auto params = queue.front();
      queue.pop();
      if (params[2] == (1 << graph.size()) - 1) {
        ans = params[0];
        break;
      }
      for (auto neighbor: graph[params[1]]) {
        int mask = (1 << neighbor) | params[2];
        if (set.find({neighbor, mask}) == set.end()) {
          set.insert({neighbor, mask});
          queue.push({params[0] + 1, neighbor, mask});
        }
      }
    }
    return ans;
  }
};

int main() {
  Solution s;

  std::vector<std::vector<int>> example1 = {{1, 2, 3}, {0}, {0}, {0}};
  assert(s.shortestPathLength(example1) == 4);

  std::vector<std::vector<int>> example2 = {{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}};
  assert(s.shortestPathLength(example2) == 4);
}
