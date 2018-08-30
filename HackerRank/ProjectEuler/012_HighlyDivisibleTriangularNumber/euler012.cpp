/**
 * Solution to the challenge #012 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler012
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <cstdint>
#include <vector>
#include <cmath>


const int MAX_N = 1000;
std::vector<int> answers;

int GetDividersCount(uint64_t number)
{
  uint64_t sqrtN = std::sqrt(number);
  int ans = 0;
  for (int k = 1; k <= sqrtN; ++k)
  {
    if (number % k == 0)
    {
      ans += 2;
    }
  }
  if (sqrtN * sqrtN == number)
  {
    ans -= 1;
  }

  return ans;
}

void CalculateAnswers()
{
  answers.resize(1, 0);
  uint64_t n = 0;
  uint64_t factorA, factorB;
  while (answers.size() <= MAX_N)
  {
    // Observation: n-th triangular number is equal to n * (n + 1) / 2 .
    // E.g. 9th triangular number == 9 * 10 / 2 == 45.
    // Instead of calculating divisors' count for a big triangular number,
    // performance is improved by calculating divisors of its factors and multiplying them.
    ++n;
    if (n % 2 == 0)
    {
      factorA = n / 2;
      factorB = n + 1;
    }
    else
    {
      factorA = n;
      factorB = (n + 1) / 2;
    }
    int dividersCount = GetDividersCount(factorA) * GetDividersCount(factorB);
    while (dividersCount > answers.size())
    {
      answers.push_back(factorA * factorB);
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
