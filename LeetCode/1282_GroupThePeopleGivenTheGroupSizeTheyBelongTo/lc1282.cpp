#include <cassert>
#include <unordered_map>
#include <vector>

#include <iostream>

class Solution {
public:
  std::vector<std::vector<int>> groupThePeople(std::vector<int>& groupSizes) {
    std::vector<std::vector<int>> ans;
    std::unordered_map<int, std::vector<int>> current_groups;
    for (int i = 0; i < groupSizes.size(); i++) {
      int sz = groupSizes[i];
      if (current_groups.count(sz) == 0) {
        current_groups[sz] = {i};
      }
      else {
        current_groups[sz].push_back(i);
      }

      if (current_groups[sz].size() == sz) {
        ans.push_back(current_groups[sz]);
        current_groups[sz] = {};
      }
    }
    return ans;
  }
};

void print_groups(std::vector<std::vector<int>> groups) {

  std::cout << "[";
  for (const auto& group: groups) {
    std::cout << "{";
    for (const auto& elem: group) {
      std::cout << elem << ", ";
    }
    std::cout << "}, ";
  }
  std::cout << "]\n";
}

int main() {
  Solution s;

  std::vector<int> example1 = {3, 3, 3, 3, 3, 1, 3};
  std::vector<std::vector<int>> expected_ans1 = {{0, 1, 2}, {5}, {3, 4, 6}};
  auto ans1 = s.groupThePeople(example1);
  print_groups(ans1);
  assert(ans1 == expected_ans1);

  std::vector<int> example2 = {2, 1, 3, 3, 3, 2};
  std::vector<std::vector<int>> expected_ans2 = {{1}, {2, 3, 4}, {0, 5}};
  auto ans2 = s.groupThePeople(example2);
  print_groups(ans2);
  assert(ans2 == expected_ans2);
}
