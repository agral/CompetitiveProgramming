#include <cassert>
#include <numeric>
#include <vector>

/*
 * The vector holds every number in an arithmetic sequence from 0 to n,
 * where n == vector's length; except one number that is missing from the sequence.
 * Idea to get this number: calculate sum of full sequence from 0 to n
 * - given by the equation n * (n+1) / 2, also calculate the actual sum
 * of all the numbers present in the vector. The difference between these
 * must be equal to the missing number.
 * Note: 1 <= n <= 10^4, so the computation: n * (n+1) / 2 == (n^2 + n) / 2
 * will not overflow an int for every n in that range. Plain int is safe to use.
 */
class Solution {
public:
  int missingNumber(std::vector<int>& nums) {
    const int n = nums.size();
    const int vector_sum = std::accumulate(nums.begin(), nums.end(), 0);
    const int zero_to_n_sum = n * (n + 1) / 2;
    return zero_to_n_sum - vector_sum;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {3, 0, 1};
    assert(s.missingNumber(nums) == 2);
  }
  {
    std::vector<int> nums = {0, 1};
    assert(s.missingNumber(nums) == 2);
  }
  {
    std::vector<int> nums = {9, 6, 4, 2, 3, 5, 7, 0, 1};
    assert(s.missingNumber(nums) == 8);
  }
}
