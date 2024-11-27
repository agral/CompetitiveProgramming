#include <iostream>

int main(int argc, char *argv[]) {
  int T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    unsigned long long N,R;
    std::cin >> N;
    R = N * (N + 2) * ((2 * N) + 1) / 8;
    std::cout << R << std::endl;
  }
  return 0;
}
