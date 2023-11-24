#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  // Runtime complexity: O(sort) == O(nlogn)
  // Memory complexity: O(sort), i.e. whatever std::sort() uses.
  int maxCoins(std::vector<int>& piles) {
    std::sort(piles.begin(), piles.end());

    // After sorting, optimal strategy for the player is to choose lowest pile and two largest piles. Then,
    // * Alice goes first, gets the largest pile.
    // * The player goes second, gets the second-largest pile.
    // * Bob goes last, has no choice but to get the smallest pile.
    // This leads to the player getting the piles marked P from sorted vector, when picking optimally:
    // {B B B B B (...) B P A P A P A P A (...) P A}.
    // The first P is placed at length/3, i.e. right after length/3 Bob's piles.

    int ans = 0;
    for (int p = piles.size() / 3; p < piles.size(); p += 2) {
      ans += piles[p];
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> piles = {2, 4, 1, 2, 7, 8};
    assert(s.maxCoins(piles) == 7 + 2);
  }
  {
    std::vector<int> piles = {2, 4, 5};
    assert(s.maxCoins(piles) == 4);
  }
  {
    std::vector<int> piles = {9, 8, 7, 6, 5, 1, 2, 3, 4};
    assert(s.maxCoins(piles) == 8 + 6 + 4);
  }
}
