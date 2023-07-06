#include <cassert>
#include <numeric>
#include <vector>

class Solution {
public:
  int pivotIndex(std::vector<int>& nums) {
    int leftSum {0};
    int rightSum = std::accumulate(nums.cbegin() + 1, nums.cend(), 0);
    if (rightSum == 0) {
      return 0;
    }

    for (std::size_t idx = 1; idx < nums.size(); idx++) {
      leftSum += nums[idx-1];
      rightSum -= nums[idx];
      if (leftSum == rightSum) {
        return idx;
      }
    }
    return -1;
  }
};

int main() {
  Solution s;
  std::vector<int> example_1 = {1, 7, 3, 6, 5, 6};
  assert(s.pivotIndex(example_1) == 3);

  std::vector<int> example_2 = {1, 2, 3};
  assert(s.pivotIndex(example_2) == -1);

  std::vector<int> example_3 = {2, 1, -1};
  assert(s.pivotIndex(example_3) == 0);

  std::vector<int> example_4 = {-3, 2, 1, 4};
  assert(s.pivotIndex(example_4) == 3);
}
