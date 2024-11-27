#include <iostream>

unsigned int howManyCards(double targetDistance);

int main(int argc, char *argv[]) {
  double targetDistance;
  std::cin >> targetDistance;
  
  while(targetDistance > 0.0) {
    std::cout << howManyCards(targetDistance) << " card(s)" << std::endl;
    std::cin >> targetDistance;
  }
}

unsigned int howManyCards(double targetDistance) {
  int cardsCount = 0;
  double deltaDistance = 0.0;
  double totalDistance = 0.0;
  while(totalDistance < targetDistance) {
    cardsCount++;
    deltaDistance = 1.0 / (cardsCount + 1);
    totalDistance += deltaDistance;
  }
  return cardsCount;
}
