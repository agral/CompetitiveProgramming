// @author Adam "simba" Szczerbiak (szczerbiakadam@gmail.com) 
// the code for PP0502B problem on spoj.com
// Created: 2015-10-15 21:35:05

#include <iostream>

long long array[100];

int main(int argc, char *argv[]) {
  // read test cases count:
  unsigned short numOfTests; // guaranteed numOfTests <= 100
  std::cin >> numOfTests;

  for(unsigned short t = 0; t < numOfTests; t++) {

    unsigned short n; // guaranteed n <= 100
    std::cin >> n;

    for(unsigned short i = 0; i < n; i++) {
      std::cin >> array[i];
    }

    for(unsigned short i = 0; i < n; i++) {
      std::cout << array[n - i - 1] << " ";
    }
    std::cout << std::endl;
  }
}
