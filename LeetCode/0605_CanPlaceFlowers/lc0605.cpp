#include <cassert>
#include <vector>

class Solution {
public:
  bool canPlaceFlowers(std::vector<int>& flowerbed, int n) {
    // insert range guards so that corner cases become standard cases:
    flowerbed.insert(flowerbed.begin(), 0);
    flowerbed.push_back(0);

    int availablePlots = 0;
    for (std::size_t i = 1; (availablePlots < n && i < flowerbed.size() - 1); i++) {
      if (flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0) {
        availablePlots++;
        i += 1; // jump one additional spot which is guaranteed to be unpotable
      }
    }

    return availablePlots >= n;
  }
};

int main() {
  Solution s;
  std::vector<int> flowerbed1 = {1, 0, 0, 0, 1};
  assert(s.canPlaceFlowers(flowerbed1, 1) == true);
  assert(s.canPlaceFlowers(flowerbed1, 2) == false);

  std::vector<int> flowerbed2 = {1, 0, 1, 0, 0, 1, 0, 0, 0, 1};
  assert(s.canPlaceFlowers(flowerbed2, 1) == true);
  assert(s.canPlaceFlowers(flowerbed2, 2) == false);

  std::vector<int> flowerbed3 = {0, 0, 0, 0, 0, 0};
  assert(s.canPlaceFlowers(flowerbed3, 3) == true);
  assert(s.canPlaceFlowers(flowerbed3, 4) == true);

  std::vector<int> flowerbed4 = {0, 0, 0, 0, 0, 0, 0};
  assert(s.canPlaceFlowers(flowerbed4, 4) == true);
  assert(s.canPlaceFlowers(flowerbed4, 5) == true);
}
