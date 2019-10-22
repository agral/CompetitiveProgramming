/**
 * Solution to the "CommonChild" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/common-child/problem
 * Created on: 04.10.2019
 * Last modified: 22.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <string>
#include <vector>

std::size_t lcs(std::string const &first, std::string const &second)
{
  std::vector<std::vector<std::size_t>> dp;
  dp.resize(first.size() + 1);
  for (std::size_t k = 0; k <= first.size(); ++k)
  {
    dp[k].resize(second.size() + 1);
  }
  for (std::size_t f = 0; f <= first.size(); ++f)
  {
    for (std::size_t s = 0; s <= second.size(); ++s)
    {
      if (f == 0 || s == 0)
      {
        dp[f][s] = 0;
      }
      else if (first[f-1] == second[s-1])
      {
        dp[f][s] = dp[f-1][s-1] + 1;
      }
      else
      {
        dp[f][s] = std::max(dp[f-1][s], dp[f][s-1]);
      }
    }
  }

  return dp[first.size()][second.size()];
}

int main(int, char**)
{
  std::string a, b;
  std::cin >> a >> b;
  std::cout << lcs(a, b) << '\n';

  return 0;
}
