#include <iostream>

unsigned long getSquaresCount(unsigned int dim);

int main(int argc, char *argv[]) {
  unsigned int d;
  std::cin >> d;
  while(d != 0) {
    std::cout << getSquaresCount(d) << std::endl;
    std::cin >> d;
  }
}

unsigned long getSquaresCount(unsigned int dim) {
  unsigned long result = 0;
  for(int i = 1; i <= dim; i++) {
    result += (i * i);
  }
  return result;
}
