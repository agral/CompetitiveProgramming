#include <iostream>
#include <unordered_map>

// Runtime complexity: O(d)
// Auxiliary space complexity: O(d) (from the map)
// where d is the number of digits in the numerator; d ~= log10(numerator).
// Subjective level: medium.
class Solution {
public:
    std::string fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) {
            return "0";
        }
        // It is guaranteed that denominator != 0.

        std::string ans = {};
        // Handle the negative numbers' representation.
        // But also handle the case of negative numerator and negative denominator.
        // There is a minus sign if one of {numerator, denominator} is negative,
        // but not when none or both are negative.
        if ((numerator < 0) ^ (denominator < 0)) {
            ans += "-";
        }

        // use longs; as long division requires a *=10 operation,
        // which might overflow an int.
        long ln = std::labs(numerator);
        long ld = std::labs(denominator);
        ans += std::to_string(ln / ld);
        if (ln % ld == 0L) {
            return ans;
        }

        // perform the long division, keep track when it starts looping.
        ans += ".";
        std::unordered_map<long, long> seen;
        for (long result = ln % ld; result != 0; result %= ld) {
            const auto it = seen.find(result);
            if (it != seen.cend()) {
                ans.insert(it->second, 1, '(');
                ans += ')';
                break;
            }
            seen[result] = ans.size();
            //std::cout << "seen[" << result << "] = " << ans.size() << "\n";
            result *= 10;
            ans += std::to_string(result / ld);
        }

        return ans;
    }
};

int main() {
    struct Testcase {
        int numerator;
        int denominator;
        std::string expected;
    };
    Testcase testcases[] = {
        {1, 2, "0.5"},
        {2, 1, "2"},
        {4, 333, "0.(012)"},
        {0, 5, "0"},
        {0, -5, "0"},

        {1, 5, "0.2"},
        {-1, 5, "-0.2"},
        {1, -5, "-0.2"},
        {-1, -5, "0.2"},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.fractionToDecimal(tc.numerator, tc.denominator);
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.numerator << "/" << tc.denominator << " failed. Got: "
                << actual << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
