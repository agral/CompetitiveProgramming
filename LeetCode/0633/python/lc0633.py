import math

# Runtime complexity: O(sqrt(c))
# Auxiliary space complexity: O(1)
# Subjective level: medium
# Solved on: 2026-03-10
class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        a = 0
        b = math.floor(math.sqrt(c))
        result = a*a + b*b
        while a <= b:
            if result == c:
                return True
            elif result < c:
                a += 1
            else: # result > c
                b -= 1
            result = a*a + b*b
        return False
