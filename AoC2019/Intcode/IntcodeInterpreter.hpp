/**
 * Name: IntcodeInterpreter.hpp
 * Description: Defines a class representing an Intcode interpreter
 * (described in various AoC2019 challenges - https://adventofcode.com/2019/)
 * Created on: 07.12.2019
 * Last modified: 24.12.2019
 * Author: Adam Graliński (adam@gralin.ski)
 * License: CC0
 */

#ifndef INTCODEINTERPRETER_HPP
#define INTCODEINTERPRETER_HPP

#include <string>
#include <vector>

class IntcodeInterpreter
{
 public:
  static const auto OPCODE_ADD = 1LL;
  static const auto OPCODE_MUL = 2LL;
  static const auto OPCODE_INPUT = 3LL;
  static const auto OPCODE_OUTPUT = 4LL;
  static const auto OPCODE_JUMP_NOT_ZERO = 5LL;
  static const auto OPCODE_JUMP_ZERO = 6LL;
  static const auto OPCODE_LESS_THAN = 7LL;
  static const auto OPCODE_EQUALS = 8LL;
  static const auto OPCODE_ADJUST_RB = 9LL;
  static const auto OPCODE_HALT = 99LL;

 public:
  void loadProgram(const std::string& programCodeString);
  bool runProgram(std::vector<long long> inputList);
  void readArguments(int count, int multiArgModeNumber);
  void requestMemory(long long maxCellAddress);
  bool executeNextInstruction();
  std::vector<long long> output();

 private:
  std::vector<long long> program_;
  std::size_t ip_;
  long long relativeBase_;
  std::vector<long long> argumentAddresses_;
  std::vector<long long> input_;
  std::vector<long long> output_;
  bool shouldHaltNow_;
};

#endif // INTCODEINTERPRETER_HPP
