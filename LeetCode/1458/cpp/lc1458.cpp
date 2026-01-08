#include <algorithm>
#include <climits>
#include <cassert>
#include <vector>

class Solution {
public:
  int maxDotProduct(std::vector<int>& nums1, std::vector<int>& nums2) {

    // dp[x][y] stores max dot product of subsequences nums1[0...x], nums2[0...y].
    std::vector<std::vector<int>> dp(nums1.size() + 1, std::vector<int>(nums2.size() + 1, INT_MIN));
    for (int i = 0; i < nums1.size(); i++) {
      for (int j = 0; j < nums2.size(); j++) {
        dp[i+1][j+1] = std::max({dp[i][j+1], dp[i+1][j], std::max(0, dp[i][j]) + nums1[i] * nums2[j]});
      }
    }
    return dp[nums1.size()][nums2.size()];
  }
};

int main() {
  Solution s;
  std::vector<int> example1_1 = {2, 1, -2, 5};
  std::vector<int> example1_2 = {3, 0, -6};
  assert(s.maxDotProduct(example1_1, example1_2) == 18); // [2, -2] dot [3, -6]

  std::vector<int> example2_1 = {3, -2};
  std::vector<int> example2_2 = {2, -6, 7};
  assert(s.maxDotProduct(example2_1, example2_2) == 21); // [3] dot [7]

  std::vector<int> example3_1 = {-1, -1};
  std::vector<int> example3_2 = {1, 1};
  assert(s.maxDotProduct(example3_1, example3_2) == -1); // [-1] dot [1]
}
