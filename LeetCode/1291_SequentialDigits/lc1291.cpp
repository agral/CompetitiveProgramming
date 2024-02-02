#include <cassert>
#include <deque>
#include <vector>

class Solution {
public:
  /*
   * All existing sequential integers si such that 10 <= si <= 10^9:
   * 12,       23,       34,      45,     56,    67,   78,  89,
   * 123,      234,      345,     456,    567,   678,  789,
   * 1234,     2345,     3456,    4567,   5678,  6789,
   * 12345,    23456,    34567,   45678,  56789,
   * 123456,   234567,   345678,  456789,
   * 1234567,  2345678,  3456789,
   * 12345678, 23456789,
   * 123456789,
   *
   * In total there are 36 such numbers, so the optimal answer
   * would be to just iterate over a const array of them.
   *
   * In order to generate them:
   * a) create a first-in-first-out queue with initial 2-digit sequential integers,
   * b) let answer be an empty vector of integers
   * b) for each number taken from the queue:
   *   - if number > high or the queue is empty, stop. All the sequential integers have been generated.
   *   - if number >= low, push_back it to the end of the vector.
   *   - if number's last digit < 9, create a new number k = 10 * n + (last_digit) + 1,
   *     and store it at the end of the queue.
   */
  std::vector<int> sequentialDigits(int low, int high) {
    static constexpr int SEQUENTIAL_INTEGERS[] = {
      12,       23,       34,      45,     56,    67,   78,  89,
      123,      234,      345,     456,    567,   678,  789,
      1234,     2345,     3456,    4567,   5678,  6789,
      12345,    23456,    34567,   45678,  56789,
      123456,   234567,   345678,  456789,
      1234567,  2345678,  3456789,
      12345678, 23456789,
      123456789,
    };
    std::vector<int> ans{};
    for (int i = 0; i < 36 /*SEQUENTIAL_INTEGERS.size()*/; i++) {
      if (SEQUENTIAL_INTEGERS[i] >= low && SEQUENTIAL_INTEGERS[i] <= high) {
        ans.push_back(SEQUENTIAL_INTEGERS[i]);
      }
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> expected = {123, 234};
    assert(s.sequentialDigits(100, 300) == expected);
  }
  {
    std::vector<int> expected = {1234, 2345, 3456, 4567, 5678, 6789, 12345};
    assert(s.sequentialDigits(1000, 13000) == expected);
  }
}
