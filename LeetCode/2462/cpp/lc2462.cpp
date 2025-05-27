#include <iostream>
#include <queue>
#include <vector>

// Runtime complexity: O(nlogn)
// Auxiliary space complexity: O(n)
// where n is the length of `costs`.
class Solution {
public:
    long long totalCost(std::vector<int>& costs, int k, int candidates) {
        long long ans = {};
        int i = 0;
        int j = costs.size() - 1;
        std::priority_queue<int, std::vector<int>, std::greater<>> minHeapLeft{};
        std::priority_queue<int, std::vector<int>, std::greater<>> minHeapRight{};

        for (int numHired = 0; numHired < k; numHired++) {
            while ((minHeapLeft.size() < candidates) && (i <= j)) {
                minHeapLeft.push(costs[i]);
                i += 1;
            }
            while ((minHeapRight.size() < candidates) && (j >= i)) {
                minHeapRight.push(costs[j]);
                j -= 1;
            }

            if (minHeapRight.empty()) {
                ans += minHeapLeft.top();
                minHeapLeft.pop();
            } else if (minHeapLeft.empty()) {
                ans += minHeapRight.top();
                minHeapRight.pop();
            } else if (minHeapLeft.top() <= minHeapRight.top()) {
                ans += minHeapLeft.top();
                minHeapLeft.pop();
            } else {
                ans += minHeapRight.top();
                minHeapRight.pop();
            }
        }
        return ans;
    }
};

int main() {
    struct Testcase {
        std::vector<int> costs;
        int k;
        int candidates;
        long long expected;
    };
    Testcase testcases[] = {
        {{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4, 11},
        {{1, 2, 4, 1}, 3, 3, 4},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.totalCost(tc.costs, tc.k, tc.candidates);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << actual << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
