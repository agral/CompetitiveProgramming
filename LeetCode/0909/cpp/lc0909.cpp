#include <iostream>
#include <queue>
#include <vector>

// Runtime complexity: O(N_SQUARED)
// Auxiliary space complexity: O(N_SQUARED)
class Solution {
public:
    int snakesAndLadders(std::vector<std::vector<int>>& board) {
        const int N = board.size();
        const int N_SQUARED = N * N;
        std::vector<bool> is_visited(N_SQUARED + 1);

        // boustrophedon style boards? Too confusing. Let's use a 1D flat array
        // representation of a board. Also with 1-based indexing.
        std::vector<int> flat(N_SQUARED + 1);

        for (size_t i = 0; i < N; i++) {
            for (size_t j = 0; j < N; j++) {
                flat[N * (N - 1 - i) + ((N-i) % 2 == 1 ? j+1 : N-j)] = board[i][j];
            }
        }

        std::queue<int> q{{1}};
        for (int stepNum = 1; !q.empty(); stepNum++) {
            for (int s = q.size(); s > 0; --s) {
                int curr = q.front();
                q.pop();
                for (int target = curr + 1; target <= std::min(curr + 6, N_SQUARED); target++) {
                    // `dest` holds a final destination after reaching `target`;
                    // which might be different than `target`, if `target` contains a snake or a ladder.
                    int dest = flat[target] > 0 ? flat[target] : target;
                    if (!is_visited[dest]) {
                        if (dest == N_SQUARED) {
                            return stepNum;
                        }
                        q.push(dest);
                        is_visited[dest] = true;
                    }
                }
            }
        }

        return -1;
    }
};

int main() {
    struct Testcase {
        std::vector<std::vector<int>> board;
        int expected;
    };
    Testcase testcases[] = {
        {
            {{-1, -1, -1, -1, -1, -1},
             {-1, -1, -1, -1, -1, -1},
             {-1, -1, -1, -1, -1, -1},
             {-1, 35, -1, -1, 13, -1},
             {-1, -1, -1, -1, -1, -1},
             {-1, 15, -1, -1, -1, -1}}, 4
        },
        {
            {{-1, -1},
             {-1, 3}}, 1
        },
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.snakesAndLadders(tc.board);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood + numBad) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
