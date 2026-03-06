import unittest

from lc1784 import Solution

class Test_checkOnesSegment(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.checkOnesSegment("1001"), False)
        self.assertEqual(s.checkOnesSegment("10101"), False)
        self.assertEqual(s.checkOnesSegment("110"), True)
        self.assertEqual(s.checkOnesSegment("1"), True)
        self.assertEqual(s.checkOnesSegment("11111111111"), True)
        self.assertEqual(s.checkOnesSegment("10000000000"), True)
        self.assertEqual(s.checkOnesSegment("11111111110"), True)

if __name__ == "__main__":
    unittest.main()
