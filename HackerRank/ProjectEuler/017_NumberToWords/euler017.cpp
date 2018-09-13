/**
 * Solution to the challenge #017 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler017
 * Created on: 04-09-2018
 * Last Modified: 04-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <cstdint>
#include <iostream>
#include <sstream>
#include <vector>

const std::string NUMBERS[] = {
  "Zero", "One", "Two", "Three", "Four",
  "Five", "Six", "Seven", "Eight", "Nine",
  "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
  "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"
};
const std::string TENS[] = {
  "", "", "Twenty", "Thirty", "Forty",
  "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"
};
const std::string POSTFIXES[] = {
  "Hundred", "Thousand", "Million", "Billion", "Trillion"
};

std::string JoinWords(std::vector<std::string>& v)
{
  std::stringstream ss;
  for (auto it = v.begin(); it != v.end(); ++it)
  {
    if (it != v.begin())
    {
      ss << " ";
    }
    ss << *it;
  }
  return ss.str();
}

std::string StringifyTriplet(uint64_t number)
{
  std::vector<std::string> components;
  uint64_t hundreds = number / 100ULL;
  if (hundreds > 0ULL)
  {
    components.push_back(NUMBERS[hundreds]);
    components.push_back(POSTFIXES[0]);
  }
  number = number % 100ULL;
  if (number >= 20)
  {
    uint64_t tens = number / 10ULL;
    components.push_back(TENS[tens]);
    number = number % 10ULL;
  }
  if (number > 0)
  {
    components.push_back(NUMBERS[number]);
  }

  return JoinWords(components);
}

std::string ToWords(uint64_t number)
{
  std::stringstream ss;
  if (number == 0ULL)
  {
     return NUMBERS[0];
  }

  std::vector<std::string> ans;
  size_t currentPostfixIndex = 0;
  while (number != 0ULL)
  {
    std::vector<std::string> converted;
    uint64_t triplet = number % 1000;
    if (triplet != 0ULL)
    {
      converted.push_back(StringifyTriplet(triplet));
      if (currentPostfixIndex >= 1)
      {
        converted.push_back(POSTFIXES[currentPostfixIndex]);
      }
      ans.insert(ans.begin(), JoinWords(converted));
    }
    number = number / 1000;
    currentPostfixIndex += 1;
  }
  return JoinWords(ans);
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    uint64_t number;
    std::cin >> number;
    std::cout << ToWords(number) << std::endl;
  }
  return 0;
}
