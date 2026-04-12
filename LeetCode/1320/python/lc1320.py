from typing import List
import functools

# Runtime complexity: O(n * (26+1) * (26+1)) == O(n) (yes! an O(n) solution.)
# Auxiliary space complexity: O(27*27*n) == O(n)
# Subjective level: hard.
# Solved on: 2026-04-12
class Solution:
    def minimumDistance(self, word: str) -> int:
        # Initial thoughts:
        # I think it makes no sense to precompute distances between all the possible letters
        # (as it's simple abs(x1-x2) + abs(y1-y2)). But maybe I should go with that approach
        # - in the end it's only 26^2 == 676 entries, and still in O(1)...?
        # The typed word has at most 300 chars, so that seems like more than twice the necessary work.
        # For now proceeding with computing distances "live", as necessary.

        # So how to approach this? It's easy for the first two chars (put the fingers initially there).
        # But what about the following chars? For each char in word[2:]:
        # - check whether it's optimal to put the finger initially there
        #   - this requires keeping the cost of typing the preceeding prefix with only one finger.
        #   - but it might result in getting actually optimal results, e.g. cost("ABABZ") == 3.
        # - some kind of DP array...? Where one dimension is the text written so far, the others
        #   probably positions of the two fingers, maybe some further state...?

        L = len(word)
        # 1. Define the keyboard:
        # KEYS maps each individual key to its (x, y) position on the keyboard.
        FREE = '.'
        KEYS = {'A': (0, 0), 'B': (1, 0), 'C': (2, 0), 'D': (3, 0), 'E': (4, 0), 'F': (5, 0),
                'G': (0, 1), 'H': (1, 1), 'I': (2, 1), 'J': (3, 1), 'K': (4, 1), 'L': (5, 1),
                'M': (0, 2), 'N': (1, 2), 'O': (2, 2), 'P': (3, 2), 'Q': (4, 2), 'R': (5, 2),
                'S': (0, 3), 'T': (1, 3), 'U': (2, 3), 'V': (3, 3), 'W': (4, 3), 'X': (5, 3),
                'Y': (0, 4), 'Z': (1, 4)}
        def getDistance(curr: str, target: str) -> int:
            """Returns the exact distance between two letters on the challenge's keyboard.
               Handles the case of first use being free (encoded as FREE special `curr` key).
            """
            if curr == FREE:
                return 0
            x1, y1 = KEYS[curr]
            x2, y2 = KEYS[target]
            return abs(x1 - x2) + abs(y1 - y2)

        @functools.lru_cache(None)
        def dp(w: int, k1: str, k2: str):
            """Returns the minimumDistance of typing words, when words[0..w) are already written,
               the first finger is currently on k1 letter, and the second finger is on k2 letter.
            """
            if w == L: # the corner case - all the text is already written.
                return 0
            costUsingFirstFinger = getDistance(k1, word[w]) + dp(w+1, word[w], k2)
            costUsingSecondFinger = getDistance(k2, word[w]) + dp(w+1, k1, word[w])
            return min(costUsingFirstFinger, costUsingSecondFinger)

        return dp(0, FREE, FREE)
