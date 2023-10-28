#include <cassert>
#include <vector>

class Solution {
public:
  const int a = 0;
  const int e = 1;
  const int i = 2;
  const int o = 3;
  const int u = 4;
  const long long MOD = 1e9 + 7;

  int countVowelPermutation(int n) {
    if (n <= 0) {
      return 0;
    }

    // dp[k][c]: holds number of strings of length k, that end with character c.
    // here k runs from 1 to n, and c is from a to u (i.e. from 0 to 4).
    std::vector<std::vector<long long>> dp(n+1, {0, 0, 0, 0, 0});
    dp[1] = {1, 1, 1, 1, 1}; // answer for n=1
    for (int k = 2; k <= n; k++) {
      dp[k][a] = (dp[k-1][e] + dp[k-1][i] + dp[k-1][u]) % MOD;
      dp[k][e] = (dp[k-1][a] + dp[k-1][i]) % MOD;
      dp[k][i] = (dp[k-1][e] + dp[k-1][o]) % MOD;
      dp[k][o] = (dp[k-1][i]); // o can only follow an i, no need to modulo this.
      dp[k][u] = (dp[k-1][i] + dp[k-1][o]) % MOD;
    }
    long long ans = 0;
    for (int c = 0; c <= u; c++) {
      ans = (ans + dp[n][c]) % MOD;
    }
    return static_cast<int>(ans);
  };
};

int main() {
  Solution s;
  assert(s.countVowelPermutation(0) == 0);
  assert(s.countVowelPermutation(1) == 5);
  assert(s.countVowelPermutation(2) == 10);
  assert(s.countVowelPermutation(5) == 68);
}
