from typing import List

# Runtime complexity: O(V+E)
# Auxiliary space complexity: O(V+E)
# where V = numCourses, E = len(prerequisites)
# Subjective level: medium
# Solved on: 2026-04-16
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        NOT_VISITED = 0
        TEMPORARY_MARK = 1
        PERMANENT_MARK = 2

        graph = [ [] for _ in range(numCourses)]
        marks = [NOT_VISITED] * numCourses

        for (pre, course) in prerequisites:
            graph[course].append(pre)

        def hasCycle(v: int) -> bool:
            if marks[v] == TEMPORARY_MARK:
                return True
            if marks[v] == PERMANENT_MARK:
                return False
            marks[v] = TEMPORARY_MARK
            if any(hasCycle(u) for u in graph[v]):
                return True
            marks[v] = PERMANENT_MARK
            return False

        return not any(hasCycle(num) for num in range(numCourses))
