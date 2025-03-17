import unittest

from lc2206 import Solution

class Test_divideArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.divideArray([3, 2, 3, 2, 2, 2]), True)
        self.assertEqual(s.divideArray([1, 2, 3, 4]), False)

if __name__ == "__main__":
    unittest.main()
