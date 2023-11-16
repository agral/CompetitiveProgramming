#include <cassert>
#include <vector>
#include <unordered_set>

class Solution {
public:
  std::vector<std::vector<int>> findDifference(std::vector<int>& nums1, std::vector<int>& nums2) {
    std::unordered_set<int> set1{nums1.begin(), nums1.end()};
    std::unordered_set<int> set2{nums2.begin(), nums2.end()};
    std::vector<std::vector<int>> ans(2, std::vector<int>{});
    for (const auto& num1: set1) {
      if (set2.find(num1) == set2.end()) {
        // This number is in nums1, but not in nums2.
        ans[0].push_back(num1);
      }
    }
    for (const auto& num2: set2) {
      if (set1.find(num2) == set1.end()) {
        // This number exists in nums2, but not in nums1.
        ans[1].push_back(num2);
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums1 = {1, 2, 3};
    std::vector<int> nums2 = {2, 4, 6};
    std::vector<std::vector<int>> expected{{3, 1}, {6, 4}};
    auto actual = s.findDifference(nums1, nums2);
    assert(expected == actual);
  }
  {
    std::vector<int> nums1 = {1, 2, 3, 3};
    std::vector<int> nums2 = {1, 1, 2, 2};
    std::vector<std::vector<int>> expected{{3}, {}};
    auto actual = s.findDifference(nums1, nums2);
    assert(expected == actual);
  }
}
