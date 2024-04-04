#include <cassert>
#include <string>

class Solution {
public:
  int maxDepth(std::string s) {
    int ans = 0;
    int depth = 0;
    for (const char& ch: s) {
      if (ch == '(') {
        ++depth;
        ans = std::max(ans, depth);
      } else if (ch == ')') {
        --depth;
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.maxDepth("(1+(2*3)+((8)/4))+1") == 3);
  assert(s.maxDepth("(1)+((2))+(((3)))") == 3);
}
