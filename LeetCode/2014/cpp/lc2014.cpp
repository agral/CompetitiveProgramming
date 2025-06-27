#include <deque>
#include <iostream>
#include <vector>

class Solution {
public:
    std::string longestSubsequenceRepeatedK(std::string s, int k) {
        std::string ans = {};
        std::vector<char> validChars;
        std::vector<int> freq(26);
        std::deque<std::string> q{{""}};

        for (const char c: s) {
            ++freq[c - 'a'];
        }

        for (char ch{'a'}; ch <= 'z'; ch++) {
            if (freq[ch - 'a'] >= k) {
                validChars.push_back(ch);
            }
        }

        while (!q.empty()) {
            std::string currSubsequence = q.front();
            q.pop_front();
            if (currSubsequence.size() * k > s.size()) {
                return ans;
            }
            for (char ch: validChars) {
                std::string newSubsequence = currSubsequence + ch;
                if (isSub(s, k, newSubsequence)) {
                    q.push_back(newSubsequence);
                    ans = newSubsequence;
                }
            }
        }

        return ans;
    }

private:
    bool isSub(std::string& s, int k, std::string& subsequence) {
        int index = 0;
        for (const char c: s) {
            if (c == subsequence[index]) {
                index += 1;
                if (index == subsequence.size()) {
                    k -= 1;
                    if (k == 0) {
                        return true;
                    }
                    index = 0;
                }
            }
        }
        return false;
    }
};

int main() {
    struct Testcase {
        std::string s;
        int k;
        std::string expected;
    };
    Testcase testcases[] = {
        {"letsleetcode", 2, "let"},
        {"bb", 2, "b"},
        {"ab", 2, ""},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.longestSubsequenceRepeatedK(tc.s, tc.k);
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
