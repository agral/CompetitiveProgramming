#include <iostream>

void test();

int main(int argc, char *argv[]) {
  test();
}

void test() {
  int currentNumber;
  std::cin >> currentNumber;
  while (currentNumber != 42) {
    std::cout << currentNumber << std::endl;
    std::cin >> currentNumber;
  }
}
