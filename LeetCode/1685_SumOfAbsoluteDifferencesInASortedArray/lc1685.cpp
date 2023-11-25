#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> getSumAbsoluteDifferences(std::vector<int>& nums) {
    // Calculate prefix and postfix sums:
    const int N = nums.size();
    std::vector<int> presum(N);
    std::vector<int> postsum(N);
    presum[0] = nums[0];
    for (int i = 1; i < N; i++) {
      presum[i] = presum[i - 1] + nums[i];
    }
    postsum[N - 1] = nums[N - 1];
    for (int i = N - 2; i >= 0; i--) {
      postsum[i] = postsum[i + 1] + nums[i];
    }

    std::vector<int> ans(N);
    for (int i = 0; i < N; i++) {
      int smaller_sum = (i + 1) * nums[i] - presum[i];
      int bigger_sum = postsum[i] - (N - i) * nums[i];
      ans[i] = smaller_sum + bigger_sum;
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {2, 3, 5};
    std::vector<int> expected = {4, 3, 5};
    auto actual = s.getSumAbsoluteDifferences(nums);
    assert(actual == expected);
  }
  {
    std::vector<int> nums = {1, 4, 6, 8, 10};
    std::vector<int> expected = {24, 15, 13, 15, 21};
    auto actual = s.getSumAbsoluteDifferences(nums);
    assert(actual == expected);
  }
}
