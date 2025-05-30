#include <iostream>
#include <vector>

// Runtime complexity: O(n)
// Auxiliary space complexity: O(n)
class Solution {
public:
    int closestMeetingNode(std::vector<int>& edges, int node1, int node2) {
        std::vector<int> dist1 = getDistances(edges, node1);
        std::vector<int> dist2 = getDistances(edges, node2);
        int minDist = 1'000'000;
        int ans = -1;

        for (int i = 0; i < edges.size(); i++) {
            if (std::min(dist1[i], dist2[i]) >= 0) {
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
    std::vector<int> getDistances(std::vector<int>& edges, int node) {
        std::vector<int> ans(edges.size(), -1);
        int dist = 0;
        while (node != -1 && ans[node] == -1) {
            ans[node] = dist;
            dist += 1;
            node = edges[node];
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
        {std::vector<int>{2, 2, 3, -1}, 0, 1, 2},
        {std::vector<int>{1, 2, -1}, 0, 2, 2},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.closestMeetingNode(tc.edges, tc.node1, tc.node2);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << numGood+numBad << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
