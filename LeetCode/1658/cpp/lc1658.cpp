#include <cassert>
#include <vector>

class Solution {
public:
  int minOperations(std::vector<int>& nums, int x) {
    std::vector<int> one_indexed_prefix_sum(nums.size() + 1);
    one_indexed_prefix_sum[0] = 0;
    int total_prefix_sum = 0;
    for (int i = 0; i < nums.size(); i++) {
      total_prefix_sum += nums[i];
      one_indexed_prefix_sum[i + 1] = total_prefix_sum;
    }
    int target = one_indexed_prefix_sum[nums.size()] - x;
    if (target < 0) {
      return -1;
    }

    int start_idx = 0; // range starts here, inclusive. start_idx always <= end_idx.
    int end_idx = 0; // range ends here, inclusive.
    int max_width = -1;
    while (end_idx < nums.size()) {
      auto range_sum = one_indexed_prefix_sum[end_idx+1] - one_indexed_prefix_sum[start_idx];
      if (range_sum == target) {
        max_width = std::max(end_idx - start_idx + 1, max_width);
        start_idx += 1;
        end_idx += 1;
      } else if (range_sum < target) {
        end_idx += 1;
      } else {
        start_idx += 1;
      }
    }

    return (max_width == -1 ? -1 : nums.size() - max_width);
  };
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 1, 4, 2, 3};
  assert(s.minOperations(example1, 5) == 2);

  std::vector<int> example2 = {5, 6, 7, 8, 9};
  assert(s.minOperations(example2, 4) == -1);

  std::vector<int> example3 = {3, 2, 20, 1, 1, 3};
  assert(s.minOperations(example3, 10) == 5);

  std::vector<int> silly1 = {1234};
  assert(s.minOperations(silly1, 1234) == 1);
  assert(s.minOperations(silly1, 1233) == -1);
  assert(s.minOperations(silly1, 1235) == -1);

  std::vector<int> long_testcase = {6016,5483,541,4325,8149,3515,7865,2209,9623,9763,
      4052,6540,2123,2074,765,7520,4941,5290,5868,6150,6006,6077,2856,7826,9119};
  assert(s.minOperations(long_testcase, 31841) == 6);

  std::vector<int> wa1 = {1, 1};
  assert(s.minOperations(wa1, 3) == -1);
}
