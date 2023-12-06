#include <cassert>

/**
 * approach: calculate sums accumulated during whole weeks,
 * then add sum accumulated in the last non-full week (0 to 6 days).
 * 1st week:  1 + 2 + 3 + 4 + 5 + 6 + 7 == 4 * 7 == 28
 * 2nd week:  2 + 3 + 4 + 5 + 6 + 7 + 8 == 5 * 7 == 35
 * 3rd week:  (...)                     == 6 * 7 == 42
 * nth week:  (...)                     == (n + 3) * 7.
 * Sum accumulated during full weeks: (4 + 5 + 6 + ... + (n+3)) * 7.
 *
 * Then, in the last non-full week (0 to 6 days):
 * k + 1 + k + 2 + k + 3 + ... + k + 6,
 * where k is num_weeks, and the last term is k + n % 7.
 */
class Solution {
public:
  int totalMoney(int n) {
    const int num_weeks = n / 7;
    const int num_days = n % 7;
    const int accumulated_full_weeks = 7 * arithmetic_sum(4, 3 + num_weeks);
    const int accumulated_last_week = arithmetic_sum(num_weeks + 1, num_weeks + num_days);
    return accumulated_full_weeks + accumulated_last_week;
  }
private:
  // Returns an arithmetic sum of (a, a+1, a+2, ..., b-2, b-1, b).
  int arithmetic_sum(int a, int b) {
    return (b + a) * (b + 1 - a) / 2; // guaranteed to be a whole int.
  }
};

int main() {
  Solution s;
  assert(s.totalMoney(4) == 10);
  assert(s.totalMoney(10) == 37);
  assert(s.totalMoney(20) == 96);
}
