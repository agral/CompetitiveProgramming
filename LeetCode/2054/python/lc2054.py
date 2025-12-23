from typing import List

class Solution:
    # Note: this is a brute force solution which compares every possible event pair
    # -> O(n^2) runtime complexity. This version resulted in a TLE upon submitting it,
    # but it's pretty easy to understand :-)
    def maxTwoEventsBruteForceTLE(self, events: List[List[int]]) -> int:
        # events[i] = [startTime, endTime, value]
        events.sort(key=lambda x: x[0]) # sort the events list in ascending order wrt. start time

        # initially set the answer to the largest value from attending a *single* event.
        ans = max(e[2] for e in events)

        # then check every possible event pair.
        for i in range(len(events)-1):
            for j in range(i+1, len(events)):
                # do not process events that overlap:
                if events[i][1] >= events[j][0]:
                    continue
                ans = max(ans, events[i][2] + events[j][2])

        return ans

    # the actual solution with O(sort) == O(n*logn) runtime complexity.
    # this version passes the challenge verification.
    def maxTwoEvents(self, events: List[List[int]]) -> int:
        # events[i] = [startTime, endTime, value]
        # ev[i] = [time, isStarting, value]
        ev = []
        for s, e, v in events:
            ev.append( (s, True, v) ) # append an "event starts" event
            ev.append( (e+1, False, v) ) # append an "event stops" event

        # sort the events list in ascending order wrt. start time.
        # for ties wrt. start time, stop events are sorted before start events (as False < True).
        ev.sort() 
        maxValuedEventSoFar = 0
        ans = 0

        # this is a modification of a "line sweep algorithm",
        # handles attending a single event too.
        for _, isStarting, value in ev:
            if isStarting:
                ans = max(ans, value + maxValuedEventSoFar)
            else:
                maxValuedEventSoFar = max(maxValuedEventSoFar, value)
        return ans

def main():
    s = Solution()
    assert(s.maxTwoEvents([[1, 3, 2], [4, 5, 2], [2, 4, 3]]) == 4)
    assert(s.maxTwoEvents([[1, 3, 2], [4, 5, 2], [1, 5, 5]]) == 5)
    assert(s.maxTwoEvents([[1, 5, 3], [1, 5, 1], [6, 6, 5]]) == 8)
    print("All tests passed.")

if __name__ == "__main__":
    main()
