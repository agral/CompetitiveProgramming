#include <cassert>

class Solution {
public:
  // climbStairs(n) == climbStairs(n-2) + climbStairs(n-1),
  // as to get to step n the climber could have taken one step from
  // n-1, or two steps from n-2.
  // Base cases are: 1 way for 1st step, and 2 ways for 2nd step.
  // So the problem boils down to calculating (n+1)th fibonacci number.
  int climbStairs(int n) {
    int prev = 0;
    int ans = 1;
    int tmp;
    for (int k = 0; k < n; k++) {
      tmp = prev;
      prev = ans;
      ans += tmp;
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.climbStairs(1) == 1);
  assert(s.climbStairs(2) == 2);
  assert(s.climbStairs(3) == 3);
  assert(s.climbStairs(4) == 5);
  assert(s.climbStairs(5) == 8);
  assert(s.climbStairs(43) == 701408733);
  assert(s.climbStairs(44) == 1134903170);
  assert(s.climbStairs(45) == 1836311903); // note: nice, no int32 overflow up to 45.
}
