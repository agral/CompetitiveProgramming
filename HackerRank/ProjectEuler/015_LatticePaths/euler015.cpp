/**
 * Solution to the challenge #015 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler015
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

const int MODULUS = 1e9 + 7;


int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int X, Y;
    std::cin >> X >> Y;

    std::vector<std::vector<int>> dp;
    dp.resize(Y + 1);

    for (int y = 0; y <= Y; ++y)
    {
      dp[y].resize(X + 1);
    }

    // There is only 1 way to go to any point on the top and left borders of the net:
    for (int x = 0; x <= X; ++x)
    {
      dp[0][x] = 1;
    }
    for (int y = 1; y <= Y; ++y)
    {
      dp[y][0] = 1;
    }

    // For all the other points, the total sum of ways to get there is a sum of ways to get to its immediate top
    // and its immediate left neighbor:
    for (int y = 1; y <= Y; ++y)
    {
      for (int x = 1; x <= X; ++x)
      {
        dp[y][x] = (dp[y - 1][x] + dp[y][x - 1]) % MODULUS;
      }
    }

    std::cout << dp[Y][X] << std::endl;
  }

  return 0;
}
