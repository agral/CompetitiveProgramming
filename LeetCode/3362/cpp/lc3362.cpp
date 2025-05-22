#include <algorithm>
#include <iostream>
#include <queue>
#include <vector>

// Runtime complexity:
// Auxiliary space complexity:
class Solution {
public:
    int maxRemoval(std::vector<int>& nums, std::vector<std::vector<int>>& queries) {
        int qIndex = 0;
        std::priority_queue<int> available{};
        std::priority_queue<int, std::vector<int>, std::greater<>> running{};

        std::ranges::sort(queries);

        for (int i = 0; i < nums.size(); i++) {
            while (qIndex < queries.size() && queries[qIndex][0] <= i) {
                available.push(queries[qIndex][1]);
                qIndex++;
            }
            while (!running.empty() && running.top() < i) {
                running.pop();
            }
            while (nums[i] > running.size()) {
                if (available.empty() || available.top() < i) {
                    return -1;
                }
                running.push(available.top());
                available.pop();
            }
        }
        return available.size();
    }
};

int main() {
    struct Testcase {
        std::vector<int> nums;
        std::vector<std::vector<int>> queries;
        int expected;
    };
    Testcase testcases[] = {
        {
            {2, 0, 2},
            {{0, 2}, {0, 2}, {1, 1}},
            1,
        },
        {
            {1, 1, 1, 1},
            {{1, 3}, {0, 2}, {1, 3}, {1, 2}},
            2,
        },
        {
            {1, 2, 3, 4},
            {{0, 3}},
            -1,
        }
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.maxRemoval(tc.nums, tc.queries);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood+numBad+1) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
