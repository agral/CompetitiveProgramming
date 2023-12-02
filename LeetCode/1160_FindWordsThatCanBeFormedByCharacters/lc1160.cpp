#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  const int NUM_LETTERS = 'z' - 'a' + 1; // 26 or so.
  int countCharacters(std::vector<std::string>& words, std::string chars) {

    std::vector<int> count(NUM_LETTERS, 0);
    for (const char& ch: chars) {
      count[ch - 'a'] += 1;
    }

    int ans = 0;
    for (const auto& word: words) {
      std::vector<int> this_word_count(count);
      ans += word.size(); // temporarily. Will be reversed if string is not good.
      for (const char& ch: word) {
        if (--this_word_count[ch - 'a'] < 0) {
          ans -= word.size();
          break;
        }
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> words = {"cat", "bt", "hat", "tree"};
    std::string chars = "atach";
    assert(s.countCharacters(words, chars) == 6);
  }
  {
    std::vector<std::string> words = {"hello", "world", "leetcode"};
    std::string chars = "welldonehoneyr";
    assert(s.countCharacters(words, chars) == 10);
  }
}
