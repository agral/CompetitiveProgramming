class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        words = [s for s in sentence.split(" ")]
        for idx in range(len(words)):
            if words[idx].startswith(searchWord):
                return idx + 1 # 1-based indexing required in problem definition.
        return -1

def main():
    s = Solution()
    assert(s.isPrefixOfWord("i love eating burger", "burg") == 4)
    assert(s.isPrefixOfWord("this problem is an easy problem", "pro") == 2)
    assert(s.isPrefixOfWord("i am tired", "you") == -1)

if __name__ == "__main__":
    main()
