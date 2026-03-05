#include <cassert>
#include <string>

class Solution {
public:
  int minOperations(std::string s) {
    int numFlipsForZero = 0; // stores number of bit flips necessary to make s
    int numFlipsForOne = 0;  // alternating, starting with '0' or '1', respectively.
    for (int i = 0; i < s.size(); ++i) {
      // for a string starting with zero, ith bit should be set if i is odd.
      // for a string starting with one, ith bit should be set if i is even.
      char wantedZero = (i & 1) ? '1' : '0';
      char wantedOne  = (i & 1) ? '0' : '1';
      if (s[i] != wantedZero) {
        ++numFlipsForZero;
      }
      if (s[i] != wantedOne) {
        ++numFlipsForOne;
      }
    }
    return std::min(numFlipsForZero, numFlipsForOne);
  }
};

int main() {
  Solution s;
  assert(s.minOperations("0100") == 1); // 0100 -> 0101
  assert(s.minOperations("10") == 0); // 10 -> already alternating.
  assert(s.minOperations("1111") == 2); // 1111 -> 0101
}
