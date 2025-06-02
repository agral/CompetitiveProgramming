#include <numeric>
#include <vector>
#include <cassert>

class Solution {
public:
  int candy(std::vector<int>& ratings) {
    std::vector<int> candies(ratings.size(), 1);

    // Sweep left-to-right:
    for (int i = 1; i < ratings.size(); i++) {
      if (ratings[i] > ratings[i-1]) {
        candies[i] = candies[i-1] + 1;
      }
    }

    // Sweep right-to-left:
    for (int i = candies.size() - 2; i >= 0; i--) {
      if (ratings[i] > ratings[i+1]) {
        candies[i] = std::max(candies[i], candies[i+1] + 1);
      }
    }

    // Return sum of all candies:
    return std::accumulate(candies.begin(), candies.end(), 0);
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 0, 2};
  assert(s.candy(example1) == 5); // min valid allocation: {2, 1, 2}
  
  std::vector<int> example2 = {1, 2, 2};
  assert(s.candy(example2) == 4); // {1, 2, 1}

  std::vector<int> test1 = {1, 2, 3, 42, 42, 42, 3, 2, 1}; // self-made test
  assert(s.candy(test1) == 21);

  std::vector<int> test2 = {6, 5, 4, 3, 4, 5, 6, 5, 4, 3, 2, 2}; // self-made test
  assert(s.candy(test2) == 31);

  std::vector<int> test3 = {1, 2, 3, 4, 3}; // self-made test.
  // verifies if right-to-left round overwrites candies with lower values (it should not!)
  assert(s.candy(test3) == 11);
}
