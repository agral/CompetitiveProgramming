#include <iostream>
#include <queue>
#include <unordered_map>

class Solution {
public:
    std::string getHappyString(int n, int k) {
        const std::unordered_map<char, std::string> letterToNext = {
            {'a', "bc"},
            {'b', "ac"},
            {'c', "ab"}
        };

        std::queue<std::string> q{{"a", "b", "c"}};

        while (q.front().size() != n) {
            std::string base = q.front();
            q.pop();
            for (char postfix: letterToNext.at(base.back())) { // note: .at() fixes "no viable operator[]".
                q.push(base + postfix);
            }
        }
        if (q.size() < k) {
            return "";
        }
        for (int i = 1; i < k; i++) { // note: starting from 1, so that (k-1) entries are removed. Then k-th is the one on top.
            q.pop();
        }
        return q.front();

    }
};

int main() {
    struct Testcase {
        int n;
        int k;
        std::string expected;
    };
    Testcase testcases[] = {
        {1, 3, "c"},
        {1, 4, ""},
        {3, 9, "cab"},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (const Testcase& tc: testcases) {
        auto actual = s.getHappyString(tc.n, tc.k);
        if (actual != tc.expected) {
            std::cout << "Testcase (n=" << tc.n << ", k=" << tc.k << ") failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
