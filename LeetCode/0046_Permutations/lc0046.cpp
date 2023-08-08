#include <algorithm>
#include <vector>
#include <iostream>

class Solution {
public:
  std::vector<std::vector<int>> permute(std::vector<int>& nums) {
    m_ans.clear();
    // push the initial permutation:
    m_ans.push_back(nums);

    // push (n! - 1) remaining permutations:
    for (int i = 1; i < factorial(nums.size()); i++) {
      std::next_permutation(nums.begin(), nums.end());
      m_ans.push_back(nums);
    }
    print_ans();
    return m_ans;
  };

private:
  int factorial(int n) {
    int ans = 1;
    for (int k = 2; k <= n; k++) {
      ans *= k;
    }
    return ans;
  }

  void print_ans() {
    for (const auto& permutation: m_ans) {
      std::cout << '[';
      for (const auto& elem: permutation) {
        std::cout << elem << ' ';
      }
      std::cout << "]\n";
    }
    std::cout << "---\n";
  }

  std::vector<std::vector<int>> m_ans;
};

int main() {
  Solution s{};
  std::vector<int> nums1 = {1, 2, 3};
  s.permute(nums1);
  std::vector<int> nums2 = {0, 1};
  s.permute(nums2);
  std::vector<int> nums3 = {1};
  s.permute(nums3);
};
