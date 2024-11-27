#include <iostream>

void printCardCount(int level) {
  if(level == 1) {
    std::cout << "2";
  }
  else {
    unsigned long long cardsCount = 2; // for 1st lev.
    unsigned long long delta = 2;
    for(int l = 2; l <= level; l++) {
      delta += 3;
      cardsCount += delta;
    }

    // include "N mod 1000007" part in the solution:
    cardsCount = cardsCount % 1000007;

    std::cout << cardsCount;
  }
  std::cout << std::endl;
}

int main(int argc, char *argv[]) {
  int T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    int level;
    std::cin >> level;
    printCardCount(level);
  }

  return 0;
}
