class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        kbd = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }

        ans = [];
        curr = [];

        def backtrack(idx):
            if idx == len(digits):
                ans.append("".join(curr))
            else:
                for letter in kbd[digits[idx]]:
                    curr.append(letter)
                    backtrack(idx + 1)
                    curr.pop()

        backtrack(0)
        return ans

