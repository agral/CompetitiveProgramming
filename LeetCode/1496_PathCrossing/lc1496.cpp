#include <cassert>
#include <set>
#include <string>
#include <utility>

class Solution {
public:
  bool isPathCrossing(std::string path) {
    std::set<std::pair<int, int>> visited;
    std::pair<int, int> location = {0, 0};
    visited.insert(location);
    for (const auto& ch: path) {
      if (ch == 'N') {
        location.second += 1;
      }
      else if (ch == 'S') {
        location.second -= 1;
      }
      else if (ch == 'E') {
        location.first += 1;
      }
      else if (ch == 'W') {
        location.first -= 1;
      }
      if (visited.find(location) != visited.end()) {
        return true;
      }
      visited.insert(location);
    }
    return false;
  }
};

int main() {
  Solution s;
  assert(s.isPathCrossing("NES") == false);
  assert(s.isPathCrossing("NESWW") == true);
}
