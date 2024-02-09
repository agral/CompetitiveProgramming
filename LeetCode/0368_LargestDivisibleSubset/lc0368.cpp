#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> largestDivisibleSubset(std::vector<int>& nums) {
    const int NUMS_SIZE = nums.size();
    int currIndex = -1;
    int maxSubsetSize = 0;

    // maxSubsetSizeAt[k]: the maximum size of a subset that starts at nums[0] and ends at nums[k-1].
    std::vector<int> maxSubsetSizeAt(NUMS_SIZE, 1);
    std::vector<int> previousIndex(NUMS_SIZE, -1);

    std::ranges::sort(nums);
    // now no need to check both nums[i] % nums[j] and nums[j] % nums[i] -
    // as the vector is sorted, it is sufficient to check nums[i] % nums[j]
    // for (i, j) pairs such that i is strictly greater than j.

    for (int i = 0; i < NUMS_SIZE; ++i) {
      for (int j = i - 1; j >= 0; --j) {
        if ((nums[i] % nums[j] == 0) && (maxSubsetSizeAt[i] - maxSubsetSizeAt[j] < 1)) {
          maxSubsetSizeAt[i] = maxSubsetSizeAt[j] + 1;
          previousIndex[i] = j;
        }
      }

      // update maxSubsetSize when larger subset size is found:
      if (maxSubsetSizeAt[i] > maxSubsetSize) {
        maxSubsetSize = maxSubsetSizeAt[i];
        currIndex = i;
      }
    }

    // Finally, push items back following the trail of previousIndex indices:
    std::vector<int> ans;
    while (currIndex != -1) {
      ans.push_back(nums[currIndex]);
      currIndex = previousIndex[currIndex];
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {1, 2, 3};
    std::vector<int> expected1 = {3, 1};
    std::vector<int> expected2 = {2, 1};
    std::vector<int> actual = s.largestDivisibleSubset(nums);
    assert(actual == expected1 || actual == expected2);
  }
  {
    std::vector<int> nums = {1, 2, 4, 8};
    std::vector<int> expected = {8, 4, 2, 1};
    assert(s.largestDivisibleSubset(nums) == expected);
  }
}
