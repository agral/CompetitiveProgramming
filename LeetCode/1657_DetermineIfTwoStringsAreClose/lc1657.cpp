#include <algorithm>
#include <cassert>
#include <set>
#include <string>
#include <unordered_map>
#include <vector>

class Solution {
public:
  bool closeStrings(std::string word1, std::string word2) {
    if (word1.size() != word2.size()) {
      return false;
    }
    // maps character to number of times it occurs in word1/word2:
    std::unordered_map<char, int> charToCount1, charToCount2;
    // count of unique characters in word1/word2 (i.e. the "int" part of map above):
    std::vector<int> count1, count2;
    // unique characters present in word1/word2:
    std::set<char> chars1, chars2;

    for (const char ch: word1) {
      ++charToCount1[ch];
    }
    for (const char ch: word2) {
      ++charToCount2[ch];
    }

    for (const auto& [ch1, val1]: charToCount1) {
      chars1.insert(ch1);
      count1.push_back(val1);
    }
    for (const auto& [ch2, val2]: charToCount2) {
      chars2.insert(ch2);
      count2.push_back(val2);
    }
    if (chars1 != chars2) {
      // Rule2 allows exchanging of letter frequencies. This means that in order for
      // word1 to be transmutable into word2, both strings must be composed of
      // the same set of letters (e.g. "abccc" -> "aaabc", but "abc" -!> "def").
      return false;
    }

    std::ranges::sort(count1);
    std::ranges::sort(count2);

    return count1 == count2;
  }
};

int main() {
  Solution s;
  assert(s.closeStrings("abc", "bca") == true);
  assert(s.closeStrings("a", "aa") == false);
  assert(s.closeStrings("cabbba", "abbccc") == true);
  assert(s.closeStrings("aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff") == false);
  assert(s.closeStrings("abbzzca", "babzzcz") == false);
  assert(s.closeStrings("uau", "ssx") == false);
  assert(s.closeStrings("abc", "def") == false);
  assert(s.closeStrings("abccc", "aaabc") == true);
}
