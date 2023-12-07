#include <cassert>
#include <string>

class Solution {
public:
  std::string largestOddNumber(std::string num) {
    int idx = num.size();
    while (idx > 0 && !(num[idx-1] & 1)) {
      --idx;
    }
    return num.substr(0, idx);
  };
};

int main() {
  Solution s;
  assert(s.largestOddNumber("52") == "5");
  assert(s.largestOddNumber("4206") == "");
  assert(s.largestOddNumber("35427") == "35427");
}
