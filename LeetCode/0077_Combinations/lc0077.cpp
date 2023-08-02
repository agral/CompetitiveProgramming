#include <cassert>
#include <vector>
#include <iostream>

class Solution {
public:
  std::vector<std::vector<int>> combine(int n, int k) {
    m_ans.clear();
    generate({}, n, k, 1);
    return m_ans;
  };

private:
  void generate(std::vector<int> prefix, int n, int k, int start) {
    if (prefix.size() == k) {
      m_ans.push_back(prefix);
    }
    else {
      for (int i = start; i <= n; i++) {
        std::vector<int> extended{prefix};
        extended.push_back(i);
        generate(extended, n, k, i+1);
      }
    }
  }

  std::vector<std::vector<int>> m_ans;
};

int main() {
  Solution s{};
  auto act = s.combine(5, 2);
  for (const auto& entry: act) {
    std::cout << '[';
    for (const auto& num: entry) {
      std::cout << num << ' ';
    }
    std::cout << "]\n";
  }

  std::vector<std::vector<int>> ans1 = {{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}};
  assert(s.combine(4, 2) == ans1);

  std::vector<std::vector<int>> ans2 = {{1}};
  assert(s.combine(1, 1) == ans2);
};
