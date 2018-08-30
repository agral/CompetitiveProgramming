/**
 * Solution to the challenge #016 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler016
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>
#include "BigInteger.hpp"

const int MAX_POWER = 1e4;

std::vector<int> answers;

void CalculateAnswers()
{
  answers.resize(MAX_POWER + 1);
  BigInteger bi{1}; // That is the value of 2^0, unused.
  for (int i = 1; i <= MAX_POWER; ++i)
  {
    bi.doubleSelf();
    answers[i] = bi.sumOfDigits();
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
