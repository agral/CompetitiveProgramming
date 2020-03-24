/**
 * Solution to the "Staircase" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/staircase/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>

int main(int, char**)
{
  int height;
  std::cin >> height;

  for (auto h = height - 1; h >= 0; --h) {
    for (auto s = 0; s < h; ++s) {
      std::cout << " ";
    }
    for (auto b = 0; b < height - h; ++b)
    {
      std::cout << "#";
    }
    std::cout << "\n";
  }

  return 0;
}
