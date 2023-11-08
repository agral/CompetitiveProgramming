#include <algorithm>
#include <cassert>
#include <cmath>

class Solution {
public:
  bool isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
    // Minimum time to reach (fx, fy) from (sx, sy) is max(|fx-sx|, |fy-sy|).
    //   |START|     |     |     |     |
    // --+-----+-----+-----+-----+-----+--
    //   |     |  1  |     |     |     |
    // --+-----+-----+-----+-----+-----+--
    //   |     |     |  2  |  3  | END4|
    // --+-----+-----+-----+-----+-----+--
    //   |     |     |     |     |     |
    // --+-----+-----+-----+-----+-----+--
    //   |     |     |     |     |     |
    //
    // Here, let's assume START=(1, 5) and END=(5, 3). Then min_t = max(|3-5|, |5-1|) = 4.
    // Now, if a cell is reachable in time T, then it's also reachable in time T+1 by taking a detour,
    // and also in time T+2 by taking a step back:
    //
    //   |START|     |     |     |     |              |START|     |     |     |     |
    // --+-----+-----+-----+-----+-----+--          --+-----+-----+-----+-----+-----+--
    //   |     |  1  |     |     |     |              |     | 1,3 |     |     |     |
    // --+-----+-----+-----+-----+-----+--          --+-----+-----+-----+-----+-----+--
    //   |     |  2  |  3  |  4  | END5|              |     |     | 2,4 |  5  | END6|
    // --+-----+-----+-----+-----+-----+--          --+-----+-----+-----+-----+-----+--
    //   |     |     |     |     |     |              |     |     |     |     |     |
    // --+-----+-----+-----+-----+-----+--          --+-----+-----+-----+-----+-----+--
    //   |     |     |     |     |     |              |     |     |     |     |     |
    //       Reaching END in T_min+1                      Reaching END in T_min+2
    //
    // It is easily possible to reach END from START in any time T > T_min
    // by combining these two techniques.
    // So, the END is reachable at time T from start if and only if T is >= T_min.
    // In order to answer, need to find T_min according to the formula above and compare.
    //
    // However, if T_min == 0 (i.e. starting point and ending point are the same,
    // it's only reachable if t =/= 1 (go around starting place, then go back during last turn.
    // It is impossible to go away and back during just one turn.
    int t_min = std::max(std::abs(sx - fx), std::abs(sy - fy));
    if (t_min == 0) {
      return (t != 1); // can go back to starting point in any number of turns other than 1.
    }
    return t >= t_min;
  }
};

int main() {
  Solution s;
  assert(s.isReachableAtTime(2, 4, 7, 7, 6) == true);
  assert(s.isReachableAtTime(2, 4, 7, 7, 5) == true);
  assert(s.isReachableAtTime(2, 4, 7, 7, 4) == false);

  assert(s.isReachableAtTime(3, 1, 7, 3, 3) == false);
  assert(s.isReachableAtTime(3, 1, 7, 3, 4) == true);

  assert(s.isReachableAtTime(0, 0, 0, 0, 1) == false);
  assert(s.isReachableAtTime(0, 0, 0, 0, 3) == true);
  assert(s.isReachableAtTime(0, 0, 0, 0, 4) == true);
}
