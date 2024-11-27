#include <iostream>
#include <cstdlib> // abs

int main(int argc, char *argv[]) {
  int n;
  long sum, totalDelta;
  int numbers[12000];

  std::cin >> n;
  while(n != -1) {
    sum = 0;
    for(int i = 0; i < n; i++) {
      std::cin >> numbers[i];
      sum += numbers[i];
    }
    if(sum % n != 0) {
      std::cout << "-1" << std::endl;
    }
    else {
      int average;
      average = sum / n;

      totalDelta = 0;
      for(int i = 0; i < n; i++) {
        totalDelta += abs(numbers[i] - average);
      }
      std::cout << totalDelta / 2 << std::endl;
    }

    std::cin >> n;
  }

  return 0;
}
