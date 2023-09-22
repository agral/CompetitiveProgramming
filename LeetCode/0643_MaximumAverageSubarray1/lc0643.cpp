#include <cassert>
#include <vector>

class Solution {
public:
  double findMaxAverage(std::vector<int>& nums, int k) {
    // guaranteed that k <= nums.size(), no need to check that.

    int current_sum = 0;
    for (int i = 0; i < k; i++) {
      current_sum += nums[i];
    }
    int max_sum = current_sum;
    for (int i = 1; i + k <= nums.size(); i++) {
      current_sum -= nums[i-1];
      current_sum += nums[i+k-1];
      if (current_sum > max_sum) {
        max_sum = current_sum;
      }
    }
    return static_cast<double>(max_sum) / k;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 12, -5, -6, 50, 3};
  assert(s.findMaxAverage(example1, 4) == 12.75);

  std::vector<int> example2 = {5};
  assert(s.findMaxAverage(example2, 1) == 5);

}
