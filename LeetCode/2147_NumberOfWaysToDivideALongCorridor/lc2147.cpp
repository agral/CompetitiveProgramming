#include <cassert>
#include <string>

class Solution {
public:
  int numberOfWays(std::string corridor) {
    const int MOD = 1'000'000'007;
    long ans = 1;
    int lastSeat = -1;
    int totalSeats = 0;
    for (int i = 0; i < corridor.size(); i++) {
      if ('S' == corridor[i]) {
        ++totalSeats;
        if ((totalSeats & 1) && (totalSeats >= 3)) {
          ans = (ans * (i - lastSeat)) % MOD;
        }
        lastSeat = i;
      }
    }

    // Corner case: if totalSeats is even, there is no way to put the odd one to use.
    // Corner case: if there are no seats at all, there can be no segments at all.
    if ((totalSeats & 1) == 1 || totalSeats == 0) {
      ans = 0;
    }

    return ans;
  }
};

int main() {
  Solution s;
  assert(s.numberOfWays("SSPPSPS") == 3);
  assert(s.numberOfWays("PPSPSP") == 1);
  assert(s.numberOfWays("S") == 0);
}
