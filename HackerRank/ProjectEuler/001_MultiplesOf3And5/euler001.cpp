/**
 * Name: euler001.cpp
 * Description: Solves challenge #001 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler001
 * Created on: 17-08-2018
 * Last Modified: 17-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <cstdint>

uint64_t ArithmeticSum(uint64_t n0, uint64_t numberOfElements)
{
  return n0 * (1 + numberOfElements) * numberOfElements / 2;
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N, n;
    std::cin >> N;
    n = N - 1; // We sum up to, but *not including* N - which is the same as summing up to and including N - 1.
    uint64_t sumOfMultiplesOf3 = ArithmeticSum(3, n / 3);
    uint64_t sumOfMultiplesOf5 = ArithmeticSum(5, n / 5);
    uint64_t sumOfMultiplesOf15 = ArithmeticSum(15, n / 15);
    uint64_t ans = sumOfMultiplesOf3 - sumOfMultiplesOf15 + sumOfMultiplesOf5;
    std::cout << ans << std::endl;
  }

  return 0;
}
