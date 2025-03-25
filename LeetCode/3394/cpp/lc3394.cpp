#include <algorithm>
#include <iostream>
#include <vector>

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space complexity: O(n) for holding the sorted coords.
class Solution {
public:
    bool checkValidCuts(int n, std::vector<std::vector<int>>& rectangles) {
        std::vector<std::pair<int, int>> xCoords;
        std::vector<std::pair<int, int>> yCoords;

        for (std::vector<int>& r: rectangles) {
            xCoords.emplace_back(r[0], r[2]);
            yCoords.emplace_back(r[1], r[3]);
        }
        return countDisjoined(xCoords) >= 3 or countDisjoined(yCoords) >= 3;
    }

private:
    int countDisjoined(std::vector<std::pair<int, int>>& intervals) {
        int ans = 0;
        int lastEnd = 0;
        std::sort(intervals.begin(), intervals.end());

        for (const auto& [start, end]: intervals) {
            if (start < lastEnd) {
                lastEnd = std::max(lastEnd, end);
            } else {
                ++ans;
                lastEnd = end;
            }
        }

        return ans;
    }
};

int main() {
    struct Testcase {
        int n;
        std::vector<std::vector<int>> rectangles;
        bool expected;
    };
    Testcase testcases[] = {
        {5, {{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}, true},
        {4, {{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}, true},
        {4, {{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}, false},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.checkValidCuts(tc.n, tc.rectangles);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << numGood + numBad << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
