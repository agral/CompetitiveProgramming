from typing import List

class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        NUM_LETTERS = ord('Z') - ord('A') + 1
        count = [0] * NUM_LETTERS
        for char in tiles:
            count[ord(char) - ord('A')] += 1

        def dfs() -> int:
            ans = 0
            for c in range(NUM_LETTERS):
                if count[c] > 0:
                    # Put the character with the offset c at the current position (A=0, B=1 etc.).
                    # This enumerates all the possible distinct sequences, but operates on offsets
                    # instead of actual characters, for simplicity.
                    count[c] -= 1
                    ans += (1 + dfs())
                    count[c] += 1
            return ans

        return dfs()


def main():
    s = Solution()
    assert(s.numTilePossibilities("AAB") == 8)
    assert(s.numTilePossibilities("AAABBC") == 188)
    assert(s.numTilePossibilities("V") == 1)
    print("All testcases passed successfully.")

    from typing import List

if __name__ == "__main__":
    main()
