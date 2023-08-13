#include <vector>
#include <cassert>
#include <iostream>

class Solution {
public:
  bool searchMatrix(std::vector<std::vector<int>>& matrix, int target) {
    std::size_t num_rows = matrix.size();
    std::size_t num_cols = matrix[0].size();

    // perform a binary search along the entire 2D array.
    std::size_t range_start = 0;
    std::size_t range_end = num_rows * num_cols - 1;
    while (range_start < range_end) {
      std::size_t range_mid = range_start + (range_end - range_start) / 2;
      int val_mid = matrix[range_mid / num_cols][range_mid % num_cols];
      if (val_mid == target) {
        return true;
      }

      if (val_mid < target) {
        // target surely won't be found in the lower half of the range.
        range_start = range_mid + 1;
      }
      else {
        // target certainly won't be found in the upper half of the range.
        range_end = range_mid;
      }
    }
    // It should be enough to return false here, but need last extra check - 
    // for the case of a one-by-one array, where both range_start and range_end are 0 from the beginning,
    // so the entire body of while-loop is skipped.
    //return false; (would return wrong values for testcase2)
    return matrix[range_start / num_cols][range_start % num_cols] == target;
  }
};

int main() {
  Solution s{};
  std::vector<std::vector<int>> testcase1 = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
  assert(s.searchMatrix(testcase1, 3) == true);
  assert(s.searchMatrix(testcase1, 2) == false);
  assert(s.searchMatrix(testcase1, 13) == false);

  std::vector<std::vector<int>> testcase2 = {{1}};
  assert(s.searchMatrix(testcase2, 1) == true);
  assert(s.searchMatrix(testcase2, 2) == false);
}
