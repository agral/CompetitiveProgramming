#include <cassert>

const int SIZE = 37 + 1;

struct Trib {
  constexpr Trib() : ans() {
    ans[0] = 0;
    ans[1] = 1;
    ans[2] = 1;
    for (int i = 3; i < SIZE; i++) {
      ans[i] = ans[i-3] + ans[i-2] + ans[i-1];
    }
  }
  unsigned ans[SIZE];
};

class Solution {
public:
  int tribonacci(int n) {
    return trib.ans[n];
  }
  static constexpr auto trib = Trib();
};

int main() {
  Solution s;
  assert(s.tribonacci(4) == 4);
  assert(s.tribonacci(25) == 1389537);
}
