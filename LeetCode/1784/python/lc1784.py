from typing import List

# Runtime complexity: O(n), where n is the length of the binary string.
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-03-06
class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        # 1. Skip to the first one
        # 1a. Note: not necessary, constraints: `s[0] is '1'`.
        i = 0
        L = len(s)

        # 2. Skip to the next character that is not '1' (so is '0').
        # If reached the end of string, there is exactly one segment of ones.
        while i < L and s[i] == '1':
            i += 1

        # 3. Now i points at the first '0' of the binary string.
        # Skip to the next '1' or to the end of the string:
        while i < L and s[i] == '0':
            i += 1

        # 4. If i points to the end of the string, there was just one segment of ones.
        # Otherwise, i points to the first '1' that forms the second segment of ones.
        return i == L
