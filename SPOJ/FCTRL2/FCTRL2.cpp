#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> factorials;

void printFactorial(int n);
void prepareList();

int main(int argc, char *argv[]) {
  int noOfTests;
  int n;

  std::cin >> noOfTests;

  for(int currentTest = 0; currentTest < noOfTests; currentTest++) {
    std::cin >> n;
    printFactorial(n);
  }
}

void printFactorial(int n) {
  std::cout << factorials[n] << std::endl;
}

void prepareList() {
  factorials.resize(101);
  
}
