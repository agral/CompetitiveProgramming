import unittest

from lc1758 import Solution

class Test_minOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minOperations("0100"), 1)
        self.assertEqual(s.minOperations("10"), 0)
        self.assertEqual(s.minOperations("1111"), 2)

if __name__ == "__main__":
    unittest.main()
