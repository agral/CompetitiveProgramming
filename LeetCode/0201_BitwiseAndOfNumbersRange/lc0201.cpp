#include <cassert>

class Solution {
public:
  int rangeBitwiseAnd(int left, int right) {
    int bitsToShift = 0;
    while (left < right) {
      left >>= 1;  // shift bits in left & right by 1 position to the right,
      right >>= 1; // which is the same as integer division by 2.
      ++bitsToShift;
    }
    return (right << bitsToShift);
  }
};

int main() {
  Solution s;
  assert(s.rangeBitwiseAnd(5, 7) == 4);
  assert(s.rangeBitwiseAnd(0, 0) == 0);
  assert(s.rangeBitwiseAnd(1, 2147483647) == 0);
}
