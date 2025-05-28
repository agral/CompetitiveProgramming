from typing import List

# Runtime complexity: O(k * (|edges1| + |edges2|))
# Auxiliary space complexity: O(|edges1| + |edges2|)
class Solution:
    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        graph1 = self._makeGraph(edges1)
        graph2 = self._makeGraph(edges2)
        maxReachableNodesIn2 = 0

        if k > 0:
            for i in range(len(edges2) + 1):
                maxReachableNodesIn2 = max(maxReachableNodesIn2, self._dfs(graph2, i, -1, k-1))
        ans = []
        for i in range(len(edges1) + 1):
            ans.append(maxReachableNodesIn2 + self._dfs(graph1, i, -1, k))
        return ans

    def _makeGraph(self, edges: List[List[int]]) -> List[List[int]]:
        graph = [[] for _ in range(len(edges) + 1)]
        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)
        return graph

    def _dfs(self, graph: List[List[int]], node: int, prev: int, k: int) -> int:
        if k == 0:
            return 1
        ans = 1
        for v in graph[node]:
            if v != prev:
                ans += self._dfs(graph, v, node, k-1)
        return ans
