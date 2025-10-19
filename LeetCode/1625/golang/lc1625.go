package lc1625

func added(s string, a int) string {
	ans := make([]byte, len(s))
	for i := range s {
		if i%2 == 0 {
			ans[i] = s[i]
		} else {
			ans[i] = '0' + (s[i]-'0'+byte(a))%10
		}
	}
	return string(ans)
}

func rotated(s string, b int) string {
	L := len(s)
	return s[L-b:] + s[0:L-b]
}

// Runtime complexity: O(n^2), I think. Quite hard to estimate.
// - main dfs algorithm is O(2n^2) + n*O(n) for copying n strings; that's still O(n^2).
//
// Auxiliary space: O(n^2) - would be O(1) if strings could be modified, but that's not possible in Golang.
// - so O(n) times 2n string copies (for added & rotated versions which are bound by initial string length).
// - that's O(n^2) in the end.
//
// Subjective level: medium/hard
func findLexSmallestString(s string, a int, b int) string {
	seen := map[string]bool{}
	ans := s
	var dfs func(str string)
	dfs = func(str string) {
		_, isSeen := seen[str]
		if isSeen {
			return
		}
		seen[str] = true
		ans = min(ans, str)

		dfs(added(str, a))
		dfs(rotated(str, b))
	}
	dfs(s)
	return ans
}
