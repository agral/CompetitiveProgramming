#include <algorithm>
#include <cassert>
#include <vector>
#include <iostream>

class Solution {
public:
  int eraseOverlapIntervals(std::vector<std::vector<int>>& intervals) {
    if (intervals.size() <= 1) {
      return 0;
    }
    // Sort the intervals in ascending order w.r.t. each interval start:
    std::sort(intervals.begin(), intervals.end(),
              [](const std::vector<int>& lhs, const std::vector<int>& rhs) {
                  return lhs[0] < rhs[0];
              });
    std::size_t numRemoved = 0;
    std::size_t lastAccepted = 0;

    for (std::size_t i = 1; i < intervals.size(); i++) {
      // Case A: interval lA ends before this interval begins.
      // ---lA---
      //           ---i---
      // best action: this interval becomes lastAccepted, numRemoved stays unchanged.
      if (intervals[lastAccepted][1] <= intervals[i][0]) {
        lastAccepted = i;
      }

      // Case B: interval lA ends before this interval ends.
      // ---lA---
      //    ---i---
      // best action: remove i interval, leave lastAccepted unchanged.
      // ties do not matter.
      else if (intervals[lastAccepted][1] < intervals[i][1]) {
        numRemoved += 1;
      }

      // Case C: interval lA ends after this interval ends.
      // ------lA------
      //    ---i---
      // best option: remove lastAccepted interval, i becomes lastAccepted now.
      else {
        lastAccepted = i;
        numRemoved += 1;
      }
    }

    return numRemoved;
  }
};

int main() {
  Solution s{};
  std::vector<std::vector<int>> test1 = {{1, 2}, {2, 3}, {3, 4}, {1, 3}};
  assert(s.eraseOverlapIntervals(test1) == 1);

  std::vector<std::vector<int>> test2 = {{1, 2},};
  assert(s.eraseOverlapIntervals(test2) == 0);

  std::vector<std::vector<int>> test3 = {};
  assert(s.eraseOverlapIntervals(test3) == 0);
}
