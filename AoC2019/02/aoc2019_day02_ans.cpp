/**
 * Name: aoc2019_day02_ans.cpp
 * Description: Answers the AoC2019 day#02 challenge
 * Created on: 02.12.2019
 * Last modified: 02.12.2019
 * Author: Adam Graliński (adam@gralin.ski)
 * License: CC0
 */

#include <vector>
#include <string>
#include <iostream>
#include <fstream>
#include <sstream>

const long long EXIT_CODE = 99;
const long long OPCODE_ADD = 1;
const long long OPCODE_MUL = 2;
const char CHAR_COMMA = ',';
const long long DESIRED_OUTPUT = 19690720;

std::vector<long long> loadProgram(const std::string& line)
{
  std::vector<long long> program;
  std::stringstream ss(line);
  for (long long i; ss >> i;)
  {
    program.push_back(i);
    if (ss.peek() == CHAR_COMMA)
    {
      ss.ignore();
    }
  }

  return program;
}

bool runProgram(std::vector<long long>& program)
{
  std::size_t ip = 0;
  while (program[ip] != EXIT_CODE)
  {
    if (program[ip] == OPCODE_ADD)
    {
      long long target = program[ip + 3];
      long long lhs = program[ip + 1];
      long long rhs = program[ip + 2];

      program[target] = program[lhs] + program[rhs];
    }
    else if (program[ip] == OPCODE_MUL)
    {
      long long target = program[ip + 3];
      long long lhs = program[ip + 1];
      long long rhs = program[ip + 2];

      program[target] = program[lhs] * program[rhs];
    }
    else
    {
      std::cout << "PANIC: unknown opcode [" << program[ip] << ", at position " << ip << "\n";
      return false;
    }

    ip = ip + 4;
  }
  return true;
}

void prettyPrint(const std::vector<long long>& program)
{
  for (std::size_t ip = 0; ip < program.size(); ++ip)
  {
    std::cout << program[ip] << "\t";
    if (ip % 4 == 3)
    {
      std::cout << "\n";
    }
  }
  std::cout << "\n";
}

void bruteForceSearchForOutput(const std::string& line, long long desired_output)
{
  for (long long noun = 0; noun < 100; ++noun)
  {
    for (long long verb = 0; verb < 100; ++verb)
    {
      auto program = loadProgram(line);
      program[1] = noun;
      program[2] = verb;
      runProgram(program);
      if (program[0] == desired_output)
      {
        std::cout << "\nDesired output (" << desired_output << ") obtained for NOUN=" << noun
                  << ", VERB=" << verb << "\n";
        return;
      }
    }
  }
}

int main()
{
  const std::string challengeInputName{"aoc2019_day02_input.txt"};
  //const std::string challengeInputName{"simple_programs.txt"};
  std::vector<long long> program;
  std::ifstream inputFile;
  inputFile.open(challengeInputName);
  if (inputFile.is_open())
  {
    std::string line;
    while (std::getline(inputFile, line))
    {
      auto program = loadProgram(line);
      program[1] = 12;
      program[2] = 2;
      runProgram(program);
      std::cout << "Value at position 0 after program terminates: " << program[0] << "\n";

      bruteForceSearchForOutput(line, DESIRED_OUTPUT);
    }
  }
  else
  {
    std::cerr << "Could not open file\n";
  }

  return 0;
}
