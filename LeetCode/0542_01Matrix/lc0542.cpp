#include <limits>
#include <cassert>
#include <deque>
#include <vector>
#include <utility>

class Solution {
public:
  std::vector<std::vector<int>> updateMatrix(std::vector<std::vector<int>>& mat) {
    int HEIGHT = mat.size();
    int WIDTH = mat[0].size();
    std::vector<std::vector<int>> ans;
    ans.resize(HEIGHT);
    for (auto& row: ans) {
      row.resize(WIDTH, std::numeric_limits<int>::max());
    }

    std::deque<std::pair<int, int>> next;

    for (int h = 0; h < HEIGHT; h++) {
      for (int w = 0; w < WIDTH; w++) {
        if (mat[h][w] == 0) {
          ans[h][w] = 0;
          next.push_back({h, w});
        }
      }
    }

    while (!next.empty()) {
      std::pair<int, int> coords = next.front();
      next.pop_front();
      int h = coords.first;
      int w = coords.second;
      int dist = ans[h][w] + 1;

      // visit top neighbor
      if (h > 0) {
        if (dist < ans[h-1][w]) {
          ans[h-1][w] = dist;
          next.push_back({h-1, w});
        }
      }
      // visit right neighbor
      if (w < WIDTH - 1) {
        if (dist < ans[h][w+1]) {
          ans[h][w+1] = dist;
          next.push_back({h, w+1});
        }
      }
      // visit bottom neighbor
      if (h < HEIGHT - 1) {
        if (dist < ans[h+1][w]) {
          ans[h+1][w] = dist;
          next.push_back({h+1, w});
        }
      }
      // visit left neighbor
      if (w > 0) {
        if (dist < ans[h][w-1]) {
          ans[h][w-1] = dist;
          next.push_back({h, w-1});
        }
      }
    }

    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<std::vector<int>> testcase1 = {{0, 0, 0}, {0, 1, 0}, {0, 0, 0}};
  std::vector<std::vector<int>> expected_ans1 = {{0, 0, 0}, {0, 1, 0}, {0, 0, 0}};
  assert(s.updateMatrix(testcase1) == expected_ans1);
}
