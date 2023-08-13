#include <cassert>
#include <vector>

class Solution {
public:
  int uniquePathsWithObstacles(std::vector<std::vector<int>>& obstacleGrid) {
    // Approach: create a 2D array holding number of ways to get from given field to bottom-right.
    // Base cases: exactly one way from bottom-right (you're already there);
    // For bottom row: there is exactly one way from any bottom cell to bottom-right corner (just go right),
    // but that's just until a first right-most obstacle is met. To the left of that obstacle there are
    // no ways reaching bottom-right corner.
    //
    // Similar logic can be applied to fill right-most column. Starting from bottom-right corner,
    // there is exactly one way to reach bottom-right (go down). When first obstacle is met,
    // this cell and the rest of rightmost column can't reach the bottom-right corner.
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   |   |   | X |      |   |   |   |   | 0 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   |   |   |   |  ==> |   |   |   |   | 1 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    // |   |   | X |   |   |      | 0 | 0 | 0 | 1 | 1 |
    // +---+---+---+---+---+      +---+---+---+---+---+
    //
    // Then, the rest of the array can be filled in the bottom-up manner. For any cell
    // not on bottom row / rightmost column, number of ways to reach bottom-right corner is equal to sum of
    // ways for its bottom and right neighbors. One exception is if this cell contains an obstacle, then no cell
    // can be reached from there - 0 should be set instead.
    // After the entire array is filled, return top-left cell as answer.

    int NUM_ROWS = obstacleGrid.size();
    int NUM_COLUMNS = obstacleGrid[0].size();

    // Note: problem states that the answer will not exceed 2e9, i.e. should fit in a signed int type.
    // BUT: that's just LeetCode's trolling; while the answer will fit there, a particular partial answer
    // from the DP array may exceed this (and then not contribute to the final answer due to smart
    // placement of the obstacles by the testcase's author(s). So, long long it is.
    long long ways[NUM_ROWS][NUM_COLUMNS];


    // bottom row:
    bool isObstacle = false;
    for (int c = NUM_COLUMNS - 1; c >= 0; c--) {
      if (isObstacle) {
        ways[NUM_ROWS - 1][c] = 0;
      } else if (obstacleGrid[NUM_ROWS - 1][c] == 1) {
        isObstacle = true;
        ways[NUM_ROWS - 1][c] = 0;
      } else {
        ways[NUM_ROWS - 1][c] = 1;
      }
    }

    // rightmost column:
    isObstacle = false;
    for (int r = NUM_ROWS - 1; r >= 0; r--) {
      if (isObstacle) {
        ways[r][NUM_COLUMNS - 1] = 0;
      } else if (obstacleGrid[r][NUM_COLUMNS - 1] == 1) {
        isObstacle = true;
        ways[r][NUM_COLUMNS - 1] = 0;
      } else {
        ways[r][NUM_COLUMNS - 1] = 1;
      }
    }

    // Rest of array:
    for (int r = NUM_ROWS - 2; r >= 0; r--) {
      for (int c = NUM_COLUMNS - 2; c >= 0; c--) {
        if (obstacleGrid[r][c] == 1) {
          ways[r][c] = 0;
        } else {
          ways[r][c] = ways[r + 1][c] + ways[r][c + 1];
        }
      }
    }
    
    return ways[0][0];
  }
};

int main() {
  Solution s{};
  std::vector<std::vector<int>> testcase1 = {{0, 0, 0}, {0, 1, 0}, {0, 0, 0}};
  assert(s.uniquePathsWithObstacles(testcase1) == 2);
  std::vector<std::vector<int>> testcase2 = {{0, 1}, {0, 0}};
  assert(s.uniquePathsWithObstacles(testcase2) == 1);
}
