#include <algorithm>
#include <cassert>
#include <vector>
#include <cmath>

class Solution {
public:
  int eliminateMaximum(std::vector<int>& dist, std::vector<int>& speed) {
    // First, calculate number of turns that each monster will need in order to reach the city.
    // It is guaranteed that each monster requires at least one turn to get there.
    // (1 <= dist[i], speed[i] <= 1e5)
    std::vector<int> numTurnsBeforeReachingCity(dist.size(), -1);
    for (int i = 0; i < dist.size(); i++) {
      numTurnsBeforeReachingCity[i] = std::ceil(static_cast<double>(dist[i]) / speed[i]);
    }

    // Sort the monsters in ascending order with regard to turns to reach city:
    std::sort(numTurnsBeforeReachingCity.begin(), numTurnsBeforeReachingCity.end());

    // Check how much monsters can be obliterated before any of them reach the city
    // or all are defeated. Assume that at time t=0, we kill the fastest one
    // (in terms of number of turns to reach the city). The player can continue killing
    // next monsters as long as their numTurnsBeforeReachingCity remains higher than
    // the current turn.
    int numKilled = 1;
    for (int k = 1; k < numTurnsBeforeReachingCity.size(); k++) {
      if (numTurnsBeforeReachingCity[k] <= k) {
        break; // can't kill that one in time, this monster reaches the city.
      }
      // otherwise, k'th monster is slain.
      numKilled++;
    }
    return numKilled;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> dist = {1, 3, 4};
    std::vector<int> speed = {1, 1, 1};
    assert(s.eliminateMaximum(dist, speed) == 3);
  }
  {
    std::vector<int> dist = {1, 1, 2, 3};
    std::vector<int> speed = {1, 1, 1, 1};
    assert(s.eliminateMaximum(dist, speed) == 1);
  }
  {
    std::vector<int> dist = {3, 2, 4};
    std::vector<int> speed = {5, 3, 2};
    assert(s.eliminateMaximum(dist, speed) == 1);
  }
  {
    std::vector<int> dist = {4, 3, 3, 3, 4};
    std::vector<int> speed = {1, 1, 1, 1, 4};
    assert(s.eliminateMaximum(dist, speed) == 3);
  }

}
