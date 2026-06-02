package lc3633

// Runtime complexity: O(W) + O(L) ~= O(n) in general. This looks optimal!
// Auxiliary space: O(1)
// Subjective level: easy, also a massive waste of time.
// Solved on: 2026-06-02
func earliestFinishTime(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {

	L := len(landStartTime)
	W := len(waterStartTime)

	// 1. Find the land ride which ends the earliest;
	// then find the water ride that starts at or after landEndTime and ends the fastest.
	bestLandIdx := 0
	bestLandEndTime := landStartTime[bestLandIdx] + landDuration[bestLandIdx]
	for l := 1; l < L; l++ {
		endTime := landStartTime[l] + landDuration[l]
		if endTime < bestLandEndTime {
			bestLandEndTime = endTime
			bestLandIdx = l
		}
	}
	plan1 := 71826576
	for w := range W {
		if waterStartTime[w] <= bestLandEndTime {
			plan1 = min(plan1, bestLandEndTime+waterDuration[w])
		} else {
			plan1 = min(plan1, waterStartTime[w]+waterDuration[w])
		}
	}

	// 2. Find the water ride which ends the earliest; then find the land ride
	// that starts at or after waterEndTime and ends the fastest.
	bestWaterIdx := 0
	bestWaterEndTime := waterStartTime[bestWaterIdx] + waterDuration[bestWaterIdx]
	for w := 1; w < W; w++ {
		endTime := waterStartTime[w] + waterDuration[w]
		if endTime < bestWaterEndTime {
			bestWaterEndTime = endTime
			bestWaterIdx = w
		}
	}
	plan2 := 71826576
	for l := range L {
		if landStartTime[l] <= bestWaterEndTime {
			plan2 = min(plan2, bestWaterEndTime+landDuration[l])
		} else {
			plan2 = min(plan2, landStartTime[l]+landDuration[l])
		}
	}

	// The answer is the minimum of the above candidates.
	return min(plan1, plan2)
}
