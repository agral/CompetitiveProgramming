#include <algorithm>
#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  std::string convertToTitle(int columnNumber) {
    std::vector<char> ans;
    while (columnNumber > 0) {
      int remainder = columnNumber % 26;
      int letterOffset = (remainder + 25) % 26;
      ans.push_back('A' + letterOffset);
      columnNumber -= (remainder == 0 ? 26 : remainder);
      columnNumber /= 26;
    }
    std::reverse(ans.begin(), ans.end());
    return std::string(ans.begin(), ans.end());
  }
};

int main() {
  Solution s{};

  assert(s.convertToTitle(1) == "A");
  assert(s.convertToTitle(2) == "B");
  assert(s.convertToTitle(25) == "Y");
  assert(s.convertToTitle(26) == "Z");
  assert(s.convertToTitle(27) == "AA");
  assert(s.convertToTitle(51) == "AY");
  assert(s.convertToTitle(52) == "AZ");
  assert(s.convertToTitle(53) == "BA");
  assert(s.convertToTitle(701) == "ZY");
  assert(s.convertToTitle(702) == "ZZ");
  assert(s.convertToTitle(703) == "AAA");
}
