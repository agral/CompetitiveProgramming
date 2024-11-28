#include <cassert>
#include <climits>
#include <queue>
#include <vector>

// holds r, y, x:
// r: number of removed obstacles
// y, x: position in the grid.
using Tp = std::tuple<int, int, int>;

// note: grid is indexed [y][x], not [x][y] as with some previous problems.
class Solution {
public:
    int minimumObstacles(std::vector<std::vector<int>>& grid) {
        const int DIRS[4][2] = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
        const int HEIGHT = grid.size();
        const int WIDTH = grid[0].size();

        std::priority_queue<Tp, std::vector<Tp>, std::greater<>> minHeap;
        // Holds minimum currently known number of obstacle removals necessary to get to [y][x].
        std::vector<std::vector<int>> removals(HEIGHT, std::vector<int>(WIDTH, INT_MAX));
        removals[0][0] = grid[0][0];
        minHeap.emplace(grid[0][0], 0, 0); // start walking from [y=0][x=0].

        while (!minHeap.empty()) {
            const auto [r, y, x] = minHeap.top();
            minHeap.pop();
            if ((y == HEIGHT - 1) && (x == WIDTH - 1)) {
                return r;
            }

            for (const auto& [dy, dx]: DIRS) {
                int yy = y + dy;
                int xx = x + dx;
                // only proceed with valid indices, inside the grid.
                if ((xx >= 0) && (xx < WIDTH) && (yy >= 0) && (yy < HEIGHT)) {
                    // the new distance is the old distance incremented by 1 if this step is on an obstacle.
                    int rr = r + grid[yy][xx];
                    // do not waste time considering moves that take the walker here with >= removals
                    // as is already known:
                    if (rr < removals[yy][xx]) {
                        removals[yy][xx] = rr; // a cheaper way to get to [y][x] has just been found!
                        minHeap.emplace(rr, yy, xx);
                    }
                }
            }
        }
        return removals[HEIGHT - 1][WIDTH - 1];
    }
};

int main() {
    Solution s;
    {
        std::vector<std::vector<int>> map = {{0, 1, 1}, {1, 1, 0}, {1, 1, 0}};
        assert(s.minimumObstacles(map) == 2);
    }
    {
        std::vector<std::vector<int>> map = {{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}};
        assert(s.minimumObstacles(map) == 0);
    }
}
