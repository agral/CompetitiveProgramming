import unittest

from lc3379 import Solution

class Test_constructTransformedArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.constructTransformedArray([3, -2, 1, 1]), [1, 1, 1, 3])
        self.assertEqual(s.constructTransformedArray([-1, 4, -1]), [-1, -1, 4])

if __name__ == "__main__":
    unittest.main()
