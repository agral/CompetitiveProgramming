#include <cassert>
#include <unordered_set>
#include <vector>

class Solution {
public:
  int findDuplicate(std::vector<int>& nums) {
    std::unordered_set<int> set;
    for (const auto& num: nums) {
      if (set.count(num) > 0) {
        return num;
      } else {
        set.insert(num);
      }
    }
    throw; // in case someone passed a vector without duplicates
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {1, 3, 4, 2, 2};
  assert(s.findDuplicate(example1) == 2);
  std::vector<int> example2 = {3, 1, 3, 4, 2};
  assert(s.findDuplicate(example2) == 3);
}
