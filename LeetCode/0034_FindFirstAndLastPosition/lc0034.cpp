#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> searchRange(std::vector<int>& nums, int target) {
    return {do_search(nums, target, false), do_search(nums, target, true)};
  }

private:
  int do_search(std::vector<int>& nums, int target, bool get_rightmost) {
    // get_rightmost: if true, find rightmost target value. Otherwise find leftmost target value.
    int left = 0;
    int right = nums.size() - 1;
    int last_target = -1;
    while (left <= right) {
      int mid = left + (right - left) / 2;
      if (nums[mid] < target) {
        left = mid + 1;
      }
      else if (nums[mid] > target) {
        right = mid - 1;
      } else {
        last_target = mid;
        if (get_rightmost) {
          left = mid + 1;
        }
        else {
          right = mid - 1;
        }
      }
    }
    return last_target;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {5, 7, 7, 8, 8, 10};
  assert(s.searchRange(example1, 8) == std::vector<int>({3, 4}));
  assert(s.searchRange(example1, 6) == std::vector<int>({-1, -1}));

  std::vector<int> example3 = {};
  assert(s.searchRange(example3, 0) == std::vector<int>({-1, -1}));

  std::vector<int> test_edgecases = {1, 1, 1, 5, 7, 7, 8, 8, 8};
  assert(s.searchRange(test_edgecases, 1) == std::vector<int>({0, 2}));
  assert(s.searchRange(test_edgecases, 8) == std::vector<int>({6, 8}));
}
