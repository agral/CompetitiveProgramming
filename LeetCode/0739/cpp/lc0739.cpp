#include <cassert>
#include <deque>
#include <vector>

class Solution {
public:
  std::vector<int> dailyTemperatures(std::vector<int>& temperatures) {
    std::deque<int> indices_stack;
    std::vector<int> ans(temperatures.size());

    for (int i = 0; i < temperatures.size(); ++i) {
      if (!indices_stack.empty()) {
        while (
          !indices_stack.empty() &&
          temperatures[indices_stack.back()] < temperatures[i]
        ) {
          const int idx = indices_stack.back();
          indices_stack.pop_back();
          ans[idx] = i - idx;
        }
      }
      indices_stack.push_back(i);
    }
    return ans;
  }
};

int main() {
  Solution s;

  {
    std::vector<int> temperatures = {73, 74, 75, 71, 69, 72, 76, 73};
    std::vector<int> expected =     { 1,  1,  4,  2,  1,  1,  0,  0};
    assert(s.dailyTemperatures(temperatures) == expected);
  }
  {
    std::vector<int> temperatures = {30, 40, 50, 60};
    std::vector<int> expected =     { 1,  1,  1,  0};
    assert(s.dailyTemperatures(temperatures) == expected);
  }
  {
    std::vector<int> temperatures = {30, 60, 90};
    std::vector<int> expected =     { 1,  1,  0};
    assert(s.dailyTemperatures(temperatures) == expected);
  }
}
