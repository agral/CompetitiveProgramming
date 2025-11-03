#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  int minCost(std::string colors, std::vector<int>& neededTime) {
    int ans = 0;
    int time = neededTime[0];
    for (int i = 1; i < colors.size(); ++i) {
      if (colors[i-1] == colors[i]) {
        // One of the two balloons has to go.
        ans += std::min(time, neededTime[i]);
        time = std::max(time, neededTime[i]);
      }
      else {
        // Hold the current balloon in hand, in case next one is same color:
        time = neededTime[i];
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::string colors = "abaac";
    std::vector<int> neededTime = {1, 2, 3, 4, 5};
    assert(s.minCost(colors, neededTime) == 3);
  }
  {
    std::string colors = "abc";
    std::vector<int> neededTime = {1, 2, 3};
    assert(s.minCost(colors, neededTime) == 0);
  }
  {
    std::string colors = "aabaa";
    std::vector<int> neededTime = {1, 2, 3, 4, 1};
    assert(s.minCost(colors, neededTime) == 2);
  }
}
