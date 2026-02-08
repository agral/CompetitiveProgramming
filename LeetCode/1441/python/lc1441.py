from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: easy+
# Solved on: 2026-02-08
class Solution:
    def buildArray(self, target: List[int], n: int) -> List[int]:
        stack = 1
        len_target = len(target)
        len_ans = 0
        ops = []
        while len_ans != len_target:
            ops.append("Push")
            if stack == target[len_ans]:
                len_ans += 1
            else:
                ops.append("Pop")
            stack += 1
        return ops
