#include <cassert>
#include <vector>

class Solution {
public:
  int maximalNetworkRank(int n, std::vector<std::vector<int>>& roads) {
    std::vector<std::vector<bool>> edges;
    edges.resize(n);
    for (auto& row: edges) {
      row.resize(n, false);
    }
    std::vector<int> numEdges;
    numEdges.resize(n, 0);

    // Read in and process the edges (roads connecting cities):
    for (const auto& road: roads) {
      numEdges[road[0]] += 1;
      numEdges[road[1]] += 1;
      edges[road[0]][road[1]] = true;
      edges[road[1]][road[0]] = true;
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
      for (int j = i + 1; j < n; j++) {
        ans = std::max(ans, numEdges[i] + numEdges[j] + (edges[i][j] ? -1 : 0));
      }
    }

    return ans;
  };
};

int main() {
  Solution s{};
  std::vector<std::vector<int>> testcase1 = {{0, 1}, {0, 3}, {1, 2}, {1, 3}};
  assert(s.maximalNetworkRank(4, testcase1) == 4);

  std::vector<std::vector<int>> testcase2 = {{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}};
  assert(s.maximalNetworkRank(5, testcase2) == 5);

  std::vector<std::vector<int>> testcase3 = {{0,1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}};
  assert(s.maximalNetworkRank(8, testcase3) == 5);
};
