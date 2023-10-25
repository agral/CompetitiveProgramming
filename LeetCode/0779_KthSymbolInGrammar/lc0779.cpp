#include <cassert>

/*
 * 0
 * 0 1
 * 01 10
 * 0110 1001
 * 01101001 10010110
 * ...
 */

class Solution {
public:
  int kthGrammar(int n, int k) {
    if (n == 1) {
      return 0;
    }
    // for anything else, value of current node can be expressed as in a recursive relation to its
    // parent node in a binary tree described by this task's grammar: [0]->01, [1]->10:
    //    0       1
    //   / \     / \
    //  0   1   1   0
    //  here, the "address" of parent's node is (n-1, (k+1)/2).
    //  obviously the "side" - left or right - matters, but also the value of the parent -
    //  as it "inverts" the values of child nodes.
    //  The overall recursive expression is:
    //  kthGrammar(n, k) = [(k % 2) + kthGrammar(n-1, (k + 1) / 2) + 1] mod 2
    //                     [  side  +    value-of-parent-node      + 1] mod 2
    return ((k & 1) + kthGrammar(n-1, (k+1) / 2) + 1) & 1;
  }
};

int main() {
  Solution s;
  assert(s.kthGrammar(1, 1) == 0); // a base case
  assert(s.kthGrammar(2, 1) == 0); // 2nd row
  assert(s.kthGrammar(2, 2) == 1);
  assert(s.kthGrammar(3, 1) == 0); // 3rd row
  assert(s.kthGrammar(3, 2) == 1);
  assert(s.kthGrammar(3, 3) == 1);
  assert(s.kthGrammar(3, 4) == 0);
  assert(s.kthGrammar(5, 8) == 1); // did not want to go further than 5th row with pen & paper
  assert(s.kthGrammar(5, 9) == 1);
  assert(s.kthGrammar(5, 10) == 0);
}
