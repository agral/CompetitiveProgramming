#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> rearrangeArray(std::vector<int>& nums) {
    const int SIZE = nums.size();
    const int HALF_SIZE = nums.size() / 2;
    std::vector<int> positives;
    positives.reserve(HALF_SIZE);
    std::vector<int> negatives;
    negatives.reserve(HALF_SIZE);
    for (const int num: nums) {
      if (num > 0) {
        positives.push_back(num);
      } else {
        negatives.push_back(num);
      }
    }

    std::vector<int> ans;
    for (int i = 0; i < HALF_SIZE; ++i) {
      ans.push_back(positives[i]);
      ans.push_back(negatives[i]);
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {3, 1, -2, -5, 2, -4};
    std::vector<int> expected = {3, -2, 1, -5, 2, -4};
    assert(s.rearrangeArray(nums) == expected);
  }
  {
    std::vector<int> nums = {-1, 1};
    std::vector<int> expected = {1, -1};
    assert(s.rearrangeArray(nums) == expected);
  }
}
