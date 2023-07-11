#include <algorithm>
#include <cassert>
#include <vector>

class Solution {
public:
  std::vector<bool> kidsWithCandies(std::vector<int>& candies, int extraCandies) {
    int max = *std::max_element(candies.begin(), candies.end());
    std::vector<bool> ans{};
    ans.resize(candies.size());
    for (std::size_t i = 0; i < candies.size(); i++) {
      ans[i] = candies[i] + extraCandies >= max;
    }
    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<int> candies1 = {2, 3, 5, 1, 3};
  std::vector<bool> expectedAnswer1 = {true, true, true, false, true};
  std::vector<bool> answer1 = s.kidsWithCandies(candies1, 3);
  assert(answer1 == expectedAnswer1);

  std::vector<int> candies2 = {4, 2, 1, 1, 2};
  std::vector<bool> expectedAnswer2 = {true, false, false, false, false};
  std::vector<bool> answer2 = s.kidsWithCandies(candies2, 1);
  assert(answer2 == expectedAnswer2);

  std::vector<int> candies3 = {12, 1, 12};
  std::vector<bool> expectedAnswer3 = {true, false, true};
  std::vector<bool> answer3 = s.kidsWithCandies(candies3, 10);
  assert(answer3 == expectedAnswer3);
}
