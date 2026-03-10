import unittest

from lc0633 import Solution

class Test_judgeSquareSum(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.judgeSquareSum(0), True) # 0*0 + 0*0 == 0
        self.assertEqual(s.judgeSquareSum(1), True) # 0*0 + 1*1 == 1
        self.assertEqual(s.judgeSquareSum(2), True) # 1*1 + 1*1 == 2
        self.assertEqual(s.judgeSquareSum(3), False) # 3 can not be expressed as sum of two squares.
        self.assertEqual(s.judgeSquareSum(4), True) # 0*0 + 2*2 == 4
        self.assertEqual(s.judgeSquareSum(5), True) # 1^2 + 2^2 = 5
        self.assertEqual(s.judgeSquareSum(6), False) # 6 can not be expressed as sum of two squares.
        self.assertEqual(s.judgeSquareSum(7), False) # 7 can not be expressed as sum of two squares.
        self.assertEqual(s.judgeSquareSum(8), True) # 2*2 + 2*2= 8
        self.assertEqual(s.judgeSquareSum(9), True) # 0*0 + 3*3 = 9
        self.assertEqual(s.judgeSquareSum(10), True) # 1*1 + 3*3 = 10
        self.assertEqual(s.judgeSquareSum(11), False) # 11 can not be expressed as sum of two squares.


if __name__ == "__main__":
    unittest.main()
