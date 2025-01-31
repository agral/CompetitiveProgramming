#include <cassert>
#include <iostream>
#include <unordered_set>
#include <vector>

class Solution {
private:
    int m_height;
    int m_width;
    int m_color;
    std::vector<int> m_areas;
public:
    int largestIsland(std::vector<std::vector<int>>& grid) {
        m_height = grid.size();
        m_width = grid[0].size();
        m_areas = {0, 0}; // areas starts pre-filled with two invalid values at indices 0, 1.
        m_color = 2;

        for (int h = 0; h < m_height; h++) {
            for (int w = 0; w < m_width; w++) {
                if (grid[h][w] == 1) {
                    m_areas.push_back(paint(grid, h, w, m_color));
                    m_color++;
                }
            }
        }

        int maxArea = 0;
        for (int h = 0; h < m_height; h++) {
            for (int w = 0; w < m_width; w++) {
                if (grid[h][w] == 0) {
                    std::unordered_set<int> neighboringColors{
                        safelyGetColor(grid, h-1, w),
                        safelyGetColor(grid, h+1, w),
                        safelyGetColor(grid, h, w-1),
                        safelyGetColor(grid, h, w+1)
                    };
                    // Try turning this singular water tile to a land tile.
                    int currArea = 1 + getArea(neighboringColors);
                    maxArea = std::max(maxArea, currArea);
                }
            }
        }

        // Special case: there may be maps where there's no water tiles whatsoever;
        // then the answer is that the entire grid is one connected island!
        if (maxArea == 0) { // only possible when there were no water tiles.
            return m_height * m_width;
        }
        // otherwise the calculated maxArea is indeed the correct answer.
        return maxArea;
    }

private:
    int paint(std::vector<std::vector<int>>& grid, int h, int w, int color) {
        // Don't paint out of bounds, don't paint on water or already painted cells:
        if (h < 0 || h >= m_height || w < 0 || w >= m_width || grid[h][w] != 1) {
            return 0;
        }
        grid[h][w] = m_color;
        return 1 // for this cell, just painted
            + paint(grid, h-1, w, m_color)
            + paint(grid, h+1, w, m_color)
            + paint(grid, h, w-1, m_color)
            + paint(grid, h, w+1, m_color);
    }

    int safelyGetColor(const std::vector<std::vector<int>>& grid, int h, int w) {
        if (h < 0 || h >= m_height || w < 0 || w >= m_width) {
            return 0;
        }
        return grid[h][w];
    }

    int getArea(const std::unordered_set<int>& neighboringColors) {
        int area = 0;
        for (const int color: neighboringColors) {
            area += m_areas[color];
        }
        return area;
    }
};

int main() {
    struct Testcase {
        std::vector<std::vector<int>> grid;
        int expected;
    };
    std::vector<Testcase> testcases {
        { {{1, 0}, {0, 1}}, 3 },
        { {{1, 1}, {1, 0}}, 4 },
        { {{1, 1}, {1, 1}}, 4 },
        { {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 9 },
        { {{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, 2 },
        { {{1, 1, 0}, {0, 0, 0}, {0, 0, 1}}, 3 },
        { {{1, 1, 1}, {0, 0, 0}, {0, 0, 1}}, 5 },
    };
    Solution s;

    for (Testcase& tc: testcases) {
        assert(s.largestIsland(tc.grid) == tc.expected);
    }
    std::cout << "All testcases have passed successfully\n";
}
