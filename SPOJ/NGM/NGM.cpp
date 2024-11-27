#include <iostream>

int main(int argc, char *argv[]) {
  int number;

  std::cin >> number;
  int remainder = number % 10;

  if(remainder == 0) {
    std::cout << "2" << std::endl;
  } else {
    std::cout << "1" << std::endl << remainder << std::endl;
  }

  return 0;
}
