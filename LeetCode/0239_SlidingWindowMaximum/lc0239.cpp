#include <vector>
#include <deque>
#include <cassert>

class Solution {
public:
  std::vector<int> maxSlidingWindow(std::vector<int>& nums, int k) {
    int ans_size = nums.size() - k + 1;
    std::deque<int> biggest_elems;

    // Analyze the beginning sliding window:
    biggest_elems.push_back(nums[0]);
    for (auto i = 1; i < k; i++) {
      while (!biggest_elems.empty() && biggest_elems.back() < nums[i]) {
        biggest_elems.pop_back();
      }
      biggest_elems.push_back(nums[i]);
    }

    std::vector<int> ans;
    ans.resize(ans_size);
    // Answer the initial floating window:
    ans[0] = biggest_elems.front();

    for (int offset = 1; offset < ans_size; offset++) {
      // Move the floating window one step to the right.
      // If value that just stopped being inside the window was biggest, remove it from the queue:
      if (biggest_elems.front() == nums[offset-1]) {
        biggest_elems.pop_front();
      }

      // Process the value that just started being inside the floating window:
      int next_val = nums[k - 1 + offset];
      while (!biggest_elems.empty() && biggest_elems.back() < next_val) {
        biggest_elems.pop_back();
      }
      biggest_elems.push_back(next_val);

      // Finally, the max-value of this floating window is in front of biggest_elems:
      ans[offset] = biggest_elems.front();
    }

    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {1, 3, -1, -3, 5, 3, 6, 7};
  std::vector<int> expected_ans1 = {3, 3, 5, 5, 6, 7};
  auto ans1 = s.maxSlidingWindow(testcase1, 3);
  assert(ans1 == expected_ans1);

  std::vector<int> testcase2 = {1};
  std::vector<int> expected_ans2 = {1};
  auto ans2 = s.maxSlidingWindow(testcase2, 1);
  assert(ans2 == expected_ans2);
}
