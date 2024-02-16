#include <cassert>
#include <queue>
#include <unordered_map>
#include <vector>

class Solution {
public:
  int findLeastNumOfUniqueInts(std::vector<int>& arr, int k) {
    std::priority_queue<int, std::vector<int>, std::greater<int>> min_heap;
    std::unordered_map<int, int> count;
    for (const int num: arr) {
      ++count[num];
    }

    for (const auto& [_ /*num*/, tally]: count) {
      min_heap.push(tally);
    }

    while (k > 0) {
      k -= min_heap.top();
      min_heap.pop();
    }

    // Count a number that has not been fully removed (i.e. three '5's, when k is 2, so only 2 can be removed):
    int extraTally = k < 0 ? 1 : 0;
    return min_heap.size() + extraTally;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> arr = {5, 5, 4};
    assert(s.findLeastNumOfUniqueInts(arr, 1) == 1);
  }
  {
    std::vector<int> arr = {4, 3, 1, 1, 3, 3, 2};
    assert(s.findLeastNumOfUniqueInts(arr, 3) == 2);
  }
}
