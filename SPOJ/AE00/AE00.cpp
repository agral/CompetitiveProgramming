#include <iostream>
#include <cmath>

int main(int argc, char *argv[]) {
  int squaresCount;
  std::cin >> squaresCount;

  int rectanglesCount = 0;
  int iLim = static_cast<int>(sqrt(squaresCount));
  for(int i = 1; i <= iLim; i++) {
    for(int k = i; k <= squaresCount; k++) {
      if(i * k <= squaresCount) {
        rectanglesCount++;
      }
    }
  }
  std::cout << rectanglesCount << std::endl;
}
