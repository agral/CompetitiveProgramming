/**
 * Name: aoc2019_day04_ans
 * Description: Answers the AoC2019 day#01 challenge
 * Created on: 04.12.2019
 * Last modified: 04.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <iostream>
#include <cstdint>
#include <vector>

const auto RANGE_MIN = 265275;
const auto RANGE_MAX = 781584;
const auto DIGITS_COUNT = 6;

void printPossiblePasswordsInRangeCount(int min, int max)
{
  auto count = 0;
  auto count2 = 0;
  for (auto candidate = min; candidate <= max; ++candidate)
  {
    auto number = candidate;
    std::vector<int> digits(DIGITS_COUNT);
    for (auto digitIndex = DIGITS_COUNT - 1; digitIndex >= 0; --digitIndex)
    {
      digits[digitIndex] = number % 10;
      number /= 10;
    }
    if ((std::is_sorted(digits.begin(), digits.end())))
    {
      bool hasConsecutiveRepeatedDigits = false;
      for (auto digitIndex = 0; ((!hasConsecutiveRepeatedDigits) && (digitIndex < DIGITS_COUNT - 1)); ++digitIndex)
      {
        if (digits[digitIndex] == digits[digitIndex + 1])
        {
          hasConsecutiveRepeatedDigits = true;
          count += 1;
        }
      }

      bool hasConsecutiveRepeatedDigits2 = false;
      // Inserts two impossible digits to ignore border cases:
      digits.insert(digits.begin(), 99);
      digits.push_back(99);
      for (auto d = 1; ((!hasConsecutiveRepeatedDigits2) && (d < DIGITS_COUNT)); ++d)
      {
        if ((digits[d-1] != digits[d]) && (digits[d] == digits[d + 1]) && (digits[d + 1] != digits[d + 2]))
        {
          hasConsecutiveRepeatedDigits2 = true;
        }
      }

      if (hasConsecutiveRepeatedDigits2)
      {
        count2 += 1;
      }
    }
  }
  std::cout << "Valid passwords in range A: " << count << "\n";
  std::cout << "Valid passwords in range B: " << count2 << "\n";
}

int main(int, char**)
{
  printPossiblePasswordsInRangeCount(RANGE_MIN, RANGE_MAX);
  return 0;
}
