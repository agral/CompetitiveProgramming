#include <iostream>

void calculate(unsigned long n);

int main(int argc, char *argv[]) {
  int noOfTests;
  unsigned long n;

  std::cin >> noOfTests;
  for(int currentTest = 0; currentTest < noOfTests; currentTest++) {
    std::cin >> n;
    calculate(n);
  }
}

void calculate(unsigned long n) {
  unsigned long sumOfZeroes = 0;
  unsigned long deltaSumOfZeroes = 0;
  unsigned long currentDenominator = 5;
  int k = 1;
  deltaSumOfZeroes = n / currentDenominator;
  while(deltaSumOfZeroes > 0) {
    sumOfZeroes += deltaSumOfZeroes;
    k++;
    currentDenominator *= 5;
    deltaSumOfZeroes = n / currentDenominator;
  }

  std::cout << sumOfZeroes << std::endl;

}
