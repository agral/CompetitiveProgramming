#include <algorithm>
#include <cassert>
#include <deque>
#include <vector>

class Solution {
public:
  std::vector<int> deckRevealedIncreasing(std::vector<int>& deck) {
    std::sort(deck.begin(), deck.end(), std::greater<int>());
    std::deque<int> queue;
    queue.push_front(deck[0]);

    for (int i = 1; i < deck.size(); ++i) {
      queue.push_front(queue.back());
      queue.pop_back();
      queue.push_front(deck[i]);
    }

    return std::vector<int>{queue.begin(), queue.end()};
  }
};

int main() {
  Solution s;
  {
    std::vector<int> deck = {17, 13, 11, 2, 3, 5, 7};
    std::vector<int> expected = {2, 13, 3, 11, 5, 17, 7};
    assert(s.deckRevealedIncreasing(deck) == expected);
  }
}
