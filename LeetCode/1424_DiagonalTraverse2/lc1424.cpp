#include <cassert>
#include <unordered_map>
#include <vector>

class Solution {
public:
  std::vector<int> findDiagonalOrder(std::vector<std::vector<int>>& nums) {
    std::vector<int> ans;
    // idea: the numbers on diagonal lines have the same sum of height+width indices.
    std::unordered_map<int, std::vector<int>> diagonals;
    int highest_diagonal = -1;
    for (int h = 0; h < nums.size(); h++) {
      for (int w = 0; w < nums[h].size(); w++) {
        int diagonal_index = h + w;
        highest_diagonal = std::max(highest_diagonal, diagonal_index);
        diagonals[diagonal_index].push_back(nums[h][w]);
      }
    }
    for (int diag = 0; diag <= highest_diagonal; diag++) {
      // Diagnoal numbers are stored in reversed order. We could use push_front,
      // but it's less optimal (slower) than pushing to the back of vector.
      // so need to iterate in reversed order again when constructing the answer:
      //
      for (int i = diagonals[diag].size() - 1; i >= 0; i--) {
        ans.push_back(diagonals[diag][i]);
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> nums = {
      {1, 2, 3},
      {4, 5, 6},
      {7, 8, 9}};
    std::vector<int> ans = s.findDiagonalOrder(nums);
    std::vector<int> expected = {1, 4, 2, 7, 5, 3, 8, 6, 9};
    assert(ans == expected);
  }
  {
    std::vector<std::vector<int>> nums = {
      {1, 2, 3, 4, 5},
      {6, 7},
      {8},
      {9, 10, 11},
      {12, 13, 14, 15, 16}};
    std::vector<int> ans = s.findDiagonalOrder(nums);
    std::vector<int> expected = {1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16};
    assert(ans == expected);
  }
}
