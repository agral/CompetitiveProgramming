#include <cassert>
#include <vector>

class Solution {
public:
  long long minimumReplacement(std::vector<int>& nums) {
    long long ans = 0LL;
    int minimum = nums[nums.size() - 1];
    for (int i = nums.size() - 2; i >= 0; i--) {
      if (nums[i] <= minimum) {
        minimum = nums[i];
      } else {
        // Need to split this number into values that each do not exceed minimum.
        // Formula for that is floor((number - 1) / minimum).
        int num_splits = (nums[i] - 1) / minimum; 
        ans += num_splits;

        // The smallest number from splitting nums[i] into as equal as possible parts
        // becomes a new minimum. This may mean that the minimum decreases (likely)
        // or stays the same (unlikely).
        minimum = nums[i] / (num_splits + 1);
      }
    }
    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {3, 9, 3};
  assert(s.minimumReplacement(testcase1) == 2);
  std::vector<int> testcase2 = {1, 2, 3, 4, 5};
  assert(s.minimumReplacement(testcase2) == 0);

  std::vector<int> testcase3 = {821, 881, 275};
  assert(s.minimumReplacement(testcase3) == 6);
}
