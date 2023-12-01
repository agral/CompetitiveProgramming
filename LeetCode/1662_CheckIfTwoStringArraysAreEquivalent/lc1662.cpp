#include <cassert>
#include <string>
#include <sstream>
#include <vector>

class Solution {
public:
  bool arrayStringsAreEqual(std::vector<std::string>& word1, std::vector<std::string>& word2) {
    std::stringstream ss1, ss2;
    for (const auto& w1: word1) {
      ss1 << w1;
    }
    for (const auto& w2: word2) {
      ss2 << w2;
    }

    return ss1.str() == ss2.str();
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> word1 = {"ab", "c"};
    std::vector<std::string> word2 = {"a", "bc"};
    assert(s.arrayStringsAreEqual(word1, word2) == true);
  }
  {
    std::vector<std::string> word1 = {"a", "cb"};
    std::vector<std::string> word2 = {"ab", "c"};
    assert(s.arrayStringsAreEqual(word1, word2) == false);
  }
  {
    std::vector<std::string> word1 = {"abc", "d", "defg"};
    std::vector<std::string> word2 = {"abcddefg"};
    assert(s.arrayStringsAreEqual(word1, word2) == true);
  }
}
