#include <cassert>
#include <vector>

class Solution {
public:
  int timeRequiredToBuy(std::vector<int>& tickets, int k) {
    int ans = 0;
    for (int i = 0; i < tickets.size(); ++i) {
      ans += std::min(tickets[i], (i > k) ? tickets[k] - 1 : tickets[k]);
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> tickets = {2, 3, 2};
    int k = 2;
    assert(s.timeRequiredToBuy(tickets, k) == 6);
  }
  {
    std::vector<int> tickets = {5, 1, 1, 1};
    int k = 0;
    assert(s.timeRequiredToBuy(tickets, k) == 8);
  }
}
