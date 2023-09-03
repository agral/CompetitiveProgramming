#include <cassert>
#include <vector>

class Solution {
public:
  int uniquePaths(int m, int n) {
    // Approach: similar to what I did in 0063/UniquePaths2, except there are no obstacles in this case.
    // Approach: create a 2D array holding number of ways to get from given field to bottom-right.
    // Base cases: exactly one way from bottom-right (you're already there);
    // For bottom row: there is exactly one way from any bottom cell to bottom-right corner (just go right).
    //
    // Similar logic can be applied to fill right-most column. Starting from bottom-right corner,
    // there is exactly one way to reach bottom-right (go down). 
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   |   |   |   |      |   |   |   |   | 1 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   |   |   |   |  ==> |   |   |   |   | 1 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   |   |   |   |      | 1 | 1 | 1 | 1 | 1 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    //
    // Then, the rest of the array can be filled in the bottom-up manner. For any cell
    // not on bottom row / rightmost column, number of ways to reach bottom-right corner is equal to sum of
    // ways for its bottom and right neighbors.
    // After the entire array is filled, return top-left cell as answer.

    std::vector<std::vector<int>> num_ways;
    num_ways.resize(m);
    for (auto& row: num_ways) {
      row.resize(n);
    }

    // bottom row:
    for (int c = n - 1; c >= 0; c--) {
      num_ways[m - 1][c] = 1;
    }

    // rightmost column:
    for (int r = m - 2; r >= 0; r--) { // m-2 instead of m-1, as bottom-right was already filled by "bottom row" pass.
      num_ways[r][n - 1] = 1;
    }

    // Rest of array:
    for (int r = m - 2; r >= 0; r--) {
      for (int c = n - 2; c >= 0; c--) {
        num_ways[r][c] = num_ways[r + 1][c] + num_ways[r][c + 1];
      }
    }
    
    return num_ways[0][0];
  }
};

int main() {
  Solution s{};

  assert(s.uniquePaths(3, 7) == 28);
  assert(s.uniquePaths(3, 2) == 3);
}
