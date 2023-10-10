#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int minOperations(std::vector<int>& nums) {
    // Make nums sorted and with duplicates removed:
    int initial_length = nums.size(); 
    std::sort(nums.begin(), nums.end());
    nums.erase(std::unique(nums.begin(), nums.end()), nums.end());

    int ans = initial_length;
    int right_idx = 0;
    for (int left_idx = 0; left_idx < nums.size(); left_idx++) {
      while (right_idx < nums.size() && nums[right_idx] < nums[left_idx] + initial_length) {
        right_idx++;
      }
      ans = std::min(ans, initial_length - (right_idx - left_idx));
    }
    return ans;
  }
};

int main() {
  Solution s;

  std::vector<int> example1 = {4, 2, 5, 3};
  assert(s.minOperations(example1) == 0);
  std::vector<int> example2 = {1, 2, 3, 5, 6};
  assert(s.minOperations(example2) == 1);
  std::vector<int> example3 = {1, 10, 100, 1000};
  assert(s.minOperations(example3) == 3);
  std::vector<int> duplicated = {1, 1, 1, 4, 5, 5};
  assert(s.minOperations(duplicated) == 3);
}
