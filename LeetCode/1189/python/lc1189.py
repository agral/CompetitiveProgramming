from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy, but a fun one.
# Solved on: 2026-06-22
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        #NUM_CHARS = 'z' - 'a' + 1
        #charCount = [0] * NUM_CHARS
        # let's do this one differently; count only the 'b', 'a', 'l', 'o' & 'n' letters.
        count_b = count_a = count_l = count_o = count_n = 0
        for ch in text:
            if ch == 'a':
                count_a += 1
            elif ch == 'b':
                count_b += 1
            elif ch == 'l':
                count_l += 1
            elif ch == 'n':
                count_n += 1
            elif ch == 'o':
                count_o += 1

        # need two 'l' and two 'o' letters to form a "balloon".
        count_l //= 2
        count_o //= 2

        ans = min(count_b, count_a, count_l, count_o, count_n)
        return ans
