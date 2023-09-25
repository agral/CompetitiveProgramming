#include <cassert>
#include <string>

class Solution {
public:
  char findTheDifference(std::string s, std::string t) {
    int charsum = 0;
    for (const auto& ch: t) {
      charsum += static_cast<int>(ch);
    }
    for (const auto& ch: s) {
      charsum -= static_cast<int>(ch);
    }
    return static_cast<char>(charsum);
  }
};

int main() {
  Solution s;
  assert(s.findTheDifference("abcd", "abcde") == 'e');
  assert(s.findTheDifference("", "y") == 'y');
}
