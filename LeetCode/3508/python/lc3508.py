from bisect import bisect_left, bisect_right
from dataclasses import dataclass
from typing import List
from collections import deque, defaultdict, Counter

@dataclass(frozen=True)
class Packet:
    source: int
    destination: int
    timestamp: int


class Router:
    def __init__(self, memoryLimit: int) -> None:
        self.memoryLimit = memoryLimit
        self.packets = set()
        self.packetQueue = deque()
        self.destinationTimestamps = defaultdict(list)
        self.processedIndices = Counter()

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        packet = Packet(source, destination, timestamp)
        if packet in self.packets:
            return False

        if len(self.packetQueue) == self.memoryLimit:
            self.forwardPacket()

        self.packetQueue.append(packet)
        self.packets.add(packet)

        if destination not in self.destinationTimestamps:
            self.destinationTimestamps[destination] = []
        self.destinationTimestamps[destination].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.packetQueue:
            return []
        p = self.packetQueue.popleft()
        self.packets.remove(p)
        self.processedIndices[p.destination] += 1
        return [p.source, p.destination, p.timestamp]

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        if destination not in self.destinationTimestamps:
            return 0
        t = self.destinationTimestamps[destination]
        idx = self.processedIndices.get(destination, 0)
        lowerBound = bisect_left(t, startTime, idx)
        upperBound = bisect_right(t, endTime, idx)
        return upperBound - lowerBound
