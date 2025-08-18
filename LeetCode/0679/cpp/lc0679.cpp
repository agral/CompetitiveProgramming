#include <iostream>
#include <vector>

// Runtime complexity: O(2^4) == O(1)
// Auxiliary space complexity: O(2^4) == O(1)
class Solution {
public:
    bool judgePoint24(std::vector<int>& cards) {
        std::vector<double> values;
        for (const int card: cards) {
            values.push_back(card);
        }
        return dfs(values);
    }

private:
    bool dfs(std::vector<double>& values) {
        if (values.size() == 1) {
            return std::abs(values[0] - 24) < 0.001;
        }

        for (int i = 0; i < values.size(); i++) {
            for (int j = 0; j < i; j++) {
                for (const double v: getValues(values[i], values[j])) {
                    std::vector<double> nextValues{v};
                    for (int k = 0; k < values.size(); k++) {
                        if (k != i && k != j) {
                            nextValues.push_back(values[k]);
                        }
                    }
                    if (dfs(nextValues)) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    std::vector<double> getValues(double lhs, double rhs) {
        return std::vector<double>{lhs+rhs, lhs-rhs, rhs-lhs, lhs*rhs, lhs/rhs, rhs/lhs};
    }
};

int main() {
    struct Testcase {
        std::vector<int> cards;
        bool expected;
    };
    Testcase testcases[] = {
        {{4, 1, 8, 7}, true}, // true, (8-4)*(7-1)
        {{1, 2, 1, 2}, false},
        {{1, 1, 1, 1}, false},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.judgePoint24(tc.cards);
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
