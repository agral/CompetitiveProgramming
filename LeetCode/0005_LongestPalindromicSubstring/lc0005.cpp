#include <cassert>
#include <string>

class Solution {
public:
  std::string longestPalindrome(std::string s) {
    int ansL = 0, ansSize = 0;
    for (int i = 0; i < s.size(); i++) {
      // Check odd-size palindromes:
      int left = i, right = i;
      while (left >= 0 && right < s.size() && s[left] == s[right]) {
        if (right - left + 1 > ansSize) {
          ansL = left;
          ansSize = right - left + 1;
        }
        left -= 1;
        right += 1;
      }

      // Also check even-sized palindromes:
      left = i, right = i + 1;
      while (left >= 0 && right < s.size() && s[left] == s[right]) {
        if (right - left + 1 > ansSize) {
          ansL = left;
          ansSize = right - left + 1;
        }
        left -= 1;
        right += 1;
      }
    }
    return s.substr(ansL, ansSize);
  };
};

int main() {
  Solution s;
  std::string ans1 = s.longestPalindrome("babad");
  assert(ans1 == "bab" || ans1 == "aba");

  std::string ans2 = s.longestPalindrome("cbbd");
  assert(ans2 == "bb");

  assert(s.longestPalindrome("") == "");
}
