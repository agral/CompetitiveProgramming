#include <iostream>
#include <vector>
#include <cmath>

void printPrimes(unsigned long rangeStart, unsigned long rangeEnd);

int main(int argc, char *argv[]) {
  int noOfTests;
  std::cin >> noOfTests;
  for(int currentTest = 1; currentTest <= noOfTests; currentTest++) {
    unsigned long rs, re;
    std::cin >> rs >> re;
    printPrimes(rs, re);
  }
}

void printPrimes(unsigned long rangeStart, unsigned long rangeEnd) {
  if(rangeStart < 2) {
    rangeStart = 2;
  }
  unsigned long rangeLength = rangeEnd + 1 - rangeStart;
  std::vector<bool> sieve;
  sieve.resize(rangeLength, true);
  int threshold = static_cast<int>(sqrt(rangeEnd)) + 1;
  //if(threshold > 1000) threshold = 1000;

  //std::cout << "Range length: " << rangeLength << ", threshold: " 
  //  << threshold << std::endl;

  for(unsigned int i = 2; i <= threshold; ++i) {
      
    // set p to first multiplicity of i such that p >= rangeStart and p > i:
    unsigned int p = rangeStart;
    while(p % i != 0) {
      p += 1;
    }
    if(p == i) {
      p += i;
    }

    //std::cout << "i = " << i << ", first multiplicity: " << p << std::endl;
    while(p <= rangeEnd) {
      //std::cout << "Erasing: " << p 
        //<< "(position: " << p - rangeStart << ")" << std::endl;
      sieve.at(p - rangeStart) = false;
      p += i;
    }
  }
  for(unsigned int p = rangeStart; p <= rangeEnd; p++) {
    if(sieve.at(p - rangeStart) == true) {
      std::cout << p << std::endl;
    }
  }

  // list should be separated by empty line:
  std::cout << std::endl;
}
