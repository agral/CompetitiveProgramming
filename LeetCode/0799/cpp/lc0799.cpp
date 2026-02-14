#include <cassert>
#include <vector>

class Solution {
public:
  double champagneTower(int poured, int query_row, int query_glass) {
    // initially all the champagne gets poured directly into the top glass
    std::vector<double> flow_previous_row(1, poured);

    for (std::size_t num_row = 1; num_row <= query_row; num_row++) {
      // there are i+1 glasses in ith row (since rows are 0-indexed), all initially empty.
      std::vector<double> flow_current_row(num_row + 1, 0);

      // for each glass in parent row, calculate how much of champagne spills equally to both children:
      for (std::size_t g = 0; g < num_row; g++) {
        // 1 glass of champagne stays in a top glass, the excess spills equally to two glasses below:
        double excess_flow = 0.5 * (flow_previous_row[g] - 1.0);
        if (excess_flow > 0.0) {
          flow_current_row[g] += excess_flow;
          flow_current_row[g+1] += excess_flow;
        }
      }

      // Just before exiting this loop, current flow becomes previous flow.
      flow_previous_row = flow_current_row;
    }
    return std::min(1.0, flow_previous_row[query_glass]);
  };
};

int main() {
  Solution s;
  assert(s.champagneTower(1, 1, 1) == 0.0);
  assert(s.champagneTower(2, 1, 1) == 0.5);
  assert(s.champagneTower(100000009, 33, 17) == 1.0);
}
