#include <iostream>
#include <string>
#include <vector>
#include <cstdint> // for uint8_t decl

class BigInteger {
  public:
    std::vector<uint8_t> digits;
    //bool isNegative;

    BigInteger();
    BigInteger(int n);

    char* toString() const;

    BigInteger operator+(const BigInteger& other) const;
};

int main(int argc, char *argv[]) {

}



BigInteger::BigInteger() {
  digits.push_back(0);
}

BigInteger::BigInteger(int n) {
  while(n != 0) {
    int lastDigit = n % 10;
    digits.insert(digits.begin(), lastDigit);
    n /= 10;
  }
}

char* BigInteger::toString() const {
  char buffer[digits.size() + 1];
  int pos = 0;
  for(auto i = digits.begin(); i != digits.end(); i++) {
    buffer[pos] = static_cast<char>(digits[i]);
    pos++;
  }
  buffer[pos] = '\0';
  return buffer;
}

BigInteger BigInteger::operator+(const BigInteger& other) const {

  for(auto i = digits.rbegin(); i != digits.rend(); ++it) {
    uint8_t 
  }
}
