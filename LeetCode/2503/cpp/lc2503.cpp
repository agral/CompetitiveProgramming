#include <iostream>
#include <vector>
#include <sstream>

// Runtime complexity:
// Auxiliary space complexity:

class Solution {
 public:
  std::vector<int> maxPoints(std::vector<std::vector<int>>& grid,
                             std::vector<int>& queries) {
    std::vector<int> ans = {};
    return ans;
  }
};

template<typename T>
std::string vectorToString(std::vector<T> vec) {
    std::string separator = "";
    std::stringstream ss;
    ss << "[";
    for (const auto& elem: vec) {
        ss << separator << elem;
        separator = ", ";
    }
    ss << "]";
    return ss.str();
}

int main() {
    struct Testcase {
        std::vector<std::vector<int>> grid;
        std::vector<int> queries;
        std::vector<int> expected;
    };
    Testcase testcases[] = {
        {/*grid=*/{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, /*queries=*/{5, 6, 2}, /*expected=*/{5, 8, 1}},
        {/*grid=*/{{5, 2, 1}, {1, 1, 2}}, /*queries=*/{3}, /*expected=*/{0}},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.maxPoints(tc.grid, tc.queries);
        if (actual != tc.expected) {
            std::cout << "Testcase " << vectorToString(tc.queries)
                << " failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
