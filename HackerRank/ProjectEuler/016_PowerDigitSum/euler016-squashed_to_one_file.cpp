/**
 * Solution to the challenge #016 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler016
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <sstream>
#include <string>
#include <vector>

class BigInteger
{
 public:
  BigInteger(int number = 0)
  {
    // Number 0 has to be handled separately - inserts a digit '0' and terminates:
    if (number == 0)
    {
      digits.push_back(0);
      return;
    }

    while (number != 0)
    {
      digits.insert(digits.begin(), number % 10);
      number /= 10;
    }
  }

  void doubleSelf()
  {
    int carry = 0, id = digits.size() - 1;
    for(; ((carry > 0) || (id >= 0)); --id)
    {
      int td = (id >= 0 ? digits[id] : 0);
      int sum = carry + td + td;

      if (id >= 0)
      {
        this->digits[id] = sum % 10;
      }
      else
      {
        this->digits.insert(this->digits.begin(), sum % 10);
      }

      carry = sum / 10;
    }
  }

  int sumOfDigits() const
  {
    int ans = 0;
    for (auto& d : digits)
    {
      ans += d;
    }

    return ans;
  }

 private:
  std::vector<int> digits;
};


const int MAX_POWER = 1e4;

std::vector<int> answers;

void CalculateAnswers()
{
  answers.resize(MAX_POWER + 1);
  BigInteger bi{1}; // That is the value of 2^0, unused.
  for (int i = 1; i <= MAX_POWER; ++i)
  {
    bi.doubleSelf();
    answers[i] = bi.sumOfDigits();
  }
}

int main()
{
  CalculateAnswers();

  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    std::cout << answers[N] << std::endl;
  }
  return 0;
}
