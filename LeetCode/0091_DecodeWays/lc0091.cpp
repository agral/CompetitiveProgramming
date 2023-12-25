#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  int numDecodings(std::string s) {
    const int sz = s.size();
    std::vector<int> dp(sz + 1); // dp[i]: number of ways to decode s[i, sz>
    dp[sz] = 1;
    dp[sz-1] = (s[sz-1] == '0' ? 0 : 1);
    for (int i = sz - 2; i >= 0; i--) {
      if (s[i] != '0') {
        dp[i] += dp[i+1];
      }
      if ((s[i] == '1') || (s[i] == '2' && s[i+1] <= '6')) {
        dp[i] += dp[i+2];
      }
    }
    return dp[0];
  }
};

int main() {
  Solution s;
  assert(s.numDecodings("12") == 2);
  assert(s.numDecodings("226") == 3);
  assert(s.numDecodings("06") == 0);
}
