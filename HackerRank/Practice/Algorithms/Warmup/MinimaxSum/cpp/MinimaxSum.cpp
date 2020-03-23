/**
 * Solution to the "Mini-Max Sum" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/mini-max-sum/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <array>
#include <iostream>
#include <numeric>

const std::size_t SIZE = 5;

int main(int, char**)
{
  std::array<long long, SIZE> input;
  for (std::size_t i = 0; i < SIZE; ++i) {
    std::cin >> input[i];
  }
  std::sort(input.begin(), input.end());

  long long min = std::accumulate(input.begin(), input.end() - 1, 0LL);
  long long max = std::accumulate(input.begin() + 1, input.end(), 0LL);
  std::cout << min << " " << max << "\n";
  return 0;
}
