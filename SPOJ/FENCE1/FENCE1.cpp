#include <iostream>
#include <iomanip>

int main(int argc, char *argv[]) {
  std::cout << std::setprecision(2) << std::fixed;
  unsigned short length;
  std::cin >> length;
  while(length > 0) {
    double area = (length * length) / 6.2831852;
    std::cout << area << std::endl;
    std::cin >> length;
  }
}
