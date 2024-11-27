#include <iostream>
#include <string>

int main(int argc, char *args[]) {
  int T;
  std::cin >> T;
  for(auto t = 0; t < T; t++) {
    std::string pt;
    std::cin >> pt;
    std::string ct = "";
    for(char &c : pt) {
      ct += static_cast<char>( ((c - 'A' + 23) % 26) + 'A' );
      ct += static_cast<char>( ((c - 'A' + 4) % 26) + 'A' );
    }
    std::cout << ct << std::endl;
  }
}
