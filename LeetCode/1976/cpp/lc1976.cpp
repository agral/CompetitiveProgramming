#include <climits>
#include <iostream>
#include <queue>
#include <vector>

class Solution {
public:
    int countPaths(int n, std::vector<std::vector<int>>& roads) {
        m_graph = std::vector<std::vector<std::pair<int, int>>>(n);
        for (std::vector<int>& road: roads) {
            // All roads are bi-directional, encode both directions in the graph:
            m_graph[road[0]].emplace_back(road[1], road[2]);
            m_graph[road[1]].emplace_back(road[0], road[2]);
        }
        return dijkstra(0, n-1);
    }

private:
    std::vector<std::vector<std::pair<int, int>>> m_graph;
    int dijkstra(int start, int end) {
        const int MOD = 1'000'000'007;
        std::vector<long> ways(m_graph.size());
        std::vector<long> dist(m_graph.size(), LONG_MAX);
        ways[start] = 1;
        dist[start] = 0;

        using LI = std::pair<long, int>; // pair(distance, intersection)
        std::priority_queue<LI, std::vector<LI>, std::greater<>> minHeap;
        minHeap.emplace(dist[start], start);
        while (!minHeap.empty()) {
            const auto [d, i] = minHeap.top();
            minHeap.pop();
            if (d <= dist[i]) {
                for (const auto& [destination, time]: m_graph[i]) {
                    if (d + time < dist[destination]) {
                        dist[destination] = d + time;
                        ways[destination] = ways[i];
                        minHeap.emplace(dist[destination], destination);
                    } else if (d + time == dist[destination]) {
                        ways[destination] += ways[i];
                        ways[destination] %= MOD;
                    }
                }
            }
        }

        return ways[end];
    }
};

int main() {
    struct Testcase {
        int n;
        std::vector<std::vector<int>> roads;
        int expected;
    };
    Testcase testcases[] = {
        {7, {{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}, 4},
        {2, {{1, 0, 10}}, 1},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.countPaths(tc.n, tc.roads);
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.n << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
