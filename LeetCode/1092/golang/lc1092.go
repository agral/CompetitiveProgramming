package lc1092

import "strings"

func longestCommonSubsequence(first string, second string) string {
	lf := len(first)
	ls := len(second)

	// dp[i][j] is the longestCommonSubsequence of substrings: first[0..i), second[0..j).
	dp := make([][]string, lf+1)
	for i := range lf + 1 {
		dp[i] = make([]string, ls+1)
	}

	for f := 1; f <= lf; f++ {
		for s := 1; s <= ls; s++ {
			if first[f-1] == second[s-1] {
				dp[f][s] = dp[f-1][s-1] + string(first[f-1])
			} else {
				if len(dp[f-1][s]) > len(dp[f][s-1]) {
					dp[f][s] = dp[f-1][s]
				} else {
					dp[f][s] = dp[f][s-1]
				}
			}
		}
	}

	return dp[lf][ls]
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	i1, i2 := 0, 0 // current indices of str1 and str2, respectively.
	lcs := longestCommonSubsequence(str1, str2)
	var ans strings.Builder

	for _, ch := range lcs {
		// Append all the letters that are not part of the LCS from both strings.
		// There is zero or more of them:
		for str1[i1] != byte(ch) {
			ans.WriteByte(str1[i1])
			i1 += 1
		}
		for str2[i2] != byte(ch) {
			ans.WriteByte(str2[i2])
			i2 += 1
		}
		// Append the letter of the LCS (present in both input strings):
		ans.WriteRune(ch)
		i1 += 1
		i2 += 1

		// Note to self: Why am I forced to use WriteRune & WriteByte - when iterating over strings?
		// (why is it treated as byte when indexed via [], but as a rune when range-looping?)

	}

	// The LCS is already fully added. There may still be some unused letters in str1 or str2; append them now:
	ans.WriteString(str1[i1:])
	ans.WriteString(str2[i2:])
	return ans.String()
}
