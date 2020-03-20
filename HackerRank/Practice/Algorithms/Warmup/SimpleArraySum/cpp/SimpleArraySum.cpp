/**
 * Solution to the "SimpleArraySum" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/simple-array-sum/problem
 * Created on: 18.03.2020
 * Last modified: 18.03.2020
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

int main(int, char**)
{
  std::size_t N;
  std::cin >> N;
  std::vector<unsigned long long> v(N);
  unsigned long long sum = 0ULL;
  for (std::size_t n = 0; n < N; ++n)
  {
    std::cin >> v[n];
    sum += v[n];
  }
  std::cout << sum << "\n";
  return 0;
}
