/**
 * Solution to the "Birthday Cake Candles" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/birthday-cake-candles/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>

int main(int, char**)
{
  std::size_t tallestHeight{}, count{}, N, height;
  std::cin >> N;
  for (std::size_t n = 0; n < N; ++n) {
    std::cin >> height;
    if (height > tallestHeight) {
      tallestHeight = height;
      count = 1;
    }
    else if (height == tallestHeight) {
      count += 1;
    }
  }
  std::cout << count << "\n";

  return 0;
}
