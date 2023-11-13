#include <cassert>
#include <vector>

// Given a binary array `nums` and an integer `k`, return the maximum number
// of consecutive `1`'s in the array if you can flip at most `k` `0`'s.

class Solution {
public:
  int longestOnes(std::vector<int>& nums, int k) {
    int rs = 0;
    int re = 0;
    int ans = 0;
    while (re < nums.size()) {
      if (nums[re] == 0) {
        k--;
      }
      while (k < 0) {
        if (nums[rs] == 0) {
          k++;
        }
        rs++;
      }
      if (re - rs + 1 > ans) {
        ans = re - rs + 1;
      }
      re++;
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> example1 = {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0};
    assert(s.longestOnes(example1, 2) == 6);
  }
  {
    std::vector<int> example2 = {0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1};
    assert(s.longestOnes(example2, 3) == 10);
  }
}
