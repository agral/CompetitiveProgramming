/**
 * Solution to the "A Very Big Sum" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/a-very-big-sum/problem
 * Created on: 20.03.2020
 * Last modified: 20.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>

int main(int, char**)
{
  unsigned long long sum{};
  std::size_t N;
  std::cin >> N;
  for (std::size_t n = 0; n < N; ++n)
  {
    unsigned long long part;
    std::cin >> part;
    sum += part;
  }
  std::cout << sum << "\n";
  return 0;
}
