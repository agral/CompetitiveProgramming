#include <cassert>
#include <vector>

class Solution {
public:
  int findJudge(int n, std::vector<std::vector<int>>& trust) {
    std::vector<int> confidence(n + 1);
    for (const auto& t: trust) {
      --confidence[t[0]]; // necessary, this will result in a "fake judge" trusting someone else not marked as valid judge.
      ++confidence[t[1]];
    }

    for (int k = 1; k <= n; ++k) {
      if (confidence[k] == n - 1) {
        return k;
      }
    }
    return -1;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> trust = {{1, 2}};
    assert(s.findJudge(2, trust) == 2);
  }
  {
    std::vector<std::vector<int>> trust = {{1, 3}, {2, 3}};
    assert(s.findJudge(3, trust) == 3);
  }
  {
    std::vector<std::vector<int>> trust = {{1, 3}, {2, 3}, {3, 1}};
    assert(s.findJudge(3, trust) == -1);
  }
}
