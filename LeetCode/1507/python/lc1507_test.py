import unittest

from lc1507 import Solution

class Test_reformatDate(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.reformatDate("20th Oct 2052"), "2052-10-20")
        self.assertEqual(s.reformatDate("6th Jun 1933"), "1933-06-06")
        self.assertEqual(s.reformatDate("26th May 1960"), "1960-05-26")

if __name__ == "__main__":
    unittest.main()
