#include <cassert>
#include <queue>
#include <string>
#include <unordered_set>
#include <vector>

class Solution {
public:
    static const int WIDTH = 3;
    static const int HEIGHT = 2;
    static constexpr std::string TARGET = "123450";

    // Four ways to make a swap with zero-tile:
    static constexpr int DELTAS[][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    int slidingPuzzle(std::vector<std::vector<int>>& board) {
        // convert the board from 2D vector of chars to a string representation:
        std::string state = toString(board);

        // handle the degenerated case of the board being already solved from the start:
        if (state == TARGET) {
            return 0;
        }

        // make a hashset holding already known/analyzed states,
        // and a queue of states to continue inspecting.
        // Initially there's only one state; the one we're starting with.
        std::unordered_set<std::string> known{state};
        std::queue<std::string> q{{state}};

        // breadth first search a (kind-of-)graph of all the possible moves:
        for (int stepNum = 1; !q.empty(); ++stepNum) {
            // only handle the *initial* size; the queue will grow with new states,
            // these need to be handled later, with increased stepNum.
            // also, caution: can not use q.size() inside this for-loop, as size() changes inside loop body.
            // I've settled with remembering size at each step's beginning,
            // another option would be to iterate from size()-1 back to 0.
            const int NUM_VARIANTS_TO_CHECK_AT_THE_BEGINNING_OF_STEP = q.size();
            for (int sz = 0; sz < NUM_VARIANTS_TO_CHECK_AT_THE_BEGINNING_OF_STEP; ++sz) {
                std::string s = q.front();
                q.pop();
                // get the x/y position of the blank space:
                int zeroPos = s.find('0');
                int zeroX = zeroPos % WIDTH;
                int zeroY = zeroPos / WIDTH;
                for (const auto& [dx, dy]: DELTAS) {
                    int swappedX = zeroX + dx;
                    int swappedY = zeroY + dy;
                    // Proceed only with valid swaps,
                    // i.e. those contained in the board's size:
                    if (swappedX >= 0 && swappedX < WIDTH && swappedY >= 0 && swappedY < HEIGHT) {
                        int swappedPos = swappedY * WIDTH + swappedX;
                        std::swap(s[swappedPos], s[zeroPos]);
                        // if target string reached, puzzle is solvable and this
                        // is guaranteed to be the least number of steps needed to solve it.
                        if (TARGET == s) {
                            return stepNum;
                        }

                        // if a new state (not seen before) is reached,
                        // add it to the queue so it will be further iterated from
                        if (!known.contains(s)) {
                            known.insert(s);
                            q.push(s);
                        }

                        // swap back to "undo" the last swap. There are 4 directions
                        // in total to consider.
                        std::swap(s[swappedPos], s[zeroPos]);
                    }
                }
            }
        }

        return -1;
    }

    std::string toString(std::vector<std::vector<int>>& board) {
        std::vector<char> state(WIDTH * HEIGHT);
        for (int h = 0; h < HEIGHT; ++h) {
            for (int w = 0; w < WIDTH; ++w) {
                state[h * WIDTH + w] = '0' + board[h][w];
            }
        }
        return std::string(state.begin(), state.end());
    }

    // Hashes the current board state into an int.
    // Range: 123450 to 543210
    // Note: unused in the final solution, chose to go with a string representation.
    int boardToInt(std::vector<std::vector<int>>& board) {
        int ans = 0;
        for (int h = 0; h < HEIGHT; ++h) {
            for (int w = 0; w < WIDTH; ++w) {
                ans *= 10;
                ans += board[h][w];
            }
        }
        return ans;
    }
};

int main() {
    Solution s;
    { // testcase example-1:
        std::vector<std::vector<int>> board {{1, 2, 3}, {4, 0, 5}};
        assert(s.slidingPuzzle(board) == 1);
        //assert(s.slidingPuzzle(board) == 1);
    }
    { // testcase example-2:
        std::vector<std::vector<int>> board {{1, 2, 3}, {5, 4, 0}};
        assert(s.slidingPuzzle(board) == -1);
    }
    { // testcase example-3:
        std::vector<std::vector<int>> board {{4, 1, 2}, {5, 0, 3}};
        assert(s.slidingPuzzle(board) == 5);
    }
}
