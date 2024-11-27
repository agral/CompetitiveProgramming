#include <iostream>

int main(int argc, char *argv[]) {
  int tCount;
  std::cin >> tCount;
  for(int t = 0; t < tCount; t++) {
    unsigned long long k, result;
    std::cin >> k;
    result = (k * 250ULL) - 58ULL;
    std::cout << result << std::endl;
  }
}
