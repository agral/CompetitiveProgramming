#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int getWinner(std::vector<int>& arr, int k) {
    // No need to run the sim more times than array's size minus one.
    k = std::min(k, static_cast<int>(arr.size()));
    int score = 0;
    int pos = 1;
    while (score < k) {
      if (arr[pos] > arr[0]) {
        // move the new winner to 0th index. Reset the score.
        std::swap(arr[pos], arr[0]);
        score = 1;
      }
      else {
        score++;
      }
      // Move the challenger pointer to the right; circle back to 1st index if end of array is reached.
      pos++;
      if (pos == arr.size()) {
        pos = 1;
      }
    }
    return arr[0];
  }
};

int main() {
  Solution s;

  {
    std::vector<int> input1 = {2, 1, 3, 5, 4, 6, 7};
    assert(s.getWinner(input1, 2) == 5);
  }
  {
    std::vector<int> input2 = {3, 2, 1};
    assert(s.getWinner(input2, 10) == 3);
  }
  {
    std::vector<int> input_self1 = {7, 6, 5, 4, 3, 2, 1, 100};
    assert(s.getWinner(input_self1, 7) == 100);
    std::vector<int> input_self2 = {7, 6, 5, 4, 3, 2, 1, 100};
    assert(s.getWinner(input_self2, 1) == 7);
    std::vector<int> input_self3 = {7, 6, 5, 4, 3, 2, 1, 100};
    assert(s.getWinner(input_self3, 1'000'000'000) == 100);
  }
}
