/**
 * Solution to the challenge #018 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler018
 * Created on: 14-09-2018
 * Last Modified: 14-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <cstdint>
#include <vector>

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::vector<std::vector<int>> pyramid;
    std::vector<std::vector<int>> maxSumHere;

    // Reads the numbers, then computes the max sum at each step of the pyramid, starting from the top.
    pyramid.resize(N);
    maxSumHere.resize(N);
    for (int n = 0; n < N; ++n)
    {
      pyramid[n].resize(n + 1);
      maxSumHere[n].resize(n + 1);
      for (int k = 0; k <= n; ++k)
      {
        std::cin >> pyramid[n][k];
      }
    }
    for (int n = 0; n < N; ++n)
    {
      for (int k = 0; k <= n; ++k)
      {
        // The maximum possible sum here equals at least the value in the current cell.
        maxSumHere[n][k] = pyramid[n][k];
        if (k < n)
        {
          // If the number is not the last one in a row, it has a "top-right" neighbor.
          // Its actual index in the array is [n-1][k].
          maxSumHere[n][k] = std::max(maxSumHere[n][k], maxSumHere[n-1][k] + pyramid[n][k]);
        }
        if (k > 0)
        {
          // If the number is not the first one in a row, it has a "top-left" neighbor.
          // Its actual index in the array is [n-1][k-1].
          maxSumHere[n][k] = std::max(maxSumHere[n][k], maxSumHere[n-1][k-1] + pyramid[n][k]);
        }
      }
    }

    // Outputs the answer - which is the largest value in the bottom row of maxSumHere vector:
    std::cout << *std::max_element(maxSumHere[N-1].begin(), maxSumHere[N-1].end()) << std::endl;
  }
  return 0;
}
