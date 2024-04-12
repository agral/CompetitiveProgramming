#include <cassert>
#include <vector>

class Solution {
public:
  int trap(std::vector<int>& height) {
    const int SZ = height.size();
    int ans = 0;
    std::vector<int> left(SZ);
    std::vector<int> right(SZ);
    for (int i = 0; i < SZ; ++i) {
      left[i] = (i == 0) ? height[i] : std::max(height[i], left[i - 1]);
    }
    for (int i = 0; i < SZ; ++i) {
      right[SZ - 1 - i] = (i == 0) ? height[SZ - 1 - i] : std::max(height[SZ - 1 - i], right[SZ - i]);
    }
    for (int i = 0; i < SZ; ++i) {
      ans += std::min(left[i], right[i]) - height[i];
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> heights = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
    assert(s.trap(heights) == 6);
  }
  {
    std::vector<int> heights = {4, 2, 0, 3, 2, 5};
    assert(s.trap(heights) == 9);
  }
}
