#include <iostream>

bool willItStop(long long int n) {
  while( (n > 1) && (n % 2 == 0) ) {
    n /= 2;
  }
  return (n == 1);
}

int main(int argc, char *argv[]) {
  long long int N;
  std::cin >> N;
  if(willItStop(N)) {
    std::cout << "TAK";
  }
  else {
    std::cout << "NIE";
  }
  std::cout << std::endl;
}
