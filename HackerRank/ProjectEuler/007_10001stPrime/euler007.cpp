/**
 * Solution to the challenge #007 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler007
 * Created on: 21-08-2018
 * Last Modified: 21-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <cstdint>
#include <vector>
#include <cmath>

const int MAX_N = 1e4;
const int SIEVE_SIZE = 1e6; // 10000th prime is 104729, SIEVE_SIZE has to be greater than that.
std::vector<bool> sieve;
std::vector<uint64_t> orderedPrimes;

void CalculateOrderedPrimes()
{
  sieve.resize(SIEVE_SIZE + 1, true);
  for (int k = 4; k < SIEVE_SIZE + 1; k += 2)
  {
    sieve[k] = false;
  }
  int divisor = 3;
  while (divisor <= std::sqrt(SIEVE_SIZE + 1) + 1)
  {
    for (int k = 2 * divisor; k < SIEVE_SIZE + 1; k += divisor)
    {
      sieve[k] = false;
    }
    divisor += 2;
  }
  orderedPrimes.resize(MAX_N + 1);
  int counter = 1;
  for (int k = 2; ((k < SIEVE_SIZE + 1) && (counter <= MAX_N)); ++k)
  {
    if (sieve[k])
    {
      orderedPrimes[counter] = k;
      ++counter;
    }
  }
}

int main()
{
  CalculateOrderedPrimes();
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::cout << orderedPrimes[N] << std::endl;
  }
  return 0;
}
