#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> sortByBits(std::vector<int>& arr) {
    std::sort(arr.begin(), arr.end(), [](const int& lhs, const int& rhs){
              auto bl = __builtin_popcount(lhs);
              auto br = __builtin_popcount(rhs);
              return (bl == br) ? lhs < rhs : bl < br;
    });
    return arr;
  }
};

int main() {
  Solution s;

  std::vector<int> example1 = {0, 1, 2, 3, 4, 5, 6, 7, 8};
  std::vector<int> expected_ans1 = {0, 1, 2, 4, 8, 3, 5, 6, 7};
  assert(s.sortByBits(example1) == expected_ans1);

  std::vector<int> example2 = {1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1};
  std::vector<int> expected_ans2 = {1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024};
  assert(s.sortByBits(example2) == expected_ans2);
}
