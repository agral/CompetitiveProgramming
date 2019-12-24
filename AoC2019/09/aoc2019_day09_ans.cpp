/**
 * Name: aoc2019_day09_ans
 * Description: Answers the AoC2019 day#09 challenge
 * Created on: 24.12.2019
 * Last modified: 24.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include "../Intcode/IntcodeInterpreter.hpp"

int main(int, char**)
{
  std::string line;
  while (std::getline(std::cin, line))
  {
    IntcodeInterpreter ic;
    auto arguments = std::vector<long long>{1};
    ic.loadProgram(line);
    ic.runProgram(arguments);

    auto output = ic.output();
    for (const auto& o : output)
    {
      std::cout << o << ",";
    }
    std::cout << "\n";

    IntcodeInterpreter ic2;
    auto arguments2 = std::vector<long long>{2};
    ic2.loadProgram(line);
    ic2.runProgram(arguments2);

    auto output2 = ic2.output();
    for (const auto& o : output2)
    {
      std::cout << o << ",";
    }
    std::cout << "\n";
  }

  return 0;
}
