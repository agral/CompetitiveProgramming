#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int maxOperations(std::vector<int>& nums, int k) {
    std::sort(nums.begin(), nums.end());
    int idx_left = 0;
    int idx_right = nums.size() - 1;
    int ans = 0;
    while (idx_left < idx_right) {
      int sum = nums[idx_left] + nums[idx_right];
      if (sum == k) {
        ans += 1;
        idx_left++;
        idx_right--;
      }
      else if (sum < k) {
        // Sum too small, move left pointer forward - this may increase the sum, never decreases it.
        idx_left++;
      }
      else {
        // Sum too high, move right pointer backward - this may decrease the sum, never increases it.
        idx_right--;
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 2, 3, 4};
  assert(s.maxOperations(example1, 5) == 2);

  std::vector<int> example2 = {3, 1, 3, 4, 3};
  assert(s.maxOperations(example1, 6) == 1);
}
