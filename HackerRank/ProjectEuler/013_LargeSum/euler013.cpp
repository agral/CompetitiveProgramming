/**
 * Solution to the challenge #013 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler013
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <string>
#include "BigInteger.hpp"

const int DIGITS_LIMIT = 10;

int main()
{
  int N;
  std::cin >> N;
  BigInteger sum{0};
  for (int n = 0; n < N; ++n)
  {
    std::string line;
    std::cin >> line;
    BigInteger delta{line};
    sum.add(delta);
  }

  std::string ans = sum.toString();
  for (int i = 0; i < DIGITS_LIMIT; ++i)
  {
    std::cout << ans[i];
  }
  std::cout << std::endl;

  return 0;
}
