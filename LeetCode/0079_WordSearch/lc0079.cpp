#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  bool exist(std::vector<std::vector<char>>& board, std::string word) {
    for (int h = 0; h < board.size(); ++h) {
      for (int w = 0; w < board[0].size(); ++w) {
        if (dfs(board, word, 0, h, w)) {
          return true;
        }
      }
    }
    return false;
  }

private:
  static constexpr char CHARACTER_ALREADY_USED = '@';
  bool dfs(std::vector<std::vector<char>>& board, const std::string& word, int word_idx, int row, int column) {
    if (row < 0 || row >= board.size() || column < 0 || column >= board[0].size()) {
      return false;
    }
    if (board[row][column] != word[word_idx]) {
      return false;
    }
    if (word_idx == word.size() - 1) {
      return true;
    }

    const char memo = board[row][column];
    board[row][column] = CHARACTER_ALREADY_USED;
    const bool isMatched =
        dfs(board, word, word_idx + 1, row + 1, column    ) ||
        dfs(board, word, word_idx + 1, row - 1, column    ) ||
        dfs(board, word, word_idx + 1, row    , column + 1) ||
        dfs(board, word, word_idx + 1, row    , column - 1);
    board[row][column] = memo;
    return isMatched;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<char>> board = {
      {'A', 'B', 'C', 'E'},
      {'S', 'F', 'C', 'S'},
      {'A', 'D', 'E', 'E'},
    };
    assert(s.exist(board, "ABCCED") == true);
  }
  {
    std::vector<std::vector<char>> board = {
      {'A', 'B', 'C', 'E'},
      {'S', 'F', 'C', 'S'},
      {'A', 'D', 'E', 'E'},
    };
    assert(s.exist(board, "SEE") == true);
  }
  {
    std::vector<std::vector<char>> board = {
      {'A', 'B', 'C', 'E'},
      {'S', 'F', 'C', 'S'},
      {'A', 'D', 'E', 'E'},
    };
    assert(s.exist(board, "ABCB") == false);
  }
}
