#include <cassert>
#include <string>
#include <sstream>
#include <deque>

class Solution {
public:
  std::string removeKdigits(std::string num, int k) {
    const size_t SIZE = num.size();
    if (k == SIZE) {
      return {"0"};
    }
    std::stringstream ans;
    std::deque<char> stack;

    for (int i = 0; i < SIZE; ++i) {
      while (k > 0 && !stack.empty() && stack.back() > num[i]) {
        stack.pop_back();
        --k;
      }
      stack.push_back(num[i]);
    }

    while (k > 0) {
      stack.pop_back();
      --k;
    }

    bool isEmpty = true;
    for (const char ch: stack) {
      if (!(isEmpty && ch == '0')) {
        isEmpty = false;
        ans << ch;
      }
    }
    if (isEmpty) {
      return {"0"};
    }
    return ans.str();
  }
};

int main() {
  Solution s;
  assert(s.removeKdigits("1432219", 3) == "1219");
  assert(s.removeKdigits("10200", 1) == "200");
  assert(s.removeKdigits("10", 2) == "0");
}
