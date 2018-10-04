/**
 * Solution to the challenge #026 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler026
 * Created on: 27-09-2018
 * Last Modified: 27-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <cstdint>
#include <iostream>
#include <unordered_map>
#include <vector>

const int MAX_N = 10000;
std::vector<int> answers;

// Calculates the value of 1/n (long division) until the process terminates (length == 0)
// or a cycle is detected.
int CalculateLengthOfReciprocalCycle(int n)
{
  int step = 0;
  // Maps the dividend to the step (a.k.a. the index of the digit) it was computed in.
  std::unordered_map<int, int> previousDividends;
  int dividend = 1;
  int divisor = n;
  while (true)
  {
    while (dividend < divisor)
    {
      dividend *= 10;
      step += 1;
    }

    dividend %= divisor;

    if (dividend == 0)
    {
      return 0; // If the diviser divides the dividend wholly, the fraction is aperiodic.
    }
    if (previousDividends.find(dividend) != previousDividends.end())
    {
      return (step - previousDividends[dividend]);
    }

    previousDividends[dividend] = step;
  }
}

void CalculateAnswers()
{
  answers.clear();
  answers.resize(MAX_N + 1);
  answers[1] = 0; // 1 / 1 == 1, aperiodic.
  int longestRCycleSoFarVal = 0;
  int longestRCycleSoFarNum = 1;
  for (int i = 2; i <= MAX_N; ++i)
  {
    int rCycle = CalculateLengthOfReciprocalCycle(i);
    if (rCycle > longestRCycleSoFarVal)
    {
      longestRCycleSoFarVal = rCycle;
      longestRCycleSoFarNum = i;
    }
    answers[i] = longestRCycleSoFarNum;
  }
}

int main()
{
  CalculateAnswers();
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int n;
    std::cin >> n;

    // The task is to find the lowest denominator < n that produces the longest cycle.
    // Since the calculated answers assume <= n, n has to be lowered by one.
    std::cout << answers[n-1] << std::endl;
  }
  return 0;
}
