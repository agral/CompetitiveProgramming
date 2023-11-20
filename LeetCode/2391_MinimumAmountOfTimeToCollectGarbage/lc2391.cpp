#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  int calc(const std::vector<std::string>& garbage, char garbageType, const std::vector<int> travelSum) {
    int furthestLocation = -1; // Note: initial location is always 0.
    int totalGarbage = 0;
    for (int i = 0; i < garbage.size(); i++) {
      const std::string& strGarbage = garbage[i];
      for (const char& ch: strGarbage) {
        if (ch == garbageType) {
          totalGarbage++;
          furthestLocation = i;
        }
      }
    }
    // Total time spent by this truck is:
    // * 1 minute each for collecting each piece of garbage
    // * total cost of getting to the farthest point.
    // Also note that it costs 0 to travel to 0th point (already there)
    return totalGarbage + (furthestLocation > 0 ? travelSum[furthestLocation - 1] : 0);
  }

  int garbageCollection(std::vector<std::string>& garbage, std::vector<int>& travel) {
    std::vector<int> travelSum(travel.size());
    travelSum[0] = travel[0];
    for (int i = 1; i < travel.size(); i++) {
      travelSum[i] = travelSum[i - 1] + travel[i];
    }

     return calc(garbage, 'M', travelSum) +
            calc(garbage, 'P', travelSum) +
            calc(garbage, 'G', travelSum);
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> garbage{{"G", "P", "GP", "GG"}};
    std::vector<int> travel{2, 4, 3};
    assert(s.garbageCollection(garbage, travel) == 21);
  }
  {
    std::vector<std::string> garbage{{"MMM", "PGM", "GP"}};
    std::vector<int> travel{3, 10};
    assert(s.garbageCollection(garbage, travel) == 37);
  }
}
