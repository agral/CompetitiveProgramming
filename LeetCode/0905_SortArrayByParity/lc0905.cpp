#include <algorithm>
#include <vector>

#include <iostream>

class Solution {
public:
  std::vector<int> sortArrayByParity(std::vector<int>& nums) {
    std::sort(nums.begin(), nums.end(), [](const int& lhs, const int& rhs){return (lhs & 1) < (rhs & 1);});
    return nums;
  }
};

void print_vec(std::vector<int>& v) {
  std::cout << "[";
  for (const auto& i: v) {
    std::cout << i << ", ";
  }
  std::cout << "]\n";
}

int main() {
  Solution s;
  std::vector<int> nums1 = {3, 1, 2, 4};
  auto ans1 = s.sortArrayByParity(nums1);
  print_vec(ans1);

  std::vector<int> nums2 = {0};
  auto ans2 = s.sortArrayByParity(nums2);
  print_vec(ans2);
}
