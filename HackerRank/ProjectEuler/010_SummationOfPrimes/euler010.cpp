/**
 * Solution to the challenge #010 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler010
 * Created on: 23-08-2018
 * Last Modified: 23-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cmath>
#include <cstdint>
#include <iostream>
#include <vector>

const int MAX_N = 1e6;

std::vector<bool> sieve;
std::vector<int> primes;
std::vector<uint64_t> answers; // note: sum of all primes up to a million does *well* exceed int32.

void CalculateAnswers()
{
  sieve.resize(MAX_N + 1, true);
  answers.resize(MAX_N + 1, 0);
  primes.clear();

  // An optimized version of a sieve is run - does nothing with multiples of 2.
  // 2 has to be manually added to the Primes vector.
  primes.push_back(2);
  for (int k = 3; k <= std::sqrt(MAX_N) + 1; k += 2)
  {
    for (int m = 2 * k; m <= MAX_N; m += k)
    {
      sieve[m] = false;
    }
  }

  for (int k = 3; k <= MAX_N; k += 2)
  {
    if (sieve[k])
    {
      primes.push_back(k);
    }
  }

  answers[1] = 0;
  int last_prime_index = -1;
  for (int k = 2; k <= MAX_N; ++k)
  {
    answers[k] = answers[k - 1];
    if ((last_prime_index + 1 < static_cast<int>(primes.size())) && (primes[last_prime_index + 1] <= k))
    {
      answers[k] += static_cast<uint64_t>(primes[last_prime_index + 1]);
      ++last_prime_index;
    }
  }
}

int main()
{
  CalculateAnswers();
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::cout << answers[N] << std::endl;
  }
  return 0;
}
