#include <cassert>
#include <string>
#include <stack>

class Solution {
public:
  bool backspaceCompare(std::string s, std::string t) {
    std::stack<char> stack_s;
    for (const auto& ch: s) {
      if (ch == '#') {
        if (!stack_s.empty()) {
          stack_s.pop();
        }
      } else {
        stack_s.push(ch);
      }
    }

    std::stack<char> stack_t;
    for (const auto& ch: t) {
      if (ch == '#') {
        if (!stack_t.empty()) {
          stack_t.pop();
        }
      } else {
        stack_t.push(ch);
      }
    }

    return stack_s == stack_t;
  };
};

int main() {
  Solution s;
  assert(s.backspaceCompare("ab#c", "ad#c") == true);
  assert(s.backspaceCompare("ab##", "c#d#") == true);
  assert(s.backspaceCompare("a#c", "b") == false);
  assert(s.backspaceCompare("y#fo##f", "y#f#o##f") == true);
}
