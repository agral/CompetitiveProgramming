#include <cassert>
#include <vector>
#include <iostream>

class Solution {
public:
  void moveZeroes(std::vector<int>& nums) {
    int idx_nonzero = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] != 0) {
        std::swap(nums[i], nums[idx_nonzero]);
        idx_nonzero += 1;
      }
    }
  }
};

int main() {
  Solution s{};
  std::vector<int> test1 = {0, 1, 0, 3, 12};
  s.moveZeroes(test1);
  assert(test1 == std::vector<int>({1, 3, 12, 0, 0}));

  std::vector<int> test2 = {0};
  s.moveZeroes(test2);
  assert(test2 == std::vector<int>({0}));
}
