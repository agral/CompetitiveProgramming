#include <cassert>
#include <string>

class Solution {
public:
  int countHomogenous(std::string s) {
    const long long MOD = 1e9 + 7;
    long long ans = 0;
    int hs = 0; // homogenous range start, inclusive.
    int he = 1; // homogenous range end, not inclusive.
    while (true) {
      // grow homogenous range until next-char does not match (or end of string is reached)
      while (he < s.size() && (s[he] == s[hs])) {
        he++;
      }
      // Add this homogenous substring's homo-score to total score:
      long long length = (he - hs);
      long long hscore = (length * (length + 1)) / 2; // length can be at most 10^5, so hscore will always fit in a long long
      ans = (ans + hscore) % MOD;

      // If end-of-string is reached, return.
      if (he == s.size()) {
        return static_cast<int>(ans);
      }

      // Otherwise move start-of-range and end-of-range over the next letter, and continue the main loop.
      hs = he;
      he++;
    }
  }
};

int main() {
  Solution s;
  assert(s.countHomogenous("abbcccaa") == 13);
  assert(s.countHomogenous("xy") == 2);
  assert(s.countHomogenous("zzzzz") == 15);
}
