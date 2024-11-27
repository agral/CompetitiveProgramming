#include <iostream>
#include <string>

int main(int argc, char *argv[]) {
  
  short T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    std::string buf;
    std::cin >> buf;
    int p = 0;
    while(p < buf.length() / 2) {
      std::cout << buf[p];
      p += 2;
    }
    std::cout << std::endl;
  }
}
