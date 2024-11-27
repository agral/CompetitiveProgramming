#include <iostream>
#include <algorithm>

int friendsOffers[2000];
int friendsCount;
int howManyStampsNeeded;

int main(int argc, char *argv[]) {
  int testLimit;
  std::cin >> testLimit;

  for(int test = 1; test <= testLimit; test++) {
    std::cin >> howManyStampsNeeded >> friendsCount;
    for(int f = 0; f < friendsCount; f++) {
      std::cin >> friendsOffers[f];
    }

    std::sort(friendsOffers, friendsOffers + friendsCount,
        [](int a, int b){return b < a;} );

    std::cout << "Scenario #" << test << ":" << std::endl;
    int sumOfStamps = 0;
    int currentFriend = 0;
    while((sumOfStamps < howManyStampsNeeded) && 
        (currentFriend < friendsCount)) {
      sumOfStamps += friendsOffers[currentFriend];
      currentFriend++;
    }

    if(sumOfStamps < howManyStampsNeeded) {
      std::cout << "impossible" << std::endl;
    }
    else {
      std::cout << currentFriend << std::endl;
    }

    std::cout << std::endl;
  }
}
