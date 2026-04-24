# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-04-24
class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        countL, countR, countAny = 0, 0, 0
        for ch in moves:
            if ch == 'L':
                countL += 1
            elif ch == 'R':
                countR += 1
            else:
                countAny += 1
        return abs(countL-countR) + countAny
