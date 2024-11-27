#include <iostream>

void printFrame(unsigned int height, unsigned int width) {
  for(int h = 1; h <= height; h++) {
    if((h == 1) || (h == height)) {
      for(int w = 1; w <= width; w++) {
        std::cout << '*';
      }
    }
    else {
      std::cout << '*';
      for(int w = 2; w < width; w++) {
        std::cout << '.';
      }
      if(width != 1) {
        std::cout << '*';
      }
    }
    std::cout << std::endl;
  }
  std::cout << std::endl;
}

int main(int argc, char *argv[]) {
  int T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    unsigned int l, c;
    std::cin >> l >> c;
    printFrame(l, c);
  }
}
