import unittest

from lc0022 import Solution

class Test_generateParenthesis(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.generateParenthesis(1), ["()"])
        self.assertEqual(sorted(s.generateParenthesis(2)), sorted(["(())", "()()"]))
        self.assertEqual(sorted(s.generateParenthesis(3)),
            sorted(["((()))", "(()())", "()(())", "(())()", "()()()"]))
        self.assertEqual(sorted(s.generateParenthesis(4)),
            sorted(["(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
                    "(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
                    "()(()())", "()(())()", "()()(())", "()()()()"]))

if __name__ == "__main__":
    unittest.main()
