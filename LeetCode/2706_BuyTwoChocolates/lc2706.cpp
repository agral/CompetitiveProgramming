#include <cassert>
#include <vector>

class Solution {
public:
  int buyChoco(std::vector<int>& prices, int money) {
    int cheapest = (prices[0] < prices[1] ? prices[0] : prices[1]);
    int cheap = (prices[0] < prices[1] ? prices[1] : prices[0]);
    for (int i = 2; i < prices.size(); ++i) {
      if (prices[i] < cheapest) {
        cheap = cheapest;
        cheapest = prices[i];
      }
      else if (prices[i] < cheap) {
        cheap = prices[i];
      }
    }

    int result = money - cheapest - cheap;
    return (result >= 0 ? result : money);
  }
};

int main() {
  Solution s;
  {
    std::vector<int> prices = {1, 2, 2};
    assert(s.buyChoco(prices, 3) == 0);
  }
  {
    std::vector<int> prices = {3, 2, 3};
    assert(s.buyChoco(prices, 3) == 3);
  }
  {
    std::vector<int> prices = {74, 31, 38, 24, 25, 24, 5};
    assert(s.buyChoco(prices, 79) == 50);
  }
}
