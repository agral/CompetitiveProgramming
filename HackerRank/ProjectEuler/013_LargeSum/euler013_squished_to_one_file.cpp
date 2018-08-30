/**
 * Solution to the challenge #013 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler013
 * Created on: 30-08-2018
 * Last Modified: 30-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <sstream>
#include <string>
#include <vector>

const int DIGITS_LIMIT = 10;

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

  BigInteger(const std::string& repr)
  {
    for (auto& d : repr)
    {
      digits.push_back(static_cast<int>(d - '0'));
    }
  }

  void add(const BigInteger& other)
  {
    int carry = 0, itd = digits.size() - 1, iod = other.digits.size() - 1;
    for (; ((carry > 0) || (itd >= 0) || (iod >= 0)); --itd, --iod)
    {
      int td = (itd >= 0 ? digits[itd] : 0 ); // the digit of this BigInteger, or 0 if not available
      int od = (iod >= 0 ? other.digits[iod] : 0 ); // the digit of the other BigInteger, or 0 if not available
      int sum = carry + td + od;

      if (itd >= 0)
      {
        digits[itd] = sum % 10;
      }
      else
      {
        digits.insert(digits.begin(), sum % 10);
      }

      carry = sum / 10;
    }
  }

  std::string toString() const
  {
    std::stringstream ss;
    for (auto& d : digits)
    {
      ss << d;
    }
    return ss.str();
  }

 private:
  std::vector<int> digits;
};


int main()
{
  int N;
  std::cin >> N;
  BigInteger sum{0};
  for (int n = 0; n < N; ++n)
  {
    std::string line;
    std::cin >> line;
    BigInteger delta{line};
    sum.add(delta);
  }

  std::string ans = sum.toString();
  for (int i = 0; i < DIGITS_LIMIT; ++i)
  {
    std::cout << ans[i];
  }
  std::cout << std::endl;

  return 0;
}
