#include <iostream>
#include <queue>
#include <vector>

const int BIG_DIST = 10'000'000;

// Runtime complexity: O(n)
// Auxiliary space complexity: O(n). This is because each vertex has at most 1 edge.
class Solution {
public:
    int closestMeetingNode(std::vector<int>& edges, int node1, int node2) {
        const int NUM_EDGES = edges.size();

        // Make a graph that uses adjacency lists:
        std::vector<std::vector<int>> graph(NUM_EDGES);
        for (int i = 0; i < NUM_EDGES; i++) {
            if (edges[i] != -1) {
                graph[i].push_back(edges[i]);
            }
        }

        std::vector<int> dist1 = getDistances(graph, node1);
        std::vector<int> dist2 = getDistances(graph, node2);
        int minDist = BIG_DIST;
        int ans = -1;

        for (int i = 0; i < NUM_EDGES; i++) {
            if (dist1[i] >= 0 && dist2[i] >= 0) {
                int score = std::max(dist1[i], dist2[i]);
                if (score < minDist) {
                    minDist = score;
                    ans = i;
                }
            }
        }
        return ans;
    }

private:
    std::vector<int> getDistances(std::vector<std::vector<int>>& graph, int node) {
        std::vector<int> ans(graph.size(), BIG_DIST);
        ans[node] = 0;
        std::queue<int> q{{node}};
        while (!q.empty()) {
            int u = q.front();
            q.pop();
            for (int v: graph[u]) {
                if (ans[v] == BIG_DIST) {
                    ans[v] = ans[u] + 1;
                    q.push(v);
                }
            }
        }
        return ans;
    }
};

int main() {
    struct Testcase {
        std::vector<int> edges;
        int node1;
        int node2;
        int expected;
    };
    Testcase testcases[] = {
        {{2, 2, 3, -1}, 0, 1, 2},
        {{1, 2, -1}, 0, 2, 2},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.closestMeetingNode(tc.edges, tc.node1, tc.node2);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood + numBad) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
