#include <iostream>

unsigned long long factorial(unsigned int n) {
  if((n == 0) || (n == 1)) {
    return 1ULL;
  }

  unsigned long long result = 1;
  for(int k = 2; k <= n; k++) {
    result *= k;
  }
  return result;
}

unsigned int twoTo(unsigned int n) {
  unsigned int result = 1;
  for(int i = 1; i <= n; i++) {
    result *= 2;
  }
  return result;
}

int main(int argc, char* args[]) {
  int N;
  std::cin >> N;
  unsigned long long series = factorial(N) + twoTo(N) - N;
  std::cout << series << std::endl;
}
