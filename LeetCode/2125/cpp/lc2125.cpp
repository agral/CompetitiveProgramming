#include <algorithm>
#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  int numberOfBeams(std::vector<std::string>& bank) {
    int last = -1;
    int ans = 0;
    for (const std::string& row: bank) {
      const int count = std::ranges::count(row, '1');
      if (count > 0) {
        if (last == -1) {
          last = count;
        }
        else {
          ans += last * count;
          last = count;
        }
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> bank = {"011001", "000000", "010100", "001000"};
    assert(s.numberOfBeams(bank) == 8);
  }
  {
    std::vector<std::string> bank = {"000", "111", "000"};
    assert(s.numberOfBeams(bank) == 0);
  }
};
