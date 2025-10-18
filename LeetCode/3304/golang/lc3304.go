package lc3304

// Runtime complexity: O(k)
// Auxiliary space: O(k)
func kthCharacter(k int) byte {
	s := "a"
	for len(s) < k {
		nextrunes := make([]rune, len(s))
		for i, ch := range s {
			// note: can Z be ever achieved in at most 500 steps? Don't think so;
			// a new letter appears every 2^n steps; so it will overflow from z back to a
			// after 2^25 ~= 10^8 chars. I.e. definitely not happening in 500 chars.
			/*
				nextch := 'a'
				if ch != 'z' {
					nextch = rune(int(ch) + 1)
				}
			*/
			nextrunes[i] = rune(int(ch) + 1)
		}
		s += string(nextrunes)
	}
	// note: k is 1-indexed in challenge description, but 0-indexed in golang and any sane language.
	// Final note: this daily question is subjectively of rather poor quality.
	return byte(s[k-1])
}
