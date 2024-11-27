#include <iostream>
#include <cmath>

bool isValidBeehive(long long cellsCount) {
  double r = std::sqrt( (4 * cellsCount - 1) / 3.0 );

  // if this result is a whole integer, then (and only then)
  // such a number of cells form a valid beehive.
  //
  // this checks if a double is actually a whole int:
  double intpart;
  return(std::modf(r, &intpart) == 0.0);
}

int main(int argc, char *argv[]) {
  long long c;
  std::cin >> c;

  while(c != -1) {
    if(isValidBeehive(c)) {
      std::cout << "Y";
    }
    else {
      std::cout << "N";
    }
    std::cout << std::endl;
    std::cin >> c;
  }

  return 0;
}

