package lc1979

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-07-18
func findGCD(nums []int) int {
	smallest := nums[0]
	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < smallest {
			smallest = nums[i]
		} else if nums[i] > largest {
			largest = nums[i]
		}
	}
	return gcd(smallest, largest)
}
