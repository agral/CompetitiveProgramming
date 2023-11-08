#include <algorithm>
#include <cassert>
#include <string>
#include <array>
#include <iostream>

class Solution {
public:
  int lengthOfLongestSubstring(std::string s) {
    if (s.empty()) { return 0; } // handle corner case.
    int ans = 1;
    int idx_start = 0; // where substring starts, inclusively
    int idx_end = 1; // where substring ends, non-inclusively. I.e. "racecar"[1, 4] -> "ace".
    m_containedChars = {false};
    m_containedChars[s[0]] = true;
    while (idx_end < s.size()) {
      // If next character in a string is not yet contained in the substring,
      // grow the substring (move end-pointer).
      if (!m_containedChars[s[idx_end]]) {
        m_containedChars[s[idx_end]] = true;
        idx_end++;
        ans = std::max(ans, (idx_end - idx_start));
      }
      // otherwise, reduce the substring by moving the start-pointer until there are no duplicates.
      else {
        m_containedChars[s[idx_start]] = false;
        idx_start++;
      }
    }
    return ans;
  }
private:
  std::array<bool, 256> m_containedChars;
};

int main() {
  Solution s;
  assert(s.lengthOfLongestSubstring("abcabcbb") == 3);
  assert(s.lengthOfLongestSubstring("bbbbb") == 1);
  assert(s.lengthOfLongestSubstring("pwwkew") == 3);

  assert(s.lengthOfLongestSubstring("abcbdefba") == 5); // cbdef
  assert(s.lengthOfLongestSubstring("abac") == 3); // bac
  assert(s.lengthOfLongestSubstring("") == 0);
  assert(s.lengthOfLongestSubstring(" ") == 1);
  assert(s.lengthOfLongestSubstring("     ") == 1);

  assert(s.lengthOfLongestSubstring("zabczd") == 5); // abczd
}
