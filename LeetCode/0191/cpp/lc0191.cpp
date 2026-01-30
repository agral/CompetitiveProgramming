#include <cassert>
#include <cstdint>

class Solution {
public:
  // O(number-of-1-bits) -> O(n), but n here is at most 32, so technically O(1).
  // Can only be faster with processor builtins, e.g. __builtin_popcount(), a true O(1).
  int hammingWeight(uint32_t n) {
    int ans = (n == 0) ? 0 : 1;
    while ((n = n & (n - 1)) != 0) {
      ++ans;
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.hammingWeight(11 /*0b1011*/) == 3);
  assert(s.hammingWeight(128 /*0b10000000*/) == 1);
  assert(s.hammingWeight(0b11111111111111111111111111111101) == 31); // thirty one '1' bits.
}
