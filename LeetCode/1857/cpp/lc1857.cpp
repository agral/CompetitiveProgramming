#include <deque>
#include <iostream>
#include <vector>

class Solution {
public:
    int largestPathValue(std::string colors, std::vector<std::vector<int>>& edges) {
        const int SZ = colors.size();
        const int NUM_LETTERS = 'z' - 'a' + 1;
        int ans = {};
        int numProcessed = 0;
        std::vector<std::vector<int>> graph(SZ);
        std::vector<int> indegrees(SZ);
        std::deque<int> q;
        std::vector<std::vector<int>> count(SZ, std::vector<int>(NUM_LETTERS));

        for (const auto& edge: edges) {
            graph[edge[0]].push_back(edge[1]);
            indegrees[edge[1]]++;
        }

        // topological sorting:
        for (int i = 0; i < SZ; i++) {
            if (indegrees[i] == 0) {
                q.push_back(i);
            }
        }
        while (!q.empty()) {
            int f = q.front();
            q.pop_front();
            ++numProcessed;
            int s = colors[f] - 'a';
            count[f][s] += 1;
            ans = std::max(ans, count[f][s]);

            for (int in: graph[f]) {
                for (int offset = 0; offset < NUM_LETTERS; offset++) {
                    count[in][offset] = std::max(count[in][offset], count[f][offset]);
                }
                indegrees[in] -= 1;
                if (indegrees[in] == 0) {
                    q.push_back(in);
                }
            }
        }
        if (numProcessed != SZ) {
            return -1;
        }
        return ans;
    }
};

int main() {
    struct Testcase {
        std::string colors;
        std::vector<std::vector<int>> edges;
        int expected;
    };
    Testcase testcases[] = {
        {"abaca", {{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
        {"a", {{0, 0}}, -1}, // there is a cycle from node 0 to node 0.
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.largestPathValue(tc.colors, tc.edges);
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.colors << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
