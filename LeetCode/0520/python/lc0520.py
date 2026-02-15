# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy+ (some corner cases that are pretty sneaky)
# Solved on: 2026-02-15
class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        isFirstCapital = word[0] >= 'A' and word[0] <= 'Z'
        isAllCapital = isFirstCapital
        for i in range(1, len(word)):
            if word[i] >= 'A' and word[i] <= 'Z':
                if not isAllCapital:
                    # Fail on the first capital letter seen after a non-capital letter
                    # has already been seen.
                    return False
            else:
                # When the first non-capital letter is seen when isAllCapital holds, it's only okay if
                # it's the second letter (so for example "Abc" is valid, but "AAAbc" is not).
                if isAllCapital and i >= 2:
                    return False
                # So when it's the second letter, all the remaining letters must be non-capital too.
                # Mark this by setting all-capital flag to false.
                isAllCapital = False
        return True
