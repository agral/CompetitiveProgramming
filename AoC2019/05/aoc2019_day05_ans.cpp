/**
 * Name: aoc2019_day05_ans
 * Description: Answers the AoC2019 day#05 challenge
 * Created on: 05.12.2019
 * Last modified: 07.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <map>

static std::string opcodeToName[] = {"-", "ADD", "MUL", "INPUT", "OUTPUT", "JNZ", "JZ", "LESSTHAN", "EQUALS"};

class IntcodeComputer
{
 public:
  static const int OPCODE_ADD = 1;
  static const int OPCODE_MUL = 2;
  static const int OPCODE_INPUT = 3;
  static const int OPCODE_OUTPUT = 4;
  static const int OPCODE_JUMP_IF_TRUE = 5;
  static const int OPCODE_JUMP_IF_FALSE = 6;
  static const int OPCODE_LESS_THAN = 7;
  static const int OPCODE_EQUALS = 8;
  static const int OPCODE_HALT = 99;
  static const char COMMA = ',';


  void loadProgram(const std::string& programCode)
  {
    std::stringstream pc(programCode);
    for (int code; pc >> code;)
    {
      program_.push_back(code);
      if (pc.peek() == COMMA)
      {
        pc.ignore();
      }
    }
  }

  void runProgram(std::vector<int> inputList)
  {
    ip_ = 0;
    input_ = inputList;
    output_.clear();
    shouldHaltNow_ = false;

    std::cout << "[IC] Starting to interpret program. Input list: { ";
    for (const auto& i : input_)
    {
      std::cout << i << " ";
    }
    std::cout << "}\n\n";

    while (!shouldHaltNow_)
    {
      /*
      for (int i = 0; i < static_cast<int>(program_.size()); ++i)
      {
        auto opcode = program_[i];
        if (opcode >= 100) { opcode = opcode % 100; }
        std::cout << "#" << i << "\t[" << program_[i] << "]";
        std::cout << "\t" << (((opcode >= 1) && (opcode <= 8)) ? opcodeToName[opcode] : "UNKNOWN OPCODE");
        if (ip_ == i)
        {
          std::cout << "\t<--IP--";
        }
        std::cout << "\n";
      }
      std::cout << "Press ENTER to execute...";
      std::string buf;
      std::getline(std::cin, buf);
      */
      executeIntcode();
    }
    for (std::size_t o = 0; o < output_.size(); ++o)
    {
      std::cout << "[IC] output #" << o << ": [" << output_[o] << "]\n";
    }
    std::cout << "[IC] Program has successfully halted.\n";
  }

  void readArgument(int count, int multiArgParameterNumber)
  {
    for (auto i = 0; i < count; ++i)
    {
      int mode = multiArgParameterNumber % 10;
      multiArgParameterNumber /= 10;
      argumentAddresses_.push_back(mode == 1 ? ip_ : program_[ip_]);
      ++ip_;
    }
  }

  void executeIntcode()
  {
    int instruction = program_[ip_];
    ++ip_;
    int opcode = instruction % 100;
    int multiArgParamNum = instruction / 100;

    if (opcode == OPCODE_ADD)
    {
      readArgument(3, multiArgParamNum);
      program_[argumentAddresses_[2]] = program_[argumentAddresses_[0]] + program_[argumentAddresses_[1]];
    }
    else if (opcode == OPCODE_MUL)
    {
      readArgument(3, multiArgParamNum);
      program_[argumentAddresses_[2]] = program_[argumentAddresses_[0]] * program_[argumentAddresses_[1]];
    }
    else if (opcode == OPCODE_INPUT)
    {
      readArgument(1, multiArgParamNum);
      program_[argumentAddresses_[0]] = input_[0];
      input_.erase(input_.begin());
    }
    else if (opcode == OPCODE_OUTPUT)
    {
      readArgument(1, multiArgParamNum);
      output_.push_back(program_[argumentAddresses_[0]]);
    }
    else if (opcode == OPCODE_JUMP_IF_TRUE)
    {
      readArgument(2, multiArgParamNum);
      if (program_[argumentAddresses_[0]] != 0)
      {
        ip_ = program_[argumentAddresses_[1]];
      }
    }
    else if (opcode == OPCODE_JUMP_IF_FALSE)
    {
      readArgument(2, multiArgParamNum);
      if (program_[argumentAddresses_[0]] == 0)
      {
        ip_ = program_[argumentAddresses_[1]];
      }
    }
    else if (opcode == OPCODE_LESS_THAN)
    {
      readArgument(3, multiArgParamNum);
      program_[argumentAddresses_[2]] = (program_[argumentAddresses_[0]] < program_[argumentAddresses_[1]] ? 1 : 0);
    }
    else if (opcode == OPCODE_EQUALS)
    {
      readArgument(3, multiArgParamNum);
      program_[argumentAddresses_[2]] = (program_[argumentAddresses_[0]] == program_[argumentAddresses_[1]] ? 1 : 0);
    }
    else if (opcode == OPCODE_HALT)
    {
      shouldHaltNow_ = true;
    }
    else
    {
      std::cerr << "[IC] PANIC: encountered unknown opcode <" << instruction << "> , at pos: " << ip_ - 1 << "\n";
      return;
    }

    argumentAddresses_.clear();
    std::cout << "Executed opcode " << opcode << ". IP is now " << ip_ << ".\n";
  }

 private:
  int ip_;
  bool shouldHaltNow_;
  std::vector<int> program_;
  std::vector<int> argumentAddresses_;
  std::vector<int> input_;
  std::vector<int> output_;
};

int main(int, char**)
{
  const auto challengeInputName = std::string{"aoc2019_day05_input.txt"};
  std::ifstream inputFile;
  inputFile.open(challengeInputName);
  if (inputFile.is_open())
  {
    std::string line;
    while (std::getline(inputFile, line))
    {
      IntcodeComputer computer;
      computer.loadProgram(line);
      // Part A:
      //computer.runProgram(std::vector<int>{1});

      // Part B:
      computer.runProgram(std::vector<int>{5});
    }
  }
  else
  {
    std::cerr << "Could not open file: " << challengeInputName << " for reading.\n";
  }

  return 0;
}
