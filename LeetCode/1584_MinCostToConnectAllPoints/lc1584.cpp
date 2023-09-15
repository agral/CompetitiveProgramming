#include <algorithm>
#include <cassert>
#include <vector>
#include <set>
#include <map>

struct EdgeCost {
  int nodeA;
  int nodeB;
  int cost;
};

class Solution {
public:
  int minCostConnectPoints(std::vector<std::vector<int>>& points) {
    // Calculate distances between all possible pairs of edges:
    std::vector<EdgeCost> edgeCosts;
    for (int i = 0; i < points.size() - 1; i++) {
      for (int j = i + 1; j < points.size(); j++) {
        edgeCosts.push_back({i, j, abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])});
      }
    }
    int ans = 0;

    // Sort edgeCosts in ascending order with regard to distance value:
    std::sort(edgeCosts.begin(), edgeCosts.end(), [](const EdgeCost& lhs, const EdgeCost& rhs){ return lhs.cost < rhs.cost; });

    // Proceed with Kruskal's algorithm. Initially there are num_points separate tree stubs.
    std::vector<int> which_tree(points.size()); // maps a point to a tree it belongs to.
    for (int i = 0; i < points.size(); i++) {
      which_tree[i] = i;
    }

    int num_joins = 0; // we expect to perform exactly num_points-1 join operations to create a minimum spanning tree.
    int currentEdgeIdx = 0;
    while (num_joins < points.size() - 1) {
      // consider the next edge.
      int nodeA = edgeCosts[currentEdgeIdx].nodeA;
      int nodeB = edgeCosts[currentEdgeIdx].nodeB;

      // If its nodes belong to same tree, skip this edge - don't create cycles.
      if (which_tree[nodeA] != which_tree[nodeB]) {
        // Otherwise (i.e. nodes belong to distinct trees), join both trees:
        int search_for = which_tree[nodeB];
        int sub_with = which_tree[nodeA];
        for (int i = 0; i < points.size(); i++) {
          if (which_tree[i] == search_for) {
            which_tree[i] = sub_with;
          }
        }

        // and accumulate this edge cost:
        ans += edgeCosts[currentEdgeIdx].cost;
        num_joins += 1;
      }
      currentEdgeIdx++;
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<std::vector<int>> example1 = {{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}};
  assert(s.minCostConnectPoints(example1) == 20);

  std::vector<std::vector<int>> example2 = {{3, 12}, {-2, 5}, {-4, 1}};
  assert(s.minCostConnectPoints(example2) == 18);
}
