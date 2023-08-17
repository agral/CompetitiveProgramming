#include <cassert>
#include <vector>
#include <optional>
#include <iostream>

class Solution {
public:
  bool increasingTriplet(std::vector<int>& nums) {
    // Scan the entire vector sequentially. Keep track of:
    // * low - lowest number seen so far,
    // * mid - lowest number so far that is bigger than low.
    // when these two are set and a number greater than mid is seen, return true.
    // Return false upon reaching the end of the vector.
    std::optional<int> low = std::nullopt;
    std::optional<int> mid = std::nullopt;

    for (const auto& number: nums) {
      if (!low || number < low) {
        low = std::make_optional<int>(number);
      }
      else if ((number > low) && (!mid || number < mid)) {
        mid = std::make_optional<int>(number);
      }
      else if (mid && number > mid) {
        return true;
      }
    }

    return false;
  }
};

int main() {
  Solution s{};
  std::vector<int> testcase1 = {1, 2, 3, 4, 5};
  assert(s.increasingTriplet(testcase1) == true);

  std::vector<int> testcase2 = {5, 4, 3, 2, 1};
  assert(s.increasingTriplet(testcase2) == false);

  std::vector<int> testcase3 = {2, 1, 5, 0, 4, 6};
  assert(s.increasingTriplet(testcase3) == true);

  std::vector<int> testcase4 = {20, 100, 10, 12, 5, 13};
  assert(s.increasingTriplet(testcase4) == true);

  std::vector<int> testcase5 = {1, 1, -2, 6};
  assert(s.increasingTriplet(testcase5) == false);
}
