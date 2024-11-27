#include <iostream>

void printNumber(int x, int y);

int main(int argc, char *argv[]) {
  unsigned int n;
  unsigned int x, y;
  std::cin >> n;

  for(unsigned int i = 0; i < n; i++) {
    std::cin >> x >> y;
    printNumber(x, y);
  }

  return 0;
}

void printNumber(int x, int y) {
  unsigned int result;
  if(x == y) {
    if(x % 2 == 0) {
      result = x + y;
    }
    else { // for odd x == y: result = x+y-1
      result = x + y - 1;
    }
    std::cout << result << std::endl;
  }
  else if(x == y + 2) {
    if(x % 2 == 0) {
      result = x + y;
    }
    else {
      result = x + y - 1;
    }
    std::cout << result << std::endl;
  }
  else {
    std::cout << "No Number" << std::endl;
  }
}
