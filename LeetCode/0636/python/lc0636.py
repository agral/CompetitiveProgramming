from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium+; hard to understand what's expected here. The problem itself is easy+.
# Solved on: 2026-02-08
class Solution:
    def exclusiveTime(self, n: int, logs: List[str]) -> List[int]:
        ans = [0] * n
        stack = [] # holds: [oldest_id, ..., freshest_id]
        lastTime = 0

        for log in logs:
            function_id, operation_type, timestamp = log.split(":")
            function_id = int(function_id)
            timestamp = int(timestamp)
            if operation_type == "start":
                if len(stack) > 0:
                    ans[stack[-1]] += timestamp - lastTime
                stack.append(function_id)
                lastTime = timestamp
            else: # in that case, it is guaranteed that operation_type == "end".
                ans[stack[-1]] += timestamp + 1 - lastTime
                stack.pop()
                lastTime = timestamp + 1

        return ans
