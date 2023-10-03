#include <cassert>
#include <vector>
#include <map>

class Solution {
public:
  int numIdenticalPairs(std::vector<int>& nums) {
    std::map<int, int> count;
    for (std::size_t i = 0; i < nums.size(); i++) {
      count[nums[i]]++;
    }
    int ans = 0;
    for (const auto& [k, v]: count) {
      if (v >= 2) {
        ans += v * (v-1) / 2;
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 2, 3, 1, 1, 3};
  assert(s.numIdenticalPairs(example1) == 4);
  std::vector<int> example2 = {1, 1, 1, 1};
  assert(s.numIdenticalPairs(example2) == 6);
  std::vector<int> example3 = {1, 2, 3};
  assert(s.numIdenticalPairs(example3) == 0);
}
