package lc1790

func areAlmostEqual(s1 string, s2 string) bool {
	diffIndices := make([]int, 0, 2)
	numDiffIndices := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			numDiffIndices++
			if numDiffIndices > 2 {
				return false
			}
			diffIndices = append(diffIndices, i)
		}
	}
	switch numDiffIndices {
	case 0:
		return true
	case 1:
		return false
	case 2:
		return s1[diffIndices[0]] == s2[diffIndices[1]] && s1[diffIndices[1]] == s2[diffIndices[0]]
	}

	return true
}
