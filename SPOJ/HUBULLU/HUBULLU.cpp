#include <iostream>

int main(int argc, char *argv[]) {
  int T;
  std::cin >> T;
  for(int t = 0; t < T; t++) {
    int noOfPieces, firstPlayer;
    std::cin >> noOfPieces >> firstPlayer;
    if(firstPlayer == 0) {
      std::cout << "Airborne";
    }
    else {
      std::cout << "Pagfloyd";
    }
    std::cout << " wins." << std::endl;
  }
  return 0;
}
