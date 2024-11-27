#include <iostream>

void printChessboard(short l, short c) {
  bool isNextLineWhite = false;
  for(short line = 0; line < l; line++) {
    bool isNextFieldWhite = isNextLineWhite;
    for(short column = 0; column < c; column++) {
      if(isNextFieldWhite) {
        std::cout << '.';
      }
      else {
        std::cout << '*';
      }
      isNextFieldWhite = !isNextFieldWhite;
    }
    isNextLineWhite = !isNextLineWhite;
    std::cout << std::endl;
  }
}

int main(int argc, char *argv[]) {
  short T;
  std::cin >> T;
  for(short t = 0; t < T; t++) {
    short l, c;
    std::cin >> l >> c;
    printChessboard(l, c);
    std::cout << std::endl;
  }
}
