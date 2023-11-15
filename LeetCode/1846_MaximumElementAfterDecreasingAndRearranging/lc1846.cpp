#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  int maximumElementAfterDecrementingAndRearranging(std::vector<int>& arr) {
    std::sort(arr.begin(), arr.end());
    arr[0] = 1; // this must happen, forced by task's 1st condition.
    for (int idx = 1; idx < arr.size(); idx++) {
      arr[idx] = std::min(arr[idx], arr[idx-1] + 1); // 2nd condition.
      // Note: no need to check whether arr[idx] was less than arr[idx-1];
      // won't happen on sorted array of non-negative integers.
    }
    return arr[arr.size()-1];
  }
};

int main() {
  Solution s;
  {
    std::vector<int> testcase = {2, 2, 1, 2, 1};
    assert(s.maximumElementAfterDecrementingAndRearranging(testcase) == 2);
  }
  {
    std::vector<int> testcase = {100, 1, 1000};
    assert(s.maximumElementAfterDecrementingAndRearranging(testcase) == 3);
  }
  {
    std::vector<int> testcase = {1, 2, 3, 4, 5}; // a NOP test, already satisfies all conditions.
    assert(s.maximumElementAfterDecrementingAndRearranging(testcase) == 5);
  }
}
