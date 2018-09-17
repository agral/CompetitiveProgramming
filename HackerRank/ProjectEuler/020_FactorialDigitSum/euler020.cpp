/**
 * Solution to the challenge #020 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler020
 * Created on: 17-09-2018
 * Last Modified: 17-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <cstdint>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

const int MAX_N = 1000;

class BigInteger
{
 public:
  BigInteger(unsigned long long int number = 0);
  BigInteger(const std::string& repr);
  BigInteger(const BigInteger& other);

  BigInteger& operator+=(const BigInteger& other);
  BigInteger operator+(const BigInteger& other) const;
  BigInteger operator*(const BigInteger& other) const;

  int sumOfDigits() const;
  std::string toString() const;

 private:
  std::vector<int> digits;
};

BigInteger::BigInteger(unsigned long long int number)
{
  // Number 0 has to be handled separately - inserts a digit '0' and terminates:
  if (number == 0ULL)
  {
    return;
  }

  while (number != 0ULL)
  {
    digits.insert(digits.begin(), number % 10ULL);
    number /= 10ULL;
  }
}

BigInteger::BigInteger(const std::string& repr)
{
  for (auto& d : repr)
  {
    digits.push_back(static_cast<int>(d - '0'));
  }
}

BigInteger::BigInteger(const BigInteger& other)
{
  digits.resize(other.digits.size());
  for (std::size_t d = 0; d != other.digits.size(); ++d)
  {
    digits[d] = other.digits[d];
  }
}

BigInteger& BigInteger::operator+=(const BigInteger& other)
{
  int carry = 0, itd = digits.size() - 1, iod = other.digits.size() - 1;
  for (; ((carry > 0) || (itd >= 0) || (iod >= 0)); --itd, --iod)
  {
    int td = (itd >= 0 ? digits[itd] : 0); // the digit of this BigInteger, or 0 if not available
    int od = (iod >= 0 ? other.digits[iod] : 0); // the digit of the other BigInteger, or 0 if not available
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

  return *this;
}

BigInteger BigInteger::operator+(const BigInteger& other) const
{
  BigInteger bi{*this};
  bi += other;
  return bi;
}

/*
BigInteger& BigInteger::operator*=(const BigInteger& other)
{

}
*/

BigInteger BigInteger::operator*(const BigInteger& other) const
{
  BigInteger result{0};
  for (int iod = other.digits.size() - 1; iod >= 0; --iod)
  {
    int carry = 0;
    BigInteger component;
    for (int itd = digits.size() - 1; itd >= 0; --itd)
    {
      int mult = digits[itd] * other.digits[iod] + carry;
      component.digits.insert(component.digits.begin(), mult % 10);
      carry = mult / 10;
    }
    if (carry > 0)
    {
      component.digits.insert(component.digits.begin(), carry);
    }

    // Applies the multiplication of result by 10 for each digit except the last one:
    for (std::size_t k = iod; k < other.digits.size() - 1; ++k)
    {
      component.digits.push_back(0);
    }
    result += component;
  }

  return result;
}

int BigInteger::sumOfDigits() const
{
  int ans = 0;
  for (auto& d : digits)
  {
    ans += d;
  }

  return ans;
}

std::string BigInteger::toString() const
{
  if (digits.empty())
  {
    return "0";
  }

  std::stringstream ss;
  for (auto& d : digits)
  {
    ss << d;
  }
  return ss.str();
}


std::vector<int> answers;

void CalculateAnswers()
{
  answers.resize(MAX_N + 1);
  answers[0] = 1;
  answers[1] = 1;
  answers[2] = 2;
  BigInteger last{2};
  for (unsigned long long n = 3ULL; n <= MAX_N; ++n)
  {
    last = last * BigInteger{n};
    answers[n] = last.sumOfDigits();
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
