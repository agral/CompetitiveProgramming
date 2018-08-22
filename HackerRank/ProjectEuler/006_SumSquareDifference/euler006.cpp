/**
 * Solution to the challenge #006 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler006
 * Created on: 21-08-2018
 * Last Modified: 21-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

/**
 * Find the absolute difference between the sum of the squares of the first N numbers
 * and the square of their sum.
 *
 * Limits:
 * 1 <= T <= 1e4   - the number of testcases
 * 1 <= N <= 1e4
 */

#include <cstdint>
#include <cstdlib>
#include <iostream>
#include <vector>

const std::size_t MAX_N  = 1e4;

struct Answer
{
  uint64_t n;
  uint64_t sum_n_squared;
  uint64_t squared_sum_n;
  uint64_t answer;
};

std::vector<Answer> answers;

void CalculateAllAnswers()
{
  answers.resize(MAX_N + 1);
  uint64_t sum_n_squared = 0;
  uint64_t sum_n = 0;
  for (uint64_t n = 1; n <= MAX_N; ++n)
  {
    sum_n_squared += (n * n);
    sum_n += n;
    answers[n].n = n;
    answers[n].sum_n_squared = sum_n_squared;
    answers[n].squared_sum_n = sum_n * sum_n;
    uint64_t temp = answers[n].sum_n_squared - answers[n].squared_sum_n;
    answers[n].answer = std::labs(temp);
  }
}

int main()
{
  CalculateAllAnswers();
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::cout << answers[N].answer << std::endl;
  }
  return 0;
}
