/**
 * Solution to the challenge #023 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler023
 * Created on: 21-09-2018
 * Last Modified: 21-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cmath>
#include <cstdint>
#include <iostream>
#include <vector>

const unsigned int THRESHOLD = 28123;
const unsigned int FIRST_ABUNDANT_NUMBER = 12; // 1 + 2 + 3 + 4 + 6 == 16 > 12

std::vector<unsigned int> abundantNumbers;
std::vector<bool> isSumOfTwoAbundantNumbers;

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

void FindAllAbundantNumbers()
{
  abundantNumbers.clear();
  for (unsigned int n = FIRST_ABUNDANT_NUMBER; n <= THRESHOLD; ++n)
  {
    if (SumOfProperDivisors(n) > n)
    {
      abundantNumbers.push_back(n);
    }
  }
}

void CalculateAnswers()
{
  FindAllAbundantNumbers();
  isSumOfTwoAbundantNumbers.clear();
  isSumOfTwoAbundantNumbers.resize(THRESHOLD + 1, false);

  // Marks the sums of two same abundant numbers:
  for (auto num = abundantNumbers.begin(); num != abundantNumbers.end(); ++num)
  {
    unsigned int sum = *num + *num;
    if (sum <= THRESHOLD)
    {
      isSumOfTwoAbundantNumbers[sum] = true;
    }
  }

  // Marks the sums of two distinct abundant numbers:
  for (auto first = abundantNumbers.begin(); first != abundantNumbers.end() - 1; ++first)
  {
    for (auto second = first + 1; second != abundantNumbers.end(); ++second)
    {
      unsigned int sum = *first + *second;
      if (sum <= THRESHOLD)
      {
        isSumOfTwoAbundantNumbers[sum] = true;
      }
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
    unsigned int N;
    std::cin >> N;
    bool answer = (N > THRESHOLD ? true : isSumOfTwoAbundantNumbers[N]);
    std::cout << (answer ? "YES" : "NO") << std::endl;
  }
  return 0;
}

