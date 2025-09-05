from typing import List

# Runtime complexity: O(61) == O(1)
# Auxiliary space complexity: O(1)
# Subjective level: medium
class Solution:
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        for numOperation in range(61):
            result = num1 - (numOperation * num2)
            if numOperation <= result:
                bitcount = result.bit_count()
                if bitcount <= numOperation:
                    return numOperation
        return -1
