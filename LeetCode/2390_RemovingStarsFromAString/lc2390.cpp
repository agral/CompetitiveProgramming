#include <cassert>
#include <string>
#include <deque>

class Solution {
public:
  std::string removeStars(std::string s) {
    std::deque<char> stack;
    for (const char ch: s) {
      if (ch == '*') {
        stack.pop_back();
      } else {
        stack.push_back(ch);
      }
    }

    return std::string(stack.begin(), stack.end());
  }
};

int main() {
  Solution s;
  assert(s.removeStars("leet**cod*e") == "lecoe");
  assert(s.removeStars("erase*****") == "");
}
