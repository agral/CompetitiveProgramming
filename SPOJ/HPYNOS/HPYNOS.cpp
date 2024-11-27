#include <iostream>
#include <vector>
#include <algorithm>

int main(int argc, char *argv[]) {
  int n;
  std::cin >> n;

  std::vector<int> stages;

  bool shouldRepeat = true;

  while( shouldRepeat ) {
    
    int newn = 0;
    while(n > 0) {
      int digit = n % 10;
      newn += (digit * digit);
      n = n / 10;
    }

    n = newn;

    if(std::find(stages.begin(), stages.end(), n) != stages.end()) {
      // stages contains newn - a loop has been found:
      std::cout << "-1" << std::endl;
      shouldRepeat = false;
    }
    else if(n == 1) {
      std::cout << stages.size() + 1 << std::endl;
      shouldRepeat = false;
    }

    stages.push_back(n);
  }

  return 0;
}
