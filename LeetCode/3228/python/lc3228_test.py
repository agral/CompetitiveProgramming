import unittest

from lc3228 import Solution

class Test_maxOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxOperations("1001101"), 4)
        self.assertEqual(s.maxOperations("00111"), 0)

if __name__ == "__main__":
    unittest.main()
