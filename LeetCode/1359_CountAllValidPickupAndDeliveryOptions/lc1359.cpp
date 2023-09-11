#include <cassert>

const int BIG_PRIME = 1e9 + 7;
const int MAX_N = 500;

class Solution {
public:
  Solution() {
    /* Note: f(x) == 1 for x=1, or else:
    *                f(x-1) * x * (2x-1).
    * (this relation can be computed using pen and paper).
    * Approach: precompute all answers, then answer any queries.
    */
    precomputed_answers[1] = 1;
    for (int i = 2; i <= MAX_N; i++) {
      precomputed_answers[i] = (precomputed_answers[i - 1] * (2 * i - 1) * i) % BIG_PRIME;
    }
  }

  int countOrders(int n) {
    return precomputed_answers[n];
  };

private:
  long long precomputed_answers[MAX_N + 1];
};

int main() {
  Solution s;

  assert(s.countOrders(1) == 1);
  assert(s.countOrders(2) == 6);
  assert(s.countOrders(3) == 90);
}
