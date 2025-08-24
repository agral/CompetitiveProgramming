#include <cassert>
#include <vector>

class Solution {
public:
  int longestSubarray(std::vector<int>& nums) {
    int ans = 0;
    int numZeroes = 0;
    int rs = 0; // range start
    int re = 0; // range end
    while (re < nums.size()) {
      if (nums[re] == 0) {
        ++numZeroes;
      }
      while (numZeroes == 2) {
        if (nums[rs] == 0) {
          --numZeroes;
        }
        ++rs;
      }

      if (re - rs > ans) {
        ans = re - rs;
      }
      ++re;
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> query = {1, 1, 0, 1};
    assert(s.longestSubarray(query) == 3);
  }
  {
    std::vector<int> query = {0, 1, 1, 1, 0, 1, 1, 0, 1};
    assert(s.longestSubarray(query) == 5);
  }
  {
    std::vector<int> query = {1, 1, 1};
    assert(s.longestSubarray(query) == 2);
  }
}
