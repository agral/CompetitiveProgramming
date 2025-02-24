#include <climits>
#include <iostream>
#include <vector>

using Tree = std::vector<std::vector<int>>;
using Edge = std::vector<int>;

class Solution {
public:
  int mostProfitablePath(Tree &edges, int bob, std::vector<int> &amount) {
    const int SZ = amount.size();
    Tree tree(SZ);
    std::vector<int> parent(SZ);
    std::vector<int> aliceDist(SZ, -1);

    // Construct the actual tree from a list of edges:
    for (const Edge &edge : edges) {
      const int u = edge[0];
      const int v = edge[1];
      tree[u].push_back(v);
      tree[v].push_back(u);
    }

    dfs(tree, 0, -1, 0, parent, aliceDist);

    // Trace the path of Bob from his node to node 0. For each node:
    //    - if Bob reaches the node before Alice, set its cost to 0.
    //    - if both reach this node on same turn, set the amount to half the
    //    original.
    for (int u = bob, bobDist = 0; u != 0; bobDist++, u = parent[u]) {
      if (bobDist < aliceDist[u]) {
        amount[u] = 0;
      } else if (bobDist == aliceDist[u]) {
        amount[u] /= 2;
      }
    }
    return getBill(tree, 0, -1, amount);
  }

private:
  void dfs(Tree &tree, int u, int prev, int d, Edge &parent, Edge &dist) {
    // Fills parent and dist vectors via a dfs traversal, implemented
    // recursively.
    parent[u] = prev;
    dist[u] = d;
    for (int v : tree[u]) {
      if (dist[v] == -1) {
        dfs(tree, v, u, d + 1, parent, dist);
      }
    }
  }

  int getBill(Tree &tree, int u, int prev, std::vector<int> &amount) {
    if (tree[u].size() == 1 && tree[u][0] == prev) {
      // it is a leaf node.
      return amount[u];
    }
    int maxPath = INT_MIN;
    for (int v : tree[u]) {
      if (v != prev) {
        maxPath = std::max(maxPath, getBill(tree, v, u, amount));
      }
    }
    return amount[u] + maxPath;
  }
};

int main() {
  struct Testcase {
    Tree edges;
    int bob;
    std::vector<int> amount;
    int expected;
  };
  Testcase testcases[] = {
      {Tree{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3,
       std::vector<int>{-2, 4, 2, -4, 6}, 6},
      {Tree{{0, 1}}, 1, std::vector<int>{-7280, 2350}, -7280},
  };
  Solution s{};
  int numGood = 0, numBad = 0;
  for (Testcase &tc : testcases) {
    int actual = s.mostProfitablePath(tc.edges, tc.bob, tc.amount);
    if (actual != tc.expected) {
      std::cout << "Testcase failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
      ++numBad;
    } else {
      ++numGood;
    }
  }
  std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " " << numGood << "/"
            << (numBad + numGood) << " testcases passed successfully.\n";
}
