from typing import List
from sortedcontainers import SortedList

class UnionFind:
    def __init__(self, n):
        self.roots = list(range(n))

    def find(self, node):
        if self.roots[node] != node:
            self.roots[node] = self.find(self.roots[node])
        return self.roots[node]

    def union(self, lhs, rhs):
        rootL = self.find(lhs)
        rootR = self.find(rhs)
        if rootL != rootR:
            if rootL < rootR:
                self.roots[rootR] = rootL
            else:
                self.roots[rootL] = rootR

# Runtime complexity: ?
# Auxiliary space complexity: ?
# Subjective level: hard.
class Solution:
    def processQueries(self, c: int, connections: List[List[int]], queries: List[List[int]]) -> List[int]:
        # Create a "map" of connected stations; all connected stations share the same root station
        # which is the connected station with the lowest ID.
        uf = UnionFind(c+1)
        for lhs, rhs in connections:
            uf.union(lhs, rhs)

        # Store all the connected stations' info for all the root stations.
        # These include the root stations themselves.
        gridMembers = [SortedList() for _ in range(c+1)]
        for i in range(1, c+1): # all stations are 1-indexed, not 0-indexed. This offset is necessary.
            rootStation = uf.find(i)
            gridMembers[rootStation].add(i)

        ans = []
        for queryType, station in queries:
            rootStation = uf.find(station)
            if queryType == 1:
                if station in gridMembers[rootStation]:
                    ans.append(station)
                elif len(gridMembers[rootStation]) > 0:
                    ans.append(gridMembers[rootStation][0]) # which is not necessarily a rootStation!
                else:
                    ans.append(-1)
            else: # queryType is either 1 or 2, so it has to be 2 in this case.
                gridMembers[rootStation].discard(station)
        return ans
