#include <cassert>
#include <string>

class Solution {
public:
    bool canMakeSubsequence(std::string str1, std::string str2) {
        int idx2 = 0;

        for (const char c: str1) {
            if ((str2[idx2] == c) || (str2[idx2] == 'a' + ((c - 'a' + 1) % 26))) {
                idx2 += 1;
                if (idx2 == str2.size()) {
                    return true;
                }
            }
        }
        return false;
    }
};

int main() {
    Solution s;
    assert(s.canMakeSubsequence("abc", "ad") == true);
    assert(s.canMakeSubsequence("zc", "ad") == true);
    assert(s.canMakeSubsequence("ab", "d") == false);
}
