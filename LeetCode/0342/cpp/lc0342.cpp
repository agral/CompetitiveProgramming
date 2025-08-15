#include <cassert>

class Solution {
public:
  bool isPowerOfFour(int n) {
    if (n < 0) {
      // because of course they will test it with INT_MIN...
      return false;
    }

    // 1st check: verify that n is a power of two.
    // Powers of two have only one bit set.
    if ((n & (n - 1)) != 0) {
      return false;
    }

    // 2nd check: verify that the only set bit is at odd index.
    // Consecutive powers of four have the following binary representations:
    // 4^0 ->  0b00000000000000000000000000000001
    // 4^1 ->  0b00000000000000000000000000000100
    // 4^2 ->  0b00000000000000000000000000010000
    // etc. This might be achievable with another trivia (e.g. it looks
    // that every power of four is == 1 (mod 3)), but using a mask works
    // just as well.
    int mask = 0b01010101010101010101010101010101;
    return (n & mask);
  };
};

int main() {
  Solution s;
  assert(s.isPowerOfFour(1) == true); // this is 4^0.
  assert(s.isPowerOfFour(2) == false); // power of two, but not of four.
  assert(s.isPowerOfFour(3) == false);
  assert(s.isPowerOfFour(4) == true);
  assert(s.isPowerOfFour(9) == false);
  assert(s.isPowerOfFour(16) == true);
  assert(s.isPowerOfFour(64) == true);
  assert(s.isPowerOfFour(256) == true);
  assert(s.isPowerOfFour(512) == false);
  assert(s.isPowerOfFour(1024) == true);
  assert(s.isPowerOfFour(4096) == true);
}
