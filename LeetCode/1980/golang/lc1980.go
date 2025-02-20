package lc1980

func findDifferentBinaryString(nums []string) string {
	length := len(nums)
	ans := make([]byte, 0, length)
	for idx := 0; idx < length; idx++ {
		if nums[idx][idx] == '0' {
			ans = append(ans, '1')
		} else {
			ans = append(ans, '0')
		}
	}
	return string(ans)
}
