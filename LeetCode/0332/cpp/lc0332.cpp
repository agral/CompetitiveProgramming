#include <iostream>
#include <vector>
#include <sstream>

// Runtime complexity:
// Auxiliary space complexity:
class Solution {
public:
    std::vector<std::string> findItinerary(std::vector<std::vector<std::string>>& tickets) {
        return {};
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
        std::vector<std::vector<std::string>> tickets;
        std::vector<std::string> expected;
    };
    Testcase testcases[] = {
        {
            /*tickets=*/{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
            /*expected=*/{"JFK", "MUC", "LHR", "SFO", "SJC"},
        },
        {
            /*tickets=*/{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
            /*expected=*/{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
        },
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.findItinerary(tc.tickets);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
