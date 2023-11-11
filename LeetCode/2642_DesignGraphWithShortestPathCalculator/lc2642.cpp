#include <cassert>
#include <vector>
#include <queue>
#include <utility>
#include <limits>

using PII = std::pair<int, int>;

class Graph {
public:
  Graph(int n, std::vector<std::vector<int>>& edges) {
    m_graph.clear();
    m_graph.resize(n);
    for (const auto& edge: edges) {
      addEdge(edge);
    }
  }

  void addEdge(std::vector<int> edge) {
      m_graph[edge[0]].emplace_back(edge[1], edge[2]);
  }

  int shortestPath(int node1, int node2) {
    // Perform Dijkstra's algorithm to find min-distance to target node2.
    std::priority_queue<PII, std::vector<PII>, std::greater<>> pq;
    std::vector<int> distances(m_graph.size(), std::numeric_limits<int>::max());
    distances[node1] = 0;
    pq.emplace(distances[node1], node1);
    while (!pq.empty()) {
      const auto [dist, vertex] = pq.top();
      pq.pop();
      if (vertex == node2) {
        return dist;
      }
      for (const auto& [v, w]: m_graph[vertex]) {
        if (dist + w < distances[v]) {
          distances[v] = dist + w;
          pq.emplace(distances[v], v);
        }
      }
    }
    return -1;
  }
private:
  std::vector<std::vector<std::pair<int, int>>> m_graph;
};

int main() {
  std::vector<std::vector<int>> edges1 = {{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}};
  Graph example1{4, edges1};
  assert(example1.shortestPath(3, 2) == 6);
  assert(example1.shortestPath(0, 3) == -1);
  example1.addEdge({1, 3, 4});
  assert(example1.shortestPath(0, 3) == 6);
}
