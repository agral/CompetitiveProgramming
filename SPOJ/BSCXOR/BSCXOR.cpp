#include <iostream>

int main(int argc, char* argv[]) {
  int first, second;
  bool f, s;

  std::cin >> first >> second;
  f = (first == 1);
  s = (second == 1);

  if(f ^ s) {
    std::cout << "1" << std::endl;
  } else {
    std::cout << "0" << std::endl;
  }

  return 0;
}
