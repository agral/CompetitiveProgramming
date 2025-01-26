#include <cassert>
#include <iostream>
#include <vector>
#include <deque>

enum class State { PRISTINE, VISITING, PROCESSED };

class Solution {
public:
    int maximumInvitations(std::vector<int>& favorite) {
        const int SZ = favorite.size();
        std::vector<std::vector<int>> graph(SZ);
        std::vector<int> inDegree(SZ);
        std::deque<int> q;
        std::vector<int> maxChainLength(SZ, 1);
        int sumChainLength = 0;

        // Construct the graph and calculate each node's in-degree:
        for (int i = 0; i < SZ; ++i) {
            graph[i].push_back(favorite[i]);
            ++inDegree[favorite[i]];
        }

        // Topological sort:
        for (int i = 0; i < SZ; ++i) {
            if (inDegree[i] == 0) {
                q.push_back(i);
            }
        }
        while (!q.empty()) {
            int u = q.front();
            q.pop_front();
            for (const int v: graph[u]) {
                if (--inDegree[v] == 0) {
                    q.push_back(v);
                }
                maxChainLength[v] = std::max(maxChainLength[v], maxChainLength[u] + 1);
            }
        }
        // Handle cycles of length 2:
        for (int i = 0; i < SZ; ++i) {
            if (favorite[favorite[i]] == i) {
                sumChainLength += maxChainLength[i] + maxChainLength[favorite[i]];
            }
        }

        int maxCycleLength = 0;
        std::vector<bool> seen(SZ, false);
        std::vector<int> parent(SZ, -1);
        std::vector<State> state(SZ);

        for (int i = 0; i < SZ; ++i) {
            if (!seen[i]) {
                findCycle(graph, i, seen, parent, state, maxCycleLength);
            }
        }

        return std::max(sumChainLength / 2, maxCycleLength);
    }

private:
    void findCycle(const std::vector<std::vector<int>>& graph, int u,
                   std::vector<bool>& seen, std::vector<int>& parent,
                   std::vector<State>& state, int& maxCycleLength) {
        state[u] = State::VISITING;
        seen[u] = true;
        for (int v: graph[u]) {
            if (!seen[v]) {
                parent[v] = u;
                findCycle(graph, v, seen, parent, state, maxCycleLength);
            } else if (state[v] == State::VISITING) {
                // vertex v is part of a cycle; find its length:
                int vertex = u;
                int length = 1;
                while (vertex != v) {
                    vertex = parent[vertex];
                    length += 1;
                }
                maxCycleLength = std::max(maxCycleLength, length);
            }
        }

        state[u] = State::PROCESSED;
    }
};

int main() {
    Solution s{};

    struct Testcase {
        std::vector<int> favorite;
        int expected;
    };

    std::vector<Testcase> testcases = {
        {{2, 2, 1, 2}, 3,},
        {{1, 2, 0}, 3},
        {{3, 0, 1, 4, 1}, 4},
    };

    int successes = 0;
    for (int i = 0; i < testcases.size(); ++i) {
        int actual = s.maximumInvitations(testcases[i].favorite);
        if (actual != testcases[i].expected) {
            std::cout << "Testcase #" << i+1 << " failed. Want " << testcases[i].expected
                << ", got " << actual << "\n";
        } else {
            ++successes;
        }
    }
    std::cout << successes << " / " << testcases.size() << " testcases successfully passed.\n";
}
