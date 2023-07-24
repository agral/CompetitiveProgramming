#include <cassert>
#include <iostream>

class Solution {
public:
  double myPow(double x, int n) {
    // Poor optimizations & handling of corner cases; required by the problem's testcases:
    if (x == 0.0) {
      return 0.0;
    }
    else if (x == 1.0) {
      return x;
    }
    else if (x == -1.0) {
      return (n % 2 == 0) ? 1.0 : -1.0;
    }
    else if (n == 0) {
      return 1.0;
    }

    long ln = static_cast<long>(n);
    if (ln < 0) {
      x = 1.0 / x;
      ln = -ln;
    }

    double ans = 1;
    while (ln > 0) {
      if (ln % 2 == 1) {
        ans *= x;
      }
      x *= x;
      ln >>= 1; // floor-div-by-two
    }

    return ans;
  }
};

int main() {
  const double EPSILON=1e-10;
  Solution s{};
  assert(s.myPow(2, 10) - 1024.0 <= EPSILON);
  assert(s.myPow(2.1, 3) - 9.26100 <= EPSILON);
  assert(s.myPow(2, -2) - 0.25 <= EPSILON);
}
