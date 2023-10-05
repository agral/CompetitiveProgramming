#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<int> majorityElement(std::vector<int>& nums) {
    if (nums.size() <= 1) {
      return nums;
    }

    // There are at most two majority elements.
    int dominantA = 0; // can be anything, initial value does not matter.
    int dominantB = 1; // can be anything except dominantA, initial value does not matter.
    int countA = 0;
    int countB = 0;

    for (const auto& n: nums) {
      if (n == dominantA) {
        countA++;
      }
      else if (n == dominantB) {
        countB++;
      }
      else if (countA == 0) {
        dominantA = n;
        countA = 1;
      }
      else if (countB == 0) {
        dominantB = n;
        countB = 1;
      }
      else {
        countA--;
        countB--;
      }
    }
    std::vector<int> ans;
    if (std::count(nums.begin(), nums.end(), dominantA) > nums.size() / 3) {
      ans.push_back(dominantA);
    }
    if (std::count(nums.begin(), nums.end(), dominantB) > nums.size() / 3) {
      ans.push_back(dominantB);
    }

    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {3, 2, 3};
  assert(s.majorityElement(example1) == std::vector<int>{3});

  std::vector<int> example2 = {1};
  assert(s.majorityElement(example2) == std::vector<int>{1});

  std::vector<int> example3 = {1, 2};
  assert(s.majorityElement(example3) == std::vector<int>({2, 1}));

  std::vector<int> empty = {};
  assert(s.majorityElement(empty) == std::vector<int>{});
}
