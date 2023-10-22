#include <cassert>
#include <vector>

class Solution {
public:
  int maximumScore(std::vector<int>& nums, int k) {
    int l = k;
    int r = k;
    int ans = nums[k];
    int current_minimum = nums[k];
    while (l > 0 || r < nums.size() - 1) {
      int left_val = (l > 0) ? nums[l - 1] : 0;
      int right_val = (r < nums.size() - 1) ? nums[r + 1] : 0;
      if (left_val > right_val) {
        l--;
        current_minimum = std::min(current_minimum, left_val);
      }
      else {
        r++;
        current_minimum = std::min(current_minimum, right_val);
      }
      ans = std::max(ans, current_minimum * (r - l + 1));
    }

    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> nums1 = {1, 4, 3, 7, 4, 5};
  assert(s.maximumScore(nums1, 3) == 15);
  std::vector<int> nums2 = {5, 5, 4, 5, 4, 1, 1, 1};
  assert(s.maximumScore(nums2, 0) == 20);
}
