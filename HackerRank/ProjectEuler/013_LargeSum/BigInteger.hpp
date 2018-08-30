#ifndef BIGINTEGER_HPP
#define BIGINTEGER_HPP

#include <vector>
#include <iostream>
#include <string>
#include <sstream>

class BigInteger
{
 public:
  BigInteger(int number = 0);
  BigInteger(const std::string& repr);

  void add(const BigInteger& other);
  std::string toString() const;

 private:
  std::vector<int> digits;
};

#endif // BIGINTEGER_HPP
