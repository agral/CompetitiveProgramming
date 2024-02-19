#include <cassert>

/*
 * Powers of three:
 * 3^0 |  1  | 0b00000001
 * 3^1 |  3  | 0b00000011
 * 3^2 |  9  | 0b00001001
 * 3^3 |  27 | 0b00011011
 * 3^4 |  81 | 0b01010001
 * 3^5 | 243 | 0b11110011
 * (there is no clear pattern here)
 *
 * Max power of three that fits in an int: 3^19 == 1'162'261'467
 */

class Solution {
public:
  bool isPowerOfThree(int n) {
    return (n > 0) && (1'162'261'467 % n == 0);
  }
};

int main() {
  Solution s;
  assert(s.isPowerOfThree(0) == false);
  assert(s.isPowerOfThree(1) == true);
  assert(s.isPowerOfThree(2) == false);
  assert(s.isPowerOfThree(3) == true);
  assert(s.isPowerOfThree(4) == false);
  assert(s.isPowerOfThree(5) == false);
  assert(s.isPowerOfThree(6) == false);
  assert(s.isPowerOfThree(7) == false);
  assert(s.isPowerOfThree(8) == false);
  assert(s.isPowerOfThree(9) == true);
  assert(s.isPowerOfThree(-9) == false);
}
