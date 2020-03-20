/**
 * Solution to the "Compare the Triplets" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/compare-the-triplets/problem
 * Created on: 20.03.2020
 * Last modified: 20.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <array>

const std::size_t SIZE = 3;

int main(int, char**)
{
  int alice[SIZE];
  int scoreAlice{}, scoreBob{};

  for (std::size_t n = 0; n < SIZE; ++n) {
    std::cin >> alice[n];
  }
  for (std::size_t n = 0; n < SIZE; ++n) {
    int bob;
    std::cin >> bob;
    if (alice[n] > bob) {
      scoreAlice += 1;
    }
    else if (alice[n] < bob) {
      scoreBob += 1;
    }
  }
  std::cout << scoreAlice << " " << scoreBob << "\n";

  return 0;
}
