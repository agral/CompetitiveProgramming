/**
 * Solution to the "Bigger is Greater" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/bigger-is-greater/problem
 * Created on: 02.09.2019
 * Last modified: 02.09.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
#include "impl.hpp"

int main(int, char**)
{
  std::string buf;
  std::getline(std::cin, buf);
  std::size_t T = std::stoi(buf);
  for (std::size_t t = 0; t < T; ++t)
  {
    std::string input;
    std::getline(std::cin, input);
    std::cout << getNextLexicographicPermutation(input) << std::endl;
  }
  return 0;
}
