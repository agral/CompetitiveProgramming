#include <cassert>
#include <string>

class Solution {
public:
  static constexpr int NUM_LETTERS = 'z' - 'a' + 1;

  bool isAnagram(std::string s, std::string t) {
    int diff[NUM_LETTERS] = {0};
    for (int i = 0; i < s.size(); ++i) {
      ++diff[s[i] - 'a'];
    }
    for (int i = 0; i < t.size(); ++i) {
      --diff[t[i] - 'a'];
    }
    for (int k = 0; k < NUM_LETTERS; k++) {
      if (diff[k] != 0) {
        return false;
      }
    }
    return true;
  }
};

int main() {
  Solution s;
  assert(s.isAnagram("anagram", "nagaram") == true);
  assert(s.isAnagram("rat", "car") == false);
  assert(s.isAnagram("abcde", "abcdef") == false);
  assert(s.isAnagram("abcde", "abcde") == true);
  assert(s.isAnagram("abcde", "abcd") == false);
}
