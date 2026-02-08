import unittest

from lc0150 import Solution

class Test_evalRPN(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.evalRPN(["2", "1", "+", "3", "*"]), 9)
        self.assertEqual(s.evalRPN(["4", "13", "5", "/", "+"]), 6)
        self.assertEqual(s.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]), 22)

if __name__ == "__main__":
    unittest.main()
