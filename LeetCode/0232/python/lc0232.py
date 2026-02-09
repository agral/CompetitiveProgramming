from typing import List

# Subjective level: easy
# Solved on: 2026-02-09
class MyQueue:
    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(n), for storing n items in the queue. Initially, the queue is empty.
    def __init__(self):
        self.queue = []

    # Runtime complexity: O(1) up to O(n), same as for pop().
    def push(self, x: int) -> None:
        self.queue.append(x)

    # Runtime complexity: O(1), for most ops. Depending on the size of queue,
    # might require reallocations; up to O(n) in the negative case. O(1) amortized, probably?
    def pop(self) -> int:
        ans = self.queue[0]
        self.queue = self.queue[1:]
        return ans

    # Runtime complexity: O(1)
    def peek(self) -> int:
        return self.queue[0]

    # Runtime complexity: O(1)
    def empty(self) -> bool:
        return len(self.queue) == 0
