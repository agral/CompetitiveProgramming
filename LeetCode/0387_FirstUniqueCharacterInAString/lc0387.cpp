#include <cassert>
#include <string>

class Solution {
public:
  int firstUniqChar(std::string s) {
    // 1. Count the total time each letter occurs in the string s:
    static constexpr std::size_t NUM_LETTERS = 'z' - 'a' + 1;
    int count[NUM_LETTERS] = {};
    for (const char ch: s) {
      count[ch - 'a'] += 1;
    }

    // 2. Iterate over letters of string s. If any letter occurs exactly once, return its index:
    for (int i = 0; i < s.size(); ++i) {
      if (count[s[i] - 'a'] == 1) {
        return i;
      }
    }

    // 3. After reaching the end of the string, no letter is unique. Return -1 as per task requirements:
    return -1;
  }
};

int main() {
  Solution s;
  assert(s.firstUniqChar("leetcode") == 0);     // 'l', 0th char, is unique
  assert(s.firstUniqChar("loveleetcode") == 2); // 'v', 2nd char, is unique
  assert(s.firstUniqChar("aabb") == -1);        // no unique chars present
}
