/**
 * Solution to the "Plus Minus" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/plus-minus/problem
 * Created on: 22.03.2020
 * Last modified: 22.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <iomanip>

int main(int, char**)
{
  int N;
  std::cin >> N;
  int number, negative{}, zeroes{}, positive{};
  for (int n = 0; n < N; ++n) {
    std::cin >> number;
    if (number > 0) {
      positive += 1;
    }
    else if (number == 0) {
      zeroes += 1;
    }
    else {
      negative += 1;
    }
  }

  std::cout << std::setprecision(6) << std::fixed
            << static_cast<double>(positive) / N << "\n"
            << static_cast<double>(negative) / N << "\n"
            << static_cast<double>(zeroes) / N << "\n";
  return 0;
}
