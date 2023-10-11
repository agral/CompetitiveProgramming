import heapq

class Solution:
    def fullBloomFlowers(self, flowers, people):
        people = [(person, idx) for (idx, person) in enumerate(people)]
        ans = [0] * len(people)
        bloom_count = 0
        bloom_start = [flower[0] for flower in flowers]
        bloom_end = [flower[1] for flower in flowers]
        heapq.heapify(bloom_start)
        heapq.heapify(bloom_end)
        for person, idx in sorted(people):
            while bloom_start and bloom_start[0] <= person:
                heapq.heappop(bloom_start)
                bloom_count += 1
            while bloom_end and bloom_end[0] < person:
                heapq.heappop(bloom_end)
                bloom_count -= 1
            ans[idx] = bloom_count;
        return ans
