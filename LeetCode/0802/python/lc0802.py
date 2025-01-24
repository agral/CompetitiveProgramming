from typing import List

class Solution:
    STATE_PRISTINE = 0
    STATE_VISITING = 1
    STATE_DONE = 2

    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        NUM_NODES = len(graph)
        states = [self.STATE_PRISTINE] * NUM_NODES

        def hasCycle(node_number: int) -> bool:
            if states[node_number] == self.STATE_DONE:
                return False # 
            if states[node_number] == self.STATE_VISITING:
                return True # a self-cycle; this node has a path from itself directly back to itself.
            states[node_number] = self.STATE_VISITING
            for n in graph[node_number]:
                if hasCycle(n):
                    return True
            states[node_number] = self.STATE_DONE
            return False

        return [v for v in range(NUM_NODES) if not hasCycle(v)]

def main():
    s = Solution()
    assert(s.eventualSafeNodes([[1,2],[2,3],[5],[0],[5],[],[]]) == [2,4,5,6])
    assert(s.eventualSafeNodes([[1,2,3,4],[1,2],[3,4],[0,4],[]]) == [4])
    print("All testcases passed successfully.")

if __name__ == "__main__":
    main()
