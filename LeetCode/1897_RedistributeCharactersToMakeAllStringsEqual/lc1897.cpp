#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  bool makeEqual(std::vector<std::string>& words) {
    const int NUM_LETTERS = 'z' - 'a' + 1;
    std::vector<int> count(NUM_LETTERS, 0);
    for (const auto& word: words) {
      for (const char ch: word) {
        ++count[ch - 'a'];
      }
    }
    for (int i = 0; i < NUM_LETTERS; i++) {
      if (count[i] % words.size() != 0) {
        return false;
      }
    }
    return true;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> words = {"abc", "aabc", "bc"};
    assert(s.makeEqual(words) == true);
  }
  {
    std::vector<std::string> words = {"ab", "a"};
    assert(s.makeEqual(words) == false);
  }
}
