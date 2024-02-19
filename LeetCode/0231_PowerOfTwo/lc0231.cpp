#include <cassert>

/*
 * This is a nice trick that works without requiring popcount() or other architecture-specific tricks.
 * powers of two have just one set bit - the highest one.
 * powers of two minus one have all the bits set, except the highest one set in the power of two.
 * so, AND-ing these together yields 0, but only for an n that is a power of two.
 * Example:
 * n = 8 == 0b1000. n-1 = 7 == 0b0111. n & (n-1) == 0b0000.
 * for a non-power-of-two:
 * n = 6 == 0b110. n-1= 5 == 0b101. n & (n-1) = 0b100 == 4.
 */
class Solution {
public:
  bool isPowerOfTwo(int n) {
    if (n <= 0) {
      return false;
    }
    return (n & (n-1)) == 0;
  }
};

int main() {
  Solution s;
  assert(s.isPowerOfTwo(1) == true);
  assert(s.isPowerOfTwo(2) == true);
  assert(s.isPowerOfTwo(3) == false);
  assert(s.isPowerOfTwo(4) == true);
  assert(s.isPowerOfTwo(5) == false);
  assert(s.isPowerOfTwo(6) == false);
  assert(s.isPowerOfTwo(7) == false);
  assert(s.isPowerOfTwo(8) == true);
  assert(s.isPowerOfTwo(9) == false);
  assert(s.isPowerOfTwo(10) == false);
  assert(s.isPowerOfTwo(11) == false);
  assert(s.isPowerOfTwo(12) == false);
  assert(s.isPowerOfTwo(13) == false);
  assert(s.isPowerOfTwo(14) == false);
  assert(s.isPowerOfTwo(15) == false);
  assert(s.isPowerOfTwo(16) == true);
}
