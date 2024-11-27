#include <iostream>

int main(int argc, char *argv[]) {
  short count;
  std::cin >> count;
  std::cout << "W";
  for(int i = 0; i < count; i++) {
    std::cout << "o";
  }
  std::cout << "w" << std::endl;

  return 0;
}
