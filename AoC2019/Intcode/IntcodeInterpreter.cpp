/**
 * Name: IntcodeInterpreter.cpp
 * Description: Implements a class representing an Intcode interpreter
 * (described in various AoC2019 challenges - https://adventofcode.com/2019/)
 * Created on: 07.12.2019
 * Last modified: 24.12.2019
 * Author: Adam Graliński (adam@gralin.ski)
 * License: CC0
 */

#include "IntcodeInterpreter.hpp"

#include <fstream>
#include <sstream>
#include <iostream>

const char COMMA = ',';

void IntcodeInterpreter::loadProgram(const std::string& programCodeString)
{
  std::stringstream pcs(programCodeString);
  for (long long number; pcs >> number;)
  {
    program_.push_back(number);
    if (pcs.peek() == COMMA)
    {
      pcs.ignore();
    }
  }
}

bool IntcodeInterpreter::runProgram(std::vector<long long> inputList)
{
  ip_ = 0;
  relativeBase_ = 0;
  input_ = inputList;
  output_.clear();
  shouldHaltNow_ = false;

  while (!shouldHaltNow_)
  {
    if (!executeNextInstruction())
    {
      std::cerr << "[IC] Program has been killed.\n";
      return false; // execution was not successful - unknown opcode passed as Intcode instruction.
    }
  }

  return true; // execution was successful - program halted normally.
}

void IntcodeInterpreter::readArguments(int count, int multiArgModeNumber)
{
  for (auto i = 0; i < count; ++i)
  {
    auto mode = multiArgModeNumber % 10;
    multiArgModeNumber /= 10;
    if (mode == 0)
    {
      // Normal mode
      argumentAddresses_.push_back(program_[ip_]);
    }
    else if (mode == 1)
    {
      argumentAddresses_.push_back(ip_);
    }
    else if (mode == 2)
    {
      argumentAddresses_.push_back(relativeBase_ + program_[ip_]);
    }
    else
    {
      std::cerr << "[IC] Unknown mode: " << mode << "\n";
      throw -1;
    }
    ++ip_;
  }
}

void IntcodeInterpreter::requestMemory(long long maxCellAddress)
{
  if (program_.size() < maxCellAddress + 1)
  {
    std::cout << "[IP] Growing program memory to accomodate " << maxCellAddress << " nodes.\n";
    program_.resize(maxCellAddress + 1);
  }
}

bool IntcodeInterpreter::executeNextInstruction()
{
  auto instruction = program_[ip_];
  ++ip_;
  auto opcode = instruction % 100;
  auto multiArgModeNumber = instruction / 100;

  //std::cout << "Executing opcode: " << opcode << ", ip: " << ip_-1 << "\n";

  if (opcode == OPCODE_ADD)
  {
    readArguments(3, multiArgModeNumber);
    requestMemory(argumentAddresses_[2]);
    program_[argumentAddresses_[2]] = program_[argumentAddresses_[0]] + program_[argumentAddresses_[1]];
  }
  else if (opcode == OPCODE_MUL)
  {
    readArguments(3, multiArgModeNumber);
    requestMemory(argumentAddresses_[2]);
    program_[argumentAddresses_[2]] = program_[argumentAddresses_[0]] * program_[argumentAddresses_[1]];
  }
  else if (opcode == OPCODE_INPUT)
  {
    if (!input_.empty())
    {
      readArguments(1, multiArgModeNumber);
      requestMemory(argumentAddresses_[0]);
      program_[argumentAddresses_[0]] = input_[0];
      input_.erase(input_.begin());
    }
    else
    {
      std::cerr << "[IC] PANIC: input needed, but all the input is already exhausted.\n";
      return false;
    }
  }
  else if (opcode == OPCODE_OUTPUT)
  {
    readArguments(1, multiArgModeNumber);
    output_.push_back(program_[argumentAddresses_[0]]);
  }
  else if (opcode == OPCODE_JUMP_NOT_ZERO)
  {
    readArguments(2, multiArgModeNumber);
    if (program_[argumentAddresses_[0]] != 0)
    {
      ip_ = program_[argumentAddresses_[1]];
    }
  }
  else if (opcode == OPCODE_JUMP_ZERO)
  {
    readArguments(2, multiArgModeNumber);
    if (program_[argumentAddresses_[0]] == 0)
    {
      ip_ = program_[argumentAddresses_[1]];
    }
  }
  else if (opcode == OPCODE_LESS_THAN)
  {
    readArguments(3, multiArgModeNumber);
    requestMemory(argumentAddresses_[2]);
    program_[argumentAddresses_[2]] = (program_[argumentAddresses_[0]] < program_[argumentAddresses_[1]] ? 1 : 0);
  }
  else if (opcode == OPCODE_EQUALS)
  {
    readArguments(3, multiArgModeNumber);
    requestMemory(argumentAddresses_[2]);
    program_[argumentAddresses_[2]] = (program_[argumentAddresses_[0]] == program_[argumentAddresses_[1]] ? 1 : 0);
  }
  else if (opcode == OPCODE_ADJUST_RB)
  {
    readArguments(1, multiArgModeNumber);
    relativeBase_ += program_[argumentAddresses_[0]];
  }
  else if (opcode == OPCODE_HALT)
  {
    shouldHaltNow_ = true;
  }
  else
  {
    std::cerr << "[IC] PANIC: encountered unknown instruction <" << instruction << ">, at pos: " << ip_ - 1 << ".\n";
    return false;
  }
  argumentAddresses_.clear();

  return true;
}

std::vector<long long> IntcodeInterpreter::output()
{
  return output_;
}
