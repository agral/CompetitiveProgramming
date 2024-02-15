#include <algorithm>
#include <cassert>
#include <numeric>
#include <vector>

class Solution {
public:
  long long largestPerimeter(std::vector<int>& nums) {
    int numSides = 0;
    long long perimeter = 0LL;

    std::ranges::sort(nums);

    // Calculate sum of the entire vector:
    long long sum_edges = std::accumulate(nums.begin(), nums.end(), 0LL);

    // Find the largest nums[i] that can be used as the longest side of a polygon.
    // This is only possible when ith edge length is less than the sum of all previous edge lengths.
    for (int edge_idx = nums.size() - 1; edge_idx >= 2; --edge_idx) {
      sum_edges -= nums[edge_idx]; // as it was already counted in
      if (sum_edges > nums[edge_idx]) {
        // Can create a polygon where this edge is the largest one.
        return sum_edges + nums[edge_idx];
      }
    }

    // Can not create any polygon with that set of edges.
    return -1;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> nums = {5, 5, 5};
    assert(s.largestPerimeter(nums) == 15);
  }
  {
    std::vector<int> nums = {1, 12, 1, 2, 5, 50, 3};
    assert(s.largestPerimeter(nums) == 12);
  }
  {
    std::vector<int> nums = {5, 5, 50};
    assert(s.largestPerimeter(nums) == -1);
  }
  {
    std::vector<int> nums = {1, 1, 12, 12, 12, 100};
    assert(s.largestPerimeter(nums) == 38);
  }
}
