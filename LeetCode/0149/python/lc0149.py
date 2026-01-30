import collections
from typing import List

# No need; using tuple[int, int] for now.
# class Point:
#    x: int
#    y: int

def recursive_gcd(x: int, y: int) -> int:
    if y == 0:
        return x
    return recursive_gcd(y, x%y)

# Runtime complexity: O(log(max(x, y)))
# Auxiliary space: O(1)
def gcd(x: int, y: int) -> int:
    while y != 0:
        x, y = y, x%y
    return x

# Runtime complexity: O(1)+O(gcd()) == O(log(max(x, y)))
# Auxiliary space: O(1)
# def getSlope(p1: List[int], p2: List[int]) -> List[int]:
def getSlope(p1: List[int], p2: List[int]) -> tuple[int, int]:
    dx, dy = p1[0]-p2[0], p1[1]-p2[1]
    if dx == 0:
        return (0, p1[0])
    if dy == 0:
        return (p1[1], 0)
    g = gcd(dx, dy)
    return (dx//g, dy//g)

class Solution:
    # Runtime complexity: O(n^2)*O(getSlope), which is still O(n^2) in the end.
    # Auxiliary space: O(n)
    # Subjective level: hard
    # Solved on: 2026-01-30
    def maxPoints(self, points: List[List[int]]) -> int:
        ans = 0
        for i1, p1 in enumerate(points):
            slopeToCount = collections.defaultdict(int)
            numRepeatedPoints = 1
            numSameSlopePoints = 0
            for i2 in range(i1+1, len(points)):
                p2 = points[i2]
                if p1 == p2:
                    numRepeatedPoints += 1
                else:
                    slope = getSlope(p1, p2)
                    slopeToCount[slope] += 1
                    numSameSlopePoints = max(numSameSlopePoints, slopeToCount[slope])
            ans = max(ans, numRepeatedPoints + numSameSlopePoints)
        return ans
