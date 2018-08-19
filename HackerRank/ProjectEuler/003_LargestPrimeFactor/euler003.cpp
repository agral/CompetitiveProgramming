/**
 * Solution to the challenge #003 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler003
 * Created on: 19-08-2018
 * Last Modified: 19-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

/**
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of a given number N?
 *
 * First line of input contains T, the number of test cases,
 * followed by T lines each containing an integer N.
 * Constraints:
 *     1 <= T <= 10
 *     10 <= N <= 10^12
 *
 * For each test case, display the largest prime factor of N.
 */

#include <iostream>
#include <cstdint>
#include <cmath>
#include <vector>

// An unoptimized factorization - returns a list of all the prime factors:
std::vector<uint64_t> Factorize(uint64_t number)
{
  std::vector<uint64_t> factors;
  while (number % 2 == 0)
  {
    factors.push_back(2);
    number /= 2;
  }
  uint64_t currentFactor = 3;
  while (number > 1)
  {
    while (number % currentFactor == 0)
    {
      factors.push_back(currentFactor);
      number /= currentFactor;
    }
    currentFactor += 2;
  }
  return factors;
}

// An optimized version of Factorize which returns only the greatest factor instead of all the factors:
uint64_t CalculateBiggestPrimeFactor(uint64_t number)
{
  int biggestPrimeFactor = 2; // bpf is 1 for N=1, but in this challenge N >= 10 is guaranteed.
  while (number % 2 == 0)
  {
    number /= 2;
  }
  uint64_t currentFactor = 3;
  uint64_t ceilSqrtNumber = std::sqrt(number) + 1;
  while ((number > 1) && (currentFactor <= ceilSqrtNumber))
  {
    while (number % currentFactor == 0)
    {
      number /= currentFactor;
      biggestPrimeFactor = currentFactor;
    }
    currentFactor += 2;
  }

  // An optimization to the prime-finding algorithm has been made as follows:
  // Instead of checking whether the number is divided by every odd number <= N,
  // it is being checked only for every odd number <= ceil(sqrt(N)).
  // This lets the code omit a *lot* of unnecessary checks for big primes.
  // Because of that, if all the possible factors <= ceil(sqrt(N)) have been exhausted
  // and the value of number is still greater than 1, it *is* a prime number.
  return (number == 1) ? biggestPrimeFactor : number;
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    uint64_t N;
    std::cin >> N;
    std::cout << CalculateBiggestPrimeFactor(N) << std::endl;
  }
  return 0;
}
