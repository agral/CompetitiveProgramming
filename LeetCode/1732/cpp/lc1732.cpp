#include <cassert>
#include <vector>

class Solution {
public:
  int largestAltitude(std::vector<int>& gain) {
    int curr = 0;
    int max = 0;
    for (auto it = gain.begin(); it != gain.end(); it++) {
      curr += *it;
      if (curr > max) {
        max = curr;
      }
    }
    return max;
  }
};

int main() {
  Solution s;
  std::vector<int> example1 = {-5, 1, 5, 0, -7};
  assert(s.largestAltitude(example1) == 1);
  std::vector<int> example2 = {-4, -3, -2, -1, 4, 3, 2};
  assert(s.largestAltitude(example2) == 0);
}
