class Solution:
    def countOfSubstrings(self, word: str, k: int) -> int:
        def countSubstringsWithUpTo(k: int) -> int:
            """Return the total number of substrings of `word` that contain
            every vowel with at most `k` consonants"""
            if k < 0:
                return 0

            VOWELS = ['a', 'e', 'i', 'o', 'u']
            ans = 0
            numVowels = 0
            numUniqueVowels = 0
            mostRecentVowel = {}
            windowStartIdx = 0
            for windowEndIdx, ch in enumerate(word):
                if ch in VOWELS:
                    numVowels += 1
                    if ch not in mostRecentVowel or mostRecentVowel[ch] < windowStartIdx:
                        numUniqueVowels += 1
                    mostRecentVowel[ch] = windowEndIdx

                while windowEndIdx + 1 - windowStartIdx - numVowels > k:
                    # Shrink the window by moving the start index until the requirement
                    # of the window containing at most k consonants is met again:
                    if word[windowStartIdx] in VOWELS:
                        numVowels -= 1
                        if mostRecentVowel[word[windowStartIdx]] == windowStartIdx:
                            numUniqueVowels -= 1
                    windowStartIdx += 1

                if numUniqueVowels == len(VOWELS):
                    # Substrings with every vowel and at most k consonants are:
                    # word[s, e]; word[s+1, e], ..., word[{min(mostRecentVowel[v]) for every vowel}, e]
                    # where s is windowStartIdx, e is windowEndIdx.
                    ans += min(mostRecentVowel[v] for v in VOWELS) + 1 - windowStartIdx
            return ans

        return countSubstringsWithUpTo(k) - countSubstringsWithUpTo(k-1)
