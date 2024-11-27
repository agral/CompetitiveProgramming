#include <iostream>

const int answers[10][4] = {
  0, 0, 0, 0, // 0
  1, 1, 1, 1, // 1
  6, 2, 4, 8, // 2
  1, 3, 9, 7, // 3
  6, 4, 6, 4, // 4
  5, 5, 5, 5, // 5
  6, 6, 6, 6, // 6
  1, 7, 9, 3, // 7
  6, 8, 4, 2, // 8
  1, 9, 1, 9, // 9
};

int main(int argc, char *argv[]) {
  short int t; 
  std::cin >> t;

  for(short int j = 0; j < t; j++) {
    unsigned int base, exponent;
    std::cin >> base >> exponent;

    if(exponent == 0) {
      std::cout << "1" << std::endl;
    }
    else {
      base %= 10;
      int modular = exponent % 4;
      std::cout << answers[base][modular] << std::endl;
    }
  }

  return 0;
}
