from typing import List

# Runtime complexity: O(len(digits)) == O(n).
# Auxiliary space complexity: O(n) (for the case of digits consisting entirely of `9`s, otherwise O(1).
# Subjective level: easy.
# Solved on: 2026-02-09
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        LEN = len(digits)

        # 1. Check if `digits` contain only `9`s.
        isOnlyNines = True
        for i in range(LEN):
            if digits[i] != 9:
                isOnlyNines = False
                break

        # 2a. If `digits` contain only nines, the answer is `1` followed by `L` `0`s.
        if isOnlyNines:
            return [1] + [0] * LEN

        # 2b. Otherwise, the answer fits in the original array.
        # Modify it by performing the standard long addition of 1,
        # then return it.
        carry = 1
        for i in range(LEN-1, -1, -1):
            digits[i] += carry
            if digits[i] == 10:
                digits[i] = 0 # and carry remains 1, to be used again in the next iteration.
            else:
                break # break this for-loop, as the addition of one is finally done.

        return digits
