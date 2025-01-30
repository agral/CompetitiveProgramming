import queue
from typing import List

class UnionFind:
    def __init__(self, numElements: int):
        self.id = list(range(numElements))
        self.rank = [0] * numElements

    def search(self, node: int) -> int:
        if self.id[node] != node:
            self.id[node] = self.search(self.id[node])
        return self.id[node]

    def union_by_rank(self, u: int, v: int) -> bool:
        a, b = self.search(u), self.search(v)
        if a == b:
            return False
        if self.rank[a] < self.rank[b]:
            self.id[a] = b
        elif self.rank[b] < self.rank[a]:
            self.id[b] = a
        else: # self.rank[a] == self.rank[b]
            self.id[a] = b
            self.rank[b] += 1
        return True


class Solution:
    def magnificentSets(self, n: int, edges: List[List[int]]) -> int:
        uf = UnionFind(len(edges) + 1)
        graph = [[] for _ in range(len(edges) + 1)]
        for edge in edges:
            u, v = edge[0], edge[1]
            graph[u].append(v)
            graph[v].append(u)
            uf.union_by_rank(u, v)

        rootToGroupSize = dict()
        for i in range(n):
            newGroupSize = self.bfs(graph, i)
            if newGroupSize == -1:
                return -1
            root_id = uf.search(i)
            if root_id not in rootToGroupSize:
                rootToGroupSize[root_id] = 0
            rootToGroupSize[root_id] = max(rootToGroupSize[root_id], newGroupSize)

        ans = 0
        for _,v in rootToGroupSize.items():
            ans += v
        return ans

    def bfs(self, graph: List[List[int]], u: int) -> int:
        ans = 0
        q = queue.Queue()
        q.put(u)
        nodeToStep = {}
        nodeToStep[u] = 1

        while not q.empty():
            ans += 1
            u = q.get()
            for v in graph[u]:
                if not v in nodeToStep.keys():
                    nodeToStep[v] = ans + 1
                    q.put(v)
                elif nodeToStep[v] == ans:
                    return -1 # odd number of edges in this cycle; graph is not bipartite.
        return ans

def main():
    s = Solution()
    assert(s.magnificentSets(6, [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]) == 4)
    print("All testcases successfully passed.")

if __name__ == "__main__":
    main()

