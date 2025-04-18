from collections import defaultdict

class Solution:
    def findItinerary(self, tickets: list[list[str]]) -> list[str]:
        flights = defaultdict(list)

        for departure, destination in sorted(tickets):
            flights[departure].append(destination)

        step = []
        result = []
        curr = "JFK"

        while curr:
            if flights[curr]:
                step.append(curr)
                next = flights[curr].pop(0)
                curr = next
            else:
                result.append(curr)
                curr = step.pop() if step else None
        return result[::-1]
