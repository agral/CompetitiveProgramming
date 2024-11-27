#include <iostream>
#include <algorithm> // sorting!

int main(int argc, char *argv[]) {
  int t;
  int n;
  unsigned short menHotness[1000];
  unsigned short womenHotness[1000];
  unsigned int totalHotness;

  std::cin >> t;
  for(int i = 0; i < t; i++) {
    std::cin >> n;
    totalHotness = 0;

    for(int j = 0; j < n; j++) {
      std::cin >> menHotness[j];
    }
    for(int j = 0; j < n; j++) {
      std::cin >> womenHotness[j];
    }

    std::sort(menHotness, menHotness + n);
    std::sort(womenHotness, womenHotness + n);
    
    for(int j = 0; j < n; j++) {
      totalHotness += menHotness[j] * womenHotness[j];
    }

    std::cout << totalHotness << std::endl;
  }

  return 0;
}
