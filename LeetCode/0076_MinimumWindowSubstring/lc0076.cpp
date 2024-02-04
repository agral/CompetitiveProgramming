#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  std::string minWindow(std::string s, std::string t) {
    std::vector<int> count(256, 0);
    int left = -1;
    int remainingCharsCount = t.size();
    int minSize = s.size() + 1;

    // Calculate individual char count of target string:
    for (const char c: t) {
      ++count[c];
    }

    for (int l = 0, r = 0; r < s.size(); ++r) {
      if (--count[s[r]] >= 0) {
        --remainingCharsCount;
      }

      while (remainingCharsCount == 0) {
        if (r + 1 - l < minSize) {
          left = l;
          minSize = r + 1 - l;
        }
        if (++count[s[l++]] > 0) {
          ++remainingCharsCount;
        }
      }
    }

    if (left == -1) {
      return {};
    }

    return s.substr(left, minSize);
  }
};

int main() {
  Solution s;
  assert(s.minWindow("ADOBECODEBANC", "ABC") == "BANC");
  assert(s.minWindow("a", "a") == "a");
  assert(s.minWindow("a", "aa") == "");
}
