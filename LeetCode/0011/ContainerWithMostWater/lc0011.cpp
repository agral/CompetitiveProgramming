#include <cassert>
#include <vector>

class Solution {
public:
  int maxArea(std::vector<int>& height) {
    int left_idx = 0;
    int right_idx = height.size() - 1;
    int max_area = 0;
    while (left_idx < right_idx) {
      int area = ((right_idx - left_idx) * std::min(height[left_idx], height[right_idx]));
      max_area = std::max(area, max_area);
      if (height[left_idx] < height[right_idx]) {
        left_idx += 1;
      }
      else if (height[right_idx] < height[left_idx]) {
        right_idx -= 1;
      }
      else {
        left_idx += 1;
        right_idx -= 1;
      }
    }

    return max_area;
  };
};

int main() {
  Solution s;
  std::vector<int> testcase1 = {1, 8, 6, 2, 5, 4, 8, 3, 7};
  assert(s.maxArea(testcase1) == 49);

  std::vector<int> testcase2 = {1, 1};
  assert(s.maxArea(testcase2) == 1);

  std::vector<int> testcase3 = {2, 3, 4, 5, 18, 17, 6};
  assert(s.maxArea(testcase3) == 17);
}
