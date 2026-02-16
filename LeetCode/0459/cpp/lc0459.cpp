#include <cassert>
#include <string>

class Solution {
public:
  bool repeatedSubstringPattern(std::string s) {
    if (s.size() <= 1) {
      // Edge cases. Strings of at most 1 character length cannot be further divided.
      return false;
    }
    /*
    std::string repeated = s + s;
    std::string testString = repeated.substr(1, 2 * s.size() - 2);
    return testString.find(s) != std::string::npos;
    */
    // Same as above, but an one-liner!:
    return (s + s).substr(1, 2 * s.size() - 2).find(s) != std::string::npos;
  }
};

int main() {
  Solution s{};
  assert(s.repeatedSubstringPattern("abab") == true); // "ab" * 2
  assert(s.repeatedSubstringPattern("aba") == false);
  assert(s.repeatedSubstringPattern("abcabcabcabc") == true); // "abc * 4 | "abcabc" * 2
}
