#include <cassert>
#include <numeric>
#include <queue>
#include <vector>

// Time complexity: O(q(q+n)), as one needs to re-check at most q+n ways each time a way from query q is added.
// Space complexity: O(q+n), as it would be n-1 ways, but each query adds another way.
class Solution {
public:
    std::vector<int> shortestDistanceAfterQueries(int n, std::vector<std::vector<int>>& queries) {
        std::vector<std::vector<int>> graph(n);
        std::vector<int> distances(n);
        std::iota(distances.begin(), distances.end(), 0);
        std::vector<int> ans;

        for (int i = 0; i < n-1; ++i) {
            // there is a road from city 0 to city 1, from city 1 to city 2, ..., and from city n-2 to city n-1.
            graph[i].push_back(i+1);
        }

        for (const std::vector<int>& query: queries) {
            int src = query[0];
            int dst = query[1];
            graph[src].push_back(dst);
            if (distances[src] + 1 < distances[dst]) {
                distances[dst] = distances[src] + 1;
                bfs(graph, dst, distances);
            }
            ans.push_back(distances[n-1]);
        }
        return ans;
    }

private:
    void bfs(const std::vector<std::vector<int>>& graph, int start_idx, std::vector<int>& distances) {
        std::queue<int> q{{start_idx}};
        while (!q.empty()) {
            int src = q.front();
            q.pop();
            for (int dst: graph[src]) {
                if (distances[src] + 1 < distances[dst]) {
                    distances[dst] = distances[src] + 1;
                    q.push(dst);
                }
            }
        }
    }
};

int main() {
    Solution s;
    {
        std::vector<std::vector<int>> queries = {{2, 4}, {0, 2}, {0, 4}};
        std::vector<int> expected = {3, 2, 1};
        assert(s.shortestDistanceAfterQueries(5, queries) == expected);
    }
    {
        std::vector<std::vector<int>> queries = {{0, 3}, {0, 2}};
        std::vector<int> expected = {1, 1};
        assert(s.shortestDistanceAfterQueries(4, queries) == expected);
    }
}
