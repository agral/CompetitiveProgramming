#include <algorithm>
#include <cassert>
#include <queue>
#include <vector>

struct Reservation {
  int roomNumber;
  long endTime;
  // Why long? It's possible to construct a testcase in such a way that int overflow occurs.
  // Leetcode does that, obviously...
};

class Solution {
public:
  int mostBooked(int n, std::vector<std::vector<int>>& meetings) {
    std::ranges::sort(meetings);
    auto cmp = [](const Reservation& lhs, const Reservation& rhs) {
      return (lhs.endTime == rhs.endTime) ?
          lhs.roomNumber > rhs.roomNumber : lhs.endTime > rhs.endTime;
    };

    std::priority_queue<int, std::vector<int>, std::greater<int>> availableRooms;
    std::priority_queue<Reservation, std::vector<Reservation>, decltype(cmp)> reserved(cmp);
    std::vector<int> useCount(n);

    // Initially all the rooms are available:
    for (int id = 0; id < n; ++id) {
      availableRooms.push(id);
    }

    for (const std::vector<int>& meeting: meetings) {
      const int start = meeting[0];
      const int stop = meeting[1];

      // Handle meetings that have ended before the currently processed meeting starts,
      // by moving them from reserved back to availableRooms queue:
      // NOTE: start is inclusive, but end is not. This was fun to debug.
   // while (!reserved.empty() && reserved.top().endTime < start) {
      while (!reserved.empty() && reserved.top().endTime <= start) {
        availableRooms.push(reserved.top().roomNumber);
        reserved.pop();
      }

      // Handle case where no rooms are currently available - push meetings to later date.
      // This extends the reserved state of the room that is due to become available first:
      if (availableRooms.empty()) {
        const auto [reservedRoom, reservedEnd] = reserved.top();
        reserved.pop();
        reserved.push(Reservation{reservedRoom, reservedEnd + stop - start});
        ++useCount[reservedRoom];
      }
      // Otherwise simply use the lowest numbered currently available room:
      else {
        const int availableRoom = availableRooms.top();
        availableRooms.pop();
        reserved.push(Reservation{availableRoom, stop});
        ++useCount[availableRoom];
      }
    }

    return std::ranges::max_element(useCount) - useCount.begin();
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> meetings = {{0, 10}, {1, 5}, {2, 7}, {3, 4}};
    assert(s.mostBooked(2, meetings) == 0);
  }
  {
    std::vector<std::vector<int>> meetings = {{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}};
    assert(s.mostBooked(3, meetings) == 1);
  }
}
