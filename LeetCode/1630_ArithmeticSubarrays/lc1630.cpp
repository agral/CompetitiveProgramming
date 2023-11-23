#include <cassert>
#include <limits>
#include <unordered_set>
#include <vector>

class Solution {
public:
  std::vector<bool> checkArithmeticSubarrays(std::vector<int>& nums,
                                             std::vector<int>& l,
                                             std::vector<int>& r) {
    std::vector<bool> ans;
    for (int i = 0; i < l.size(); i++) {
      ans.push_back(isSubarrayArithmetic(nums, l[i], r[i]));
    }
    return ans;
  }

private:
  bool isSubarrayArithmetic(std::vector<int>& nums, int l, int r) {
    if (l + 2 > r) { // corner case, small -> always arithmetic
      return true;
    }

    int min_value = std::numeric_limits<int>::max();
    int max_value = std::numeric_limits<int>::min();
    std::unordered_set<int> unique_nums;
    for (int k = l; k <= r; k++) {
      min_value = std::min(min_value, nums[k]);
      max_value = std::max(max_value, nums[k]);
      unique_nums.insert(nums[k]);
    }

    if ((max_value - min_value) % (r - l) != 0) {
      // not evenly divisible by subarray's size -> not arithmetic,
      // as delta would be a fraction, and we're dealing with ints here.
      return false;
    }

    int delta = (max_value - min_value) / (r - l);
    for (int d = 1; d <= (r - l); d++) {
      // could also sort the interval and check - but this approach has same
      // complexity, and less hassle. Also, note: 0th elem can be safely skipped.
      if (unique_nums.find(min_value + d * delta) == unique_nums.end()) {
        // d'th expected element of sequence not in actual sequence -> not arithmetic.
        return false;
      }
    }
    return true;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {4, 6, 5, 9, 3, 7};
    std::vector<int> l = {0, 0, 2};
    std::vector<int> r = {2, 3, 5};
    std::vector<bool> expected = {true, false, true};
    auto actual = s.checkArithmeticSubarrays(nums, l, r);
    assert(expected == actual);
  }
  {
    std::vector<int> nums = {-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10};
    std::vector<int> l = {0, 1, 6, 4, 8, 7};
    std::vector<int> r = {4, 4, 9, 7, 9, 10};
    std::vector<bool> expected = {false, true, false, false, true, true};
    auto actual = s.checkArithmeticSubarrays(nums, l, r);
    assert(expected == actual);
  }
}
