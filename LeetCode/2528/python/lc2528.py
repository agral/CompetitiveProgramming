from typing import List

# Runtime complexity: ?, but less than O(N*(N+k)).
# Auxiliary space complexity: O(N).
# Subjective level: hard.
class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        N = len(stations)
        S = sum(stations)
        left = min(stations)
        right = k + 1 + S

        def checkForMinPower(stations: List[int], extraStations: int, minPower: int) -> bool:
            """Returns True iff every city is provided with at least minPower; False otherwise."""
            power = sum(stations[:r])

            for n in range(N):
                if n+r < N:
                    power += stations[n+r]

                # All the cities need to have at least minPower; otherwise this solution is insufficient.
                if power < minPower:
                    neededExtraPower = minPower - power
                    # Check whether extra power stations would satisfy neededExtraPower.
                    # If so, plant them in the farthest place to the right (the greedy part):
                    if extraStations < neededExtraPower:
                        return False
                    farRightIdx = min(n+r, N-1)
                    extraStations -= neededExtraPower
                    stations[farRightIdx] += neededExtraPower
                    power += neededExtraPower

                # Subtract power from stations that will no longer be reachable
                # (too far to the left):
                if (n >= r):
                    power -= stations[n-r]

            # When reached all the way to the end, all the cities are sufficiently powered.
            return True

        while left < right:
            mid = (left + right) // 2
            if checkForMinPower(stations.copy(), k, mid):
                left = mid + 1
            else:
                right = mid

        return left-1
