#include <cassert>
#include <vector>
#include <iostream>

class Solution {
public:
  // How many combinations of coins exist that sum up to given amount?
  // e.g. amount=5, coins=[1,2,5]: 4 combinations.
  // 5=5, 5=2+2+1, 5=2+1+1+1, 5=1+1+1+1+1
  int change(int amount, std::vector<int>& coins) {
    int dp[amount+1][coins.size()];
    // Fill in the base case: exactly 1 way to get amount of 0, by using 0 of each coins.
    for (std::size_t c = 0; c < coins.size(); c++) {
      dp[0][c] = 1;
    }
    // Fill in the rest of the dp array in a bottom-up manner:
    for (int a = 1; a <= amount; a++) {
      for (int c = 0; c < coins.size(); c++) {
        int val = coins[c];
        // Number of solutions that sum up to a and INCLUDE coins[c]:
        int including = (a - val >= 0) ? dp[a - val][c] : 0;
        
        // Number of solutions that sum up to a and NOT INCLUDE coins[c]:
        int excluding = (c == 0) ? 0 : dp[a][c-1];

        dp[a][c] = including + excluding;
      }
    }
    /*
    for (std::size_t a = 0; a <= amount; a++) {
      for (std::size_t c = 0; c < coins.size(); c++) {
        std::cout << dp[a][c] << "\t";
      }
      std::cout << "\n";
    }
    */
    return dp[amount][coins.size() - 1];
  }
};

int main() {
  Solution s;
  std::vector<int> testcase1 = {1, 2, 5};
  assert(s.change(5, testcase1) == 4);
  std::vector<int> testcase2 = {2};
  assert(s.change(3, testcase2) == 0);
  std::vector<int> testcase3 = {10};
  assert(s.change(10, testcase3) == 1);

}
