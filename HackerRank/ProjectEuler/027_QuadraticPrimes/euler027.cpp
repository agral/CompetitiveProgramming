/**
 * Solution to the challenge #027 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler027
 * Created on: 05-10-2018
 * Last Modified: 05-10-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <cmath>
#include <cstdint>
#include <iostream>
#include <unordered_set>
#include <utility>
#include <vector>

const int N_MIN = 42;
const int N_MAX = 2000;
const int HUGE_LIMIT = 1e6;

std::unordered_set<uint64_t> sPrimes;
std::vector<std::pair<int, int>> answers;

inline bool IsPrime(uint64_t number) { return sPrimes.find(number) != sPrimes.end(); }

void CalculatePrimesSet(uint64_t limit)
{
  sPrimes.clear();
  sPrimes.insert(1); // Has to be added manually, since the rest of the function works in range <2:limit>.
  std::vector<bool> sieve(limit + 1, true);

  uint64_t sqrtLimit = std::sqrt(limit);
  for (uint64_t k = 2; k <= sqrtLimit; ++k)
  {
    for (uint64_t m = 2 * k; m <= limit; m += k)
    {
      sieve[m] = false;
    }
  }

  for (uint64_t k = 2; k <= limit; ++k)
  {
    if (sieve[k])
    {
      sPrimes.insert(k);
    }
  }
}

int GetConsecutivePrimesCount(int a, int b)
{
  if (!IsPrime(b))
  {
    // "n^2 + an + b" reduces to "b" for n=0. If b is not a prime number,
    // the sequence generates no consecutive prime numbers.
    return 0;
  }
  if (!IsPrime(a + b + 1))
  {
    // "n^2 + an + b" reduces to "1+a+b" for n=1. If a+b+1 is not a prime number,
    // the sequence terminates here - returning just one prime for n=0.
    return 1;
  }
  int counter = 2;
  while (IsPrime((counter * counter) + (a * counter) + b))
  {
    counter += 1;
  }

  return counter - 1;
}

void CalculateAnswers()
{
  answers.clear();
  answers.resize(N_MAX + 1);
  answers[42] = std::make_pair(-1, 41);
  int mostPrimesGeneratedSoFar = 40;
  int primesCount;
  for (int n = 43; n <= N_MAX; ++n)
  {
    // The minimum answer for any N is the value obtained for N-1 (it can't be worse than that):
    answers[n] = answers[n-1];

    for (int i = -n; i <= n; ++i)
    {
      // Note: b has to be positive, since for n=0 the generated number has to be greater than zero.
      // This rules out calling GetConsecutivePrimesCount(i, -n), leaving us with
      // (i, n), (-n, i) and (n, i) pairs to check (6N checks in total for each checked N)

      primesCount = GetConsecutivePrimesCount(i, n);
      if (primesCount > mostPrimesGeneratedSoFar)
      {
        answers[n] = std::make_pair(i, n);
        mostPrimesGeneratedSoFar = primesCount;
      }

      primesCount = GetConsecutivePrimesCount(-n, i);
      if (primesCount > mostPrimesGeneratedSoFar)
      {
        answers[n] = std::make_pair(-n, i);
        mostPrimesGeneratedSoFar = primesCount;
      }

      primesCount = GetConsecutivePrimesCount(n, i);
      if (primesCount > mostPrimesGeneratedSoFar)
      {
        answers[n] = std::make_pair(n, i);
        mostPrimesGeneratedSoFar = primesCount;
      }
    }
  }
}

int main()
{
  CalculatePrimesSet(HUGE_LIMIT);
  CalculateAnswers();
  int N;
  std::cin >> N;
  std::cout << answers[N].first << " " << answers[N].second << std::endl;
  return 0;
}
