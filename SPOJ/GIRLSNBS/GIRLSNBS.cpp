#include <iostream>
#include <cmath>
#include <algorithm>

int main(int argc, char *argv[]) {
  int g;
  int b;
  std::cin >> g >> b;
  while((g != -1) && (b != -1)) {
    if(g == 0) {
      if(b == 0) {
        std::cout << "0" << std::endl;
      }
      else {
        std::cout << b << std::endl;
      }
    }
    else if(b == 0) {
      std::cout << g << std::endl;
    }
    else {
      int res = ceil( static_cast<double>(std::max(g, b)) / (std::min(g, b) + 1) );
      std::cout << res << std::endl;
    }

    std::cin >> g >> b;
  }
}
