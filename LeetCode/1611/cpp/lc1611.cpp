#include <cassert>

// *** Note from agral:
// After doing it, I *hate* this, and it's a daily question for 30 Nov '23.
// There's no way this tests anything in a candidate,
// except having grinded a lot of programming challenges before attempting this one.

// Crucial observation (pen & paper, but this may take some time to spot):
// solving 2^k -> 0 requires (2^(k+1))-1 operations.
// No. | Op | Num | Notes
// 0   |    | 100 | 2^2, will need 2^3-1 == 7 operations
// 1   | 1  | 101 |
// 2   | 2  | 111 |
// 3   | 1  | 110 |
// 4   | 2  | 010 | this is 2^1, it will require 3 ops
// 5   | 1  | 011 |
// 6   | 2  | 001 | this is 2^0, requires 1 ops
// 7   | 1  | 110 |

// So, for n = 1XXXX, the approach is somewhat recursive:
// - if the 2nd bit is 0, then the "extra" cost of flipping 2nd bit to 1 has to be added,
// - if the 2nd bit is 1, there's no extra cost.
class Solution {
public:
  int minimumOneBitOperations(int n) {
    if (n == 0) {
      // A nasty corner case
      return 0;
    }

    int p = 1;
    while (p << 1 <= n) { // while 2*p <= n
      p <<= 1; // p *= 2, but fancier.
    }
    // Now p holds the largest power-of-two <= n.
    return minimumOneBitOperations(n ^ (p | p >> 1)) + p;
    // *** Note 2: No way I'd come to the relation aboveduring an interview.
    // Took solid 30 mins of pen-and-paper work, when not stressed.
  }
};

int main() {
  Solution s;
  assert(s.minimumOneBitOperations(3) == 2);
  assert(s.minimumOneBitOperations(6) == 4);
}
