#include <iostream>
#include <sstream>
#include <stack>
#include <string>

class Solution {
public:
    std::string smallestNumber(std::string pattern) {
        std::stringstream ss;
        std::stack<char> stack{ {'1'} };

        for (const char ch: pattern) {
            char maxChar = stack.top();
            if (ch == 'I') {
                while (!stack.empty()) {
                    maxChar = std::max(maxChar, stack.top());
                    ss << stack.top();
                    stack.pop();
                }
            }
            stack.push(maxChar + 1);
        }

        while (!stack.empty()) {
            ss << stack.top();
            stack.pop();
        }

        return ss.str();
    }
};

int main() {
    struct Testcase {
        std::string pattern;
        std::string expected;
    };
    Testcase testcases[] = {
        {"IIIDIDDD", "123549876"},
        {"DDD", "4321"},
    };
    Solution s{};
    bool is_ok = true;
    for (const Testcase& tc: testcases) {
        auto actual = s.smallestNumber(tc.pattern);
        if (actual != tc.expected) {
            is_ok = false;
            std::cout << "Testcase " << tc.pattern << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
        }
    }
    if (is_ok) {
        std::cout << "[OK] All testcases passed successfully.\n";
    } else {
        std::cout << "[FAIL] Some testcases are failing.\n";
    }
    return 0;
}
