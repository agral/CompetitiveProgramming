#include <iostream>
#include <string>

int main(int argc, char *args[]) {
  int T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    std::string num;
    std::cin >> num;
    
    int sumOfCircles = 0;
    for(int p = 0; p < num.length(); p++) {
      char digit = num[p];
      if((digit == '6') || (digit == '9') || (digit == '0')) {
        sumOfCircles++;
      }
      else if(digit == '8') {
        sumOfCircles += 2;
      }
    }

    std::cout << sumOfCircles << std::endl;
  }
  return 0;
}
