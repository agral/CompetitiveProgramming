#include <cassert>
#include <limits>
#include <vector>
#include <stack>

class Solution {
public:
  bool find132pattern(std::vector<int>& nums) {
    // need to find a sequence where ni < nk < nj
    std::stack<int> stack; // holds candidates for nj
    int nk = std::numeric_limits<int>::min();
    for (int a = nums.size() - 1; a >= 0; a--) {
      if (nums[a] < nk) {
        // ni is always smaller than nk, so 1-3-2 triplet has just been found.
        return true;
      }
      while (!stack.empty() && nums[a] > stack.top()) {
        nk = stack.top();
        stack.pop();
      }
      stack.push(nums[a]);
    }
    return false;
  }
};

int main() {
  Solution s;
  std::vector<int> nums1 = {1, 2, 3, 4};
  assert(s.find132pattern(nums1) == false);

  std::vector<int> nums2 = {3, 1, 4, 2};
  assert(s.find132pattern(nums2) == true);

  std::vector<int> nums3 = {-1, 3, 2, 0};
  assert(s.find132pattern(nums3) == true);
}
