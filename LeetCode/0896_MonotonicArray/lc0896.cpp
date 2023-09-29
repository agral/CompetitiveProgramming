#include <cassert>
#include <vector>

class Solution {
public:
  bool isMonotonic(std::vector<int>& nums) {
    bool isAscending = true;
    bool isDescending = true;

    for (int i = 1; (isAscending | isDescending) && i < nums.size(); i++) {
      if (nums[i - 1] > nums[i]) {
        isAscending = false;
      }
      else if (nums[i - 1] < nums[i]) {
        isDescending = false;
      }
    }
    return isAscending | isDescending;
  }
};

int main() {
  Solution s;
  std::vector<int> nums1 = {1, 2, 2, 3};
  assert(s.isMonotonic(nums1) == true);
  std::vector<int> nums2 = {6, 5, 4, 4};
  assert(s.isMonotonic(nums2) == true);
  std::vector<int> nums3 = {1, 3, 2};
  assert(s.isMonotonic(nums3) == false);
}
