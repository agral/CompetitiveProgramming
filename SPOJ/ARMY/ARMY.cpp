#include <iostream>

int main(int argc, char *argv[]) {
  int t;
  std::cin >> t;
  for(int tc = 0; tc < t; tc++) {
    int countG, countM;
    int str;
    std::cin >> countG >> countM;

    int strongestG = 0, strongestM = 0;
    for(int g = 0; g < countG; g++) {
      std::cin >> str;
      if(str > strongestG) {
        strongestG = str;
      }
    }
    for(int m = 0; m < countM; m++) {
      std::cin >> str;
      if(str > strongestM) {
        strongestM = str;
      }
    }

    if(strongestM > strongestG) {
      std::cout << "Mecha";
    }
    std::cout << "Godzilla" << std::endl;
  }

  return 0;
}
