package lc2140

func mostPoints(questions [][]int) int64 {
	SZ := len(questions)
	dp := make([]int64, SZ+1)

	for i := SZ - 1; i >= 0; i-- {
		points, brainpower := int64(questions[i][0]), questions[i][1]
		numNextQuestion := i + brainpower + 1
		nextQuestionPoints := int64(0)
		if numNextQuestion < SZ {
			nextQuestionPoints = dp[numNextQuestion]
		}
		dp[i] = max(points+nextQuestionPoints, dp[i+1])
	}
	return dp[0]
}
