#include <algorithm>
#include <cassert>
#include <map>
#include <stack>
#include <string>
#include <sstream>
#include <set>

class Solution {
public:
  std::string removeDuplicateLetters(std::string s) {
    std::map<char, int> char_count{};
    for (const char& ch: s) {
      char_count[ch] += 1;
    }

    std::stack<char> stack;
    std::set<char> known_chars;

    for (const char& ch: s) {
      char_count[ch] -= 1;

      if (known_chars.count(ch) > 0) {
        continue;
      }

      while (!stack.empty() && ch < stack.top() && char_count[stack.top()] > 0) {
        known_chars.erase(stack.top());
        stack.pop();
      }
      stack.push(ch);
      known_chars.insert(ch);
    }

    std::stringstream ans;
    while (!stack.empty()) {
      ans << stack.top();
      stack.pop();
    }

    std::string k = ans.str();
    std::reverse(k.begin(), k.end());
    return k;
  }
};

int main() {
  Solution s;
  assert(s.removeDuplicateLetters("bcabc") == "abc");
  assert(s.removeDuplicateLetters("cbacdcbc") == "acdb");
}
