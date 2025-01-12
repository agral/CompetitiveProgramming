package lc2116

func canBeValid(s string, locked string) bool {
	reverse := func(rn []rune) string {
		for l, r := 0, len(rn)-1; l < r; l, r = l+1, r-1 {
			rn[l], rn[r] = rn[r], rn[l]
		}
		return string(rn)
	}

	isValid := func(isGoingForward bool) bool {
		numChangeable := 0
		numOpen, numClosed := 0, 0
		for i := 0; i < len(s); i++ {
			if locked[i] == '0' {
				numChangeable++ // and don't count it as neither Open nor Closed.
			} else if s[i] == '(' {
				numOpen++
			} else { //if s[i] == ')', which is certain at that point.
				numClosed++
			}
			if isGoingForward && numChangeable+numOpen-numClosed < 0 {
				return false
			}
			if !isGoingForward && numChangeable+numClosed-numOpen < 0 {
				return false
			}
		}
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	if !isValid(true) {
		return false
	}

	s = reverse([]rune(s))
	locked = reverse([]rune(locked))
	return isValid(false)
}
