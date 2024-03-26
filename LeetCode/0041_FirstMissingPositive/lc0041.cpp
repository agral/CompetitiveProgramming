#include <cassert>
#include <vector>
#include <cmath>

class Solution {
public:
  /**
   * By pigeonhole principle, an array of length SZ can not hold SZ+1 distinct positive integers -
   * at least one will be missing.
   * Approach:
   * 1st step: normalize the array by converting non-positive numbers to SZ+1.
   * Thus all the numbers in the array are now positive integers.
   * 2nd step: for each number k in vector <= SZ, set v[k-1] to negative of itself.
   * 3rd step: iterate over entire vector again. First index containing a positive number
   * corresponds to the smallest non-present number; ans = k+1.
   *
   * Time complexity: O(3n) = O(n), as each above step is O(n) individually.
   * Space complexity: O(1), a.k.a. constant; the input array is reused.
   */
  int firstMissingPositive(std::vector<int>& nums) {
    const int SZ = nums.size();
    // Step 1: normalization of the array
    for (int i = 0; i < SZ; ++i) {
      if (nums[i] <= 0) {
        nums[i] = SZ + 1;
      }
    }

    // Step 2: poor man's hashmap:
    for (int i = 0; i < SZ; ++i) {
      const int absnum = std::abs(nums[i]) - 1;
      if (absnum < SZ) {
        if (nums[absnum] > 0) {
          nums[absnum] = -nums[absnum];
        }
      }
    }

    // Step 3: final search for lowest non-present integer:
    for (int i = 0; i < SZ; ++i) {
      if (nums[i] > 0) {
        return i + 1;
      }
    }
    return SZ + 1;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {1, 2, 0};
    assert(s.firstMissingPositive(nums) == 3);
  }
  {
    std::vector<int> nums = {3, 4, -1, 1};
    assert(s.firstMissingPositive(nums) == 2);
  }
  {
    std::vector<int> nums = {7, 8, 9, 11, 12};
    assert(s.firstMissingPositive(nums) == 1);
  }
  {
    std::vector<int> nums = {1, 2, 3, 4};
    assert(s.firstMissingPositive(nums) == 5);
  }
}
