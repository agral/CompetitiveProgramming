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
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        uf = UnionFind(len(edges)+1)

        for edge in edges:
            if not uf.union_by_rank(edge[0], edge[1]):
                return edge
        raise RuntimeError("Redundant edge has not been detected")


def main():
    s = Solution()
    assert(s.findRedundantConnection([[1, 2], [1, 3], [2, 3]]) == [2, 3])
    assert(s.findRedundantConnection([[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]]) == [1, 4])
    print("All testcases successfully passed.")

if __name__ == "__main__":
    main()
