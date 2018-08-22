/**
 * Solution to the challenge #008 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler008
 * Created on: 22-08-2018
 * Last Modified: 22-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>
#include <string>

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N, K;
    std::cin >> N >> K;

    std::string input;
    std::cin >> input;

    // Converts the string to a vector of ints (representing individual digits):
    std::vector<int> digits;
    for (auto& c : input)
    {
      digits.push_back(c - '0');
    }

    // Calculates the maximum product:
    int maxProduct = 0;
    for (int n = 0; n <= N - K; ++n)
    {
      int product = 1;
      for (int k = 0; k < K; ++k)
      {
        product *= digits[n+k];
      }

      if (product > maxProduct)
      {
        maxProduct = product;
      }
    }

    // Prints out the answer:
    std::cout << maxProduct << std::endl;
  }
  return 0;
}
