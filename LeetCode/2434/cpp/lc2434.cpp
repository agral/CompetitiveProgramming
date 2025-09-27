#include <iostream>
#include <sstream>
#include <stack>
#include <vector>

// Runtime complexity: O(n)
// Auxiliary space complexity: O(n)
// Subjective level: medium/hard
class Solution {
public:
    static const int NUM_CHARS = 'z' - 'a' + 1;

    std::string robotWithString(std::string s) {
        std::vector<int> count(NUM_CHARS);
        for (const char ch: s) {
            ++count[ch-'a'];
        }

        std::stringstream ans;
        std::stack<char> stack;
        for (const char ch: s) {
            stack.push(ch);
            --count[ch-'a'];
            const char lowestAvailChar = getLowestAvailableChar(count);
            while (!stack.empty() && stack.top() <= lowestAvailChar) {
                ans << stack.top();
                stack.pop();
            }
        }

        // Use up the remaining entries on a stack:
        while (!stack.empty()) {
            ans << stack.top();
            stack.pop();
        }

        return ans.str();
    }
private:
    char getLowestAvailableChar(const std::vector<int>& count) {
        for (int i = 0; i < NUM_CHARS; i++) {
            if (count[i] > 0) {
                return 'a'+i;
            }
        }
        return '!';
    }
};

int main() {
    struct Testcase {
        std::string s;
        std::string expected;
    };
    Testcase testcases[] = {
        {"zza", "azz"},
        {"bac", "abc"},
        {"bdda", "addb"},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.robotWithString(tc.s);
        if (actual != tc.expected) {
            std::cout << "Testcase " << tc.s << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
