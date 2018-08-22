/**
 * Solution to the challenge #005 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler005
 * Created on: 20-08-2018
 * Last Modified: 20-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cstdint>
#include <iostream>
#include <numeric>

const int maxN = 40;

/**
 * This answer is valid in c++17, but HackerRank's most recent supported version is still c++14...
 */
/*
int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    uint64_t lcm = 1;
    for (int n = 2; n <= N; ++n)
    {
      lcm = std::lcm(lcm, n);
    }
    std::cout << lcm << std::endl;
  }
  return 0;
}
*/

/**
 * std::lcm() and std::gcd are introduced in c++17.
 * The below code calculates gcd(a, b) manually and uses that
 * to calculate the lcm(a,b) == a * b / gcd(a, b).
 */

uint64_t Gcd(uint64_t a, uint64_t b)
{
  if ((a == 0) || (b == 0))
  {
    return 0;
  }
  if (b > a)
  {
    return Gcd(b, a);
  }
  while (b != 0)
  {
    uint64_t amodb = a % b;
    a = b;
    b = amodb;
  }
  return a;
}

uint64_t Lcm(uint64_t a, uint64_t b)
{
  return a * b / Gcd(a, b);
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    uint64_t N;
    std::cin >> N;
    uint64_t lcm = 1;
    for (uint64_t n = 2; n <= N; ++n)
    {
      lcm = Lcm(lcm, n);
    }
    std::cout << lcm << std::endl;
  }

  return 0;
}
