/**
 * Solution to the "The Coin Change Problem" challenge form HackerRank:
 * https://www.hackerrank.com/challenges/coin-change/problem
 * Created on: 23.10.2019
 * Last modified: 23.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

long long countWays(long long target, long long noOfCoins, std::vector<long long> const& values)
{
  std::vector<std::vector<long long>> dp(target + 1);
  for (auto &row : dp)
  {
    row.resize(values.size() + 1);
  }

  // Fills in the entries for the base case (one where target == 0):
  for (auto colId = 0; colId <= noOfCoins; ++colId)
  {
    dp[0][colId] = 1;
  }

  // Fills in the rest of the table:
  for (auto rowId = 1; rowId <= target; ++rowId)
  {
    for (auto colId = 0; colId < noOfCoins; ++colId)
    {
      // Number of distinct solutions that include rowId's coin:
      auto inclusive = (rowId - values.at(colId) >= 0) ? dp.at(rowId - values.at(colId)).at(colId) : 0;

      // Number of distinct solutions that exclude rowId's coin:
      auto exclusive = (colId > 0) ? dp[rowId][colId - 1] : 0;
      dp[rowId][colId] = inclusive + exclusive;
    }
  }

  return dp[target][noOfCoins - 1];
}

int main(int, char**)
{
  long long targetValue, noOfCoins;
  std::cin >> targetValue >> noOfCoins;
  std::vector<long long> denominations(noOfCoins);
  for (auto c = 0LL; c < noOfCoins; ++c)
  {
    std::cin >> denominations[c];
  }

  std::cout << countWays(targetValue, noOfCoins, denominations) << "\n";
  return 0;
}
