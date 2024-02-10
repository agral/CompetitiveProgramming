#include <cassert>
#include <string>

class Solution {
public:
  int countSubstrings(std::string s) {
    int ans = 0;
    for (int i = 0; i < s.size(); ++i) {
      ans += countPalindromes(s, i, i); // counts odd-length palindromes
      ans += countPalindromes(s, i, i + 1); // counts even-length palindromes
    }
    return ans;
  }

private:
  int countPalindromes(const std::string& s, int left, int right) {
    int ans = 0;
    while (left >= 0 && right < s.size() && (s[left] == s[right])) {
      --left;
      ++right;
      ++ans;
    }
    return ans;
  }
};

int main() {
  Solution s;

  assert(s.countSubstrings("abc") == 3);
  assert(s.countSubstrings("aaa") == 6);
}
