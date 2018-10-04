/**
 * Solution to the challenge #024 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler024
 * Created on: 24-09-2018
 * Last Modified: 24-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cstdint>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>

std::vector<int> ToFactoradic(uint64_t number, unsigned minDigits = 13)
{
  std::vector<int> digits{0}; // The 0!-indexed digit is always 0.
  for (uint64_t divider = 2; number != 0; ++divider)
  {
    digits.insert(digits.begin(), number % divider);
    number /= divider;
  }
  while (digits.size() < minDigits)
  {
    digits.insert(digits.begin(), 0);
  }

  return digits;
}

std::string Lehmer13(uint64_t number)
{
  // The task assumes that 1st permutation is the original string, but it is zero-indexed
  // in Lehmer code. Thus, a value of 1 is substracted from the original query:
  number -= 1;
  std::stringstream ss;

  std::vector<char> availableChars = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'};
  auto nFactoradic = ToFactoradic(number);
  for (auto& d : nFactoradic)
  {
    ss << availableChars.at(d);
    availableChars.erase(availableChars.begin() + d);
  }

  return ss.str();
}

int main()
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    uint64_t N;
    std::cin >> N;
    std::cout << Lehmer13(N) << std::endl;
  }
  return 0;
}
