#include <vector>
#include <cassert>
#include <unordered_map>

class Solution {
public:
  int minOperations(std::vector<int>& nums) {
    std::unordered_map<int, int> count;
    for (const int num: nums) {
      ++count[num];
    }

    int ans = 0;
    for (const auto& [_, cnt]: count) {
      // Every integer >= 2 can be represented as sum of threes and twos.
      // 2 -> 1x2      (1 number)
      // 3 -> 1x3      (1 number)
      // 4 -> 2x2      (2 numbers)
      // 5 and above: subtract 3 until num <= 4.
      // 5 -> 1x3,1x2  (2 numbers)
      // 6 -> 2x3      (2 numbers)
      // 7 -> 1x3,2x2  (3 numbers)
      // 8 -> 2x3,1x2  (3 numbers)
      // 9 -> 3x3      (3 numbers)
      // 10 -> 2x3,2x2 (4 numbers)
      if (cnt == 1) {
        return -1;
      }

      ans += (cnt + 2) / 3; // ceil(cnt/3.0) would work too, but this is faster.
    }

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {2, 3, 3, 2, 2, 4, 2, 3, 4};
    assert(s.minOperations(nums) == 4);
  }
  {
    std::vector<int> nums = {2, 1, 2, 2, 3, 3};
    assert(s.minOperations(nums) == -1);
  }
}
