package lc1910

import "bytes"

// Complexity: O(len_s * len_part); probably optimal
func removeOccurrences(s string, part string) string {
	len_s := len(s)
	len_part := len(part)
	part_as_slice := []byte(part)

	index_out := 0
	out := []byte(s)
	for index_s := 0; index_s < len_s; index_s++ {
		out[index_out] = s[index_s]
		index_out += 1
		if index_out >= len_part {
			if bytes.Equal(part_as_slice, out[index_out-len_part:index_out]) {
				index_out -= len_part
			}
		}
	}

	return string(out[:index_out])
}
