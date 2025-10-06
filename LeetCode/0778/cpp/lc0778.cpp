#include <iostream>
#include <queue>
#include <vector>

// Runtime complexity: O(n^2 * log(n^2)).
// Auxiliary space complexity: O(n^2).
// Subjective level: hard, but not too hard. Overall, pretty enjoyable.
class Solution {
public:
    int swimInWater(std::vector<std::vector<int>> grid) {
        const std::vector<std::vector<int>> DIRS = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
        const int SZ = grid.size(); // grid is a rectangle of size SZ x SZ.
        using T = std::tuple<
            int, // height == grid[y][x] || note: the order matters; want height first for proper sorting.
            int, // y
            int  // x
        >;
        std::vector<std::vector<bool>> isReachable(SZ, std::vector<bool>(SZ));
        std::priority_queue<T, std::vector<T>, std::greater<>> minHeap;
        minHeap.emplace(grid[0][0], 0, 0);
        isReachable[0][0] = true;

        // note: it is not guaranteed that heights[0][0] is 0!
        // it might be any different number; but then can only escape the starting tile
        // after time grid[0][0].
        int ans = grid[0][0];

        while (!minHeap.empty()) {
            const auto [height, y, x] = minHeap.top();
            //std::cout << "Processing: y=" << y << ", x=" << x << ", h=" << height << "\n";
            minHeap.pop();
            if (height > ans) {
                ans = height;
            }
            if ((y == SZ-1) && (x == SZ-1)) {
                // just reached the bottom right cell.
                return ans;
            }
            for (const auto& dir: DIRS) {
                int ny = y + dir[0];
                int nx = x + dir[1];
                if ((ny < 0) || (ny >= SZ) || (nx < 0) || (nx >= SZ) || isReachable[ny][nx]) {
                    continue;
                }
                minHeap.emplace(grid[ny][nx], ny, nx);
                isReachable[ny][nx] = true;
            }
        }

        throw -1;
    }
};

int main() {
    struct Testcase {
        std::vector<std::vector<int>> grid;
        int expected;
    };
    Testcase testcases[] = {
        {std::vector<std::vector<int>>{{0, 2}, {1, 3}}, 3},
        {std::vector<std::vector<int>>{
            { 0,  1,  2,  3,  4},
            {24, 23, 22, 21,  5},
            {12, 13, 14, 15, 16},
            {11, 17, 18, 19, 20},
            {10,  9,  8,  7,  6},
        }, 16},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.swimInWater(tc.grid);
        if (actual != tc.expected) {
            std::cout << "Testcase #" << (numGood+numBad+1) << " failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
