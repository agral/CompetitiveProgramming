#include <iostream>

int main(int argc, char* argv[]) {
  int first, second, third;
  int firstDiff, secondDiff;
  int firstRatio, secondRatio;
  int nextNumber;

  std::cin >> first >> second >> third;
  while(!((first == 0) && (second == 0) && (third == 0))) {
    firstDiff = second - first;
    secondDiff = third - second;
    if(firstDiff == secondDiff) {
      if(firstDiff == 0) {
        std::cout << "GP " << first << std::endl;
      }
      else {
        nextNumber = third + firstDiff;
        std::cout << "AP " << nextNumber << std::endl;
      }
    }
    else {
      firstRatio = second / first;
      secondRatio = third / second;
      if(firstRatio == secondRatio) {
        nextNumber = third * secondRatio;
        std::cout << "GP " << nextNumber << std::endl;
      }
      else {
        std::cout << "ERR" << std::endl;
      }
    }

    std::cin >> first >> second >> third;
  }
}
