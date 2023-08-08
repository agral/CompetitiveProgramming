#include <vector>
//#include <iostream>
#include <cassert>

class Solution {
public:
  int search(std::vector<int>& nums, int target) {
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
    /// 3a. If there exists a part without a pivot and its start < target < end,
    /// continue searching there, removing pivot containing part from further consideration.
    /// 3b. If part without a pivot can't contain target, continue with part that has pivot.
    /// 4. Repeat steps 2-3 until range is reduced to one element.
    /// This has runtime complexity of O(logn), with auxillary memory space complexity of O(1)
    /// assuming that O(n) for the array is not counted - passed by reference.
    int range_left = 0;
    int range_right = nums.size() - 1;
    while (range_left < range_right) {
      int mid = range_left + (range_right - range_left) / 2;
      //std::cout << "search, l=" << range_left << ", r=" << range_right << ", m=" << mid << "\n";
      // 3a. Is left part of range pivot-free?
      //   --> note: if it is not, then right part of range is pivot-free.
      //             (there is only at most one pivot, so it can exist in at most one given half)
      if (nums[range_left] <= nums[mid]) {
        // Left part of range is pivot-free. Does it contain target?
        if ((nums[range_left] <= target) && (target <= nums[mid])) {
          // Only the pivot-free left half of range may contain target. Continue search there.
          range_right = mid;
        }
        else {
          // Pivot-free left half of range does NOT contain target.
          // Continue search in pivoted right half of range.
          range_left = mid + 1;
        }
      } else {
        // Right half of range is pivot-free. Does it contain target? (code similar to left-part).
        if ((nums[mid+1] <= target) && (target <= nums[range_right])) {
          // Only pivot-free right half of range can contain target. Continue search there.
          range_left = mid + 1;
        }
        else {
          // Only pivot-containing left half of range might contain target.
          // Continue search in pivoted left half of range.
          range_right = mid;
        }
      }
    }
    // range_left now either points to target, or target is not present in nums.
    //std::cout << "Search complete, idx=" << range_left << ", val=" << nums[range_left];
    //std::cout << ", found=" << (nums[range_left] == target ? "YES" : "NO") << "\n\n";
    return nums[range_left] == target ? range_left : -1;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {4, 5, 6, 7, 0, 1, 2};
  assert(s.search(testcase1, 0) == 4);
  std::vector<int> testcase2 = {4, 5, 6, 7, 0, 1, 2};
  assert(s.search(testcase2, 3) == -1);
  std::vector<int> testcase3 = {1};
  assert(s.search(testcase3, 0) == -1);
  assert(s.search(testcase3, 1) == 0);

  std::vector<int> testcase_wa1 = {5, 1, 3};
  assert(s.search(testcase_wa1, 1) == 1);
}
