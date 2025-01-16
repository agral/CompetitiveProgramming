package lc2425

import "fmt"

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums1)%2 == 1 {
		for _, num := range nums2 {
			ans ^= num
		}
	}
	if len(nums2)%2 == 1 {
		for _, num := range nums1 {
			ans ^= num
		}
	}
	return ans
}

// Used to prepare custom testcases. Initial intuition seems to work OK.
func xorAllNumsTheHardWay(nums1 []int, nums2 []int) int {
	ans := 0
	for i := range nums1 {
		for j := range nums2 {
			ans ^= (nums1[i] ^ nums2[j])
		}
	}
	return ans
}

func main() {
	{
		// both even
		nums1 := []int{9, 8, 7, 6}
		nums2 := []int{1, 2, 3, 4}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
	{
		// both even, 1 < 2
		nums1 := []int{9, 8, 7, 6}
		nums2 := []int{1, 2, 3, 4, 5, 6}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
	{
		// both odd
		nums1 := []int{9, 8, 7}
		nums2 := []int{1, 2, 3}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
	{
		// both odd, 1 < 2
		nums1 := []int{9, 8, 7}
		nums2 := []int{1, 2, 3, 33, 333}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
	{
		// 1 even, 2 odd
		nums1 := []int{1, 11, 111, 1111}
		nums2 := []int{1, 2, 3, 33, 333}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
	{
		// 1 odd, 2 even
		nums1 := []int{1, 11, 111, 1111, 11111}
		nums2 := []int{3, 33, 333, 3333}
		xoredHard := xorAllNumsTheHardWay(nums1, nums2)
		fmt.Printf("xorAllNums of %v and %v is %d\n", nums1, nums2, xoredHard)
	}
}
