class Solution:
    def minimumRecolors(self, blocks: str, k: int) -> int:
        blacksInWindow = 0
        maxBlacksInWindow = 0
        for i, color in enumerate(blocks):
            if color == 'B':
                blacksInWindow += 1
            if i >= k and blocks[i-k] == 'B':
                blacksInWindow -= 1
            if blacksInWindow > maxBlacksInWindow:
                maxBlacksInWindow = blacksInWindow
        return k - maxBlacksInWindow

def main():
    pass

if __name__ == "__main__":
    main()
