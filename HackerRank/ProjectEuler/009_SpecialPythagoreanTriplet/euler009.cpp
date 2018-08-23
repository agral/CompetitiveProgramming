/**
 * Solution to the challenge #009 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler009
 * Created on: 22-08-2018
 * Last Modified: 22-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <cstdint>
#include <algorithm>


const int maxN = 3000;

int CalcMaxProduct(int n)
{
  int maxProduct = -1;
  for (int a = 1; a < n / 3; ++a)
  {
    int b = ((n * n) - (2 * a * n)) / (2 * (n - a));
    int c = n - a - b;
    if ((a * a) + (b * b) == (c * c))
    {
      maxProduct = std::max(maxProduct, a * b * c);
    }
  }
  return maxProduct;
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::cout << CalcMaxProduct(N) << std::endl;
  }
  return 0;
}
