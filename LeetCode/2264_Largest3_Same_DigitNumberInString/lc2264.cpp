#include <cassert>
#include <string>

class Solution {
public:
  std::string largestGoodInteger(std::string num) {
    char best = '!'; // guaranteed lower than 0-9 digits
    for (int i = 0; i < num.size() - 2; i++) {
      if (num[i] >= '0' && num[i] <= '9' && num[i+1] == num[i] && num[i+2] == num[i]) {
        best = std::max(best, num[i]);
      }
    }
    return (best == '!' ? "" : std::string(3, best));
  }
};

int main() {
  Solution s;
  assert(s.largestGoodInteger("6777133339") == "777");
  assert(s.largestGoodInteger("2300019") == "000");
  assert(s.largestGoodInteger("42352338") == "");
}
