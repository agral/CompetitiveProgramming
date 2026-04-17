from typing import List

# Runtime complexity: O(L + n)
# Auxiliary space complexity: O(L + n)
# Subjective level: medium+
# Solved on: 2026-04-17
class Solution:
    def colorTheArray(self, n: int, queries: List[List[int]]) -> List[int]:
        array = [0] * n
        L = len(queries)
        ans = [0] * L
        matchingColor = 0

        for q, (idx, color) in enumerate(queries):
            if idx > 0:
                if array[idx-1] > 0 and array[idx-1] == array[idx]:
                    matchingColor -= 1
                if array[idx-1] == color:
                    matchingColor += 1
            if idx+1 < n:
                if array[idx+1] > 0 and array[idx] == array[idx+1]:
                    matchingColor -= 1
                if array[idx+1] == color:
                    matchingColor += 1
            array[idx] = color
            ans[q] = matchingColor

        return ans
