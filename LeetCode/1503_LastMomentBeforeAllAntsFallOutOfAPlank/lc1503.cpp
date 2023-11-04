#include <cassert>
#include <vector>

class Solution {
public:
  int getLastMoment(int n, std::vector<int>& left, std::vector<int>& right) {
    // The "collision rule" here is a red herring,
    // the answer is the same as if ants just walked straight forward
    // with no collisions whatsoever.
    // So, the only ants that matter are:
    //   a) the left-most right-facing ant,
    //   b) the right-most left-facing ant.
    // The answer is the biggest of the distances between a)'s and b)'s
    // respective end of plank.
    int leftmost = n;
    int rightmost = 0;
    for (auto it = right.begin(); it != right.end(); it++) {
      leftmost = std::min(leftmost, *it);
    }
    for (auto it = left.begin(); it != left.end(); it++) {
      rightmost = std::max(rightmost, *it);
    }

    return std::max(n - leftmost, rightmost);
  }
};

int main() {
  Solution s;
  { // example 1
    std::vector<int> left_facing = {3, 4};
    std::vector<int> right_facing = {0, 1};
    assert(s.getLastMoment(4, left_facing, right_facing) == 4);
  }
  { // example 2
    std::vector<int> left_facing = {};
    std::vector<int> right_facing = {0, 1, 2, 3, 4, 5, 6, 7};
    assert(s.getLastMoment(7, left_facing, right_facing) == 7);
  }
  { // example 3
    std::vector<int> left_facing = {0, 1, 2, 3, 4, 5, 6, 7};
    std::vector<int> right_facing = {};
    assert(s.getLastMoment(7, left_facing, right_facing) == 7);
  }
}
