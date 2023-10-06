#include <cassert>

class Solution {
public:
  int integerBreak(int n) {
    if (n == 2) {
      return 1; // break 2 into 1+1; 1*1 is 1.
    }
    if (n == 3) {
      return 2; // break 3 into 2+1, 2*1 is 2.
    }
    int ans = 1;
    while (n > 4) {
      n -= 3;
      ans *= 3;
    }
    return ans * n;
  }
};

int main() {
  Solution s;
  assert(s.integerBreak(2) == 1);
  assert(s.integerBreak(3) == 2);
  assert(s.integerBreak(4) == 4); // 2 * 2
  assert(s.integerBreak(5) == 6); // 3 * 2
  assert(s.integerBreak(6) == 9); // 3 * 3
  assert(s.integerBreak(7) == 12); // 3 * 2 * 2
  assert(s.integerBreak(8) == 18); // 3 * 3 * 2
  assert(s.integerBreak(9) == 27); // 3 * 3 * 3
  assert(s.integerBreak(10) == 36); // 3 * 3 * 2 * 2
}
