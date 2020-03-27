/**
 * Solution to the "Almost Sorted" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/almost-sorted/problem
 * Created on: 25.03.2020
 * Last modified: 27.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <vector>

const auto strYes = std::string("yes");
const auto strNo = std::string("no");
const auto strYesReverse = strYes + std::string("\nreverse ");
const auto strYesSwap = strYes + std::string("\nswap ");

int main(int, char**)
{
  std::size_t N;
  std::cin >> N;

  std::vector<int> input(N);
  for (std::size_t n = 0; n < N; ++n) {
    std::cin >> input[n];
  }

  if (std::is_sorted(input.begin(), input.end())) {
    std::cout << strYes << "\n";
  }
  else {
    std::size_t l = 0, r = N - 1;
    bool foundL = false, foundR = false;

    // Finds out the index of a first element of first strictly decreasing pair in the sequence (l):
    for (std::size_t i = 0; ((i < N - 2) && (!foundL)); ++i) {
      if (input[i] > input[i + 1]) {
        l = i;
        foundL = true;
      }
    }

    // Finds out the index of a second element of last strictly decreasing pair in the sequence (r):
    for (std::size_t i = N - 1; ((i >= 1) && (!foundR)); --i) {
      if (input[i - 1] > input[i]) {
        r = i;
        foundR = true;
      }
    }

    // Checks whether the sequence with l, r swapped is sorted.
    // If so, outputs "yes\nswap l r".
    // NOTE: this challenge requires outputting 1-indexed indices!
    std::swap(input[l], input[r]);
    if (std::is_sorted(input.begin(), input.end())) {
      std::cout << strYesSwap << l + 1 << " " << r + 1 << "\n";
    }
    else {
      // Swaps back, checks whether the sequence with [l,r] reversed is sorted.
      // If so, outputs "yes\nreverse l r".
      std::swap(input[l], input[r]);
      std::reverse(input.begin() + l, input.begin() + r + 1);
      if (std::is_sorted(input.begin(), input.end())) {
        std::cout << strYesReverse << l + 1 << " " << r + 1 << "\n";
      }
      else {
        std::cout << strNo << "\n";
      }
    }
  }

  return 0;
}
