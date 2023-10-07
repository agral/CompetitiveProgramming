#include <cassert>

class Solution {
public:
  const int MOD = 1e9 + 7;
  int numOfArrays(int n, int m, int k) {
    // n: length of the array. <1 ... 50>.
    // m: values (kinda?). <1 ... 100>.
    // k: search cost in number of comparisons made by custom algo,
    // given in problem description. Can be anything in <0 ... n> range.
    long long dp[50+1][100+1][50+1] = {};

    for (int i = 0; i <= m; i++) {
      // Base case: exactly 1 way to construct an array with one sole number.
      dp[1][i][1] = 1;
    }

    for (int in = 1; in <= n; in++) {
      for (int im = 1; im <= m; im++) {
        for (int ik = 1; ik <= k; ik++) {
          long long sum = 0;
          sum = (sum + im * dp[in-1][im][ik]) % MOD;
          for (int a = 1; a < im; a++) {
            sum = (sum + dp[in-1][a][ik-1]) % MOD;
          }
          dp[in][im][ik] = (dp[in][im][ik] + sum) % MOD;
        }
      }
    }

    long long ans{0LL};
    for (int im = 1; im <= m; im++) {
      ans = (ans + dp[n][im][k]) % MOD;
    }
    return ans;
  }
};

int main() {
  Solution s;
  assert(s.numOfArrays(2, 3, 1) == 6);
  assert(s.numOfArrays(5, 2, 3) == 0);
  assert(s.numOfArrays(9, 1, 1) == 1);
}
