from collections import defaultdict
from typing import List

class Solution:
    def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
        adjacency = defaultdict(list)
        for parent, child in relations:
            adjacency[parent].append(child)

        max_time = {}

        def calc(node):
            if node in max_time:
                return max_time[node]

            ans = time[node-1]
            for neighbor in adjacency[node]:
                ans = max(ans, time[node-1] + calc(neighbor))
            max_time[node] = ans
            return ans

        for k in range(1, n+1):
            calc(k)

        return max(max_time.values())
