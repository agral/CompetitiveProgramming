// Solution to LeetCode problem #0001, "Two Sum"
// https://leetcode.com/problems/two-sum/
// by https://gralin.ski/

#include <iostream>
#include <vector>

class Solution {
  public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
      for (int i = 0; i < nums.size() - 1; i++) {
        for (int j = i + 1; j < nums.size(); j++) {
          if (nums[i] + nums[j] == target) {
            return {i, j};
          }
        }
      }

      return {-1};
    }
};

int main()
{
  Solution s;

  {
    std::cout << "Example 1\n";
    std::vector vals = {2, 7, 11, 15};
    auto rv = s.twoSum(vals, 9);
    if ((rv.at(0) == 0) && (rv.at(1) == 1)) {
      std::cout << "Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "Example 2\n";
    std::vector vals = {3, 2, 4};
    auto rv = s.twoSum(vals, 6);
    if ((rv.at(0) == 1) && (rv.at(1) == 2)) {
      std::cout << "Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "Example 3\n";
    std::vector vals = {3, 3};
    auto rv = s.twoSum(vals, 6);
    if ((rv.at(0) == 0) && (rv.at(1) == 1)) {
      std::cout << "Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
}
