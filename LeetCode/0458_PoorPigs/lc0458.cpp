#include <cassert>

class Solution {
public:
  int poorPigs(int buckets, int minutesToDie, int minutesToTest) {
    // Corner case:
    if (buckets == 1) {
      return 0;
    }
    int sliceFactor = 1 + (minutesToTest / minutesToDie);
    buckets--;
    int ans = 0;
    while (buckets > 0) {
      buckets = buckets / sliceFactor;
      ans++;
    }

    return ans;
  }
};

int main() {
  Solution s;
  assert(s.poorPigs(4, 15, 15) == 2);
  assert(s.poorPigs(4, 15, 30) == 2);
}
