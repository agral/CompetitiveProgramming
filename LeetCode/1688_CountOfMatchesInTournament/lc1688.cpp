#include <cassert>

class Solution {
public:
  int numberOfMatches(int n) {
    // Until a winner is decided, (n-1) teams have to lose.
    // So, n-1 decisive games have to be played, after which
    // one team remains undefeated, and proclaimed a winner.
    return n - 1;
  }
};

class ActuallySimulated {
public:
  int numberOfMatches(int n) {
    if (n <= 1) {
      return 0;
    }
    int matches = (n / 2); // these matches generate floor(n/2) losers, who fall out of the tournament.
    int remaining = n - matches;
    return matches + numberOfMatches(remaining);
  }
};

int main() {
  Solution s;
  ActuallySimulated as;
  for (int i = 1; i <= 200; ++i) {
    assert(s.numberOfMatches(i) == as.numberOfMatches(i));
  }
}
