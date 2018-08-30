/**
 * Solution to the challenge #012 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler012
 * Created on: 29-08-2018
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
  uint64_t triangular = 0;
  uint64_t triangular_index = 0;
  while (answers.size() <= MAX_N)
  {
    // Calculates the next triangular number:
    ++triangular_index;
    triangular += triangular_index;
    int dividersCount = GetDividersCount(triangular);
    while (dividersCount > answers.size())
    {
      answers.push_back(triangular);
    }
  }
}

int main()
{
  CalculateAnswers();
  std::cout << "Ready." << std::endl;
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
