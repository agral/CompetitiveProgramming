from typing import List

class Solution:
    # Python does not have signum in its library.
    # There's one in NumPy, but none in vanilla python...
    def sign(self, num: int) -> int:
        return 1 if num > 0 else -1 if num < 0 else 0

    def decrypt(self, code: List[int], k: int) -> List[int]:
        ans = []
        SZ = len(code)
        sign = self.sign(k)
        for i in range(SZ):
            decrypted = 0
            for offset in range(1, abs(k) + 1):
                decrypted += code[(i + sign * offset) % SZ]
            ans.append(decrypted)

        return ans

def main():
    s = Solution()
    assert(s.decrypt([5, 7, 1, 4], 3) == [12, 10, 16, 13])
    assert(s.decrypt([1, 2, 3, 4], 0) == [0, 0, 0, 0])
    assert(s.decrypt([2, 4, 9, 3], -2) == [12, 5, 6, 13])

if __name__ == "__main__":
    main()
