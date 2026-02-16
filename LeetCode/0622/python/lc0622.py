from typing import List

# Runtime complexity: O(1), all operations.
# Auxiliary space complexity: O(k): __init__(), others are O(1).
# Subjective level: medium
# Solved on: 2026-02-16
class MyCircularQueue:
    def __init__(self, k: int):
        self.MAX = k
        self.buf = [0] * k
        self.size = 0
        self.front = 0
        self.rear = k-1

    def Front(self) -> int:
        return self.buf[self.front] if self.size > 0 else -1

    def Rear(self) -> int:
        return self.buf[self.rear] if self.size > 0 else -1

    def isEmpty(self) -> bool:
        return self.size == 0

    def isFull(self) -> bool:
        return self.size == self.MAX

    def enQueue(self, value: int) -> bool:
        if self.size == self.MAX:
            return False
        self.rear = (self.rear + 1) % self.MAX
        self.buf[self.rear] = value
        self.size += 1
        return True

    def deQueue(self) -> bool:
        if self.size == 0:
            return False
        self.front = (self.front + 1) % self.MAX
        self.size -= 1
        return True
