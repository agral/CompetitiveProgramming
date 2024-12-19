#include <vector>
#include <cassert>

class Solution {
public:
  bool search(std::vector<int>& nums, int target) {
    /// The array is almost sorted: it's in ascending order, except for a pivot point
    /// where it "loops" back from max to min value and continues ascending.
    /// Also the pivot can be in the beginning, in which case the array is simply sorted.
    /// Approach: perform a binary search. Each range cut divides it in two subranges.
    /// One of these must be without a pivot --> simply sorted in ascending order.
    /// The other one might contain pivot. This can be determined by examining subrange limits.
    /// If subrange_start < subrange_end, pivot is not contained within.
    /// If subrange_start > subrange_end, piot must be contained within.
    /// So, combined approach:
    /// 1. Start with full array range, from 0 to length-1.
    /// 2. Divide range in half. Determine if each part contains pivot.
    /// 3a. If both left, mid, and right indices contain same value, it can not be determined
    /// which part contains the pivot. move left and right one index closer.
    /// 3b. If there exists a part without a pivot and its start < target < end,
    /// continue searching there, removing pivot containing part from further consideration.
    /// 3c. If part without a pivot can't contain target, continue with part that has pivot.
    /// 4. Repeat steps 2-3 until range is no longer valid.
    /// This has runtime complexity of O(logn) in average case; O(n) for pessimistic case.
    /// Space complexity: O(1) for left, right and mid indices.
    int range_left = 0;
    int range_right = nums.size() - 1;
    while (range_left <= range_right) {
      int mid = range_left + (range_right - range_left) / 2;
      if (nums[mid] == target) {
        return true;
      }
      // 3a: Shrink range iff pivot part can not be determined.
      if ((nums[range_left] == nums[mid]) && (nums[mid] == nums[range_right])) {
        ++range_left;
        --range_right;
      }
      // 3b. Is left part of range pivot-free?
      //   --> note: if it is not, then right part of range is pivot-free.
      //             (there is only at most one pivot, so it can exist in at most one given half)
      else if (nums[range_left] <= nums[mid]) {
        // Left part of range is pivot-free. Does it contain target?
        if ((nums[range_left] <= target) && (target < nums[mid])) {
          // Only the pivot-free left half of range may contain target. Continue search there.
          range_right = mid - 1;
        }
        else {
          // Pivot-free left half of range does NOT contain target.
          // Continue search in pivoted right half of range.
          range_left = mid + 1;
        }
      } else {
        // Right half of range is pivot-free. Does it contain target? (code similar to left-part).
        if ((nums[mid] < target) && (target <= nums[range_right])) {
          // Only pivot-free right half of range can contain target. Continue search there.
          range_left = mid + 1;
        }
        else {
          // Only pivot-containing left half of range might contain target.
          // Continue search in pivoted left half of range.
          range_right = mid - 1;
        }
      }
    }
    // at this point left >= right; target is guaranteed to not be present in the input array.
    return false;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {2, 5, 6, 0, 0, 1, 2};
  assert(s.search(testcase1, 0) == true);
  assert(s.search(testcase1, 3) == false);

  std::vector<int> testcase_wa1 = {1, 0, 1, 1, 1};
  assert(s.search(testcase_wa1, 0) == true);

  std::vector<int> testcase_wa2 = {3, 1};
  assert(s.search(testcase_wa2, 1) == true);
}
