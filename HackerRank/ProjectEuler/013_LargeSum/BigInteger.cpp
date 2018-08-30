#include "BigInteger.hpp"

BigInteger::BigInteger(int number)
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

BigInteger::BigInteger(const std::string& repr)
{
  for (auto& d : repr)
  {
    digits.push_back(static_cast<int>(d - '0'));
  }
}

void BigInteger::add(const BigInteger& other)
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

std::string BigInteger::toString() const
{
  std::stringstream ss;
  for (auto& d : digits)
  {
    ss << d;
  }
  return ss.str();
}
