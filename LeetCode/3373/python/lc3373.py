from typing import List

# Runtime complexity: O(l1+l2)
# Auxiliary space complexity: O(l1+l2)
class Solution:
    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]]) -> List[int]:
        graph1 = self._makeGraph(edges1)
        graph2 = self._makeGraph(edges2)
        l1 = len(edges1) + 1
        l2 = len(edges2) + 1
        parity1 = [False] * l1
        even1 = self._dfs(graph1, 0, -1, parity1, True)
        even2 = self._dfs(graph2, 0, -1, [False] * l2, True)
        odd1 = l1 - even1
        odd2 = l2 - even2
        max_even2_odd2 = max(even2, odd2)
        return [(even1 if parity1[i] else odd1) + max_even2_odd2 for i in range(l1)]

    def _makeGraph(self, edges: List[List[int]]) -> List[List[int]]:
        graph = [[] for _ in range(len(edges) + 1)]
        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)
        return graph

    """Returns the number of nodes reachable from `node` with even number of steps."""
    def _dfs(self, graph: List[List[int]], node: int, prev: int, parity: List[bool], is_even: bool) -> int:
        ans = 1 if is_even else 0
        parity[node] = is_even
        for v in graph[node]:
            if v != prev:
                ans += self._dfs(graph, v, node, parity, not is_even)
        return ans
