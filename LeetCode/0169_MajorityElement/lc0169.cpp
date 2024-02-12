#include <cassert>
#include <vector>

class Solution {
public:
  int majorityElement(std::vector<int>& nums) {
    int counter = 0;
    int ans;

    for (const int num: nums) {
      if (counter == 0) {
        ans = num;
        counter = 1;
      }
      else {
        if (num == ans) {
          ++counter;
        }
        else {
          --counter;
        }
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {3, 2, 3};
    assert(s.majorityElement(nums) == 3);
  }
  {
    std::vector<int> nums = {2, 3, 3};
    assert(s.majorityElement(nums) == 3);
  }
  {
    std::vector<int> nums = {3, 3, 2};
    assert(s.majorityElement(nums) == 3);
  }

  {
    std::vector<int> nums = {2, 2, 1, 1, 1, 2, 2};
    assert(s.majorityElement(nums) == 2);
  }
}
