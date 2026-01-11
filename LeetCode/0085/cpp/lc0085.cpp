#include <cassert>
#include <stack>
#include <vector>

class Solution {
public:
  int maximalRectangle(std::vector<std::vector<char>>& matrix) {
    const int HEIGHT = matrix.size();
    const int WIDTH = matrix[0].size();
    std::vector<int> histogram(WIDTH);
    int ans = 0;

    for (const std::vector<char>& row: matrix) {
      for (int w = 0; w < WIDTH; ++w) {
        histogram[w] = (row[w] == '0') ? 0 : histogram[w] + 1;
      }
      ans = std::max(ans, getLargestRectangleArea(histogram));
    }

    return ans;
  }

private:
  int getLargestRectangleArea(const std::vector<int>& hist) {
    int ans = 0;
    std::stack<int> stack;
    for (int i = 0; i <= hist.size(); ++i) {
      while (
        !stack.empty() &&
        (i == hist.size() || hist[stack.top()] > hist[i])
      ) {
        const int height = hist[stack.top()];
        stack.pop();
        const int width = (stack.empty()) ? i : i - stack.top() - 1;
        ans = std::max(ans, height * width);
      }
      stack.push(i);
    }
    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<char>> matrix = {
      {'1', '0', '1', '0', '0'},
      {'1', '0', '1', '1', '1'},
      {'1', '1', '1', '1', '1'},
      {'1', '0', '0', '1', '0'},
    };
    assert(s.maximalRectangle(matrix) == 6);
  }
  {
    std::vector<std::vector<char>> matrix = {
      {'0'},
    };
    assert(s.maximalRectangle(matrix) == 0);
  }
  {
    std::vector<std::vector<char>> matrix = {
      {'1'},
    };
    assert(s.maximalRectangle(matrix) == 1);
  }
}
