#include <cassert>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <queue>

class Solution {
public:
  int numBusesToDestination(std::vector<std::vector<int>> routes, int source, int target) {
    if (source == target) {
      return 0;
    }

    // map of [bus stop] -> [list of buses that stop there]
    std::unordered_map<int, std::vector<int>> stopSchedule;
    for (int busNum = 0; busNum < routes.size(); busNum++) {
      for (int stopNum = 0; stopNum < routes[busNum].size(); stopNum++) {
        stopSchedule[routes[busNum][stopNum]].push_back(busNum);
      }
    }

    std::unordered_set<int> busesTaken;
    int ans = 0;
    std::queue<int> q({source});
    while (!q.empty()) {
      ans++;
      for (int k = q.size(); k > 0; k--) {
        int this_route = q.front();
        q.pop();

        for (const auto& busNum: stopSchedule[this_route]) {
          if (busesTaken.find(busNum) == busesTaken.end()) {
            busesTaken.insert(busNum);
            for (const auto& next_route: routes[busNum]) {
              if (next_route == target) {
                return ans;
              }
              q.push(next_route);
            }
          }
        }
      }
    }

    return -1;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> routes = {{1, 2, 7}, {3, 6, 7}};
    assert(s.numBusesToDestination(routes, 1, 6) == 2);
  }
  {
    std::vector<std::vector<int>> routes = {{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}};
    assert(s.numBusesToDestination(routes, 15, 12) == -1);
  }
}
