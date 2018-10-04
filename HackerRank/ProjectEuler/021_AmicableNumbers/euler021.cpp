/**
 * Solution to the challenge #021 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler021
 * Created on: 17-09-2018
 * Last Modified: 17-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cmath>
#include <cstdint>
#include <iostream>
#include <vector>

const unsigned int MAX_N = 1e5;
std::vector<unsigned int> properDivisors;
std::vector<unsigned int> amicable;

unsigned int SumOfProperDivisors(unsigned int n)
{
  unsigned int sumOfProperDivisors = 1;
  for (int k = 2; k <= std::sqrt(n); ++k)
  {
    if (n % k == 0)
    {
      sumOfProperDivisors += k;
      if (k * k != n)
      {
        sumOfProperDivisors += (n / k);
      }
    }
  }

  return sumOfProperDivisors;
}

void FindAllAmicableNumbers()
{
  properDivisors.resize(MAX_N + 1);
  for (unsigned int n = 1; n <= MAX_N; ++n)
  {
    properDivisors[n] = SumOfProperDivisors(n);
  }

  for (unsigned int n = 1; n <=MAX_N; ++n)
  {
    if ((properDivisors[n] != n) &&
        (properDivisors[n] <= MAX_N) &&
        (properDivisors[properDivisors[n]] == n))
    {
      amicable.push_back(n);
    }
  }
}

int main()
{
  FindAllAmicableNumbers();

  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    unsigned int N;
    std::cin >> N;
    unsigned int sum = 0;
    for (auto it = amicable.begin(); ((it != amicable.end()) && (*it < N)); ++it)
    {
      sum += *it;
    }
    std::cout << sum << std::endl;
  }
  return 0;
}
