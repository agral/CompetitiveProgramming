package lc3494

// Runtime complexity: O(m*n)
// Auxiliary space: O(1)
// Subjective level: medium
func minTime(skill []int, mana []int) int64 {
	// skill and mana are given as ints, but the result needs to be int64.
	// so first let's make all the values int64 proper:
	sumSkill := int64(0) // also calculate the sum of skills of all the mages.
	skill64 := make([]int64, len(skill))
	for i, v32 := range skill {
		skill64[i] = int64(v32)
		sumSkill += skill64[i]
	}
	mana64 := make([]int64, len(mana))
	for i, m32 := range mana {
		mana64[i] = int64(m32)
	}

	prevWizardReady := sumSkill * mana64[0]
	for m := 1; m < len(mana64); m++ {
		prevPotionReady := prevWizardReady
		// calculate the minimum starting time for w'th wizard to start brewing m'th mana potion.
		for w := len(skill64) - 2; w >= 0; w-- {
			prevPotionReady -= skill64[w+1] * mana64[m-1]
			prevWizardReady -= skill64[w] * mana64[m]
			// but the wizard can only start working when not working on previous potion;
			// i.e. previous potion is done:
			if prevPotionReady > prevWizardReady {
				prevWizardReady = prevPotionReady
			}
		}
		prevWizardReady += sumSkill * mana64[m]
	}
	return prevWizardReady
}
