#include <cassert>
#include <utility>
#include <map>

const int MOD = 1e9 + 7;
class Solution {
public:
  int numWays(int steps, int arrLen) {
    m_memo.clear();

    // Account only for the range of the array that can be
    // reached and returned from in the limited number of steps:
    m_arrLen = std::min(arrLen, steps/2 + 1);
    return calc(0, steps);
  }

  int calc(int i, int steps) {
    if (steps == 0) {
      return i == 0 ? 1 : 0;
    }

    auto z = m_memo.find({i, steps});
    if (z != m_memo.end()) {
      return z->second;
    }

    // Stay in the current position:
    int ans = calc(i, steps - 1) % MOD;

    // Move to the left:
    if (i > 0) {
      ans += calc(i - 1, steps - 1);
      ans %= MOD;
    }

    // Move to the right:
    if (i < m_arrLen - 1) {
      ans += calc(i + 1, steps - 1);
      ans %= MOD;
    }

    m_memo[{i, steps}] = ans;
    return ans;
  }
  int m_arrLen;
  std::map<std::pair<int, int>, int> m_memo;
};

int main() {
  Solution s;

  assert(s.numWays(3, 2) == 4);
  assert(s.numWays(2, 4) == 2);
  assert(s.numWays(4, 2) == 8);
  assert(s.numWays(1, 1) == 1);
  assert(s.numWays(1, 10000000) == 1);
}
