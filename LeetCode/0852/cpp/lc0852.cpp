#include <vector>
#include <cassert>

class Solution {
public:
  int peakIndexInMountainArray(std::vector<int>& arr) {
    int left = 0;
    int right = arr.size() - 1;
    while (left < right) {
      int mid = left + (right - left) / 2;
      if (arr[mid] > arr[mid+1]) {
        // current window is in the DESCENDING part of a mountain.
        right = mid;
      }
      else {
        // current window is in the ASCENDING part of a mountain.
        left = mid + 1;
      }
    }
    return left;
  }
};

int main() {
  Solution s{};
  std::vector<int> test1 = {0, 1, 0};
  assert(s.peakIndexInMountainArray(test1) == 1);
  std::vector<int> test2 = {0, 2, 1, 0};
  assert(s.peakIndexInMountainArray(test2) == 1);
  std::vector<int> test3 = {0, 10, 5, 2};
  assert(s.peakIndexInMountainArray(test3) == 1);
}
