#include <algorithm>
#include <iostream>
#include <set>
#include <vector>

class SegmentTree {
public:
    explicit SegmentTree(const std::vector<int>& nums)
    : m_numSegments(nums.size() - 1)
    , m_values(nums)
    , m_coveredCount(4 * (nums.size() - 1))
    , m_coveredWidth(4 * (nums.size() - 1))
    {
    }

    // Adds value to the range: [b, e]. A public method.
    void Add(int b, int e, int value) {
        do_add(0, 0, m_numSegments-1, b, e, value);
    }

    // Returns the covered width of all the values from [0 .. numSegments-1].
    int getCoverage() {
        return m_coveredWidth[0];
    }

private:
    int m_numSegments;
    std::vector<int> m_values;
    std::vector<int> m_coveredCount;
    std::vector<int> m_coveredWidth;

    void do_add(int index, int lo, int hi, int b, int e, int value) {
        if ((b >= m_values[hi+1]) || (e <= m_values[lo])) {
            return;
        }

        if ((b <= m_values[lo]) && (e >= m_values[hi+1])) {
            m_coveredCount[index] += value;
        } else {
            int mid = (lo + hi) / 2;
            do_add(2 * index + 1, lo, mid, b, e, value);
            do_add(2 * index + 2, mid+1, hi, b, e, value);
        }

        if (m_coveredCount[index] > 0) {
            m_coveredWidth[index] = m_values[hi+1] - m_values[lo];
        } else if (lo == hi) {
            m_coveredWidth[index] = 0;
        } else {
            m_coveredWidth[index] = m_coveredWidth[2 * index + 1] + m_coveredWidth[2 * index + 2];
        }
    }
};

// A vector of (y, isStart, xStart, xEnd) tuples.
// Note: isStart has value of 1 for starts, and -1 for ends.
// This simplifies width calculation w.r.t. the approach from 3453's solution (in Golang).
using Geometry = std::vector<std::tuple<int, int, int, int>>;

// Runtime complexity: O(nlogn), where n is the number of segments in the tree;
// generally proportional to the count of the squares.
// Auxiliary space complexity: O(n)
// Subjective level: hard+.
// Solved on: 2026-01-14
class Solution {
public:
    double separateSquares(std::vector<std::vector<int>>& squares) {
        Geometry geometry;
        std::set<int> xs;

        for (const std::vector<int>& square: squares) {
            int xPos = square[0];
            int yPos = square[1];
            int length = square[2];
            geometry.emplace_back(yPos, 1, xPos, xPos + length);
            geometry.emplace_back(yPos + length, -1, xPos, xPos + length);
            xs.insert(xPos);
            xs.insert(xPos + length);
        }

        std::ranges::sort(geometry);

        SegmentTree st({xs.begin(), xs.end()});
        double halfArea = 0.5 * getArea(geometry, xs);

        // won't fit in an int, have to use a long int.
        long area = 0;
        int lastY = 0;

        for (const auto& [y, isStart, xStart, xEnd]: geometry) {
            int coverage = st.getCoverage();
            long deltaArea = coverage * static_cast<long>(y - lastY);
            if (area + deltaArea >= halfArea) {
                return lastY + (halfArea - area) / static_cast<double>(coverage);
            }

            // otherwise keep closing in on the actual half area:
            area += deltaArea;
            st.Add(xStart, xEnd, isStart);
            lastY = y;
        }

        // Should never be reached, but return a nonsense answer instead of throwing or doing other stuff.
        return -42.0;
    }

private:
    long getArea(const Geometry& geometry, const std::set<int>& xs) {
        SegmentTree st({xs.begin(), xs.end()});
        long ans = 0.0;
        int lastY = 0;
        for (const auto& [y, isStart, xStart, xEnd]: geometry) {
            ans += static_cast<long>(st.getCoverage()) * static_cast<long>(y - lastY);
            st.Add(xStart, xEnd, isStart);
            lastY = y;
        }

        return ans;
    }
};

int main() {
    struct Testcase {
        std::vector<std::vector<int>> squares;
        double expected;
    };
    Testcase testcases[] = {
        {{{0, 0, 1}, {2, 2, 1}}, 1.0},
        {{{0, 0, 2}, {1, 1, 1}}, 1.0},
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.separateSquares(tc.squares);
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
